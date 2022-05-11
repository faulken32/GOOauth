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
	ID            int64 `bun:"id,pk,autoincrement"`
	UserId        int64 `bun:"user_id,unique:compo"`
	RealmId       int64 `bun:"realm_id,unique:compo"`
}

func NewRealmUsers(ID int64, userId int64, realmId int64) *RealmUsers {
	return &RealmUsers{ID: ID, UserId: userId, RealmId: realmId}
}

func NewRealmUsersNoId(userId int64, realmId int64) RealmUsers {
	return RealmUsers{UserId: userId, RealmId: realmId}
}

// AddUserToRealm add a user to realm  does not check user or real exit
func (r RealmUsers) AddUserToRealm() (sql.Result, error) {

	ctx, db, errCreate := r.createTable()
	if errCreate != nil {
		Utils.CheckAndWarn(errCreate)

		_ = db.Close()
		return nil, errCreate
	} else {
		exec, errInsert := db.NewInsert().Model(&r).Exec(ctx)
		if errInsert != nil {
			_ = db.Close()
			return nil, errInsert
		}
		_ = db.Close()
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
