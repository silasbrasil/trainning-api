package usecases

import (
	"trainnig-api-poc/api/users/entities"
	"trainnig-api-poc/api/users/repositories"
)

func CreateUser(name string, email string) (*entities.User, error) {
	newUser := &entities.User{
		Name:  name,
		Email: email,
	}

	userRepo := repositories.NewUserRepository()
	_, err := userRepo.Save(newUser)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
