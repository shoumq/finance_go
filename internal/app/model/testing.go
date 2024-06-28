package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user2@example.org",
		Password: "password1",
	}
}
