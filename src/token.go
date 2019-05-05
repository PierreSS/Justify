package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("tictactrip")

//Liste de mails obligatoire pour avoir un token
var mail = []string{
	"foo@bar.com",
}

//Credentials pour le token
type Credentials struct {
	Email string `json:"email"`
}

//Claims pour le token
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

//Vérifie si le token du cookie est valide
func verifyToken(w http.ResponseWriter, r *http.Request) (bool, string) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false, string(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusBadRequest)
		return false, string(http.StatusBadRequest)
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false, string(http.StatusUnauthorized)
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false, string(http.StatusUnauthorized)
		}
		w.WriteHeader(http.StatusBadRequest)
		return false, string(http.StatusBadRequest)
	}
	return true, string(http.StatusOK)
}

//Vérifie qu'un mail est bien enregistré
func verify(creds *Credentials) bool {
	for x := 0; x != len(mail); x++ {
		if mail[x] == creds.Email {
			return true
		}
	}
	return false
}

//Créé un cookie contenant le token unique
func createToken(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !verify(&creds) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(2 * time.Minute)
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Path:    "/",
		Value:   tokenString,
		Expires: expirationTime,
	})
	fmt.Fprintf(w, "%s", tokenString)
}
