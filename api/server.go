package main

import (
	"log"

	"github.com/rcrowley/go-tigertonic"
	"github.com/tszpinda/mongoutil"
	"github.com/tszpinda/userMgmt/api/store"
	"github.com/tszpinda/userMgmt/api/web"
	"labix.org/v2/mgo"
)

func main() {
	mdb := initMongo()

	mux := tigertonic.NewTrieServeMux()

	userStore := store.UserStore{Db: mdb}
	userApi := web.UserResource{UserStore: userStore}

	sessionStore := store.SessionStore{Db: mdb}
	sessionApi := web.SessionResource{UserStore: userStore, SessionStore: sessionStore}

	// We'll use this CORSBuilder to set Access-Control-Allow-Origin headers
	// on certain endpoints.
	cors := tigertonic.NewCORSBuilder()
	cors.AddAllowedOrigins("*")
	cors.AddAllowedHeaders("Origin, X-Requested-With, Content-Type, Accept, AuthToken")

	mux.Handle(
		"POST",
		"/users",
		cors.Build(tigertonic.Marshaled(userApi.AddUser)))

	mux.Handle(
		"PUT",
		"/sessions/{id}",
		cors.Build(tigertonic.Marshaled(sessionApi.LoginHandler)))

	mux.Handle(
		"GET",
		"/sessions/{id}",
		cors.Build(tigertonic.Marshaled(sessionApi.GetSession)))

	server := tigertonic.NewServer(":3000", tigertonic.ApacheLogged(mux))
	err := server.ListenAndServe()

	if nil != err {
		log.Fatalln(err)
	}
}

func initMongo() *mgo.Database {
	url, dbName := mgou.GetMongoConfg("userMgmt")
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session.DB(dbName)
}
