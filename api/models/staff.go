package models

import (
	"time"
)

type Staff struct {
	ID            string       `json:"id" bson:"_id,omitempty"`
	Name          string       `json:"name" bson:"name,omitempty"`
	Email         string       `json:"email" bson:"email,omitempty"`
	EmailVerified bool         `json:"emailVerified" bson:"emailVerified,omitempty"`
	Roles         []*StaffRole `json:"roles" bson:"roles,omitempty"`
	Password      string       `json:"password" bson:"password,omitempty"`
	CreatedAt     time.Time    `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt" bson:"updatedAt,omitempty"`
}

type StaffRole struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
