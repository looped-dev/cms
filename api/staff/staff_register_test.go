package staff

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	test_setup "github.com/looped-dev/cms/api/test_setup"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	mail "github.com/xhit/go-simple-mail/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbClient *mongo.Client
var smtpClient *mail.SMTPClient

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

	var resourceSMTP *dockertest.Resource
	smtpClient, resourceSMTP, err = testContainer.NewMailTestServer(context.TODO())
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

	if err = pool.Purge(resourceSMTP); err != nil {
		log.Fatalf("Could not purge SMTP docker container: %s", err)
	}

	os.Exit(code)
}

func TestStaff_StaffRegister(t *testing.T) {
	staffInput := &model.StaffRegisterInput{
		Name:     "test",
		Email:    "johndoe@example.com",
		Password: "password",
	}
	staff := NewStaffRepository(smtpClient, dbClient)
	got, err := staff.StaffRegister(context.TODO(), staffInput)
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
	staff := NewStaffRepository(smtpClient, dbClient)
	staffMember, err := staff.StaffSendInvite(context.TODO(), staffInvite)
	assert.NoError(t, err)
	fetchStaffMember, err := staff.FetchStaffFromDB(context.TODO(), staffInvite.Email)
	assert.NoError(t, err)
	assert.Equal(t, staffMember.Email, fetchStaffMember.Email)
	assert.NotEmpty(t, fetchStaffMember.InviteCode)
	assert.NotEmpty(t, fetchStaffMember.InviteCode.Expiry)
	// check whether the code expiry time is with 24 hours
	assert.Greater(t, fetchStaffMember.InviteCode.Expiry.T, uint32(time.Now().Unix())+60*59*24)
}

func TestStaffAcceptInvite(t *testing.T) {
	staffInsert := &models.StaffMember{
		Email: "johndoeinvite@example.com",
		InviteCode: models.InviteCode{
			Code: "CODE",
			Expiry: primitive.Timestamp{
				T: uint32(time.Now().Add(time.Hour).Unix()),
			},
		},
	}
	staff := NewStaffRepository(smtpClient, dbClient)
	staffMember, err := staff.addNewStaffToDB(context.TODO(), staffInsert)
	assert.NoError(t, err)
	invite := &model.StaffAcceptInviteInput{
		Name:            "John Doe",
		Email:           staffInsert.Email,
		Code:            "CODE",
		Password:        "password",
		ConfirmPassword: "password",
	}
	staffMemberInvite, err := staff.StaffAcceptInvite(context.TODO(), invite)
	assert.NoError(t, err)
	assert.Equal(t, staffMember.ID, staffMemberInvite.ID)
	assert.Equal(t, staffMember.Email, staffMemberInvite.Email)
	assert.NotNil(t, staffMember.HashedPassword)
	assert.Empty(t, staffMemberInvite.InviteCode)
}

func TestStaff_StaffExistsExists(t *testing.T) {
	staffInsert := &models.StaffMember{
		Email: "johndoe_exists@example.com",
		InviteCode: models.InviteCode{
			Code: "CODE",
			Expiry: primitive.Timestamp{
				T: uint32(time.Now().Add(time.Hour).Unix()),
			},
		},
	}
	staff := NewStaffRepository(smtpClient, dbClient)
	_, err := staff.addNewStaffToDB(context.TODO(), staffInsert)
	assert.NoError(t, err)
	count, err := staff.StaffExists(context.TODO())
	assert.NoError(t, err)
	assert.True(t, count)
}
