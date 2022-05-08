package setting

import (
	"context"
	"strings"
	"time"

	"github.com/looped-dev/cms/api/constants"
	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSettingRepository(dbClient *mongo.Client) *SettingRepository {
	// determine database name and inject it to the repository
	dbName := db.GetDatabaseName()
	return &SettingRepository{
		DBClient: dbClient,
		dbName:   dbName,
	}
}

type SettingRepository struct {
	DBClient *mongo.Client
	dbName   string
}

// Details fetch the settings of the current settings from the database, returns
// nil if none is found.
func (setting *SettingRepository) Details(ctx context.Context) (*model.SiteSettings, error) {
	settings := &model.SiteSettings{}
	// TODO: figure out out to fetch the first record
	err := setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		FindOne(ctx, bson.M{}).
		Decode(settings)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.NewGraphQLError(404, "No settings found")
		}
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	return settings, nil
}

// Exists checks whether settings have been set in the database
func (setting *SettingRepository) Exists(ctx context.Context) (bool, error) {
	count, err := setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		// count all documents, it should return 1 record as the collection is
		// capped to a single collection.
		CountDocuments(ctx, bson.D{})
	if err != nil {
		return false, utils.NewGraphQLErrorWithError(500, err)

	}
	// should only have a single document, as it is a capped collection
	return count == 1, nil
}

// SaveSettings saves the settings to the database. If the settings already, it
// updates existing settings, otherwise it creates a new settings. Also, it
// ensures only a single record will exist in the database.
func (setting *SettingRepository) SaveSettings(ctx context.Context, input model.UpdateSiteSettingsInput) (*model.SiteSettings, error) {
	_, err := setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		InsertOne(ctx, input)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	siteSettings := &model.SiteSettings{}
	err = setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		// using find one to get the first record as this collection is capped and
		// can only contain one record
		FindOne(ctx, bson.M{}).Decode(siteSettings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	return siteSettings, nil
}

// UpdateSEOSettings updates the SEO settings for site
func (setting *SettingRepository) UpdateSEOSettings(ctx context.Context, input model.UpdateSEOSettingsInput) (*model.SiteSettings, error) {
	settings := model.SiteSettings{}
	// we only have one record in the collection, FindOne will return the first
	// record
	err := setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		FindOne(ctx, bson.M{}).
		Decode(&settings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	// Replace existing settings with update settings, this is because a capped
	// collection can not be updated if the size of the document changes, however,
	// we can add a new document and replace the existing one.
	// https://www.mongodb.com/docs/manual/core/capped-collections/#document-size
	settings.ID = primitive.NewObjectID()
	title := strings.TrimSpace(input.Title)
	description := strings.TrimSpace(input.Description)
	settings.Seo = &model.Seo{
		Title:       &title,
		Description: &description,
	}
	settings.UpdatedAt = primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	_, err = setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		InsertOne(ctx, settings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	siteSettings := &model.SiteSettings{}
	err = setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		// using find one to get the first record as this collection is capped and
		// can only contain one record
		FindOne(ctx, bson.M{}).Decode(siteSettings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	return siteSettings, nil
}

// UpdateFacebookCardSettings updates the facebook Open Graph card settings
func (setting *SettingRepository) UpdateFacebookCardSettings(ctx context.Context, input model.UpdateFacebookCardSettingsInput) (*model.SiteSettings, error) {
	settings := model.SiteSettings{}
	// we only have one record in the collection, FindOne will return the first
	// record
	err := setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		FindOne(ctx, bson.M{}).
		Decode(&settings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	// Replace existing settings with update settings, this is because a capped
	// collection can not be updated if the size of the document changes, however,
	// we can add a new document and replace the existing one.
	// https://www.mongodb.com/docs/manual/core/capped-collections/#document-size
	settings.ID = primitive.NewObjectID()
	title := strings.TrimSpace(input.Title)
	description := strings.TrimSpace(input.Description)
	cardType := strings.TrimSpace(input.Type)
	url := strings.TrimSpace(input.URL)
	settings.FacebookCard = &model.FacebookCard{
		Title:       &title,
		Description: &description,
		Type:        &cardType,
		URL:         &url,
	}
	settings.UpdatedAt = primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	_, err = setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		InsertOne(ctx, settings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	siteSettings := &model.SiteSettings{}
	err = setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		// using find one to get the first record as this collection is capped and
		// can only contain one record
		FindOne(ctx, bson.M{}).Decode(siteSettings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	return siteSettings, nil
}

// UpdateTwitterCardSettings updates the twitter card settings
func (setting *SettingRepository) UpdateTwitterCardSettings(ctx context.Context, input model.UpdateTwitterCardSettingsInput) (*model.SiteSettings, error) {
	settings := model.SiteSettings{}
	// we only have one record in the collection, FindOne will return the first
	// record
	err := setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		FindOne(ctx, bson.M{}).
		Decode(&settings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	// Replace existing settings with update settings, this is because a capped
	// collection can not be updated if the size of the document changes, however,
	// we can add a new document and replace the existing one.
	// https://www.mongodb.com/docs/manual/core/capped-collections/#document-size
	settings.ID = primitive.NewObjectID()
	title := strings.TrimSpace(input.Title)
	description := strings.TrimSpace(input.Description)
	creator := strings.TrimSpace(*input.Creator)
	site := strings.TrimSpace(*input.Site)
	card := strings.TrimSpace(input.Card)
	settings.TwitterCard = &model.TwitterCard{
		Title:       &title,
		Description: &description,
		Creator:     &creator,
		Site:        &site,
		Card:        &card,
	}
	settings.UpdatedAt = primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	_, err = setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		InsertOne(ctx, settings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	siteSettings := &model.SiteSettings{}
	err = setting.DBClient.Database(setting.dbName).
		Collection(constants.SettingsCollectionName).
		// using find one to get the first record as this collection is capped and
		// can only contain one record
		FindOne(ctx, bson.M{}).Decode(siteSettings)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	return siteSettings, nil
}
