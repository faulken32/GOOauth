package users

import (
	"GOOauth/auth/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestAsRightOn(t *testing.T) {
//	user := UserAuthRequest{
//		Login: "nicolas",
//	}
//	on, _ := user.AsRightOn("toto")
//
//	assert.True(t, on)
//}

func TestFull(t *testing.T) {

	user := NewUser("nicolas", "nicolas", "nicolas@toto.com", "toto")

	_, _ = user.CreateOne()

	r := UserAuthRequest{
		Login:    "nicolas",
		Password: "toto",
	}

	r.MapToUser().asRightOn("toto")
}

func TestNewFromRequest(t *testing.T) {
	r := dto.AuthRequest{
		Login:    "",
		Password: "",
	}

	users := NewFromRequest(r)
	assert.Equalf(t, users, UserAuthRequest{
		Login:    "",
		Password: "",
	}, "")
}
