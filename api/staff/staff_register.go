package staff

import (
	"context"
	"time"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Staff struct {
	client *mongo.Client
}

// StaffRegister creates a new staff (admin users) and returns the Staff object.
func (s *Staff) StaffRegister(input *model.RegisterInput) (*model.Staff, error) {
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
	result, err := s.client.Database("cms").Collection("staff").InsertOne(context.TODO(), staff)
	if err != nil {
		return nil, err
	}
	staff.ID = result.InsertedID.(string)
	return staff, nil
}
