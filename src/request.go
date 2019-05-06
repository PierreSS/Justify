package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var m = make(map[string]int)
var mTime = make(map[string]int64)

//Gere toute les routes du serveur HTTP
func handleRequest(router *mux.Router) {
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/api/justify", justify).Methods("POST")
	router.HandleFunc("/api/token", createToken).Methods("POST")
}

//Base route
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there, welcome to my justify api !</h1>")
}

//Justify route
func justify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	claims, ok := verifyToken(w, r)
	if !ok {
		fmt.Fprintf(w, "%s", "error")
		return
	}
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	str := justifytxt(responseString, 80, claims.Email, w)

	fmt.Fprintf(w, "%s", str)
}
