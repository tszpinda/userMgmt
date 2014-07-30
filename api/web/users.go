package web

import (
	"github.com/tszpinda/userMgmt/api/store"
	"net/http"
	"net/url"
	"log"
)

type UserResource struct {
	UserStore store.UserStore
}


func (this *UserResource) AddUser(url *url.URL, inHeaders http.Header, user *store.User) (int, http.Header, interface{}, error) {
	log.Printf("user: %+v", user)
	u := this.UserStore.AddUser(user.Email, user.Password, user.Name)
	return 200, nil, u, nil
}
