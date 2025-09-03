package main

import (
	"log"
	"trainnig-api-poc/api"
	"trainnig-api-poc/api/database"
)

func main() {
	db, err := database.Init()
	database.RunMigrations(db)

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	r := api.SetupRouter()
	r.Run(":8080")
}
