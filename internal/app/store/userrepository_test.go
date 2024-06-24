package store_test

import (
	"micros/internal/app/model"
	"micros/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositiry_Create(t *testing.T) {
	s, teardown := store.TestStore(t, DatabaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositiry_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, DatabaseURL)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
