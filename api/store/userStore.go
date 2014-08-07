package store

import (
	"crypto/sha512"
	"encoding/base64"

	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const userCollection = "userCollection"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Name     string        `json:"name"`
}

type UserStore struct {
	Db *mgo.Database
}

func (this UserStore) AddUser(email, password, name string) *User {
	u := User{Email: email, Password: hashPasswd(password), Name: name, Id: bson.NewObjectId()}
	this.Db.C(userCollection).Insert(u)
	return &u
}

func (this UserStore) FindUserByEmail(email string) (u *User, _ error) {
	return u, this.Db.C(userCollection).Find(bson.M{"email": email}).One(&u)
}

func (this UserStore) FindUserById(id string) (u *User, _ error) {
	return u, this.Db.C(userCollection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&u)
}

func (this UserStore) UpdateUser(id, email, name string) error {
	upd := bson.M{"email": email, "name": name}
	return this.updateById(id, upd)
}

func (this UserStore) UpdatePassword(id, password string) error {
	upd := bson.M{"password": hashPasswd(password)}
	return this.updateById(id, upd)
}

func (this UserStore) updateById(hexId string, update bson.M) error {
	mId := bson.ObjectIdHex(hexId)
	selector := bson.M{"_id": mId}
	return this.Db.C(userCollection).Update(selector, bson.M{"$set": update})
}

func PasswdEqual(passwdRaw, passwdHashed string) bool {
	same := hashPasswd(passwdRaw) == passwdHashed
	if !same {
		log.Println("Passwd not same")
	}
	return same
}

func hashPasswd(plain string) string {
	hasher := sha512.New()
	hasher.Write([]byte(plain))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
