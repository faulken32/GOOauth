package main

import (
	"GOOauth/Auth"
	"GOOauth/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/auth", authHandler)
	err := http.ListenAndServe(":8080", nil)
	Utils.CheckAndDie(err)

}

func authHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	authenticate := Auth.Authenticate(r)
	json.NewEncoder(w).Encode(authenticate)
	//Utils.CheckAndWarn(err)

}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "ping")
}
