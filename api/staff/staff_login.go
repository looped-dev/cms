package staff

import (
	"context"

	"github.com/looped-dev/cms/api/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// StaffLogin login staff (admin users) and returns the Staff object.
func StaffLogin(client *mongo.Client, input *model.LoginInput) (*model.Staff, error) {
	staff := &model.Staff{}
	err := client.Database("cms").Collection("staff").FindOne(
		context.TODO(),
		model.Staff{
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
