package users

import (
	"GOOauth/myDB"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TruncateTable() {
	log.Println("prepare test for user crud ----- TRUNCATE USER TABLE -------")

	db := myDB.InitDb()
	u := User{}
	_, err := db.NewTruncateTable().Model(&u).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

}

func TestCreateOne(t *testing.T) {
	TruncateTable()
	log.Println("start test db insert")

	user := NewUser("nicolas", "nicolas", "nicolas@toto.com", "toto")

	one, userError := user.CreateOne()
	if one != nil {
		user.ToJson()
	}

	if userError != nil {
		userError.ToJson()
	}

	user = NewUser("nicolas", "nicolas", "nicolas@toto.com", "toto")
	_, err := user.CreateOne()
	assert.NotNil(t, err, "assert duplicate user creation")
	if err != nil {
		err.ToJson()
	}

}

//func TestGetOne(t *testing.T) {
//	TruncateTable()
//
//	user := NewUser("nicolas", "nicolas", "nicolas@toto.com", "toto")
//	user, _ = UserRepository.CreateOne(user)
//
//	user.ToJson()
//	userFromDb := UserRepository.GetOneByLogin(user, "nicolas")
//	assert.Equal(t, user, userFromDb)
//
//}
//
//func TestUser_ValidateIdentity(t *testing.T) {
//
//	TruncateTable()
//	user := NewUser("nicolas", "nicolas", "nicolas@toto.com", "toto")
//	user, _ = UserRepository.CreateOne(user)
//	identity, userError := UserService.ValidateIdentity(user, "toto")
//
//	assert.Nil(t, userError)
//	assert.True(t, identity)
//
//	identity, userError = UserService.ValidateIdentity(user, "")
//	assert.False(t, identity)
//	assert.Error(t, userError)
//}
