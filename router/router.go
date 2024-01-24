package router

import (
	"github.com/gorilla/mux"
	"github.com/maadiab/majalisulelm/handler"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handler.ServeHome).Methods("GET")
	router.HandleFunc("/createuser", handler.CreateSystemUser).Methods("POST")
	router.HandleFunc("/getuser/{id}", handler.GetSystemUser).Methods("GET")
	router.HandleFunc("/getallusers", handler.GetSystemUsers).Methods("GET")

	return router
}
