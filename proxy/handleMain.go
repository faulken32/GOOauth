package proxy

import (
	"log"
	"net/http"
)

type ProxRequest struct {
	uri string
}

func Main(w http.ResponseWriter, r *http.Request) {

	uri := r.URL.RequestURI()
	path := r.URL.Path
	query := r.URL.Query()
	scheme := r.URL.Scheme
	host := r.URL.Host
	get, err := http.Get(scheme + "://" + host + uri)
	if err != nil {
		log.Println(err.Error())

	}

	log.Println(scheme)
	log.Println(query)
	log.Println(path)
	log.Println(get)
	log.Println(uri)
}
