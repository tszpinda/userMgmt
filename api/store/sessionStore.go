package store

import (
	"crypto/rand"
	"crypto/sha1"

	"fmt"
	"io"

	"code.google.com/p/go.crypto/bcrypt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const sessionCollection = "sessionCollection"

type Session struct {
	MId       bson.ObjectId `json:"_" bson:"_id"`
	Id        string        `json:"id" bson:"_"`
	AuthToken string        `json:"authToken" bson:"authToken"`
	UserId    bson.ObjectId `json:"_" bson:"userId"`
}

type SessionStore struct {
	Db *mgo.Database
}

func (this SessionStore) CreateAuthSession(userIdHex string) *Session {
	userId := bson.ObjectIdHex(userIdHex)
	s := Session{AuthToken: generateAuthToken(), MId: bson.NewObjectId(), UserId: userId}
	this.Db.C(sessionCollection).Insert(s)
	return &s
}

func (this SessionStore) FindSessionByAuthToken(db *mgo.Database, token string) (s *Session, _ error) {
	return s, this.Db.C(sessionCollection).Find(bson.M{"authToken": token}).One(&s)
}

// newUUID generates a random UUID according to RFC 4122
func generateAuthToken() string {
	uuid := make([]byte, 16)
	io.ReadFull(rand.Reader, uuid)
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	token := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	tokenByte, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	h := sha1.New()
	h.Write(tokenByte)
	h.Sum(nil)
	return string(tokenByte)
}
