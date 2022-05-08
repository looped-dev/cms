package setting

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/graph/model"
	test_setup "github.com/looped-dev/cms/api/test_setup"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client

func TestMain(m *testing.M) {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	testContainer := test_setup.TestContainers{
		Pool: pool,
	}

	var resource *dockertest.Resource
	dbClient, resource, err = testContainer.NewMongoContainer(context.TODO())
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// run setup
	setup := db.NewSetupRepository(dbClient)
	if err := setup.Initialize(os.Stdout, context.TODO()); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	// run tests
	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSettingRepository_Details(t *testing.T) {
	repository := NewSettingRepository(dbClient)
	settings, err := repository.SaveSettings(context.TODO(), model.UpdateSiteSettingsInput{
		SiteName: "Site Name",
		BaseURL:  "http://localhost:3000",
	})
	assert.NotNil(t, settings)
	assert.Nil(t, err)

	settingsDetails, err := repository.Details(context.TODO())
	assert.NotNil(t, settings)
	assert.Nil(t, err)
	assert.Equal(t, settingsDetails, settings)
}

func TestSettingRepository_Exists(t *testing.T) {
	repository := NewSettingRepository(dbClient)
	settings, err := repository.SaveSettings(context.TODO(), model.UpdateSiteSettingsInput{
		SiteName: "Site Name",
		BaseURL:  "http://localhost:3000",
	})
	assert.NotNil(t, settings)
	assert.Nil(t, err)

	exists, err := repository.Exists(context.TODO())
	assert.True(t, exists)
	assert.Nil(t, err)
}

func TestSettingRepository_SaveSettings(t *testing.T) {
	repository := NewSettingRepository(dbClient)
	settings, err := repository.SaveSettings(context.TODO(), model.UpdateSiteSettingsInput{
		SiteName: "Site Name",
		BaseURL:  "http://localhost:3000",
	})
	assert.NotNil(t, settings)
	assert.Nil(t, err)
}

func TestSettingRepository_UpdateSEOSettings(t *testing.T) {
	repository := NewSettingRepository(dbClient)
	initialSetting, err := repository.SaveSettings(context.TODO(), model.UpdateSiteSettingsInput{
		SiteName: "Site Name",
		BaseURL:  "http://localhost:3000",
	})
	assert.NotNil(t, initialSetting)
	assert.Nil(t, err)

	// update settings with SEO Details
	setting, err := repository.UpdateSEOSettings(context.TODO(), model.UpdateSEOSettingsInput{
		Title:       "New Title",
		Description: "New Description",
	})
	assert.NotNil(t, setting)
	assert.Nil(t, err)

	// query the db to check the details were saved successfully
	settingDetails, err := repository.Details(context.TODO())
	assert.NotNil(t, settingDetails)
	assert.Nil(t, err)

	// expect SEO to have saved the SEO details
	assert.Equal(t, *settingDetails.Seo.Title, "New Title")
	assert.Equal(t, *settingDetails.Seo.Description, "New Description")

	// expect existing settings to have not been updated
	assert.Equal(t, settingDetails.SiteName, initialSetting.SiteName)
	assert.Equal(t, settingDetails.BaseURL, initialSetting.BaseURL)
}

func TestSettingRepository_UpdateFacebookCardSettings(t *testing.T) {
	repository := NewSettingRepository(dbClient)
	initialSetting, err := repository.SaveSettings(context.TODO(), model.UpdateSiteSettingsInput{
		SiteName: "Site Name",
		BaseURL:  "http://localhost:3000",
	})
	assert.NotNil(t, initialSetting)
	assert.Nil(t, err)

	// update settings with SEO Details
	setting, err := repository.UpdateFacebookCardSettings(context.TODO(), model.UpdateFacebookCardSettingsInput{
		Title:       "New Title",
		Description: "New Description",
		Type:        "article",
		URL:         "http://localhost:3000",
	})
	assert.NotNil(t, setting)
	assert.Nil(t, err)

	// query the db to check the details were saved successfully
	settingDetails, err := repository.Details(context.TODO())
	assert.NotNil(t, settingDetails)
	assert.Nil(t, err)

	// expect SEO to have saved the SEO details
	assert.Equal(t, *settingDetails.FacebookCard.Title, "New Title")
	assert.Equal(t, *settingDetails.FacebookCard.Description, "New Description")
	assert.Equal(t, *settingDetails.FacebookCard.Type, "article")
	assert.Equal(t, *settingDetails.FacebookCard.URL, "http://localhost:3000")

	// expect existing settings to have not been updated
	assert.Equal(t, settingDetails.SiteName, initialSetting.SiteName)
	assert.Equal(t, settingDetails.BaseURL, initialSetting.BaseURL)
}

func TestSettingRepository_UpdateTwitterCardSettings(t *testing.T) {
	repository := NewSettingRepository(dbClient)
	initialSetting, err := repository.SaveSettings(context.TODO(), model.UpdateSiteSettingsInput{
		SiteName: "Site Name",
		BaseURL:  "http://localhost:3000",
	})
	assert.NotNil(t, initialSetting)
	assert.Nil(t, err)

	// update settings with SEO Details
	site := "@site"
	creator := "@creator"
	setting, err := repository.UpdateTwitterCardSettings(context.TODO(), model.UpdateTwitterCardSettingsInput{
		Title:       "New Title",
		Description: "New Description",
		Card:        "summary_large_image",
		Site:        &site,
		Creator:     &creator,
	})
	assert.NotNil(t, setting)
	assert.Nil(t, err)

	// query the db to check the details were saved successfully
	settingDetails, err := repository.Details(context.TODO())
	assert.NotNil(t, settingDetails)
	assert.Nil(t, err)

	// expect SEO to have saved the SEO details
	assert.Equal(t, *settingDetails.TwitterCard.Title, "New Title")
	assert.Equal(t, *settingDetails.TwitterCard.Description, "New Description")
	assert.Equal(t, *settingDetails.TwitterCard.Card, "summary_large_image")
	assert.Equal(t, *settingDetails.TwitterCard.Site, "@site")
	assert.Equal(t, *settingDetails.TwitterCard.Creator, "@creator")

	// expect existing settings to have not been updated
	assert.Equal(t, settingDetails.SiteName, initialSetting.SiteName)
	assert.Equal(t, settingDetails.BaseURL, initialSetting.BaseURL)
}
