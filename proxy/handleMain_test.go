package proxy

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_main(t *testing.T) {

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	type request struct {
		name string
	}

	body := request{
		name: "toto",
	}

	out, _ := json.Marshal(body)
	buffer := bytes.NewBuffer(out)
	r, _ := http.NewRequest("GET", "http://localhost/api/toto", buffer)
	w := httptest.NewRecorder()
	Main(w, r)

}
