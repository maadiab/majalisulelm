package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	Database "github.com/maadiab/majalisulelm/database"
	"github.com/maadiab/majalisulelm/handler"
)

func main() {
	Database.ConnectDB()
	Database.CreateUsersTable()
	r := mux.NewRouter()
	r.HandleFunc("/", handler.ServeHome).Methods("GET")
	r.HandleFunc("/createuser", handler.CreateSystemUser).Methods("POST")
	r.HandleFunc("/getuser/{id}", handler.GetSystemUser).Methods("GET")
	r.HandleFunc("/getallusers", handler.GetSystemUsers).Methods("GET")

	fmt.Println("server is running at port: 8080!")
	http.ListenAndServe(":8080", r)

	defer Database.DB.Close()

}
