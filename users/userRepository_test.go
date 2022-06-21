package users

import (
	"GOOauth/realms"
	"log"
	"testing"
)

func TestTruncate(t *testing.T) {

	TruncateTable()
}

func TestUser_GetUserRealm(t *testing.T) {
	TruncateTable()
	user := NewUser("nicolas", "nicolas", "nicolas@toto.com", "toto")

	user, err := user.CreateOne()
	realm := realms.NewRealm("test", "/test")
	realm.TruncateTable()
	realm, err2 := realm.CreateOneInDb()
	usersRealm := realms.RealmUsers{UserId: user.ID, RealmId: realm.ID}
	usersRealm.TruncateTable()
	toRealm, err3 := usersRealm.AddUserToRealm()
	userRealm, err4 := user.GetUserRealm()
	if err4 != nil {
		return
	}
	if err3 != nil {
		return
	}
	if err2 != nil {
		return
	}

	log.Println(err)
	log.Println(user)
	log.Println("test")
	log.Println(toRealm)
	log.Println(userRealm)

}
