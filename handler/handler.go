package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maadiab/majalisulelm/core"
	Database "github.com/maadiab/majalisulelm/database"
	"github.com/maadiab/majalisulelm/helper"
)

func CreateSystemUser(w http.ResponseWriter, r *http.Request) {
	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = helper.CreateUser(Database.DB, user)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSystemUsers(w http.ResponseWriter, r *http.Request) {

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
