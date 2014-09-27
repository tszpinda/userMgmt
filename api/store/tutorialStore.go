package store

import (
	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const tutorialCollection = "tutorialCollection"

type Step struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Text       string        `json:"text"`
	Selector   string        `json:"selector"`
	No         int           `json:"no"`
	TutorialId string        `json:"tutorial" bson:"-"`
}

type Tutorial struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name"`
	ApiKey  string        `json:"apiKey"`
	Domain  string        `json:"domain"`
	Page    string        `json:"page"`
	Steps   []Step        `json:"-" bson:"steps"`
	StepIds *[]string     `json:"steps" bson:"-"`
}

type TutorialApi struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	ApiKey  string      `json:"apiKey"`
	Domain  string      `json:"domain"`
	Page    string      `json:"page"`
	Options []OptionApi `json:"options"`
}
type OptionApi struct {
	Selector string `json:"element"`
	Text     string `json:"intro"`
}

type TutorialStore struct {
	Db *mgo.Database
}

func populateTutorailsStepListId(ts []Tutorial) []Tutorial {
	for i, t := range ts {
		ts[i] = *populateStepListId(&t)

	}
	return ts
}
func populateStepListId(t *Tutorial) *Tutorial {
	ids := make([]string, 0)
	for _, s := range t.Steps {
		ids = append(ids, s.Id.Hex())
	}
	t.StepIds = &ids
	log.Println("step ids", ids)
	return t
}

func (this TutorialStore) AddTutorial(apiKey, domain, page, name string) *Tutorial {
	u := Tutorial{Name: name, ApiKey: apiKey, Domain: domain, Page: page, Id: bson.NewObjectId()}
	this.Db.C(tutorialCollection).Insert(u)
	return populateStepListId(&u)
}

func (this TutorialStore) UpdateTutorial(id, domain, page, name string) error {
	upd := bson.M{"name": name, "page": page, "domain": domain}
	return this.updateById(id, upd)
}

func (this TutorialStore) FindApiTutorials(apiKey, domain, page string) []TutorialApi {
	r := make([]Tutorial, 0)
	results := make([]TutorialApi, 0)
	this.Db.C(tutorialCollection).Find(bson.M{"apikey": apiKey, "domain": domain, "page": page}).All(&r)
	for _, t := range r {
		opts := make([]OptionApi, 0)
		for _, step := range t.Steps {
			//selector := "document.querySelector('" + step.Selector + "')"
			selector := step.Selector
			opt := OptionApi{Selector: selector, Text: step.Text}
			opts = append(opts, opt)
		}
		//	element: document.querySelector('#email'),
		//	intro: "This is a tooltip."
		tapi := TutorialApi{Id: t.Id.Hex(), Name: t.Name, ApiKey: t.ApiKey, Domain: t.Domain, Page: t.Page, Options: opts}
		results = append(results, tapi)
	}
	return results
}

func (this TutorialStore) FindTutorials(apiKey, domain string) []Tutorial {
	r := make([]Tutorial, 0)
	this.Db.C(tutorialCollection).Find(bson.M{"apikey": apiKey, "domain": domain}).All(&r)
	return populateTutorailsStepListId(r)
}

func (this TutorialStore) FindTutorialsForApiKey(apiKey string) []Tutorial {
	r := make([]Tutorial, 0)
	log.Println("FindTutorialsForApiKey", apiKey)
	this.Db.C(tutorialCollection).Find(bson.M{"apikey": apiKey}).All(&r)
	return populateTutorailsStepListId(r)
}

func (this TutorialStore) updateById(hexId string, update bson.M) error {
	mId := bson.ObjectIdHex(hexId)
	selector := bson.M{"_id": mId}
	return this.Db.C(tutorialCollection).Update(selector, bson.M{"$set": update})
}

func (this TutorialStore) FindTutorialById(hexId string) *Tutorial {
	mId := bson.ObjectIdHex(hexId)
	t := Tutorial{}
	this.Db.C(tutorialCollection).FindId(mId).One(&t)

	return populateStepListId(&t)
}

func (this TutorialStore) AddStep(tutorialId, selector, text string, no int) *Step {
	t := this.FindTutorialById(tutorialId)
	newStep := Step{Id: bson.NewObjectId(), Selector: selector, Text: text, TutorialId: tutorialId, No: no}
	t.Steps = append(t.Steps, newStep)
	this.Db.C(tutorialCollection).UpsertId(t.Id, t)
	return &newStep
}

func (this TutorialStore) UpdateStep(stepId, selector, text string, no int) *Step {
	q := bson.M{"steps": bson.M{"$elemMatch": bson.M{"_id": bson.ObjectIdHex(stepId)}}}
	rq := bson.M{"$set": bson.M{"steps.$.selector": selector, "steps.$.text": text, "steps.$.no": no}}
	err := this.Db.C(tutorialCollection).Update(q, rq)

	if err != nil {
		panic(err)
	}

	newStep := Step{Id: bson.ObjectIdHex(stepId), Selector: selector, Text: text, No: no}
	return &newStep
}
