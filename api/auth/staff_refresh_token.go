package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/looped-dev/cms/api/constants"
	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewStaffRefreshToken(client *mongo.Client) *StaffRefreshTokenRepository {
	return &StaffRefreshTokenRepository{
		dbName:   db.GetDatabaseName(),
		DBClient: client,
	}
}

type StaffRefreshTokenRepository struct {
	DBClient *mongo.Client
	dbName   string
}

// StaffRefreshTokenCreate creates a new staff refresh token and returns the StaffRefreshToken object.
func (srt *StaffRefreshTokenRepository) CreateStaffRefreshTokenSession(ctx context.Context, staff *models.StaffMember) (*models.RefreshToken, error) {
	refreshToken := models.RefreshToken{
		UserID: staff.ID.Hex(),
		ID:     primitive.NewObjectID(),
		CreatedAt: primitive.Timestamp{
			T: uint32(time.Now().Unix()),
		},
		// expires after 1 month
		ExpiresAt: primitive.Timestamp{
			T: uint32(time.Now().Unix() + 60*60*24*30),
		},
		InvalidatedAt: primitive.Timestamp{},
	}
	_, err := srt.DBClient.Database(srt.dbName).
		Collection(constants.StaffRefreshTokenCollection).
		InsertOne(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

// StaffRefreshTokenGetByID returns a staff refresh token by ID.
func (srt *StaffRefreshTokenRepository) VerifyStaffRefreshToken(ctx context.Context, userID, refreshTokenID string) (*models.RefreshToken, error) {
	id, err := primitive.ObjectIDFromHex(refreshTokenID)
	if err != nil {
		return nil, err
	}
	refreshToken := &models.RefreshToken{}
	err = srt.DBClient.Database(srt.dbName).
		Collection(constants.StaffRefreshTokenCollection).
		FindOne(ctx, bson.M{"_id": id}).
		Decode(refreshToken)
	if err != nil {
		return nil, err
	}
	// checks if the refresh token is for the current user
	if refreshToken.UserID != userID {
		// probably consider invalidating the refresh token at this point
		return nil, fmt.Errorf("Refresh token is invalid UserID does not match")
	}
	// check if refresh token is already used
	if !refreshToken.InvalidatedAt.IsZero() {
		return nil, fmt.Errorf("Refresh token already used")
	}
	// check if refresh token has expired
	expiry := time.Unix(int64(refreshToken.ExpiresAt.T), 0)
	if time.Now().After(expiry) {
		return nil, fmt.Errorf("Refresh token has expired")
	}
	return refreshToken, nil
}

// StaffRefreshTokenGetByID invalidates refresh token by adding InvalidatedAt
// timestamp on it and returning it.
func (srt *StaffRefreshTokenRepository) InvalidateRefreshToken(ctx context.Context, refreshToken *models.RefreshToken) (*models.RefreshToken, error) {
	refreshToken.InvalidatedAt = primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	_, err := srt.DBClient.Database(srt.dbName).
		Collection(constants.StaffRefreshTokenCollection).
		UpdateOne(ctx, bson.M{"_id": refreshToken.ID}, bson.M{"$set": refreshToken})
	return refreshToken, err
}
