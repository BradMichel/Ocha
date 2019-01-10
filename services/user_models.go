package services

import (
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}

var (
	nameSession = fmt.Sprintf("%X", securecookie.GenerateRandomKey(15))
	store       = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))
)
