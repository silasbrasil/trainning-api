package usecases

import (
	"trainnig-api-poc/api/users/entities"
	"trainnig-api-poc/api/users/repositories"
)

func GetUserById(id uint64) *entities.User {
	var user *entities.User

	userRepo := repositories.NewUserRepository()
	user = userRepo.GetById(id)

	return user
}
