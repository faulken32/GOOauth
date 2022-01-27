package users

import (
	"GOOauth/auth/dto"
	"log"
)

// UserAuth interface
type UserAuth interface {
	AsRightOn(realm string) bool
}

type User struct {
	Login    string
	Password string
	Active   bool
	Realm    string
}

func (u User) AsRightOn(realm string) bool {

	log.Println(u)
	log.Println(realm)

	return true
}

func New() *User {
	return &User{}
}

func NewFromRequest(request dto.Request) User {

	return User{
		Login:    request.Login,
		Password: request.Password,
		Active:   true,
		Realm:    request.Realm,
	}

}
