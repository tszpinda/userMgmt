package web

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/tszpinda/goember"
	"github.com/tszpinda/userMgmt/api/store"
)

type UserResource struct {
	UserStore store.UserStore
}

func (this *UserResource) AddUser(url *url.URL, inHeaders http.Header, m map[string]*store.User) (int, http.Header, interface{}, error) {
	user, _ := m["user"]

	if valErr := em.Required("name", user.Name); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("email", user.Email); valErr != nil {
		return em.ValidationResponse(valErr)
	}
	if valErr := em.Required("password", user.Password); valErr != nil {
		return em.ValidationResponse(valErr)
	}

	_, err := this.UserStore.FindUserByEmail(user.Email)
	if err == nil {
		return em.ValidationResponse(em.ValidationError("email", "already exists"))
	}
	u := this.UserStore.AddUser(user.Email, user.Password, user.Name)
	m["user"] = u
	return 200, nil, m, nil
}

func (this *UserResource) GetUser(url *url.URL, inHeaders http.Header, _ map[string]interface{}, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	id := url.Query().Get("id")
	if ctx.User.Id.Hex() != id {
		log.Println("id not the same", ctx.User.Id.Hex(), "!=", id)
		return 403, nil, nil, errors.New("forbidden")
	}
	response := make(map[string]interface{})
	u, err := this.UserStore.FindUserById(id)
	if err != nil {
		log.Println("user lookup failed", id)
		return 404, nil, nil, errors.New("Not found")
	}
	response["user"] = u
	return 200, nil, response, nil
}
