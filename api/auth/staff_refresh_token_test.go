package auth

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/looped-dev/cms/api/models"
	test_setup "github.com/looped-dev/cms/api/test_setup"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// run tests
	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestStaffRefreshToken_CreateStaffRefreshTokenSession(t *testing.T) {
	refreshToken := &StaffRefreshToken{
		DBClient: dbClient,
	}
	staff := &models.StaffMember{
		ID:            primitive.NewObjectID(),
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		EmailVerified: true,
	}
	tokenRecord, err := refreshToken.CreateStaffRefreshTokenSession(context.TODO(), staff)
	assert.NoError(t, err)
	assert.NotNil(t, tokenRecord)
	assert.Equal(t, tokenRecord.UserID, staff.ID.Hex())
}

func TestStaffRefreshToken_VerifyStaffRefreshToken(t *testing.T) {
	refreshToken := &StaffRefreshToken{
		DBClient: dbClient,
	}
	staff := &models.StaffMember{
		ID:            primitive.NewObjectID(),
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		EmailVerified: true,
	}
	tokenNewResult, err := refreshToken.CreateStaffRefreshTokenSession(context.TODO(), staff)
	assert.NoError(t, err)
	assert.NotNil(t, tokenNewResult)

	// verify tokenVerifyResult
	tokenVerifyResult, err := refreshToken.VerifyStaffRefreshToken(context.TODO(), staff.ID.Hex(), tokenNewResult.ID.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, tokenVerifyResult)
	assert.EqualValues(t, tokenNewResult, tokenVerifyResult)
}

func TestStaffRefreshToken_InvalidateRefreshToken(t *testing.T) {
	refreshToken := &StaffRefreshToken{
		DBClient: dbClient,
	}
	staff := &models.StaffMember{
		ID:            primitive.NewObjectID(),
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		EmailVerified: true,
	}
	tokenNewResult, err := refreshToken.CreateStaffRefreshTokenSession(context.TODO(), staff)
	assert.NoError(t, err)
	assert.NotNil(t, tokenNewResult)
	// sanity check, invalidateAt should not be set at this point
	assert.Empty(t, tokenNewResult.InvalidatedAt)

	// verify tokenInvalidateToken
	tokenInvalidateToken, err := refreshToken.InvalidateRefreshToken(context.TODO(), tokenNewResult)
	assert.NoError(t, err)
	assert.NotNil(t, tokenInvalidateToken)
	assert.NotEmpty(t, tokenInvalidateToken.InvalidatedAt)
}
