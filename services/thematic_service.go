package services

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

func (thematic *Thematic) Get(db *gorm.DB) (responseStatus int, err error) {
	if db == nil {
		db := Conect()
		defer db.Close()
	}
	db.First(thematic, thematic.ID)
	db.Preload("Typeofthematic").Find(&thematic)
	responseStatus = http.StatusOK
	return
}

func (thematics *Thematics) Get() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()
	db.Find(&thematics.List)
	db.Preload("Typeofthematic").Find(&thematics.List)
	responseStatus = http.StatusOK
	return
}

func (typeOfThematics *Typeofthematics) Get() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()
	db.Find(&typeOfThematics.List)
	responseStatus = http.StatusOK
	return
}
