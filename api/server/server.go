package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/looped-dev/cms/api/auth"
	"github.com/looped-dev/cms/api/constants"
	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/emails"
	"github.com/looped-dev/cms/api/graph"
	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/utils"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func init() {
	// configure viper to use environment variables or looped.config.yaml file,
	// may be expanded to JSON and other config formats with time
	viper.SetConfigName("looped.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func run(ctx context.Context) error {
	port := viper.GetString("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoDbConnString := viper.GetString("MONGODB_CONNSTRING")
	if mongoDbConnString == "" {
		return fmt.Errorf("Please set the MONGODB_CONNSTRING in the config file. This is the connection string to your MongoDB instance.")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbConnString))
	if err != nil {
		return fmt.Errorf("Failed to create a db client: %v", err)
	}

	// if server is new, run initial setup
	setup := db.NewSetupRepository(client)
	if err := setup.Initialize(os.Stdout, ctx); err != nil {
		return err
	}

	mailServer, err := emails.NewSMTPClient()
	if err != nil {
		return fmt.Errorf("Failed to create a mail server: %v", err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Failed to close the db connection: %v", err)
		}
	}()

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB:         client,
					SMTPClient: mailServer,
				},
				Directives: generated.DirectiveRoot{
					HasStaffRole: func(ctx context.Context, obj interface{}, next graphql.Resolver, role models.StaffRole) (interface{}, error) {
						// implement this here
						user := ctx.Value(constants.CurrentlyAuthenticatedUserContextKey)
						if user == nil {
							return nil, utils.NewGraphQLErrorWithError(401, fmt.Errorf("Access Denied!"))
						}
						var userClaims auth.StaffJWTClaims
						var ok bool
						if userClaims, ok = user.(auth.StaffJWTClaims); !ok {
							return nil, utils.NewGraphQLErrorWithError(500, fmt.Errorf("Internal Error"))
						}
						if userClaims.Role != role.String() {
							return nil, utils.NewGraphQLErrorWithError(401, fmt.Errorf("Access Denied!"))
						}
						// or let it pass through
						return next(ctx)
					},
					IsSignedIn: func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
						user := ctx.Value(constants.CurrentlyAuthenticatedUserContextKey)
						if user == nil {
							return nil, utils.NewGraphQLErrorWithError(401, fmt.Errorf("Access Denied!"))
						}
						return next(ctx)
					},
				},
			},
		),
	)

	// customize error messages going to the fronted, adding error codes in the
	// process. For more info: https://gqlgen.com/reference/errors/
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		var customError *utils.GraphQLError
		if errors.As(e, &customError) {
			// message is for UI pressentation purpose
			err.Message = customError.Message
			// the code can be easier to use compare to the message
			err.Extensions["code"] = customError.Code
		}
		return err
	})

	httpRouter := chi.NewRouter()

	// attach authentication router
	httpRouter.Use(authenticationMiddleware(client))

	// Basic CORS
	httpRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	httpRouter.Handle("/", playground.Handler("GraphQL playground", "/query"))
	httpRouter.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, httpRouter)
}

func main() {
	if err := run(context.TODO()); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
