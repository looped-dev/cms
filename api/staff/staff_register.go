package staff

import (
	"context"
	"time"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// StaffRegister creates a new staff (admin users) and returns the Staff object.
func StaffRegister(client *mongo.Client, input *model.StaffRegisterInput) (*models.Staff, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	staff := &models.Staff{
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

// StaffSendInvite creates a new staff, with a specific role and creates an invite
// code and sends an email to the staff member.
func StaffSendInvite(client *mongo.Client, input *model.StaffInviteInput) (*models.Staff, error) {
	panic("not implemented")
}

// StaffAcceptInvite verify invite code and set the new staff password and email
// as verified.
func StaffAcceptInvite(client *mongo.Client, input *model.StaffAcceptInviteInput) (*models.Staff, error) {
	panic("not implemented")
}

// StaffUpdate updates the details of the staff i.e. Name, Email, Role.
func StaffUpdate(client *mongo.Client, input *model.StaffUpdateInput) (*models.Staff, error) {
	panic("not implemented")
}

// StaffDelete soft deletes the staff from the database by adding a delatedAt field.
func StaffDelete(client *mongo.Client, input *model.StaffDeleteInput) (*models.Staff, error) {
	panic("not implemented")
}

// StaffChangePassword update the staff password.
func StaffChangePassword(client *mongo.Client, input *model.StaffChangePasswordInput) (*models.Staff, error) {
	panic("not implemented")
}
