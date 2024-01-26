package helper

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/maadiab/majalisulelm/core"
)

// create user record

func CreateUser(db *sqlx.DB, user core.User) error {
	_, err := db.Exec("INSERT INTO users (name, mobile, email, password) VALUES ($1,$2,$3,$4)",
		user.Name, user.Mobile, user.Mobile, user.Password)
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
func GetLessonById(db *sqlx.DB, lessonID int) (core.Lesson, error) {
	var lesson core.Lesson

	err := db.Get(&lesson, "SELECT id,name,type,author,duration,time,link,location FROM lessons where id=$1", lessonID)

	if err != nil {

		log.Println("Error getting lesson from db!!!")
	}

	return lesson, err
}

// delete lesson
func DeleteLessonById(lessonId int) {

}

// create lesson
func CreateLesson(db *sqlx.DB, lesson core.Lesson) error {
	_, err := db.Exec("INSERT INTO lessons (name, type, duration, author, link, time, location) VALUES ($1,$2,$3,$4,$5,$6,$7)", lesson.Name, lesson.Type, lesson.Duration, lesson.Author, lesson.Link, lesson.Time, lesson.Location)
	if err != nil {
		log.Println("Error Adding Lesson To Database !")
	}

	return err
}
