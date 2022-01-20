package main

import (
	"GOOauth/Auth"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	var a = Auth.Authenticate()
	fmt.Println(a)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
