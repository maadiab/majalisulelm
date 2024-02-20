package router

import (
	"github.com/gorilla/mux"
	Handler "github.com/maadiab/majalisulelm/handler"
	Middleware "github.com/maadiab/majalisulelm/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Handler.ServeHome).Methods("GET")

	router.HandleFunc("/createuser", Middleware.Authenticate(Handler.CreateSystemUser)).Methods("POST")

	router.HandleFunc("/getuser/{id}", Middleware.Authenticate(Handler.GetSystemUser)).Methods("GET")
	router.HandleFunc("/getallusers", Middleware.Authenticate(Handler.GetSystemUsers)).Methods("GET")
	router.HandleFunc("/addlesson", Middleware.Authenticate(Handler.Create)).Methods("POST")
	router.HandleFunc("/getlesson/{id}", Middleware.Authenticate(Handler.Get)).Methods("GET")
	router.HandleFunc("/getlessons", Middleware.Authenticate(Handler.GetAll)).Methods("GET")
	router.HandleFunc("/deletelesson/{id}", Middleware.Authenticate(Handler.Delete)).Methods("DELETE")
	router.HandleFunc("/login", Middleware.Login).Methods("POST")

	return router
}
