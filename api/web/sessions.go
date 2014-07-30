package web

import (
	"github.com/tszpinda/userMgmt/api/store"
	"github.com/tszpinda/goember"
	"net/http"
	"net/url"
	"log"
)

type SessionResource struct {
	SessionStore store.SessionStore
	UserStore store.UserStore
}

type Session struct {
	Email string `json:"email"`
	Password string`json:"password"`
}

func (this *SessionResource) LoginHandler(url *url.URL, inHeaders http.Header, login *Session) (int, http.Header, interface{}, error) {
	log.Printf("login: '%+v'", login.Email)
	u, err := this.UserStore.FindUserByEmail(login.Email)
	log.Printf("found: %+v, %+v", u, err)
	if err != nil || !store.PasswdEqual(login.Password, u.Password) {
		return 422, nil, em.ValidationError("email", "or password is invalid"), nil
	}
	session := this.SessionStore.CreateAuthSession(u.Id.Hex())
	session.Id = "current"

	return 200, nil, session, nil
}

func (this *SessionResource) GetSession(url *url.URL, inHeaders http.Header, m map[string]interface{}) (int, http.Header, interface{}, error) {
	s := new(store.Session)
	s.Id = "current"
	return 200, nil, s, nil
}