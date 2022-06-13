package users

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"log"
)

// UserAuth interface
type (
	UserAuthRequest struct {
		Login    string
		Password string
	}
	UserAuth interface {
		MapToUser() User
	}
)

func (u User) Auth(realm string) {

	known := u.GetOneByLogin()

	if known != nil {

		validPass, err := UserService.ValidateIdentity(u, known.Password)
		if validPass && err == nil {
			on, err := u.asRightOn(realm)
			Utils.CheckAndWarn(err)
			if on {
				// generate token
			}
		}
	}
}

func (u UserAuthRequest) MapToUser() User {

	user := User{}
	if u.Login != "" {
		user.Login = u.Login
	}
	if u.Password != "" {
		user.Password = u.Password
	}

	return user
}

// AsRightOn check is user as right on a realm
func (u User) asRightOn(realm string) (bool, error) {

	userRealm, err := u.GetUserRealm()
	Utils.CheckAndWarn(err)
	log.Println(userRealm)
	//var res = ""
	//var query = "SELECT u.login from \"user\" as u " +
	//	" inner join realms_users ru on u.id = ru.user_id " +
	//	" inner join realms r on r.id = ru.realm_id" +
	//	" where u.login = ? AND r.name= ? ;"
	//db := myDB.InitDb()
	//
	//err := db.QueryRow(query, u.Login, realm).Scan(&res)

	//UserRepository(User{ })

	//if err != nil {
	//	return false, err
	//}
	//if res == u.Login {
	//	return true, nil
	//}

	return false, nil

}

//NewFromRequest create a user from an AuthRequest
func NewFromRequest(request dto.AuthRequest) UserAuthRequest {

	return UserAuthRequest{
		Login:    request.Login,
		Password: request.Password,
	}

}
