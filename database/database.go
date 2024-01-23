package Database

import (
	"database/sql"
	"fmt"
)

const (
	Host     = "127.0.0.1"
	User     = "postgres"
	DbName   = "postgres"
	Password = "passwd"
)

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", Host, User, Password, DbName)
	db, err := sql.Open("postgres", dsn)
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

}
