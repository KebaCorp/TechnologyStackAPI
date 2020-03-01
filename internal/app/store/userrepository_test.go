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

	u, err := s.User().CreateUser(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u := model.TestUser(t)
	expectedUser, _ := s.User().CreateUser(u)
	actualUser, err := s.User().FindByEmail(u.Email)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Email, actualUser.Email)
}
