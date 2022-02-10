package users

type UserCreationRequest struct {
	Login    string
	Name     string
	Email    string
	Realm    string
	Password string
}
