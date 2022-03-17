package realms

import "github.com/uptrace/bun"

type realm struct {
	bun.BaseModel `bun:",table:realms"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,unique"`
	UserId        int64  `bun:"user_id"`
}

func newRealmForUser(name string, userId int64) *realm {
	return &realm{Name: name, UserId: userId}
}
