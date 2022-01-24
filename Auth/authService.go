package Auth

import (
	"GOOauth/Auth/dto"
	"GOOauth/Utils"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

type Request struct {
	Login    string
	Password string
}

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

// Authenticate auth an user
func Authenticate(request *http.Request) dto.AuthResponse {

	var ar Request

	j := json.NewDecoder(request.Body)
	err := j.Decode(&ar)
	Utils.CheckAndWarn(err)
	token := createToken(ar.Login)
	Utils.CheckAndWarn(err)
	response := dto.New(token, "")
	log.Println(response)
	return response

}
