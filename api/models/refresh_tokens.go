package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RefreshToken struct {
	// ID is unique per refresh token and used to identify the token claims
	ID        primitive.ObjectID  `json:"id" bson:"_id"`
	UserID    string              `json:"userID" bson:"userID"`
	CreatedAt primitive.Timestamp `json:"createdAt" bson:"createdAt"`
	ExpiresAt primitive.Timestamp `json:"expiresAt,omitempty" bson:"expiresAt,omitempty"`
	// when the refresh token is invalidated, it gets an invalidated at date
	InvalidatedAt primitive.Timestamp `json:"invalidatedAt,omitempty" bson:"invalidatedAt,omitempty"`
}
