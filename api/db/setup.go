package db

import (
	"context"
	"fmt"
	"io"

	"github.com/looped-dev/cms/api/setting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewSetupRepository(dbClient *mongo.Client) *SetupRepository {
	return &SetupRepository{
		DBClient: dbClient,
	}
}

// SetupRepository the database with the default values and indexes
type SetupRepository struct {
	DBClient *mongo.Client
}

// Initialize the database with the default values and indexes
func (s *SetupRepository) Initialize(w io.ReadWriter, ctx context.Context) error {
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
	indexes := NewIndexesRepository(s.DBClient)
	if err := indexes.StaffCollectionIndexes(w, ctx); err != nil {
		return err
	}
	// 3. Create capped collection for settings
	if err := s.CreateSettingCollection(w, ctx); err != nil {
		return err
	}
	fmt.Fprintf(w, "🔨 Creating database: %s \n", DefaultDatabaseName)
	return nil
}

// CreateIndexes creates the indexes for the database
func (s *SetupRepository) CreateIndexes() error {
	return nil
}

// CreateInitialStaffMember creates the initial Staff Member for the CMS, the
// initial user is going to be setup via the API
func (s *SetupRepository) CreateInitialStaffMember() error {
	return nil
}

// if database doesn't exist create collection and add necessary indexes
func (s *SetupRepository) ShouldSetupDB(ctx context.Context) (bool, error) {
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
		if name == DefaultDatabaseName {
			isDBFound = true
			break
		}
	}
	return !isDBFound, nil
}

// CreateSettingCollection creates a capped settings collection in the database
// that only stores a single document. This is ideal for storing settings as
// you only want a single document in the collection for the CMSs settings.
func (s *SetupRepository) CreateSettingCollection(w io.ReadWriter, ctx context.Context) error {
	fmt.Fprintf(w, "🔨 creating setting collection \n")
	boolean := true
	maxDocuments := int64(1)
	cappedSize := int64(4096)
	settingsCollectionOptions := options.CreateCollectionOptions{
		Capped:       &boolean,
		MaxDocuments: &maxDocuments,
		SizeInBytes:  &cappedSize,
	}
	err := s.DBClient.Database("cms").CreateCollection(ctx, setting.SettingsCollectionName, &settingsCollectionOptions)
	if err != nil {
		fmt.Fprintf(w, "❌ Error creating capped collection: %v", err)
		return fmt.Errorf("Error creating settings capped collection: %v", err)
	}
	fmt.Fprintf(w, "✅ Setting collection created! \n")
	return nil
}
