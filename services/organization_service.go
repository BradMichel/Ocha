package services

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func (organization *Organization) Get(db *gorm.DB) (responseStatus int, err error) {
	if db == nil {
		db = Conect()
		defer db.Close()
	}

	db.First(organization, organization.ID)
	responseStatus = http.StatusOK
	return
}

func (organization *Organization) Create() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()

	if !db.NewRecord(organization) {
		responseStatus = http.StatusBadRequest
		return
	}

	db.Create(organization)

	if db.NewRecord(organization) {
		responseStatus = http.StatusBadRequest
		return
	}

	responseStatus = http.StatusOK
	return
}

func (organization *Organization) Update() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()
	db.Save(organization)
	responseStatus = http.StatusOK
	return
}

func (organization *Organization) Delete() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()
	db.Delete(organization)
	db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(organization)

	responseStatus = http.StatusOK
	return
}

func (organizations *Organizations) Get() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()

	if organizations.IDLast == 0 {
		db.Limit(organizations.Limit).Order("id DESC").Find(&organizations.List)
	} else {
		db.Where("id < ?", organizations.IDLast).Limit(organizations.Limit).Order("id DESC").Find(&organizations.List)
	}
	if err != nil {
		log.Println(err)
		responseStatus = http.StatusBadRequest
		return
	}
	log.Println(organizations)
	responseStatus = http.StatusOK
	return
}

func (organizations *Organizations) Create() (responseStatus int, err error) {
	for _, organization := range organizations.List {
		responseStatus, err = organization.Create()
		if err != nil {
			panic(err)
		}
	}
	responseStatus = http.StatusOK
	return
}

func (organizations *Organizations) Update() (responseStatus int, err error) {
	for _, organization := range organizations.List {
		responseStatus, err = organization.Update()
		if err != nil {
			panic(err)
		}
	}
	responseStatus = http.StatusOK
	return
}

func (organizations *Organizations) Delete() (responseStatus int, err error) {
	for _, organization := range organizations.List {
		responseStatus, err = organization.Delete()
		if err != nil {
			panic(err)
		}
	}
	responseStatus = http.StatusOK
	return
}
