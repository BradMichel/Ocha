package controllers

import (
	"log"
	"net/http"

	s "../services"
)

func GetThematics(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	thematics := new(s.Thematics)

	if responseStatus, err := thematics.Get(); err != nil {
		log.Println(err)
		w.WriteHeader(responseStatus)
		return
	} else {
		responseJson(w, thematics, responseStatus)
	}
}

func GetTypeOfThematics(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	typeOfThematics := new(s.Typeofthematics)

	if responseStatus, err := typeOfThematics.Get(); err != nil {
		log.Println(err)
		w.WriteHeader(responseStatus)
		return
	} else {
		responseJson(w, typeOfThematics, responseStatus)
	}
}
