package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	Database "github.com/maadiab/majalisulelm/database"
	"github.com/maadiab/majalisulelm/router"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := Database.ConnectDB(ctx)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Connection Time Out !!!")
			return
		} else if ctx.Err() == context.Canceled {
			log.Println("Connection Cancelled !!!")
			return
		} else {
			log.Println("Error Connecting Database !!!")
			return
		}
	}

	Database.CreateUsersTable()
	Database.CreateLessonsTable()
	r := router.Router()
	fmt.Println("server is running at port: 8080 ...")

	http.ListenAndServe(":8080", r)

	defer Database.DB.Close()

}
