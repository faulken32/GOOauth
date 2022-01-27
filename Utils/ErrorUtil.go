package Utils

import (
	"log"
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
