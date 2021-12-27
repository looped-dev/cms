package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/looped-dev/cms/api/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Staff struct {
	client *mongo.Client
}

// hashPassword hashes the password using bcrypt and store the has instead of
// the plain text password.
func hashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("Error hashing password: %w", err)
	}
	return string(hash), nil
}

// Register creates a new user and returns the user object.
func (s *Staff) Register(input *model.RegisterInput) (*model.Staff, error) {
	hashedPassword, err := hashPassword(input.Password)
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
