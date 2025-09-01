package main

import (
	"fmt"
	"log"
	"trainnig-api-poc/api"
	"trainnig-api-poc/api/db"
	userRepo "trainnig-api-poc/api/users/repositories"
)

func main() {
	db, err := db.Connection()
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	userRepo := userRepo.NewRepository(db)

	fmt.Println("User Repository:", userRepo)

	r := api.SetupRouter()
	r.Run(":8080")
}
