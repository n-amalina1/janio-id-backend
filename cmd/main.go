package main

import (
	"database/sql"

	routes "janio-id-backend/routes"
)

var db *sql.DB

func main() {
	routes.SetupRoutes(db)

}
