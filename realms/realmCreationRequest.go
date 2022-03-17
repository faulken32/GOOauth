package realms

type realmCreationRequest struct {
	Name   string
	UserId int64
}

func newRealmCreationRequest(name string, userId int64) *realmCreationRequest {
	return &realmCreationRequest{Name: name, UserId: userId}
}

func (r realmCreationRequest) MapToRealm() *realm {

	r2 := &realm{}

	if r.Name != "" {
		r2.Name = r.Name
	}
	if r.UserId != 0 {
		r2.UserId = r.UserId
	}

	return r2

}
