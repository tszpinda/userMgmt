package store

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/tszpinda/mongoutil"
	"labix.org/v2/mgo/bson"
)

var testDb = "userTestDb"

func TestAddUser(t *testing.T) {
	WithTestCtx(testDb, func(ctx *TestCtx) {
		//given
		AddUser(ctx.Db, "john@gmx.com", "passwd", "John")
		//when
		count, err := ctx.Db.C(userCollection).Find(bson.M{"email": "john@gmx.com"}).Count()
		//then
		assert.Nil(t, err)
		assert.NotEqual(t, count, 0, "user not added")
		assert.Equal(t, count, 1, "too many users added")
	})
}

func TestFindUserByEmail(t *testing.T) {
	WithTestCtx(testDb, func(ctx *TestCtx) {
		//given
		AddUser(ctx.Db, "john@gmx.com", "passwd", "John")
		//when
		u, err := FindUserByEmail(ctx.Db, "john@gmx.com")
		//then
		log.Printf("%+v", err)
		assert.Nil(t, err, "unexpected error")
		assert.NotNil(t, u, "user not found")
		assert.Equal(t, u.Email, "john@gmx.com")
	})
}

func TestUpdateUser(t *testing.T) {
	WithTestCtx(testDb, func(ctx *TestCtx) {
		//given
		u := AddUser(ctx.Db, "john@gmx.com", "passwd", "John")
		//when
		err := UpdateUser(ctx.Db, u.Id.Hex(), "john@gmail.com", "John Travolta")
		assert.Nil(t, err, "unexpected error")

		//then
		u, err = FindUserByEmail(ctx.Db, "john@gmail.com")
		assert.Nil(t, err, "unexpected error")
		assert.NotNil(t, u, "user not found")
		assert.Equal(t, u.Email, "john@gmail.com")
		assert.Equal(t, u.Name, "John Travolta")
	})
}

func TestUpdatePassword(t *testing.T) {
	WithTestCtx(testDb, func(ctx *TestCtx) {
		//given
		u := AddUser(ctx.Db, "john@gmx.com", "passwd", "John")
		//when
		err := UpdatePassword(ctx.Db, u.Id.Hex(), "passwd1234")
		assert.Nil(t, err, "unexpected error")

		//then
		u, err = FindUserByEmail(ctx.Db, "john@gmx.com")
		assert.Nil(t, err, "unexpected error")
		assert.NotNil(t, u, "user not found")
		assert.Equal(t, u.Password, hashPasswd("passwd1234"))
		assert.Equal(t, u.Email, "john@gmx.com")
		assert.Equal(t, u.Name, "John")
	})
}
