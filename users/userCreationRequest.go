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
		MapToUser() User
	}

	UserAuthRequest struct {
		Login    string
		Password string
		EndPoint string
	}
	UserAuth interface {
		Auth(endPoint string) bool
	}
)

func (ucr UserCreationRequest) MapToUser() User {

	user := User{}
	// TODO misssing validator

	if ucr.Name != "" {
		user.Name = ucr.Name
	}
	if ucr.Login != "" {
		user.Login = ucr.Login
	}
	if ucr.Email != "" {
		user.Email = ucr.Email
	}
	if ucr.Password != "" {
		user.Password = ucr.Password
	}
	return user
}
