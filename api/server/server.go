package server

import (
	"log"
	"net/http"

	"../pkg/mux"
)

// Router is mux NewRouter() handler
var Router = mux.NewRouter()

// Start server
func Start() {
	log.Fatal(http.ListenAndServe(":8001", Router))
}
