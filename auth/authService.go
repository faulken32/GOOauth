package auth

import (
	"GOOauth/auth/dto"
	"GOOauth/users"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

type JWTMaker struct {
	secretKey string
}

type authRealms struct {
	realms []users.QueryRes
}

type customClaims struct {
	Username   string           `json:"username"`
	Password   string           `json:"password"`
	Login      string           `json:"login"`
	AuthRealms []users.QueryRes `json:"authRealms"`
	jwt.StandardClaims
}

// token validation
func decodeAndValidateToken(token string) (bool, error, jwt.MapClaims) {
	maker := JWTMaker{secretKey: "secureSecretText"}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, nil
		}
		return []byte(maker.secretKey), nil
	}
	parse, err := jwt.Parse(token, keyFunc)
	claims := parse.Claims.(jwt.MapClaims)
	log.Println(claims)
	if err != nil {
		return false, err, nil
	}
	if parse.Valid {
		return true, nil, claims
	}
	return false, err, nil
}

// extract token from request
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	return bearToken
}

// Authenticate  a user
func Authenticate(request *http.Request) (dto.Response, dto.ErrorResponse) {

	authRequest, err := decodeRequest(request)
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

	// check user in db
	/*fromRequest := users.NewFromRequest(authRequest)
	on, err := fromRequest.AsRightOn(authRequest.Realm)

	if err != nil {
		return dto.Response{}, dto.ErrorResponse{
			HttpStatus:   500,
			ErrorMessage: err.Error(),
		}
	}

	var response dto.Response
	if on {
		return encodeAuthResponse(authRequest, response)
	} else {
		return dto.Response{}, dto.ErrorResponse{
			HttpStatus:   403,
			ErrorMessage: "You don't have right on realm",
		}
	}
	*/
	return dto.Response{}, dto.ErrorResponse{
		HttpStatus:   403,
		ErrorMessage: "You don't have right on realm",
	}
}

func decodeRequest(request *http.Request) (dto.AuthRequest, error) {
	var authRequest dto.AuthRequest

	j := json.NewDecoder(request.Body)
	err := j.Decode(&authRequest)
	return authRequest, err
}
