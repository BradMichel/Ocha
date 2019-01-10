package services

import (
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (user *User) Validation(rw http.ResponseWriter, r *http.Request) (responseStatus int, err error) {
	log.Println("Validation")

	if responseStatus, err = user.Session(rw, r); err != nil {
		log.Println(err)
	}

	return

}

func (user *User) Login(rw http.ResponseWriter, r *http.Request) (responseStatus int, err error) {
	log.Println("Login")
	password := user.Password
	db := Conect()
	defer db.Close()

	db.Where("name=?", user.Name).First(&user)
	if err != nil {
		log.Println(err)
		responseStatus = http.StatusBadRequest
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		responseStatus = http.StatusBadRequest
		return
	}

	session, err := store.Get(r, nameSession)
	if err != nil {
		log.Println(err)
		responseStatus = http.StatusInternalServerError
		return
	}
	user.Password = ""
	session.Values["IDUser"] = user.ID
	session.Save(r, rw)
	responseStatus = http.StatusOK
	return

}

func (user *User) Session(rw http.ResponseWriter, r *http.Request) (responseStatus int, err error) {
	session, err := store.Get(r, nameSession)
	if err != nil {
		responseStatus = http.StatusInternalServerError
		return
	}
	if session.Values["IDUser"] == nil {
		responseStatus = http.StatusUnauthorized
		err = errors.New("StatusUnauthorized")
		return
	}
	responseStatus = http.StatusAccepted
	return
}

func (user *User) Logout(rw http.ResponseWriter, r *http.Request) (responseStatus int, err error) {
	session, err := store.Get(r, nameSession)
	if err != nil {
		responseStatus = http.StatusInternalServerError
		return
	}
	session.Options.MaxAge = -1
	responseStatus = http.StatusAccepted
	return
}
