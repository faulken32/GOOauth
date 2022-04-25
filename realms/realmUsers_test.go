package realms

import (
	"GOOauth/myDB"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func truncateTableRealmUser() {
	log.Println("prepare test for user crud ----- TRUNCATE RealmUsers TABLE -------")

	db := myDB.InitDb()
	u := RealmUsers{}
	_, err := db.NewTruncateTable().Model(&u).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

}

func TestAddUserToRealms(t *testing.T) {
	truncateTableRealmUser()
	users := NewRealmUsersNoId("1", "1")
	realm, err := users.AddUserToRealm()
	assert.NoError(t, err)
	log.Println("res for test :  ", realm)
	if err != nil {
		return
	}
}

func TestAddUserToRealms_Error(t *testing.T) {
	TestAddUserToRealms(t)
	users := NewRealmUsersNoId("1", "1")
	_, err := users.AddUserToRealm()

	assert.Error(t, err)

	if err != nil {
		return
	}
}
