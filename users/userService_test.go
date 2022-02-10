package users

import (
	"GOOauth/myDB"
	"context"
	"log"
	"testing"
)

func truncateTable() {
	log.Println("prepare test for user crud")

	db := myDB.InitDb()
	u := UserDb{}
	_, err := db.NewTruncateTable().Model(&u).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

}

func TestCreateOne(t *testing.T) {
	truncateTable()
	log.Println("start test db insert")

	noId := NewUserDbNoId("nicolas", "nicolas", "nicolas")

	one, userError := noId.CreateOne()
	if one != nil {
		one.ToJson()
	}

	if userError != nil {
		userError.ToJson()
	}
	noId = NewUserDbNoId("nicolas", "nicolas", "nicolas")
	_, u := noId.CreateOne()
	if u.Error != nil {
		u.ToJson()
	}

}

func TestGetOne(t *testing.T) {

	u := GetOneByLogin("nicolas")

	u.ToJson()

}
