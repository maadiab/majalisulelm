package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	Database "github.com/maadiab/majalisulelm/database"
)

func main() {
	Database.ConnectDB()
	http.ListenAndServe(":8080", nil)
	fmt.Println("server is running at port: 8080!")
	defer Database.DB.Close()

}
