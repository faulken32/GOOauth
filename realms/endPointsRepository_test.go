package realms

import (
	"log"
	"testing"
)

func initTest() error {

	endpoint := &Endpoint{Name: "test", Url: "http://localhost/api/toto"}
	endpoint.TruncateTable()
	endpoint, err := endpoint.Save()
	if err != nil {
		return err
	}

	log.Println(endpoint)

	return nil
}

func TestEndpoint_FindById(t *testing.T) {

	err := initTest()
	if err != nil {
		return
	}

	endpoint := Endpoint{ID: 1}

	id, err := endpoint.FindById()
	if err != nil {
		return
	}

	log.Println(id)
}
