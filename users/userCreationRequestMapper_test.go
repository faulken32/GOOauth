package users

import "testing"

func TestMapIt(t *testing.T) {

	request := UserCreationRequest{
		Login:    "nicolas",
		Name:     "nicolas",
		Email:    "nico@",
		Realm:    "realm",
		Password: "pass",
	}

	mapIt(request)

}
