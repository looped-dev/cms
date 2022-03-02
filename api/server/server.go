package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/emails"
	"github.com/looped-dev/cms/api/graph"
	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/utils/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func run(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoDbConnString := configs.GetConfig("MONGODB_CONNSTRING")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbConnString))
	if err != nil {
		return fmt.Errorf("Failed to create a db client: %v", err)
	}

	// if server is new, run initial setup
	setup := db.NewSetup(client)
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
			},
		),
	)

	httpRouter := chi.NewRouter()

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
