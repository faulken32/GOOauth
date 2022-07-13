package auth

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func initContext() {
	Utils.ReadConfig(true)
}

func TestCreateToken(t *testing.T) {

	initContext()
	token, _ := CreateToken("nicolas", "nicolas", "toto")
	assert.NotEmpty(t, token)
	assert.NotNil(t, token)

}

func TestAuthenticate(t *testing.T) {
	initContext()
	authRequest := dto.AuthRequest{
		Login:             "nicolas",
		Password:          "toto",
		Name:              "nicolas",
		RequestedEndPoint: "/api/toto",
	}

	out, _ := json.Marshal(authRequest)
	buffer := bytes.NewBuffer(out)
	token, _ := CreateToken("nicolas", "nicolas", "toto")
	r, _ := http.NewRequest("POST", "/auth", buffer)
	r.Header.Add("Authorization", token)
	authenticate, err := Authenticate(r)
	log.Println(authenticate)
	log.Println(err)

}
