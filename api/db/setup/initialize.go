package setup

import "go.mongodb.org/mongo-driver/mongo"

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
func (s *Setup) Initialize() error {
	return nil
}

// CreateIndexes creates the indexes for the database
func (s *Setup) CreateIndexes() error {
	return nil
}

// CreateInitialStaffMember creates the initial Staff Member for the CMS
func (s *Setup) CreateInitialStaffMember() error {
	return nil
}
