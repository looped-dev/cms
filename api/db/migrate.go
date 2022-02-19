package db

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultDatabaseName = "looped"

func CreateIndexes(client *mongo.Client, ctx context.Context, dbName string) error {
	staffIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := client.Database(dbName).Collection("staff").Indexes().CreateMany(ctx, staffIndexes)
	if err != nil {
		return err
	}
	return nil
}
