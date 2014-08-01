package web

import (
	"github.com/tszpinda/userMgmt/api/store"
	"github.com/tszpinda/goember"
	"net/http"
	"net/url"
)

type SessionResource struct {
	SessionStore store.SessionStore
	UserStore store.UserStore
}

func (this *SessionResource) LoginHandler(url *url.URL, inHeaders http.Header, m map[string]*store.User) (int, http.Header, interface{}, error) {
	loginCredentials := m["session"]
	response := make(map[string]interface{})
	
	u, err := this.UserStore.FindUserByEmail(loginCredentials.Email)
	if err != nil || !store.PasswdEqual(loginCredentials.Password, u.Password) {
		return 422, nil, em.ValidationError("email", "or password is invalid"), nil
	}

	session := this.SessionStore.CreateAuthSession(u.Id.Hex())
	session.Id = "current"

	response["session"] = session
	return 200, nil, response, nil
}

func (this *SessionResource) GetSession(url *url.URL, inHeaders http.Header, m map[string]interface{}) (int, http.Header, interface{}, error) {
	s := new(store.Session)
	s.Id = "current"
	response := make(map[string]interface{})
	response["session"] = s
	return 200, nil, response, nil
}