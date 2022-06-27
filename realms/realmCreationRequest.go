package realms

type RealmCreationRequest struct {
	Name string
	Url  string
}

// NewRealmCreationRequest Create a realm
func NewRealmCreationRequest(name string, url string) *RealmCreationRequest {
	return &RealmCreationRequest{Name: name, Url: url}
}

//MapToRealm map a request to real struct
func (r RealmCreationRequest) MapToRealm() *Endpoint {
	r2 := &Endpoint{}
	if r.Name != "" {
		r2.Name = r.Name
	}

	if r.Url != "" {
		r2.Url = r.Url
	}

	return r2
}
