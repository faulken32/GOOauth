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

// AsRightOn check is user as right on a realm
func (u User) AsRightOn(realm string) bool {

	log.Println(u)
	log.Println(realm)
	//oneByLogin := GetOneByLogin(u.Login)

	return true
}

//func New() *User_id {
//	return &User_id{}
//}

//NewFromRequest create a user from an AuthRequest
func NewFromRequest(request dto.AuthRequest) User {

	return User{
		Login:    request.Login,
		Password: request.Password,
		Active:   true,
		Realm:    request.Realm,
	}

}
