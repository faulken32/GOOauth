package realms

import (
	"GOOauth/Utils"
	"testing"
)

func TestRealm_CreateOne(t *testing.T) {

	r := newRealmForUser("test", 1)

	_, s := r.CreateOneInDb()

	Utils.CheckAndWarn(s)
}
