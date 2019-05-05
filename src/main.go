package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Set le port
func balanceTonPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := balanceTonPort()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	handleRequest(r)
	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
