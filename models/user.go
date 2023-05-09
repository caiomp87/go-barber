package models

import "time"

// User means the user instance
type User struct {
	ID           string    `json:"id,omitempty" bson:"id,omitempty"`
	Name         string    `json:"Name,omitempty" bson:"Name,omitempty"`
	Email        string    `json:"Email,omitempty" bson:"Email,omitempty"`
	Password     string    `json:"Password,omitempty" bson:"Password,omitempty"`
	PasswordHash string    `json:"PasswordHash,omitempty" bson:"PasswordHash,omitempty"`
	CreatedAt    time.Time `json:"CreatedAt,omitempty" bson:"CreatedAt,omitempty"`
	UpdatedAt    time.Time `json:"UpdatedAt,omitempty" bson:"UpdatedAt,omitempty"`
}
