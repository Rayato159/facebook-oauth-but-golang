package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/Rayato159/facebook-oauth-but-go/modules/entities"
	"github.com/Rayato159/facebook-oauth-but-go/package/utils"
	"github.com/gofiber/fiber/v2"
)

type facebookController struct {
	FacebookUsecase entities.FacebookUsecase
}

func NewFacebookController(r fiber.Router, facebookUse entities.FacebookUsecase) {
	controller := &facebookController{
		FacebookUsecase: facebookUse,
	}

	// Router -> /facebook
	r.Get("/login", controller.LoginCall)
	r.Get("/callback", controller.LoginCallback)
}

func (fc *facebookController) LoginCall(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.FacebookCon, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())
	defer fmt.Printf("return:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())

	_ = ctx

	return c.Status(fiber.StatusOK).JSON(nil)
}

func (fc *facebookController) LoginCallback(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.FacebookCon, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())
	defer fmt.Printf("return:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())

	_ = ctx

	return c.Status(fiber.StatusOK).JSON(nil)
}
