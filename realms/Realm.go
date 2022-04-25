package realms

import "github.com/uptrace/bun"

type Realm struct {
	bun.BaseModel `bun:",table:realms"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,unique"`
	Url           string `bun:"url"`
}

func NewRealm(ID int64, name string) *Realm {
	return &Realm{ID: ID, Name: name}
}
