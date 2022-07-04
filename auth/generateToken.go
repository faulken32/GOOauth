package auth

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"GOOauth/users"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

// CreateToken create a jwt token
func CreateToken(name string, login string, password string) (string, error) {

	user := users.NewUser(login, name, "", password)
	userFromDb := user.GetOneByLogin()

	if user.Login == userFromDb.Login && user.Password == userFromDb.Password {

		realm, err := userFromDb.GetUserRealm()

		if err != nil {
			return "", errors.New("failed to request db")
		}

		//user.ValidateIdentity(password)
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := customClaims{
			Login:      userFromDb.Login,
			Username:   userFromDb.Name,
			AuthRealms: realm,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
				Issuer:    "nameOfWebsiteHere",
			},
		}
		secret := Utils.AppConfig.Secret
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString([]byte(secret.SecretKey))
		Utils.CheckAndWarn(err)
		return signedToken, nil
	}
	return "", errors.New("failed to auth user")
}

// build the grant response with token
func encodeAuthResponse(authRequest dto.AuthRequest, response dto.Response) (dto.Response, dto.ErrorResponse) {
	token, _ := CreateToken("", authRequest.Login, authRequest.Password)
	response = dto.NewResponse(token, "")

	return response, dto.ErrorResponse{}
}
