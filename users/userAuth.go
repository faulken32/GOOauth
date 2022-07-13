package users

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"GOOauth/myDB"
	"log"
)

func (u User) Auth(endPoint string) bool {

	known := u.GetOneByLogin()

	if known != nil {
		validPass, err := UserService.ValidateIdentity(u, known.Password)

		if validPass && err == nil {
			on, err := u.asRightOn(endPoint)
			Utils.CheckAndWarn(err)
			if on {
				return true
			}
		} else {
			Utils.CheckAndWarn(err)
		}
	}
	return false
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

	var items []map[string]interface{}
	userRealm, err := u.GetUserRealm()
	Utils.CheckAndWarn(err)
	log.Println(userRealm)
	var res = ""
	var query = "SELECT u.login from \"user\" as u " +
		" inner join realms_users ru on u.id = ru.user_id " +
		" inner join realms r on r.id = ru.realm_id" +
		" where u.login = ? AND r.name= ? ;"
	db := myDB.InitDb()

	err2 := db.QueryRow(query, u.Login, realm).Scan(&res, &items)

	//UserRepository(User{ })

	if err2 != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}
	if res == u.Login {
		return true, nil
	}

	return false, nil

}

//NewFromRequest create a user from an AuthRequest
func NewFromRequest(request dto.AuthRequest) UserAuthRequest {

	return UserAuthRequest{
		Login:    request.Login,
		Password: request.Password,
		EndPoint: request.RequestedEndPoint,
	}

}
