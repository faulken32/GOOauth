package realms

import "github.com/uptrace/bun"

type Realm struct {
	bun.BaseModel `bun:",table:realms"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,unique"`
	Url           string `bun:"url"`
}

func NewRealm(name string, url string) *Realm {

	return &Realm{Name: name, Url: url}
}
