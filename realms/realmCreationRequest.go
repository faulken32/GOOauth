package realms

type RealmCreationRequest struct {
	Name   string
	Url    string
	Method string
}

type RealmUpdateRequest struct {
	ID     int64
	Name   string
	Url    string
	Method string
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

	if r.Method != "" {
		r2.Method = r.Method
	}

	return r2
}

// MapToRealmForUpdate MapToRealm map a request to real update  struct
func (r RealmUpdateRequest) MapToRealmForUpdate() *Endpoint {
	r2 := &Endpoint{}

	if r.ID != 0 {
		r2.ID = r.ID
	}

	if r.Name != "" {
		r2.Name = r.Name
	}

	if r.Url != "" {
		r2.Url = r.Url
	}

	if r.Method != "" {
		r2.Method = r.Method
	}

	return r2
}
