package sqlstore_test

import (
	"micros/internal/app/model"
	"micros/internal/app/store"
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

func TestUserRepositiry_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepositiry_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
