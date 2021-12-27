package staff

import (
	"context"
	"testing"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestStaff_StaffRegister(t *testing.T) {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://looped:root@127.0.0.1:27017"),
	)
	if err != nil {
		t.Fatalf("Failed to create a db client: %v", err)
	}
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
	got, err := StaffRegister(client, staffInput)
	assert.NoError(t, err)
	assert.Equal(t, got.Email, staffInput.Email)
	assert.Equal(t, got.Name, staffInput.Name)
	assert.NotEqual(t, got.Password, staffInput.Password)
}
