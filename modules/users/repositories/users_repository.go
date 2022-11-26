package repositories

import "github.com/Rayato159/facebook-oauth-but-go/modules/entities"

type usersRepository struct{}

func NewUsersRepository() entities.UsersRepository {
	return &usersRepository{}
}
