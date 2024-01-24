package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/maadiab/majalisulelm/core"
	Database "github.com/maadiab/majalisulelm/database"
	"github.com/maadiab/majalisulelm/helper"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is the home page</h1>"))
}

func CreateSystemUser(w http.ResponseWriter, r *http.Request) {
	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error here from decoding json !!")
	}
	err = helper.CreateUser(Database.DB, user)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSystemUser(w http.ResponseWriter, r *http.Request) {

	// set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	var user core.User

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 32, 32)
	if err != nil {
		log.Fatal(err)
	}

	user, nil := helper.GetUser(Database.DB, int(userID))

	// convert the user struct to json
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	// write the json data to response writer
	w.Write(jsonData)
}

func GetSystemUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := helper.GetUsers(Database.DB)
	if err != nil {
		fmt.Println("error")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("error marshalling data")
	}

	w.Write(jsonData)
}

func DeleteSystemUser(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {

}

func GetAll(w http.ResponseWriter, r *http.Request) {

}

func GetById(w http.ResponseWriter, r *http.Request) {

}

func DeleteAll(w http.ResponseWriter, r *http.Request) {

}

func DeleteById(w http.ResponseWriter, r *http.Request) {

}
