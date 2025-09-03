package usecases

import (
	"trainnig-api-poc/api/users/entities"
	"trainnig-api-poc/api/users/repositories"
)

func ListUsers() []entities.User {
	userRepo := repositories.NewUserRepository()

	users := userRepo.List()

	return users
}
