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
	"go.mongodb.org/mongo-driver/mongo"
)

func addNewStaffToDB(client *mongo.Client, ctx context.Context, staff *models.Staff) (*models.Staff, error) {
	createdAt := primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	staff.CreatedAt = createdAt
	staff.UpdatedAt = createdAt
	result, err := client.Database("cms").Collection("staff").InsertOne(
		ctx,
		staff,
	)
	staff.ID = result.InsertedID.(primitive.ObjectID)
	return staff, err
}

func fetchStaffFromDB(client *mongo.Client, ctx context.Context, email string) (*models.Staff, error) {
	staff := &models.Staff{}
	err := client.Database("cms").Collection("staff").FindOne(
		ctx,
		bson.M{"email": email},
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
		bson.M{"_id": staff.ID},
		bson.M{
			"$set": staff,
		},
	)
	return err
}
