package main

import (
	"GOOauth/Utils"
	"GOOauth/auth"
	"GOOauth/proxy"
	"GOOauth/realms"
	"GOOauth/users"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	log.Println("starting app")

	rtr := mux.NewRouter()

	rtr.HandleFunc("/api/{.*}", proxy.Main)

	rtr.HandleFunc("/auth", authHandler)
	rtr.HandleFunc("/private/user/create", userHandler)
	rtr.HandleFunc("/private/realm/create", realmAddHandler)
	rtr.HandleFunc("/private/realm/add/user", realms.RealmAddUserHandler)

	http.Handle("/", rtr)
	log.Println("Listening...")
	err := http.ListenAndServe(":8090", nil)
	Utils.CheckAndDie(err)
	log.Println("app started")
}

func authHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	encoder := json.NewEncoder(w)
	success, errorResponse := auth.Authenticate(r)
	if errorResponse.HttpStatus != 0 {

		err := encoder.Encode(errorResponse)
		if errorResponse.HttpStatus == 401 {
			w.WriteHeader(http.StatusForbidden)
		}
		if errorResponse.HttpStatus == 400 {
			w.WriteHeader(http.StatusBadRequest)
		}

		Utils.CheckAndWarn(err)

	} else {
		//w.WriteHeader(http.StatusOK)
		err := encoder.Encode(success)
		Utils.CheckAndWarn(err)
	}

}

// add a realm into Db
func realmAddHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	reals := &realms.RealmCreationRequest{}
	j := json.NewDecoder(r.Body)
	err := j.Decode(reals)
	Utils.CheckAndWarn(err)
	realm, err := reals.MapToRealm().Save()

	encoder := json.NewEncoder(w)
	Utils.ReturnErrorOrHTTPResponse(w, err, encoder, realm)

}

func userHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	if r.Method != "POST" {
		err := errors.New("HTTP method not supported")
		if err != nil {
			Utils.CheckAndWarn(err)
			w.WriteHeader(500)
			_ = encoder.Encode(err.Error())

		}
	} else {
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		userCreationRequest := users.UserCreationRequest{}

		j := json.NewDecoder(r.Body)
		err := j.Decode(&userCreationRequest)
		// TODO  check token from user super admin
		Utils.CheckAndWarn(err)
		user := userCreationRequest.MapToUser()
		user.ToJson()
		one, userError := user.CreateOne()
		encoder := json.NewEncoder(w)
		if userError != nil {
			log.Println(userError)
			w.WriteHeader(http.StatusInternalServerError)
		_:
			encoder.Encode(userError.Error())

		} else {
			w.WriteHeader(http.StatusCreated)
		_:
			encoder.Encode(one)
		}
	}

}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "ping")
}
