package users

import (
	"GOOauth/auth/dto"
	"GOOauth/myDB"
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
func (u User) AsRightOn(realm string) (bool, error) {

	var res = ""
	db := myDB.InitDb()
	err := db.QueryRow("SELECT u.login from \"user\" as u "+
		" inner join realms_users ru on u.id = ru.user_id "+
		" inner join realms r on r.id = ru.realm_id"+
		" where u.login = ? AND r.name= ? ;", u.Login, realm).Scan(&res)
	if err != nil {
		return false, err
	}
	if res == u.Login {
		return true, nil
	}
	return false, nil

}

//NewFromRequest create a user from an AuthRequest
func NewFromRequest(request dto.AuthRequest) User {

	return User{
		Login:    request.Login,
		Password: request.Password,
		Active:   true,
		Realm:    request.Realm,
	}

}
