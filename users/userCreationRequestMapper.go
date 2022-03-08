package users

func MapIt(request UserCreationRequest) UserDb {

	user := UserDb{}
	// TODO misssing validator

	if request.Realm != "" {
		user.Realm = request.Realm
	}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Login != "" {
		user.Login = request.Login
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	return user
}

// try to map in dynaique FAIL
//	requestValue := reflect.ValueOf(request)
//	requestType := requestValue.Type()

//log.Println(v)

//	userValue := reflect.ValueOf(user)
//	userType := userValue.Type()

//	for i := 0; i < requestValue.NumField(); i++ {
//		for j := 0; j < userValue.NumField(); j++ {
//			if requestType.Field(i).Name == userType.Field(j).Name {
//				log.Println("it's a match ")
//				log.Println(requestType.Field(i).Name)
/*				log.Println("for value {}", requestValue.Field(i).Interface())
			//value := requestValue.Field(i).Interface()

		}
	}
}*/

//log.Println(v.Field(i).Interface())
//	log.Println(t.Field(i).Name)
