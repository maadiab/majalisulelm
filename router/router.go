package router

import (
	"github.com/gorilla/mux"
	Handler "github.com/maadiab/majalisulelm/handler"
	Middleware "github.com/maadiab/majalisulelm/middleware"
	"html/template" // Import the html/template package
//	main "github.com/maadiab/majalisulelm/main"
	"net/http")


var	Tmpl = template.Must(template.ParseGlob("templates/*.html"))
func Router() *mux.Router {
	router := mux.NewRouter()

	// router.HandleFunc("/", Handler.ServeHome).Methods("GET")


// Serve static files (CSS, images, etc.) from the "static" folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Tmpl.ExecuteTemplate(w, "base.html", nil)
	})


	router.HandleFunc("/createuser", Middleware.Authenticate(Handler.CreateSystemUser)).Methods("POST")

	router.HandleFunc("/user/{id}", Middleware.Authenticate(Handler.GetSystemUser)).Methods("GET")
	router.HandleFunc("/users", Middleware.Authenticate(Handler.GetSystemUsers)).Methods("GET")
	router.HandleFunc("/addlesson", Middleware.Authenticate(Handler.Create)).Methods("POST")
	router.HandleFunc("/lesson/{id}", Middleware.Authenticate(Handler.Get)).Methods("GET")
	router.HandleFunc("/lessons", Middleware.Authenticate(Handler.GetAll)).Methods("GET")
	router.HandleFunc("/deletelesson/{id}", Middleware.Authenticate(Handler.Delete)).Methods("DELETE")
	router.HandleFunc("/login", Middleware.Login).Methods("POST")

	return router
}
