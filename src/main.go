package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Set le port
func balanceTonPort() string {
	port := os.Getenv("PORT")
	return ":" + port
}

func main() {
	r := mux.NewRouter()
	handleRequest(r)
	log.Fatal(http.ListenAndServe(balanceTonPort(), r))
}
