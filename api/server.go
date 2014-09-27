package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/rcrowley/go-tigertonic"
	"github.com/tszpinda/mongoutil"
	"github.com/tszpinda/userMgmt/api/store"
	"github.com/tszpinda/userMgmt/api/web"
	"labix.org/v2/mgo"
)

var sessionStore store.SessionStore
var userStore store.UserStore

func main() {
	mdb := initMongo()

	mux := tigertonic.NewTrieServeMux()

	userStore = store.UserStore{Db: mdb}
	userApi := web.UserResource{UserStore: userStore}

	sessionStore = store.SessionStore{Db: mdb}
	sessionApi := web.SessionResource{UserStore: userStore, SessionStore: sessionStore}

	tutorialStore := store.TutorialStore{Db: mdb}
	tutorialApi := web.TutorialResource{TutorialStore: tutorialStore}

	// We'll use this CORSBuilder to set Access-Control-Allow-Origin headers
	// on certain endpoints.
	cors := tigertonic.NewCORSBuilder()
	cors.AddAllowedOrigins("*")
	cors.AddAllowedHeaders("Origin, X-Requested-With, Content-Type, Accept, auth-token")

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
		cors.Build(tigertonic.If(withCtxHandler, tigertonic.Marshaled(sessionApi.GetSession))))

	mux.Handle(
		"DELETE",
		"/sessions/{id}",
		cors.Build(tigertonic.Marshaled(sessionApi.LogoutHandler)))

	mux.Handle(
		"GET",
		"/users/{id}",
		cors.Build(tigertonic.If(authenticatedHandler, tigertonic.Marshaled(userApi.GetUser))))

	mux.Handle(
		"GET",
		"/xtutorials/{apiKey}/{domain}/{page}",
		cors.Build(tigertonic.Marshaled(tutorialApi.GetForPage)))

	mux.Handle(
		"GET",
		"/tutorials",
		cors.Build(tigertonic.If(authenticatedHandler, tigertonic.Marshaled(tutorialApi.GetForApiKey))))

	mux.Handle(
		"POST",
		"/tutorials",
		cors.Build(tigertonic.If(authenticatedHandler, tigertonic.Marshaled(tutorialApi.AddTutorial))))

	mux.Handle(
		"PUT",
		"/tutorials/{id}",
		cors.Build(tigertonic.If(authenticatedHandler, tigertonic.Marshaled(tutorialApi.UpdateTutorial))))

	mux.Handle(
		"GET",
		"/tutorials/{id}",
		cors.Build(tigertonic.If(authenticatedHandler, tigertonic.Marshaled(tutorialApi.GetById))))

	mux.Handle(
		"POST",
		"/steps",
		cors.Build(tigertonic.If(authenticatedHandler, tigertonic.Marshaled(tutorialApi.AddStep))))

	mux.Handle(
		"PUT",
		"/steps/{id}",
		cors.Build(tigertonic.If(authenticatedHandler, tigertonic.Marshaled(tutorialApi.UpdateStep))))

	server := tigertonic.NewServer(":3000", tigertonic.ApacheLogged(tigertonic.WithContext(mux, store.Ctx{})))
	err := server.ListenAndServe()

	if nil != err {
		log.Fatalln(err)
	}
}

func populateCtx(r *http.Request) *store.Ctx {
	token := r.Header.Get("auth-token")
	ctx := tigertonic.Context(r).(*store.Ctx)
	log.Println("authToken:", token)
	if token == "" {
		log.Println("no token")
		return ctx
	}
	session, err := sessionStore.FindSessionByAuthToken(token)
	if err != nil {
		log.Println("session not found:", err)
		return ctx
	}
	user, err := userStore.FindUserById(session.UserId.Hex())
	if err != nil {
		log.Println("user not found:", err)
		return ctx
	}
	//log.Println("user:", user)
	tigertonic.Context(r).(*store.Ctx).User = user
	tigertonic.Context(r).(*store.Ctx).ApiKey = user.Id.Hex()
	return ctx
}

func withCtxHandler(r *http.Request) (http.Header, error) {
	populateCtx(r)
	return nil, nil
}
func authenticatedHandler(r *http.Request) (http.Header, error) {
	ctx := populateCtx(r)
	if ctx.User == nil {
		return nil, authError()
	}
	return nil, nil
}

func authError() tigertonic.Forbidden {
	log.Println("authError")
	return tigertonic.Forbidden{errors.New("forbidden")}
}

func initMongo() *mgo.Database {
	url, dbName := mgou.GetMongoConfg("userMgmt")
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session.DB(dbName)
}
