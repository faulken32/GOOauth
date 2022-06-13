package auth

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"github.com/golang-jwt/jwt"
	"time"
)

// create a jwt token
func CreateToken(user string) string {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := customClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "nameOfWebsiteHere",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secureSecretText"))
	Utils.CheckAndWarn(err)
	return signedToken
}

// build the grant response with token
func encodeAuthResponse(authRequest dto.AuthRequest, response dto.Response) (dto.Response, dto.ErrorResponse) {
	token := CreateToken(authRequest.Login)
	response = dto.NewResponse(token, "")

	return response, dto.ErrorResponse{}
}
