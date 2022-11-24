package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/facebook-oauth-but-go/modules/entities"
	"github.com/Rayato159/facebook-oauth-but-go/package/utils"
)

type facebookUsecase struct {
	FacebookRepo entities.FacebookRepository
}

func NewFacebookUsecase(facebookRepo entities.FacebookRepository) *facebookUsecase {
	return &facebookUsecase{
		FacebookRepo: facebookRepo,
	}
}

func (u *facebookUsecase) InsertOneUsers(ctx context.Context, req *entities.UsersRegisterReq) (*entities.UsersCredential, error) {
	defer fmt.Printf("return:\t%v.%v\n", entities.FacebookUse, utils.GetFunctionName())
	ctx = context.WithValue(ctx, entities.FacebookUse, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookUse, utils.GetFunctionName())

	res, err := u.FacebookRepo.InsertOneUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
