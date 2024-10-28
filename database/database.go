package Database

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	Host     = "127.0.0.1"
	User     = "postgres"
	DbName   = "postgres"
	Password = "postgres"
)

var DB *sqlx.DB

// change it to init function later
func ConnectDB(ctx context.Context) (*sqlx.DB, error) {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", Host, User, Password, DbName)

	// sqlx.Open("postgres", connStr)

	db, err := sqlx.ConnectContext(ctx, "postgres", connStr)

	// open connection
	if err != nil {
		return nil, fmt.Errorf("Error Openning connection %w", err)
	}

	// test connection
	err = db.Ping()
	if err != nil {
	
		return nil, err
	}

	log.Println("connected sucsessfully to Database ...")

	DB = db

	return db, err
	// create table if not exist

	// test get user

	// userID := 1

	// user, err := helper.GetUser(db, userID)

	// if err != nil {
	// 	log.Fatal(err)

	// }

	// fmt.Println(user)

	// test add user

	// firstUser := core.User{

	// 	Name:     "mohanad",
	// 	Mobile:   550795131,
	// 	Email:    "mohanad_diab@live.com",
	// 	Password: "Aa123",
	// }

	// err = helper.CreateUser(db, firstUser)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {

	// 	fmt.Println("user added successfully !!")
	// }

}

func CreateUsersTable() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR (255),
		mobile INT,
		email VARCHAR (255),
		password VARCHAR (255)
		)
		`)

	if err != nil {

		log.Println("Error Creating User Table In Database!!!")
	} else {

		log.Println("Creating Users Table If Not Exists ...")
	}
}

func CreateLessonsTable() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS lessons (
		id SERIAL PRIMARY KEY,
		name VARCHAR (255),
		type VARCHAR (255),
		duration VARCHAR (255),
		author VARCHAR (255),
		link VARCHAR (255),
		time VARCHAR (255),
		location VARCHAR (255)
	)`)

	if err != nil {

		log.Println("Error Creating Lessons Table In Database!!!")
	} else {

		log.Println("Creating Lessons Table If Not Exists ...")
	}
}
