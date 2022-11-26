package usecases

import "github.com/Rayato159/facebook-oauth-but-go/modules/entities"

type usersUsecase struct{}

func NewUsersUsecase() entities.UsersUsecase {
	return &usersUsecase{}
}
