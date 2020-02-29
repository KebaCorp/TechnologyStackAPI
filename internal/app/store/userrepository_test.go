package store_test

import (
	"testing"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
	"github.com/KebaCorp/TechnologyStackAPI/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email:             "user@example.org",
		Username:          "TestUser",
		FirstName:         "TestFirstName",
		LastName:          "TestLastName",
		MiddleName:        "TestMiddleName",
		IsActive:          true,
		EncryptedPassword: "TestEncryptedPassword",
		CreatorUserId:     1,
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	expectedUser, _ := s.User().Create(&model.User{
		Email:             "user@example.org",
		Username:          "TestUser",
		FirstName:         "TestFirstName",
		LastName:          "TestLastName",
		MiddleName:        "TestMiddleName",
		IsActive:          true,
		EncryptedPassword: "TestEncryptedPassword",
		CreatorUserId:     1,
	})

	actualUser, err := s.User().FindByEmail("user@example.org")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, actualUser)
}
