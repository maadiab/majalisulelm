package Handler

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
	Middleware "github.com/maadiab/majalisulelm/middleware"
)

func hasPermissions(userPermissions []string, requiredPermissions []string) bool {
	for _, perm := range requiredPermissions {
		found := false
		for _, userPerm := range userPermissions {
			if perm == userPerm {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func ServeHome(w http.ResponseWriter, r *http.Request) {

	// ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	// fmt.Println(ctx)
	// defer cancel()

	helper.ServeTemplates(w, "home.page.html")
}

// // login

// func Login(w http.ResponseWriter, r *http.Request) {

// 	var user middleware.Credentials

// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		log.Println("Error Decoding user Data for login")
// 	}

// 	middleware.CheckUser(Database.DB, user)
// }

func CreateSystemUser(w http.ResponseWriter, r *http.Request) {

	requiredPermissons := []string{"read", "write"}
	claims, ok := r.Context().Value("claims").(*Middleware.Claims)
	if !ok {
		log.Println("no permissions found", ok)
		http.Error(w, "permissons not found !!! ", http.StatusInternalServerError)
		return
	}

	if !hasPermissions(claims.Permissions, requiredPermissons) {
		http.Error(w, "insufficient permissions !!!", http.StatusForbidden)
		return
	}

	var user core.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error here from decoding json !!")
	}
	helper.CreateUser(Database.DB, user)

}

func GetSystemUser(w http.ResponseWriter, r *http.Request) {

	// set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	var user core.User

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 32, 32)
	if err != nil {
		log.Println("No user found !!!")
		return
	}

	user, nil := helper.GetUser(Database.DB, int(userID))

	// convert the user struct to json
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Println("Error no user found !!!")
	}

	// write the json data to response writer
	w.Write(jsonData)
}

func GetSystemUsers(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json")

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

	var lesson core.Lesson
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error here from decoding lesson json !!")
	}
	err = helper.CreateLesson(Database.DB, lesson)
	if err != nil {
		log.Fatal(err)
	}

}

func GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	lessons := helper.GetAllLessons(Database.DB)
	data, err := json.Marshal(lessons)
	if err != nil {
		log.Println("Error Marshalling lessons json!!!")
	}

	w.Write(data)

}

func Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var lesson core.Lesson
	params := mux.Vars(r)
	lessonID, err := strconv.ParseUint(params["id"], 32, 32)
	if err != nil {
		log.Fatal(err)
	}
	lesson, nil := helper.GetLessonById(Database.DB, int(lessonID))
	if err != nil {
		log.Println(err)
	}
	data, err := json.Marshal(lesson)
	if err != nil {

		log.Println("Error Parsing Lesson")
	}
	w.Write(data)

}

func DeleteAll(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

	// requiredPermissions := []string{"admin"}
	params := mux.Vars(r)
	lessonID, err := strconv.ParseUint(params["id"], 32, 32)
	if err != nil {
		log.Fatal(err)
	}
	helper.DeleteLessonById(Database.DB, int(lessonID))

	if err != nil {
		log.Println("Error Deleting Lesson!!!")
	}

}
