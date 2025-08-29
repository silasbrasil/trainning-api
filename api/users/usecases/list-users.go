package usecases

import "trainnig-api-poc/api/users/entities"

func ListUsers() []entities.User {
	return []entities.User{
		{ID: 1, Name: "Alice", Email: "alice@gmail.com"},
		{ID: 2, Name: "Bob", Email: "bob@gmail.com"},
	}
}
