package realms

import "github.com/uptrace/bun"

type Endpoint struct {
	bun.BaseModel `bun:",table:end_points"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,unique"`
	Url           string `bun:"url"`
}

func NewEndPoint(name string, url string) *Endpoint {

	return &Endpoint{Name: name, Url: url}
}
