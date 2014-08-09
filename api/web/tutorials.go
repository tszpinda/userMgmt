package web

import (
	"net/http"
	"net/url"

	"log"

	"labix.org/v2/mgo/bson"

	"github.com/tszpinda/goember"
	"github.com/tszpinda/userMgmt/api/store"
)

type TutorialResource struct {
	TutorialStore store.TutorialStore
}

func (this *TutorialResource) GetForPage(url *url.URL, h http.Header, _ interface{}) (int, http.Header, interface{}, error) {
	apiKey := url.Query().Get("apiKey")
	domain := url.Query().Get("domain")
	page := url.Query().Get("page")

	if valErr := this.validateReqFields(apiKey, domain, page); valErr != nil {
		return em.ValidationResponse(valErr)
	}

	tutorials := this.TutorialStore.FindTutorialsForPage(apiKey, domain, page)
	m := make(map[string]interface{})
	m["tutorials"] = tutorials
	return 200, nil, m, nil
}

func (this *TutorialResource) GetForApiKey(url *url.URL, h http.Header, _ interface{}, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	apiKey := ctx.ApiKey

	if valErr := em.Required("apiKey", apiKey); valErr != nil {
		return em.ValidationResponse(valErr)
	}

	tutorials := this.TutorialStore.FindTutorialsForApiKey(apiKey)
	m := make(map[string]interface{})
	m["tutorials"] = tutorials
	steps := make([]store.Step, 0)
	for _, t := range tutorials {
		for _, s := range t.Steps {
			steps = append(steps, s)
		}
	}
	m["steps"] = steps
	return 200, nil, m, nil
}

func (this *TutorialResource) AddTutorial(url *url.URL, inHeaders http.Header, m map[string]*store.Tutorial, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	t, _ := m["tutorial"]
	if valErr := em.Required("name", t.Name); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("domain", t.Domain); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	t = this.TutorialStore.AddTutorial(ctx.ApiKey, t.Domain, t.Page, t.Name)
	t.ApiKey = ""
	m["tutorial"] = t
	return 200, nil, m, nil
}

func (this *TutorialResource) UpdateTutorial(url *url.URL, inHeaders http.Header, m map[string]*store.Tutorial, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	id := url.Query().Get("id")
	t, _ := m["tutorial"]
	if valErr := em.Required("name", t.Name); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("domain", t.Domain); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	this.TutorialStore.UpdateTutorial(id, t.Page, t.Name, t.Domain)
	t.ApiKey = ""
	t.Id = bson.ObjectIdHex(id)
	m["tutorial"] = t
	return 200, nil, m, nil
}

func (this *TutorialResource) validateReqFields(apiKey, domain, page string) *em.RestError {
	if valErr := em.Required("apiKey", apiKey); valErr != nil {
		return valErr
	}
	if valErr := em.Required("domain", domain); valErr != nil {
		return valErr
	}
	if valErr := em.Required("page", page); valErr != nil {
		return valErr
	}
	return nil
}

func (this *TutorialResource) AddStep(url *url.URL, inHeaders http.Header, m map[string]*store.Step, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	s, _ := m["step"]
	log.Printf("Adding step: %+v", s)
	if valErr := em.Required("selector", s.Selector); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("text", s.Text); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	s = this.TutorialStore.AddStep(s.TutorialId, s.Selector, s.Text)
	m["step"] = s
	return 200, nil, m, nil
}

func (this *TutorialResource) UpdateStep(url *url.URL, inHeaders http.Header, m map[string]*store.Step, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	id := url.Query().Get("id")
	s, _ := m["step"]
	if valErr := em.Required("selector", s.Selector); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("text", s.Text); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	log.Printf("Updating step: %+v", s)
	this.TutorialStore.UpdateStep(id, s.Selector, s.Text)
	s.Id = bson.ObjectIdHex(id)
	m["step"] = s
	return 200, nil, m, nil
}
