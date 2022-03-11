package auth

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateToken(t *testing.T) {

	token := createToken("nicolas")
	log.Print(token)
	assert.NotEmpty(t, token)
	assert.NotNil(t, token)

	b, err := valid(token)

	log.Println(b)
	if err != nil {
		log.Fatal(err)
	}
}
