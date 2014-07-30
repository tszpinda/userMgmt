package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tszpinda/mongoutil"
	"labix.org/v2/mgo/bson"
)

var sTestDb = "sTestDb"

func TestCreateAuthSession(t *testing.T) {
	WithTestCtx(sTestDb, func(ctx *TestCtx) {
		//given
		userId := bson.NewObjectId()
		//when
		s := CreateAuthSession(ctx.Db, userId.Hex())
		//then
		count, err := ctx.Db.C(sessionCollection).Find(bson.M{"userId": userId}).Count()
		assert.NotNil(t, s)
		assert.NotNil(t, s.AuthToken)
		assert.Nil(t, err)
		assert.NotEqual(t, count, 0, "user not added")
		assert.Equal(t, count, 1, "too many users added")
		assert.Equal(t, userId.Hex(), s.UserId.Hex())
	})
}

func TestFindSessionByAuthToken(t *testing.T) {
	WithTestCtx(sTestDb, func(ctx *TestCtx) {
		//given
		userId := bson.NewObjectId()
		s := CreateAuthSession(ctx.Db, userId.Hex())
		//when
		result, err := FindSessionByAuthToken(ctx.Db, s.AuthToken)
		//then
		assert.Nil(t, err, "unexpected error")
		assert.NotNil(t, result, "Session not found")
		assert.Equal(t, s.AuthToken, result.AuthToken)
		assert.Equal(t, userId.Hex(), result.UserId.Hex())
	})
}
