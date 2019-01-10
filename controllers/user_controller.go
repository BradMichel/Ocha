package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	s "../services"

	"github.com/gorilla/context"
)

func Authentication(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Authentication")
	userRequest := new(s.User)

	if responseStatus, err := userRequest.Validation(rw, r); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
	} else {
		context.Set(r, "user", userRequest)
		next(rw, r)
	}
}

func Login(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Login")

	user := new(s.User)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(user); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if responseStatus, err := user.Login(rw, r); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
		return
	} else {
		responseJson(rw, user, responseStatus)
	}
}

func Logout(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user := new(s.User)
	if responseStatus, err := user.Logout(rw, r); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
		return
	} else {
		responseJson(rw, user, responseStatus)
	}
}
