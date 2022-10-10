package realms

import (
	"GOOauth/Utils"
	"log"
	"testing"
)

func initTest() error {
	Utils.ReadConfig(true)
	endpoint := &Endpoint{Name: "test", Url: "http://localhost:8080/api/toto", Uri: "/api/toto", Method: "GET"}
	endpoint.TruncateTable()
	endpoint, err := endpoint.Save()
	endpoint2 := &Endpoint{Name: "test2", Url: "http://localhost:8080/api/toto", Uri: "/api/toto", Method: "GET"}
	save, err := endpoint2.Save()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	log.Println(endpoint)
	log.Println(save)

	return nil
}

func TestEndpoint_FindAll(t *testing.T) {
	err := initTest()
	endpoint := Endpoint{ID: 1}

	all, err := endpoint.FindAll()
	if err != nil {
		return
	}
	log.Println(all)
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

func TestEndpoint_FindByUrl(t *testing.T) {

	err := initTest()
	if err != nil {
		return
	}

	endpoint := Endpoint{Url: "http://localhost:8080/api/toto"}

	id, err := endpoint.FindByUrl()
	if err != nil {
		return
	}

	log.Println(id)
}
