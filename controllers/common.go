package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseJson(w http.ResponseWriter, data interface{}, responseStatus int) {
	responseJson, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(responseJson)
}
