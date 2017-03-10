package server

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"../pkg/mux"
)

// Router is mux NewRouter() handler
var Router = mux.NewRouter()

// Start server
func Start() {
	fmt.Println("server start")

	http.ListenAndServe(":8001", Router)
}

// GetSession method start session with MongoDB
func GetSession() *mgo.Session {
	fmt.Println("get session")

	session, error := mgo.Dial("mongodb://localhost")

	if error != nil {
		panic(error)
	}

	return session
}

// TODO: end session
