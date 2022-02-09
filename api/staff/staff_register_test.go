package staff

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func TestMain(m *testing.M) {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "5.0",
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=looped",
			"MONGO_INITDB_ROOT_PASSWORD=root",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// set timeout to 5 minutes
	if err := resource.Expire(300); err != nil {
		log.Fatalf("Couldn't setup resource expiration to 5 minutes: %v", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		var err error
		db, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://looped:root@localhost:%s", resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return db.Ping(context.TODO(), nil)
	})

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// create a login user for testing
	staffInput := &model.StaffRegisterInput{
		Name:     "test",
		Email:    "login_test@example.com",
		Password: "password",
	}
	_, errRegisterUser := StaffRegister(db, staffInput)

	if errRegisterUser != nil {
		log.Fatalf("Unable to create users for testing")
	}

	// run tests
	code := m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestStaff_StaffRegister(t *testing.T) {
	staffInput := &model.StaffRegisterInput{
		Name:     "test",
		Email:    "johndoe@example.com",
		Password: "password",
	}
	got, err := StaffRegister(db, staffInput)
	assert.NoError(t, err)
	assert.Equal(t, got.Email, staffInput.Email)
	assert.Equal(t, got.Name, staffInput.Name)
	assert.NotEqual(t, got.HashedPassword, staffInput.Password)
}

func TestStaffSendInvite(t *testing.T) {
	staffInvite := &model.StaffInviteInput{
		Email: "johninvite@example.com",
		Role:  models.StaffRoleEditor,
	}
	staff, err := StaffSendInvite(db, context.TODO(), staffInvite)
	assert.NoError(t, err)
	fetchStaff, err := fetchStaffFromDB(db, context.TODO(), staffInvite.Email)
	assert.NoError(t, err)
	assert.Equal(t, staff.Email, fetchStaff.Email)
	assert.NotEmpty(t, fetchStaff.InviteCode)
	assert.NotEmpty(t, fetchStaff.InviteCode.Expiry)
	// check whether the code expiry time is with 24 hours
	assert.Greater(t, fetchStaff.InviteCode.Expiry.T, uint32(time.Now().Unix())+60*59*24)
}

func TestStaffAcceptInvite(t *testing.T) {
	staffInsert := &models.Staff{
		Email: "johndoe2@example.com",
		InviteCode: models.InviteCode{
			Code: "CODE",
			Expiry: primitive.Timestamp{
				T: uint32(time.Now().Add(time.Hour).Unix()),
			},
		},
	}
	staff, err := addNewStaffToDB(db, context.TODO(), staffInsert)
	assert.NoError(t, err)
	invite := &model.StaffAcceptInviteInput{
		Name:            "John Doe",
		Email:           staffInsert.Email,
		Code:            "CODE",
		Password:        "password",
		ConfirmPassword: "password",
	}
	staffInvite, err := StaffAcceptInvite(db, context.TODO(), invite)
	assert.NoError(t, err)
	assert.Equal(t, staff.ID, staffInvite.ID)
	assert.Equal(t, staff.Email, staffInvite.Email)
	assert.NotNil(t, staff.HashedPassword)
	assert.Empty(t, staffInvite.InviteCode)
}
