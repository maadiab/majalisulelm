package Database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/maadiab/majalisulelm/core"
	"github.com/maadiab/majalisulelm/helper"
)

const (
	Host     = "127.0.0.1"
	User     = "postgres"
	DbName   = "postgres"
	Password = "passwd"
)

// change it to init function later
func ConnectDB() {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", Host, User, Password, DbName)

	db, err := sqlx.Open("postgres", connStr)

	// open connection
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected sucsessfully to Database !!")

	// create table if not exist

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR (255),
		mobile INT,
		email VARCHAR (255),
		password VARCHAR (255)
		)
		`)

	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println("users table created !!")
	}
	// test get user

	userID := 1

	user, err := helper.GetUser(db, userID)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(user)

	// test add user

	firstUser := core.User{

		Name:     "mohanad",
		Mobile:   550795131,
		Email:    "mohanad_diab@live.com",
		Password: "Aa123",
	}

	err = helper.CreateUser(db, firstUser)
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println("user added successfully !!")
	}

}
