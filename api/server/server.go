package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Router is mux NewRouter() handler
var Router = mux.NewRouter()

// Start server
func Start() {
	fmt.Println(time.Now(), "Go HTTP server start")

	http.ListenAndServe(":8001", Router)
}

// GetSession method start session with MongoDB
func GetSession() *mgo.Session {
	fmt.Println(time.Now(), "Connect with MongoDB")

	session, error := mgo.Dial("mongodb://localhost")

	if error != nil {
		panic(error)
	}

	return session
}

// FinishSession method finish session with MongoDB
func FinishSession(session *mgo.Session) {
	fmt.Println(time.Now(), "Disconnect with MongoDB")

	session.Close()
}
