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
	router.HandleFunc("/addlesson", handler.Create).Methods("POST")
	router.HandleFunc("/getlesson/{id}", handler.Get).Methods("GET")

	return router
}
