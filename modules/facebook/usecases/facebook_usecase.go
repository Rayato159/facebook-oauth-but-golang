package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/facebook-oauth-but-go/modules/entities"
	"github.com/Rayato159/facebook-oauth-but-go/package/utils"
)

type facebookUsecase struct {
	FacebookApi entities.FacebookApi
}

func NewFacebookUsecase(facebookApi entities.FacebookApi) *facebookUsecase {
	return &facebookUsecase{
		FacebookApi: facebookApi,
	}
}

func (u *facebookUsecase) GetAccessToken(ctx context.Context) error {
	defer fmt.Printf("return:\t%v.%v\n", entities.FacebookUse, utils.GetFunctionName())
	ctx = context.WithValue(ctx, entities.FacebookUse, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookUse, utils.GetFunctionName())

	return nil
}
