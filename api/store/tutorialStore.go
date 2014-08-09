package store

import (
	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const tutorialCollection = "tutorialCollection"

type Step struct {
	Text     string `json:"text"`
	Selector string `json:"selector"`
}

type Tutorial struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name"`
	ApiKey string        `json:"apiKey"`
	Domain string        `json:"domain"`
	Page   string        `json:"page"`
	Steps  []Step        `json:"steps"`
}

type TutorialStore struct {
	Db *mgo.Database
}

func (this TutorialStore) AddTutorial(apiKey, domain, page, name string) *Tutorial {
	u := Tutorial{Name: name, ApiKey: apiKey, Domain: domain, Page: page, Id: bson.NewObjectId()}
	this.Db.C(tutorialCollection).Insert(u)
	return &u
}

func (this TutorialStore) UpdateTutorial(id, domain, name, page string) error {
	upd := bson.M{"name": name, "page": page, "domain": domain}
	return this.updateById(id, upd)
}

func (this TutorialStore) FindTutorialsForPage(apiKey, domain, page string) []Tutorial {
	r := make([]Tutorial, 0)
	this.Db.C(tutorialCollection).Find(bson.M{"apikey": apiKey, "domain": domain, "page": page}).All(&r)
	return r
}

func (this TutorialStore) FindTutorials(apiKey, domain string) []Tutorial {
	r := make([]Tutorial, 0)
	this.Db.C(tutorialCollection).Find(bson.M{"apikey": apiKey, "domain": domain}).All(&r)
	return r
}

func (this TutorialStore) FindTutorialsForApiKey(apiKey string) []Tutorial {
	r := make([]Tutorial, 0)
	log.Println("FindTutorialsForApiKey", apiKey)
	this.Db.C(tutorialCollection).Find(bson.M{"apikey": apiKey}).All(&r)
	return r
}

func (this TutorialStore) updateById(hexId string, update bson.M) error {
	mId := bson.ObjectIdHex(hexId)
	selector := bson.M{"_id": mId}
	return this.Db.C(tutorialCollection).Update(selector, bson.M{"$set": update})
}
