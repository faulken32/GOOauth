package realms

import (
	"github.com/uptrace/bun"
)

type Endpoint struct {
	bun.BaseModel `bun:",table:end_points"`
	ID            int64  `bun:"id,pk,autoincrement" json:"id"`
	Name          string `bun:"name,unique" json:"name"`
	Url           string `bun:"url" json:"url"`
	Uri           string `bun:"uri" json:"uri"`
	Method        string `bun:"method" json:"method"`
}

func NewEndPoint(name string, url string) *Endpoint {

	return &Endpoint{Name: name, Url: url}
}
