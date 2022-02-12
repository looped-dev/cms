package staff

import (
	"context"
	"fmt"
	"time"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Staff) addNewStaffToDB(ctx context.Context, staffMember *models.StaffMember) (*models.StaffMember, error) {
	createdAt := primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	staffMember.CreatedAt = createdAt
	staffMember.UpdatedAt = createdAt
	result, err := s.DBClient.Database("cms").Collection("staff").InsertOne(
		ctx,
		staffMember,
	)
	staffMember.ID = result.InsertedID.(primitive.ObjectID)
	return staffMember, err
}

func (s Staff) fetchStaffFromDB(ctx context.Context, email string) (*models.StaffMember, error) {
	staffMember := &models.StaffMember{}
	err := s.DBClient.Database("cms").Collection("staff").FindOne(
		ctx,
		bson.M{"email": email},
	).Decode(staffMember)
	if err != nil {
		return nil, err
	}
	return staffMember, nil
}

func (s Staff) updateStaffInDB(ctx context.Context, staffMember *models.StaffMember, input *model.StaffAcceptInviteInput) error {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("Error hashing password: %v", err)
	}
	// update staff object with new details from the user
	staffMember.HashedPassword = hashedPassword
	staffMember.EmailVerified = true
	staffMember.Name = input.Name
	// remove invite code as now user has password
	staffMember.InviteCode = models.InviteCode{}
	staffMember.UpdatedAt = primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	_, err = s.DBClient.Database("cms").Collection("staff").UpdateOne(
		ctx,
		bson.M{"_id": staffMember.ID},
		bson.M{
			"$set": staffMember,
		},
	)
	return err
}
