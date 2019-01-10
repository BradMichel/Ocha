package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	s "../services"

	"github.com/gorilla/mux"
)

func GetOrganization(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organization := new(s.Organization)
	organization.ID, _ = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if responseStatus, err := organization.Get(nil); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
	} else {
		responseJson(rw, organization, responseStatus)
	}
}

func PostOrganization(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organization := new(s.Organization)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(organization); err != nil {
		responseStatus := http.StatusBadRequest
		rw.WriteHeader(responseStatus)
	} else {
		responseStatus, _ := organization.Create()
		responseJson(rw, organization, responseStatus)
	}
	defer r.Body.Close()
}

func UpdateOrganization(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organization := new(s.Organization)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(organization); err != nil {
		responseStatus := http.StatusBadRequest
		rw.WriteHeader(responseStatus)
	} else {
		responseStatus, _ := organization.Update()
		responseJson(rw, organization, responseStatus)
	}
	defer r.Body.Close()
}

func DeleteOrganization(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organization := new(s.Organization)
	organization.ID, _ = strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

	if responseStatus, err := organization.Delete(); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
	} else {
		responseJson(rw, organization, responseStatus)
	}
}

func GetOrganizations(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organizations := new(s.Organizations)
	organizations.Page, _ = strconv.Atoi(mux.Vars(r)["page"])

	if responseStatus, err := organizations.Get(); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
	} else {
		responseJson(rw, organizations, responseStatus)
	}
}

func PostOrganizations(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organizations := new(s.Organizations)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(organizations); err != nil {
		responseStatus := http.StatusBadRequest
		rw.WriteHeader(responseStatus)
	} else {
		responseStatus, _ := organizations.Create()
		responseJson(rw, organizations, responseStatus)
	}
	defer r.Body.Close()
}

func UpdateOrganizations(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organizations := new(s.Organizations)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(organizations); err != nil {
		responseStatus := http.StatusBadRequest
		rw.WriteHeader(responseStatus)
	} else {
		responseStatus, _ := organizations.Update()
		responseJson(rw, organizations, responseStatus)
	}
	defer r.Body.Close()
}

func DeleteOrganizations(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	organizations := new(s.Organizations)

	if responseStatus, err := organizations.Delete(); err != nil {
		log.Println(err)
		rw.WriteHeader(responseStatus)
	} else {
		responseJson(rw, organizations, responseStatus)
	}
}
