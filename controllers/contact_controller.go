package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	s "../services"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func GetContact(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("GetContact")
	log.Println(context.Get(r, "user"))
	contact := new(s.Contact)
	contact.ID, _ = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if responseStatus, err := contact.Get(); err != nil {
		log.Println(err)
		w.WriteHeader(responseStatus)
		return
	} else {
		responseJson(w, contact, responseStatus)
	}
}

func PostContact(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("PostContact")

	contact := new(s.Contact)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(contact); err != nil {
		responseStatus := http.StatusBadRequest
		w.WriteHeader(responseStatus)
		return
	} else {
		responseStatus, _ := contact.Create()
		responseJson(w, contact, responseStatus)
	}
	defer r.Body.Close()
}

func UpdateContact(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contact := new(s.Contact)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(contact); err != nil {
		responseStatus := http.StatusBadRequest
		rw.WriteHeader(responseStatus)
		return
	} else {
		responseStatus, _ := contact.Update()
		responseJson(rw, contact, responseStatus)
	}
	defer r.Body.Close()
}

func DeleteContact(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contact := new(s.Contact)
	contact.ID, _ = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	responseStatus, err := contact.Delete()

	if err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
		return
	}
}

func GetContacts(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("GetContacts")

	contacts := new(s.Contacts)
	contacts.Page, _ = strconv.Atoi(mux.Vars(r)["page"])
	contacts.Limit, _ = strconv.Atoi(mux.Vars(r)["limit"])
	contacts.IDLast, _ = strconv.Atoi(mux.Vars(r)["idLast"])

	responseStatus, err := contacts.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(responseStatus)
		return
	}
	responseJson(w, contacts, responseStatus)

}

func PostContacts(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contacts := new(s.Contacts)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(contacts); err != nil {
		responseStatus := http.StatusBadRequest
		rw.WriteHeader(responseStatus)
		return
	} else {
		responseStatus, _ := contacts.Create()
		responseJson(rw, contacts, responseStatus)
	}
	defer r.Body.Close()
}

func UpdateContacts(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contacts := new(s.Contacts)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(contacts); err != nil {
		responseStatus := http.StatusBadRequest
		rw.WriteHeader(responseStatus)
		return
	} else {
		responseStatus, _ := contacts.Update()
		responseJson(rw, contacts, responseStatus)
	}
	defer r.Body.Close()
}

func DeleteContacts(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contacts := new(s.Contacts)

	if responseStatus, err := contacts.Delete(); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
		return
	}
}
