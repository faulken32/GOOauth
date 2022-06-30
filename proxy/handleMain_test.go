package proxy

import (
	"GOOauth/auth"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_main(t *testing.T) {

	type request struct {
		name string
	}

	body := request{
		name: "toto",
	}

	out, _ := json.Marshal(body)
	buffer := bytes.NewBuffer(out)
	token, err := auth.CreateToken("nicolas", "nicolas", "toto")
	if err != nil {
		return
	}

	r, _ := http.NewRequest("GET", "http://localhost:8090/api/toto", buffer)

	w := httptest.NewRecorder()
	r.Header.Add("Authorization", token)
	Main(w, r)

	log.Println(w.Body)

}
func Test_post(t *testing.T) {

	values := map[string]string{"name": "John Doe", "occupation": "gardener"}

	out, _ := json.Marshal(values)
	buffer := bytes.NewBuffer(out)
	token, err := auth.CreateToken("nicolas", "nicolas", "toto")
	if err != nil {
		return
	}

	r, _ := http.NewRequest("POST", "http://localhost:8090/api/totoPost", buffer)

	w := httptest.NewRecorder()
	r.Header.Add("Authorization", token)
	r.Header.Add("Content-Type", "application/json")
	Main(w, r)

	log.Println(w.Body)

}
