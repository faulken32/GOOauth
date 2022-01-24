package Utils

import "fmt"

func CheckAndDie(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckAndWarn(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
