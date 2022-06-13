package users

type (
	UserCreationRequest struct {
		Login    string
		Name     string
		Email    string
		Realm    string
		Password string
	}

	Mapper interface {
		MapToUser(request UserCreationRequest) User
	}
)

func (ucr UserCreationRequest) MapToUser(request UserCreationRequest) User {

	user := User{}
	// TODO misssing validator

	if request.Name != "" {
		user.Name = ucr.Name
	}
	if request.Login != "" {
		user.Login = ucr.Login
	}
	if request.Email != "" {
		user.Email = ucr.Email
	}
	if request.Password != "" {
		user.Password = ucr.Password
	}
	return user
}
