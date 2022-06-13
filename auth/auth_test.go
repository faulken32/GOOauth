package auth

import (
	"GOOauth/auth/dto"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestCreateToken(t *testing.T) {

	token := createToken("nicolas")
	assert.NotEmpty(t, token)
	assert.NotNil(t, token)
	b, err, _ := isValid(token)
	assert.True(t, b)
	if err != nil {
		log.Fatal(err)
	}
}

func TestAuthenticate(t *testing.T) {

	authRequest := dto.AuthRequest{
		Login:    "nicolas",
		Password: "test",
		Realm:    "ttt",
	}

	out, _ := json.Marshal(authRequest)
	buffer := bytes.NewBuffer(out)
	token := createToken("nicolas")
	r, _ := http.NewRequest("POST", "/auth", buffer)
	r.Header.Add("Authorization", token)
	authenticate, err := Authenticate(r)
	log.Println(authenticate)
	log.Println(err)

}
