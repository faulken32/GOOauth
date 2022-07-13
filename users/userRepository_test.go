package users

import (
	"GOOauth/Utils"
	"GOOauth/realms"
	"log"
	"testing"
)

func TestUser_GetUserRealm(t *testing.T) {
	Utils.ReadConfig(true)
	TruncateTable()
	user := NewUser("nicolas", "nicolas", "nicolas@toto.com", "toto")

	user, err := user.CreateOne()
	realm := realms.NewEndPoint("test", "/test")
	realm.TruncateTable()
	realm, err2 := realm.Save()
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
