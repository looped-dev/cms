package staff

import (
	"context"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// StaffLogin login staff (admin users) and returns the Staff object.
func StaffLogin(client *mongo.Client, input *model.LoginInput) (*models.Staff, error) {
	staff := &models.Staff{}
	err := client.Database("cms").Collection("staff").FindOne(
		context.TODO(),
		models.Staff{
			Email: input.Email,
		},
	).Decode(staff)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}
	return staff, nil
}
