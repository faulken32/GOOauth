package realms

import (
	"GOOauth/myDB"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func truncateTable() {
	log.Println("prepare test for real crud ----- TRUNCATE real TABLE -------")

	db := myDB.InitDb()
	u := Realm{}
	_, err := db.NewTruncateTable().Model(&u).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

}

func TestRealmCreationRequest_MapToRealm(t *testing.T) {
	truncateTable()
	toRealm := NewRealmCreationRequest("nicolas", "/api").MapToRealm()
	log.Println(toRealm)

	toRealm, err := toRealm.CreateOneInDb()
	if err != nil {
		return
	}
	log.Println(toRealm)
	assert.Equal(t, toRealm.Name, "nicolas")
	assert.Equal(t, toRealm.ID, int64(1))

}
