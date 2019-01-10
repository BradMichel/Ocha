package main

import (
	"log"
	"net/http"

	"./routers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
)

func main() {
	log.Println("main")
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./static/node_modules/"))))
	http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("./static/src/"))))
	http.Handle("/", context.ClearHandler(router))
	log.Println("listening localhost:8080")
	log.Println("UI localhost:8080/app")
	http.ListenAndServe(":8080", nil)

}
