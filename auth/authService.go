package auth

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"GOOauth/users"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

type JWTMaker struct {
	secretKey string
}

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// token validation
func isValid(token string) (bool, error, jwt.MapClaims) {
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

// create a jwt token
func createToken(user string) string {
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

// extract token from request
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	return bearToken
}

// Authenticate  a user
func Authenticate(request *http.Request) (dto.Response, dto.ErrorResponse) {

	var authRequest dto.AuthRequest

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

	//token := extractToken(request)
	//	valid, err, claims := isValid(token)

	if err != nil {
		return dto.Response{}, dto.ErrorResponse{}
	}

	//if claims != nil {
	//	userName := claims["username"].(string)
	//	log.Println(userName)
	//}

	var response dto.Response

	// check user in db
	fromRequest := users.NewFromRequest(authRequest)

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
			ErrorMessage: "You don't have right on realm",
		}

	}

}
