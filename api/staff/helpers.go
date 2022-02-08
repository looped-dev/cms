package staff

import (
	"context"
	"fmt"
	"time"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func fetchStaffFromDB(client *mongo.Client, ctx context.Context, email string) (*models.Staff, error) {
	staff := &models.Staff{}
	err := client.Database("cms").Collection("staff").FindOne(
		ctx,
		models.Staff{
			Email: email,
		},
	).Decode(staff)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func updateStaffInDB(client *mongo.Client, ctx context.Context, staff *models.Staff, input *model.StaffAcceptInviteInput) error {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("Error hashing password: %v", err)
	}
	// update staff object with new details from the user
	staff.HashedPassword = hashedPassword
	staff.EmailVerified = true
	staff.Name = input.Name
	// remove invite code as now user has password
	staff.InviteCode = models.InviteCode{}
	staff.UpdatedAt = primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	_, err = client.Database("cms").Collection("staff").UpdateOne(
		ctx,
		models.Staff{
			ID: staff.ID,
		},
		staff,
	)
	return err
}

func validateInviteCode(code string, dbCode models.InviteCode) error {
	if code != dbCode.Code {
		return fmt.Errorf("Invalid invite code")
	}
	now := time.Now()
	expiry := time.Unix(int64(dbCode.Expiry.T), 0)
	if now.After(expiry) {
		return fmt.Errorf("Invite code has expired")
	}
	return nil
}
