package web

import (
	"log"
	"net/http"
	"net/url"

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
	return 200, nil, m, nil
}

func (this *TutorialResource) AddTutorial(url *url.URL, inHeaders http.Header, m map[string]*store.Tutorial, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	t, _ := m["tutorial"]
	log.Printf("in:t: %+v", t)

	if valErr := em.Required("name", t.Name); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("domain", t.Domain); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("page", t.Page); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	log.Println(t.Page)
	t = this.TutorialStore.AddTutorial(ctx.ApiKey, t.Domain, t.Page, t.Name)
	log.Printf("t: %+v", t)
	t.ApiKey = ""
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
