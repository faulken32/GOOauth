package users

import "testing"

func TestMapIt(t *testing.T) {

	request := UserCreationRequest{
		Login:    "login",
		Name:     "Name",
		Email:    "Email",
		Realm:    "realm",
		Password: "password",
	}

	mapIt(request)

}
