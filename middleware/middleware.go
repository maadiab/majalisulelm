package middleware

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var Authorized = false

func CheckUser(db *sqlx.DB, user string, password string) bool {

	_, err := db.Exec("SELECT name, password FROM users where name = $1 and password = $2", user, password)

	if err != nil {
		log.Println("Please Check Username and Password !!!")
	} else {

		Authorized = true
	}

	return Authorized
}
