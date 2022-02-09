package users

import (
	"log"
	"testing"
)

func TestCreateOne(t *testing.T) {

	log.Println("start test db insert")

	//	userFromDb := NewUserDb("nicolas", "nicolas", "toto@titi.com").CreateOne()
	//userFromDb2 := NewUserDb(2, "nicolas", "nicolas", "toto@titi.com").CreateOne()
	noId := NewUserDbNoId("nicolas", "nicolas", "nicolas")
	noId.CreateOne()
	log.Println(noId)
	//log.Println(userFromDb2)
	//log.Println(noId)
}
