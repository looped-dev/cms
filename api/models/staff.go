package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Staff struct Staff/Admin members within the CMS
type Staff struct {
	ID            string     `json:"id" bson:"_id,omitempty"`
	Name          string     `json:"name" bson:"name,omitempty"`
	Email         string     `json:"email" bson:"email,omitempty"`
	EmailVerified bool       `json:"emailVerified" bson:"emailVerified,omitempty"`
	Role          StaffRole  `json:"role" bson:"role,omitempty"`
	Password      string     `json:"password" bson:"password,omitempty"`
	InviteCode    InviteCode `json:"inviteCode" bson:"inviteCode,omitempty"`
	CreatedAt     time.Time  `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time  `json:"updatedAt" bson:"updatedAt,omitempty"`
	DeletedAt     time.Time  `json:"deletedAt" bson:"deletedAt,omitempty"`
}

type InviteCode struct {
	Code   string              `json:"code" bson:"code,omitempty"`
	Expiry primitive.Timestamp `json:"expiry" bson:"expiry,omitempty"`
}
