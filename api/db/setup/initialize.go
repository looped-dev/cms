package setup

import (
	"context"
	"fmt"
	"io"

	"github.com/looped-dev/cms/api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSetup(dbClient *mongo.Client) *Setup {
	return &Setup{
		DBClient: dbClient,
	}
}

// Setup the database with the default values and indexes
type Setup struct {
	DBClient *mongo.Client
}

// Initialize the database with the default values and indexes
func (s *Setup) Initialize(w io.ReadWriter, ctx context.Context) error {
	shouldSetupDatabse, err := s.ShouldSetupDB(ctx)
	if err != nil {
		return fmt.Errorf("Error checking whether database exists: %v", err)
	}
	if !shouldSetupDatabse {
		fmt.Fprintf(w, "✅ Database is already setup, skipping setup! \n")
		return nil
	}
	// from here, we need run the database setup, this will basically prepare the
	// database for use by looped CMS.
	// Checklist:
	// 1. Create the database
	// 2. Create the indexes
	// 3. Create capped collection for settings
	fmt.Fprintf(w, "🔨 Creating database: %s \n", db.DefaultDatabaseName)
	return nil
}

// CreateIndexes creates the indexes for the database
func (s *Setup) CreateIndexes() error {
	return nil
}

// CreateInitialStaffMember creates the initial Staff Member for the CMS, the
// initial user is going to be setup via the API
func (s *Setup) CreateInitialStaffMember() error {
	return nil
}

// if database doesn't exist create collection and add necessary indexes
func (s *Setup) ShouldSetupDB(ctx context.Context) (bool, error) {
	client := s.DBClient
	listOfDatabases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	// setup if there are not databases
	if len(listOfDatabases) == 0 {
		return true, nil
	}
	// check if the database being used by the CMS exists, if exists, do not setup
	isDBFound := false
	for _, name := range listOfDatabases {
		if name == db.DefaultDatabaseName {
			isDBFound = false
			break
		}
	}
	return isDBFound, nil
}
