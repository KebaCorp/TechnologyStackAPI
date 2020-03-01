package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:         "user@example.org",
		Username:      "TestUser",
		FirstName:     "TestFirstName",
		LastName:      "TestLastName",
		MiddleName:    "TestMiddleName",
		IsActive:      true,
		Password:      "password",
		CreatorUserId: 1,
	}
}
