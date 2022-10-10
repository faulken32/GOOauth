/*
 * Copyright (c) 2022.
 *
 * canicatti.nicolas@gmail.com
 *
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"GOOauth/Utils"
	"GOOauth/auth"
	"GOOauth/proxy"
	"GOOauth/realms"
	"GOOauth/users"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	log.Println("starting app")
	err := Utils.ReadConfig(false)
	if err != nil {
		log.Panic(err.Error())
	}

	database := Utils.AppConfig.Database
	log.Println("database : ", database.DbName, " : ", database.Host)

	rtr := mux.NewRouter()

	rtr.HandleFunc("/api/{.*}", proxy.Main)

	rtr.HandleFunc("/auth", authHandler)
	rtr.HandleFunc("/private/user/create", userHandler)
	rtr.HandleFunc("/private/realm/create", realmAddHandler)
	rtr.HandleFunc("/private/realm/update", realmUpdateHandler)
	rtr.HandleFunc("/private/realm/all", endpointsGetAll)
	rtr.HandleFunc("/private/realm/add/user", realms.RealmAddUserHandler)

	http.Handle("/", rtr)
	log.Println("Listening...")

	err = http.ListenAndServe(":8090", nil)

	Utils.CheckAndDie(err)

}

func endpointsGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	encoder := json.NewEncoder(w)
	endpoint := realms.Endpoint{}
	all, err := realms.EndPointsRepository(endpoint).FindAll()
	err2 := encoder.Encode(all)
	if err == nil || err2 == nil {
		log.Println(err2, err)
		return
	}

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reals := &realms.RealmCreationRequest{}
	j := json.NewDecoder(r.Body)
	err := j.Decode(reals)
	Utils.CheckAndWarn(err)
	realm, err := reals.MapToRealm().Save()

	encoder := json.NewEncoder(w)
	Utils.ReturnErrorOrHTTPResponse(w, err, encoder, realm)

}

func realmUpdateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	reals := &realms.RealmUpdateRequest{}
	j := json.NewDecoder(r.Body)
	err := j.Decode(reals)
	Utils.CheckAndWarn(err)
	realm, err := reals.MapToRealmForUpdate().Update()

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
