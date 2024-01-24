package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	Database "github.com/maadiab/majalisulelm/database"
	"github.com/maadiab/majalisulelm/router"
)

func main() {
	Database.ConnectDB()
	Database.CreateUsersTable()
	r := router.Router()
	fmt.Println("server is running at port: 8080!")
	http.ListenAndServe(":8080", r)

	defer Database.DB.Close()

}
