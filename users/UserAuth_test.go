package users

import (
	"GOOauth/auth/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAsRightOn(t *testing.T) {
	user := User{}
	on := user.AsRightOn("")

	if on != true {
		t.Error("eeeee")
	}
}

func TestNewFromRequest(t *testing.T) {
	r := dto.Request{
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
