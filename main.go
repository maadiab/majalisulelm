package main

import (
	_ "github.com/lib/pq"
	Database "github.com/maadiab/majalisulelm/database"
)

func main() {
	Database.ConnectDB()

}
