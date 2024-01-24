package helper

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/maadiab/majalisulelm/core"
)

// create user record

func CreateUser(db *sqlx.DB, user core.User) error {
	_, err := db.Exec("INSERT INTO users (name, mobile, email, password) VALUES ($1,$2,$3,$4)", user.Name, user.Mobile, user.Mobile, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return err
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
		log.Fatal(err)
	}
	return user, nil
}

// update user
func UpdateUser(userId int) {

}

// delete user
func DeleteUser(userId int) {

}

// get all lessons

func GetAllLessons() {

}

// delete all lessons
func DeleteAllLessons(lessonId int) {

}

// get lesson by id
func GetLessonById(lessonId int) {

}

// delete lesson
func DeleteLessonById(lessonId int) {

}

// create lesson
func CreateLesson(lesson core.Lesson) {

}
