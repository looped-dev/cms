package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/emails"
	"github.com/looped-dev/cms/api/graph"
	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/utils/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultPort = "8080"

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mongoDbConnString := configs.GetConfig("MONGODB_CONNSTRING")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbConnString))
	if err != nil {
		return fmt.Errorf("Failed to create a db client: %v", err)
	}

	err = db.CreateIndexes(client, context.Background(), db.DefaultDatabaseName)
	if err != nil {
		return err
	}

	mailServer, err := emails.NewSMTPClient()
	if err != nil {
		return fmt.Errorf("Failed to create a mail server: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
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

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
