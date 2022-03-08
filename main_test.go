package main

import (
	"GOOauth/myDB"
	"GOOauth/users"
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func truncateTable() {
	log.Println("prepare test for user crud ----- TRUNCATE USER TABLE -------")

	db := myDB.InitDb()
	u := users.UserDb{}
	_, err := db.NewTruncateTable().Model(&u).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

}

func Test_userHandler(t *testing.T) {

	truncateTable()
	body := &users.UserCreationRequest{
		Login:    "nicolas",
		Name:     "canicatti",
		Email:    "canicatti.eee@aaaaa.com",
		Realm:    "",
		Password: "",
	}

	out, _ := json.Marshal(body)
	buffer := bytes.NewBuffer(out)
	r, _ := http.NewRequest("POST", "/private/user/create", buffer)
	w := httptest.NewRecorder()

	userHandler(w, r)
	result := w.Result()
	log.Println(result.StatusCode)
	assert.Equal(t, 201, result.StatusCode)

	r, _ = http.NewRequest("GET", "/private/user/create", buffer)
	w = httptest.NewRecorder()

	userHandler(w, r)
	result = w.Result()
	log.Println(result.StatusCode)
	assert.Equal(t, 500, result.StatusCode)

}
