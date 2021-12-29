package staff

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
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

	resource, err := pool.Run("mongo", "5.0", []string{
		"MONGO_INITDB_ROOT_USERNAME=looped",
		"MONGO_INITDB_ROOT_PASSWORD=root",
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
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

	// run tests
	m.Run()

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func TestStaff_StaffRegister(t *testing.T) {
	type fields struct {
		client *mongo.Client
	}
	type args struct {
		input *model.RegisterInput
	}
	staffInput := &model.RegisterInput{
		Name:     "test",
		Email:    "johndoe@example.com",
		Password: "password",
	}
	got, err := StaffRegister(db, staffInput)
	assert.NoError(t, err)
	assert.Equal(t, got.Email, staffInput.Email)
	assert.Equal(t, got.Name, staffInput.Name)
	assert.NotEqual(t, got.Password, staffInput.Password)
}
