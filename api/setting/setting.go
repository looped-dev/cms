package setting

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const SettingsCollectionName = "settings"

func NewSetting(dbClient *mongo.Client) *Setting {
	return &Setting{
		DBClient: dbClient,
	}
}

type Setting struct {
	DBClient *mongo.Client
}

// Details fetch the settings of the current settings from the database, returns
// nil if none is found.
func (setting *Setting) Details(ctx context.Context) (*model.SiteSettings, error) {
	settings := &model.SiteSettings{}
	// TODO: figure out out to fetch the first record
	err := setting.DBClient.Database("cms").
		Collection(SettingsCollectionName).
		FindOne(ctx, bson.M{}).
		Decode(settings)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No settings found")
		}
		return nil, err
	}
	return settings, nil
}

// Exists checks whether settings have been set in the database
func (setting *Setting) Exists(ctx context.Context) (bool, error) {
	count, err := setting.DBClient.Database("cms").
		Collection(SettingsCollectionName).
		CountDocuments(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// SaveSettings saves the settings to the database. If the settings already, it
// updates existing settings, otherwise it creates a new settings. Also, it
// ensures only a single record will exist in the database.
func (setting *Setting) SaveSettings(ctx context.Context, input model.SiteSettingsInput) (*model.SiteSettings, error) {
	panic("not implemented!")
}

// CreateSettingCollection creates a capped settings collection in the database
// that only stores a single collection.
func (setting *Setting) CreateSettingCollection(ctx context.Context) error {
	boolean := true
	maxDocuments := int64(1)
	return setting.DBClient.Database("cms").
		CreateCollection(ctx, SettingsCollectionName, &options.CreateCollectionOptions{
			Capped:       &boolean,
			MaxDocuments: &maxDocuments,
		})
}
