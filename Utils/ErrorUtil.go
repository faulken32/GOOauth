package Utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func CheckAndDie(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CheckAndWarn(e error) {
	if e != nil {
		log.Println(e)
	}
}

// CheckAndWarnWfName check error and print it with function name
func CheckAndWarnWfName(e error, funcName string) {
	if e != nil {
		log.Println("from function", funcName)
		log.Println(e)
	}
}

// ReturnErrorOrHTTPResponse check error and return it or return http response  entity
func ReturnErrorOrHTTPResponse(w http.ResponseWriter, err error, encoder *json.Encoder, entity interface{}) {
	if err != nil {
		w.WriteHeader(500)
		_ = encoder.Encode(err.Error())
	} else {
		err = encoder.Encode(entity)
	}
}
