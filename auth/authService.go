package auth

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"GOOauth/users"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// create a jwt token
func createToken(user string) string {

	claims := customClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "nameOfWebsiteHere",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secureSecretText"))
	Utils.CheckAndWarn(err)
	return signedToken
}

// Authenticate  a user
func Authenticate(request *http.Request) (dto.Response, dto.ErrorResponse) {

	var authRequest dto.Request
	j := json.NewDecoder(request.Body)
	err := j.Decode(&authRequest)
	if err != nil {
		return dto.Response{}, dto.ErrorResponse{}
	}

	validatorError := authRequest.Validator(authRequest)
	if validatorError != nil {
		return dto.Response{}, dto.ErrorResponse{
			HttpStatus:   http.StatusBadRequest,
			ErrorMessage: validatorError.InternalErrorMessage,
		}
	}

	fromRequest := users.NewFromRequest(authRequest)

	var response dto.Response

	if fromRequest.AsRightOn(authRequest.Realm) {

		j := json.NewDecoder(request.Body)
		err := j.Decode(&authRequest)
		Utils.CheckAndWarn(err)
		token := createToken(authRequest.Login)
		response = dto.NewResponse(token, "")
		return response, dto.ErrorResponse{}
	} else {

		return dto.Response{}, dto.ErrorResponse{
			HttpStatus:   403,
			ErrorMessage: "You don't have right on reaml",
		}

	}

}
