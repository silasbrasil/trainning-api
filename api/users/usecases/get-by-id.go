package usecases

import "trainnig-api-poc/api/users/entities"

func GetUserById(id uint64) entities.User {
	return entities.User{
		ID:    id,
		Name:  "Alice",
		Email: "alice@gmail.com",
	}
}
