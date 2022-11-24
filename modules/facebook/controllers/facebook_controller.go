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
	r.Post("/users", controller.InsertOneUsers)
}

func (fc *facebookController) LoginCall(c *fiber.Ctx) error {
	defer fmt.Printf("return:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())
	ctx := context.WithValue(c.Context(), entities.FacebookCon, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())

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

func (fc *facebookController) InsertOneUsers(c *fiber.Ctx) error {
	ctx := context.WithValue(c.Context(), entities.FacebookCon, time.Now().UnixMicro())
	fmt.Printf("called:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())
	defer fmt.Printf("return:\t%v.%v\n", entities.FacebookCon, utils.GetFunctionName())

	req := new(entities.UsersRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.Response{
			Status:     fiber.ErrBadRequest.Message,
			StatusCode: fiber.StatusBadRequest,
			Message:    "error, can't parse a request body into the struct",
			Result: entities.Result{
				Data: nil,
			},
		})
	}
	res, err := fc.FacebookUsecase.InsertOneUsers(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Response{
			Status:     fiber.ErrInternalServerError.Message,
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
			Result: entities.Result{
				Data: nil,
			},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entities.Response{
		Status:     "Created",
		StatusCode: fiber.StatusCreated,
		Message:    "",
		Result: entities.Result{
			Data: res,
		},
	})
}
