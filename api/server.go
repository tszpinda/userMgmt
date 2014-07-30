package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"log"

	"labix.org/v2/mgo"

//	"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
	
	"github.com/rcrowley/go-tigertonic"
	"github.com/tszpinda/goember"
	"github.com/tszpinda/mongoutil"
	"github.com/tszpinda/userMgmt/api/store"
	"github.com/tszpinda/userMgmt/api/web"
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
		cors.Build(em.Marshaled(userApi.AddUser)))
	
		mux.Handle(
		"PUT", 
		"/sessions/{id}", 
		cors.Build(em.Marshaled(sessionApi.LoginHandler)))
		
		mux.Handle(
		"GET", 
		"/sessions/{id}", 
		cors.Build(em.Marshaled(sessionApi.GetSession)))
		
	
	/*mux.Handle(
		"POST", 
		"/users", 
		cors.Build(tigertonic.Marshaled(userApi.AddUser)))*/
//	mux.HandleFunc("/users", appHandler(AddUserHandler)).Methods("POST")
//	mux.HandleFunc("/sessions/{key}", appHandler(SessionHandler)).Methods("GET")
//	mux.HandleFunc("/sessions/{key}", appHandler(LoginHandler)).Methods("PUT")

//	n := negroni.Classic()
//	n.Use(negroni.HandlerFunc(CrossDomainMiddleware))
//	n.UseHandler(mux)

//	n.Run(":3000")
	server := tigertonic.NewServer(":3000", tigertonic.ApacheLogged(mux))
	err := server.ListenAndServe()

	if nil != err {
		log.Fatalln(err)
	}
}

//func CrossDomainMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
//	addHeaders(w)
//	if r.Method == "OPTIONS" {
//		log.Println("options request")
//		return
//	}
//	next(w, r)
//}
//
//func AddUserHandler(w http.ResponseWriter, req *http.Request) (err error) {
//	data := struct {
//		User *User
//	}{}
//	if err = encode(req, &data); err != nil {
//		return err
//	}
//	u := data.User
//	//u = AddUser(mdb, u.Email, u.Password, u.Name)
//
//	return em.Post(w, "user", u)
//}
//
//func SessionHandler(w http.ResponseWriter, req *http.Request) (err error) {
//	s := new(Session)
//	s.Id = "current"
//	return em.Post(w, "session", s)
//}
//


func appHandler(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addHeaders(w)
		err := fn(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Printf("handling %q: %v", r.RequestURI, err)
		}
	}
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, OPTIONS, DELETE")
	//other allowed headers		   "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, AuthToken")
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func encode(req *http.Request, entity interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(entity); err != nil {
		return err
	}
	return nil
}

func getenv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

func initMongo() *mgo.Database {
	url, dbName := mgou.GetMongoConfg("userMgmt")
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session.DB(dbName)
}
