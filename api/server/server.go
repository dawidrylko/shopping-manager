package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Router is mux NewRouter() handler
var Router = mux.NewRouter()

// Start server
func Start() {
	fmt.Println(time.Now(), "Go HTTP server start")

	alowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(alowedHeaders, allowedOrigins, allowedMethods)(Router)))
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
