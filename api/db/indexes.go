package db

import (
	"context"
	"fmt"
	"io"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultDatabaseName = "cms"

type IndexesRepository struct {
	dbClient *mongo.Client
}

func NewIndexesRepository(dbClient *mongo.Client) *IndexesRepository {
	return &IndexesRepository{
		dbClient: dbClient,
	}
}

func (i IndexesRepository) StaffCollectionIndexes(w io.ReadWriter, ctx context.Context) error {
	fmt.Fprintf(w, "🔨 creating indexes for staff collection \n")
	staffCollectionIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		// add necessary indexes for staff here here
	}
	// create indexes for staff collection
	_, err := i.dbClient.Database(DefaultDatabaseName).
		Collection("staff").
		Indexes().
		CreateMany(ctx, staffCollectionIndexes)
	if err != nil {
		fmt.Fprintf(w, "❌ Error creating indexes for staff collection: %v \n", err)
		return fmt.Errorf("Error creating staff collection indexes: %v", err)
	}
	fmt.Fprintf(w, "✅ Created indexes for staff collection successfully! \n")
	return nil
}
