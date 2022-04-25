package realms

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"database/sql"
	"github.com/uptrace/bun"
)

type RealmUserService interface {
	AddUserToRealm() (sql.Result, error)
}

type RealmUsers struct {
	bun.BaseModel `bun:",table:realms_users"`
	ID            int64  `bun:"id,pk,autoincrement"`
	UserId        string `bun:"user_id,unique:compo"`
	RealmId       string `bun:"realm_id,unique:compo"`
}

func NewRealmUsers(ID int64, userId string, realmId string) *RealmUsers {
	return &RealmUsers{ID: ID, UserId: userId, RealmId: realmId}
}

func NewRealmUsersNoId(userId string, realmId string) RealmUsers {
	return RealmUsers{UserId: userId, RealmId: realmId}
}

func (r RealmUsers) AddUserToRealm() (sql.Result, error) {

	ctx, b, errCreate := r.createTable()
	if errCreate != nil {
		Utils.CheckAndWarn(errCreate)
		return nil, errCreate
	} else {
		exec, errInsert := b.NewInsert().Model(&r).Exec(ctx)
		if errInsert != nil {
			return nil, errInsert
		}
		return exec, nil
	}
}

func (r RealmUsers) createTable() (context.Context, *bun.DB, error) {
	ctx := context.Background()
	db := myDB.InitDb()
	_, err := db.NewCreateTable().Model((*RealmUsers)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarnWfName(err, "createTable")

	return ctx, db, err
}
