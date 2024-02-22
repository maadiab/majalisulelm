package helper

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/maadiab/majalisulelm/core"
	"golang.org/x/crypto/bcrypt"
)

// serve html templates
func ServeTemplates(w http.ResponseWriter, tmpl string) {

	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Println("Error Parsing template", err)
	}

}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

// create user record

func CreateUser(db *sqlx.DB, user core.User) {

	// if user send an empty json

	// if user send incorrect json

	// if it good

	inputPassword := user.Password

	hashedPassword, err := HashPassword(inputPassword)
	if err != nil {
		log.Println("Error Hashing password !!!", err)
	}

	_, err = db.Exec("INSERT INTO users (name, mobile, email, password, permissions) VALUES ($1,$2,$3,$4,$5)",
		user.Name, user.Mobile, user.Email, hashedPassword, user.Permissions)

	if err != nil {
		log.Println("Error Creating user !!!", err)
	}

	// log.Println("You Are Not Authenticated, Please Sign In !!!")

}

// get all users
func GetUsers(db *sqlx.DB) ([]core.User, error) {

	var users []core.User
	// shoud be a loop here
	err := db.Select(&users, "SELECT name, mobile, email FROM users")
	if err != nil {
		fmt.Println("Error in function GetUsers !")
	}
	return users, err
}

// get user by id
func GetUser(db *sqlx.DB, userID int) (core.User, error) {

	var user core.User
	err := db.Get(&user, "SELECT id, name, mobile, email, password FROM users where id= $1", userID)
	if err != nil {
		log.Println("NO user found !!!")
	}
	return user, nil
}

// update user
func UpdateUser(db *sqlx.DB, userId int, userData core.User) error {

	_, err := db.Exec(`"UPDATE users
	SET name = $1, mobile =$2, email = $3, password= $4
	where id =$5
	"`, userId, userData.Name, userData.Mobile, userData.Email, userData.Password)

	if err != nil {

		log.Println("Error Updating User !!!")
	}

	return err
}

// delete user
func DeleteUser(db *sqlx.DB, userId int) error {
	_, err := db.Exec("DELETE FROM users where id=$1", userId)
	if err != nil {

		log.Println("Error Deleting user !!!")
	}

	return err
}

// get all lessons

func GetAllLessons(db *sqlx.DB) []core.Lesson {
	var lessons []core.Lesson

	err := db.Select(&lessons, "SELECT * FROM lessons")
	if err != nil {
		log.Println("Error Getting Lessons From Database!!!")
	}
	return lessons

}

// delete all lessons
func DeleteAllLessons(db *sqlx.DB) error {

	_, err := db.Exec("DELETE * FROM lessons")
	if err != nil {

		log.Println("Error Deleting All Lessons!!!")
	}
	return err
}

// get lesson by id
func GetLessonById(db *sqlx.DB, lessonID int) (core.Lesson, error) {
	var lesson core.Lesson

	err := db.Get(&lesson, "SELECT id,name,type,author,duration,time,link,location FROM lessons where id=$1", lessonID)

	if err != nil {

		log.Println("Error getting lesson from db!!!")
	}

	return lesson, err
}

// delete lesson
func DeleteLessonById(db *sqlx.DB, lessonId int) error {
	_, err := db.Exec("DELETE FROM lessons where id=$1", lessonId)
	if err != nil {
		log.Println("Error Deleting From Lessons!!!")
	}
	return err
}

// create lesson
func CreateLesson(db *sqlx.DB, lesson core.Lesson) error {
	_, err := db.Exec("INSERT INTO lessons (name, type, duration, author, link, time, location) VALUES ($1,$2,$3,$4,$5,$6,$7)", lesson.Name, lesson.Type, lesson.Duration, lesson.Author, lesson.Link, lesson.Time, lesson.Location)
	if err != nil {
		log.Println("Error Adding Lesson To Database !")
	}

	return err
}
