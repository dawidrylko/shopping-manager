package server

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"../pkg/mux"
)

// Router is mux NewRouter() handler
var Router = mux.NewRouter()

// Start server
func Start() {
	log.Fatal(http.ListenAndServe(":8001", Router))
}

func GetSession() *mgo.Session {
	fmt.Println("get session")
	session, error := mgo.Dial("mongodb://localhost")

	if error != nil {
		panic(error)
	}
	return session
}
