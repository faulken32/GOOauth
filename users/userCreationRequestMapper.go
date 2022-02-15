package users

import (
	"log"
	"reflect"
)

func mapIt(request UserCreationRequest) {

	user := UserDb{}
	requestValue := reflect.ValueOf(request)
	requestType := requestValue.Type()

	//log.Println(v)

	userValue := reflect.ValueOf(user)
	userType := userValue.Type()

	for i := 0; i < requestValue.NumField(); i++ {
		for j := 0; j < userValue.NumField(); j++ {
			if requestType.Field(i).Name == userType.Field(j).Name {
				log.Println("it's a match ")
				log.Println(requestType.Field(i).Name)
				log.Println("for value {}", requestValue.Field(i).Interface())
				userValue.Field(j).Interface()
			}
		}
	}
	//log.Println(v.Field(i).Interface())
	//	log.Println(t.Field(i).Name)

	//if request.Realm != "" {
	//	//user.
	//}
	//if request.Name != "" {
	//	user.Name = request.Name
	//}

}
