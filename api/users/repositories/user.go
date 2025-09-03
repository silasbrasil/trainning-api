package repositories

import (
	"fmt"
	"trainnig-api-poc/api/database"
	"trainnig-api-poc/api/users/entities"
	"trainnig-api-poc/api/users/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) List() []entities.User {
	db := database.GetDb()
	var userModels []models.User
	var userEntities []entities.User

	db.Find(&userModels)

	fmt.Println("Aqui")

	for _, user := range userModels {
		userEntities = append(userEntities, entities.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userEntities
}

func (r *UserRepository) GetByID(id uint) *entities.User {
	db := database.GetDb()
	var user entities.User

	if err := db.First(&user, id).Error; err != nil {
		return nil
	}

	return &user
}

func (r *UserRepository) Save(user *entities.User) (*entities.User, error) {
	db := database.GetDb()
	userModel := models.User{
		Name:  user.Name,
		Email: user.Email,
	}

	if err := db.Create(&userModel).Error; err != nil {
		return nil, err
	}

	user.ID = userModel.ID
	return user, nil
}
