package routers

import (
	"html/template"

	"log"
	"net/http"

	"../controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Route struct {
	Method         string
	Pattern        string
	Authentication bool
	f              interface{}
}

type Routes []Route

var routes = Routes{
	Route{
		"POST", "/login", false, controllers.Login,
	},
	Route{
		"GET", "/api/thematics", true, controllers.GetThematics,
	},
	Route{
		"GET", "/api/typeOfThematics", true, controllers.GetTypeOfThematics,
	},
	Route{
		"GET", "/api/contacts/page/{page}/limit/{limit}/last/{idLast}", true, controllers.GetContacts,
	},
	Route{
		"POST", "/api/contacts/", true, controllers.PostContacts,
	},
	Route{
		"UPDATE", "/api/contacts/", true, controllers.UpdateContacts,
	},
	Route{
		"DELETE", "/api/contacts/", true, controllers.DeleteContacts,
	},
	Route{
		"GET", "/api/contacts/{id}", true, controllers.GetContact,
	},
	Route{
		"POST", "/api/contacts/{name}", true, controllers.PostContact,
	},
	Route{
		"UPDATE", "/api/contacts/{id}", true, controllers.UpdateContact,
	},
	Route{
		"DELETE", "/api/contacts/{id}", true, controllers.DeleteContact,
	},
}

var homeTemplate = template.Must(template.ParseFiles("./static/src/index.html"))

func SetRoutes(router *mux.Router) *mux.Router {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("rutas del del api")
	for _, route := range routes {
		log.Println("Method:", route.Method, "Pattern", route.Pattern, "Authentication", route.Authentication)
		log.Println()
		var n *negroni.Negroni
		if route.Authentication {
			n = negroni.New(
				negroni.HandlerFunc(controllers.Authentication),
				negroni.HandlerFunc(route.f.(func(http.ResponseWriter, *http.Request, http.HandlerFunc))),
			)
		} else {
			n = negroni.New(
				negroni.HandlerFunc(route.f.(func(http.ResponseWriter, *http.Request, http.HandlerFunc))),
			)
		}

		router.Handle(route.Pattern, n).Methods(route.Method)

		//fs := http.FileServer(http.Dir("static/src"))
		//router.Handle("/", fs)

		/* router.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
			wr.Header().Set("Content-Type", "text/html; charset=utf-8")
			homeTemplate.Execute(wr, r.Host)
		}) */

	}

	return router

}
