package sqlstore_test

import (
	"micros/internal/app/model"
	"micros/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositiry_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepositiry_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.NoError(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}