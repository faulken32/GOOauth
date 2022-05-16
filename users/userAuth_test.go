package users

import (
	"GOOauth/auth/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsRightOn(t *testing.T) {
	user := User{
		Login: "nicolas",
	}
	on, _ := user.AsRightOn("toto")

	assert.True(t, on)
}

func TestNewFromRequest(t *testing.T) {
	r := dto.AuthRequest{
		Login:    "",
		Password: "",
		Realm:    "",
	}

	users := NewFromRequest(r)
	assert.Equalf(t, users, User{
		Login:    "",
		Password: "",
		Realm:    "",
		Active:   true,
	}, "")
}
