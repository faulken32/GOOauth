package main

import (
	"GOOauth/Utils"
	"GOOauth/auth"
	"GOOauth/users"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/private/user/create", userHandler)
	err := http.ListenAndServe(":8080", nil)
	Utils.CheckAndDie(err)

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

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	userCreationRequest := users.UserCreationRequest{}

	j := json.NewDecoder(r.Body)
	err := j.Decode(&userCreationRequest)
	// TODO  check token from user super admin
	Utils.CheckAndWarn(err)

}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "ping")
}
