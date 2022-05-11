package realms

import (
	"GOOauth/Utils"
	"encoding/json"
	"errors"
	"net/http"
)

type addUserRequest struct {
	UserId  int64 `json:"userId"`
	RealmId int64 `json:"realmId"`
}

func newAddUserRequest(userId int64, realmId int64) *addUserRequest {
	return &addUserRequest{UserId: userId, RealmId: realmId}
}

func newAddUserRequestEmpty() *addUserRequest {
	return &addUserRequest{}
}

func (r addUserRequest) mapToRealmUser() (RealmUsers, error) {
	if r.UserId != 0 && r.RealmId != 0 {
		return NewRealmUsersNoId(r.UserId, r.RealmId), nil
	}
	return RealmUsers{}, errors.New("invalid request")
}

// RealmAddUserHandler http handler
func RealmAddUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	request := newAddUserRequestEmpty()
	j := json.NewDecoder(r.Body)
	err := j.Decode(request)

	realm, err := request.mapToRealmUser()
	if err != nil {
		encoder := json.NewEncoder(w)
		Utils.ReturnErrorOrHTTPResponse(w, err, encoder, nil)
	} else {
		toRealm, err := realm.AddUserToRealm()

		Utils.CheckAndWarn(err)
		encoder := json.NewEncoder(w)
		Utils.ReturnErrorOrHTTPResponse(w, err, encoder, toRealm)
	}

}
