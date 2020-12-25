package models

import "github.com/kamva/mgm/v3"

// User Structure
type User struct {
	// DefaultModel add _id,created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"password" bson:"password"`
	Details          string `json:"details" bson:"details"`
}

// NewUser Generates a new User Object
func NewUser(username string, password string, details string) *User {
	return &User{
		Username: username,
		Password: password,
		Details:  details,
	}
}
