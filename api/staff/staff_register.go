package staff

import (
	"context"
	"time"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// StaffRegister creates a new staff (admin users) and returns the Staff object.
func StaffRegister(client *mongo.Client, input *model.RegisterInput) (*model.Staff, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	staff := &model.Staff{
		Name:          input.Name,
		Email:         input.Email,
		Password:      hashedPassword,
		EmailVerified: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	result, err := client.Database("cms").Collection("staff").InsertOne(context.TODO(), staff)
	if err != nil {
		return nil, err
	}
	staff.ID = result.InsertedID.(primitive.ObjectID).String()
	return staff, nil
}
