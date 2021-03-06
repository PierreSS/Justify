package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Set le port
func balanceTonPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func main() {
	addr := balanceTonPort()
	r := mux.NewRouter()
	handleRequest(r)
	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
