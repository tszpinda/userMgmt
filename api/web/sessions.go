package web

import (
	"log"
	"net/http"
	"net/url"

	"github.com/tszpinda/goember"
	"github.com/tszpinda/userMgmt/api/store"
)

type SessionResource struct {
	SessionStore store.SessionStore
	UserStore    store.UserStore
}

type CreateToken struct {
	Id        string `json:"id"`
	AuthToken string `json:"authToken"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	User      string `json:"user,omitempty"`
}

func (this *SessionResource) logout(token string) (int, http.Header, interface{}, error) {
	log.Println("logout, token present")
	//logout
	this.SessionStore.DeleteSessionByAuthToken(token)
	s := new(CreateToken)
	s.Id = "current"
	response := make(map[string]interface{})
	response["session"] = s
	return 200, nil, response, nil
}

func (this *SessionResource) LoginHandler(url *url.URL, h http.Header, m map[string]*CreateToken) (int, http.Header, interface{}, error) {
	if token := h.Get("auth-token"); token != "" {
		return this.logout(token)
	}

	createToken := m["session"]

	u, err := this.UserStore.FindUserByEmail(createToken.Email)
	if err != nil || !store.PasswdEqual(createToken.Password, u.Password) {
		return 422, nil, em.ValidationError("email", "or password is invalid"), nil
	}

	session := this.SessionStore.CreateAuthSession(u.Id.Hex())
	log.Println("Auth token created:", session.AuthToken)

	//don't send password back
	u.Password = ""
	createToken.Id = "current"
	createToken.AuthToken = session.AuthToken
	createToken.Email = ""
	createToken.Password = ""
	createToken.User = u.Id.Hex()

	response := make(map[string]interface{})
	response["session"] = createToken
	users := make([]*store.User, 0)
	users = append(users, u)
	response["users"] = users
	return 200, nil, response, nil
}

func (this *SessionResource) LogoutHandler(url *url.URL, h http.Header, _ interface{}) (int, http.Header, interface{}, error) {
	token := h.Get("auth-token")
	this.SessionStore.DeleteSessionByAuthToken(token)
	return 204, nil, nil, nil
}

func (this *SessionResource) GetSession(url *url.URL, inHeaders http.Header, m map[string]interface{}, ctx *store.Ctx) (int, http.Header, interface{}, error) {
	s := new(CreateToken)
	s.Id = "current"
	s.Email = ""
	s.Password = ""
	response := make(map[string]interface{})
	if ctx.User != nil {
		response["user"] = ctx.User
		s.User = ctx.User.Id.Hex()
	}
	response["session"] = s
	return 200, nil, response, nil
}
