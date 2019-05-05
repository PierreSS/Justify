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

	//	port, err := balanceTonPort()
	//	checkError(err)

	r := mux.NewRouter()
	handleRequest(r)
	//	log.Fatal(http.ListenAndServe(":"+env.PortWebRequest, r))
	log.Fatal(http.ListenAndServe(":8005", r))
}
