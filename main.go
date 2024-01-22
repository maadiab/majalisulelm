package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	Database "github.com/maadiab/majalisulelm/database"
)

func main() {

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", Database.Host, Database.User, Database.Password, Database.DbName)
	db, err := sql.Open("postgres", connString)
	//open connection
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//test connection

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("connected sucsessfully to Database !!")
}
