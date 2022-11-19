package servers

import (
	"fmt"
	"log"

	"github.com/Rayato159/facebook-oauth-but-go/configs"
	"github.com/Rayato159/facebook-oauth-but-go/modules/entities"
	"github.com/Rayato159/facebook-oauth-but-go/package/middlewares"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	Cfg *configs.Config
	App *fiber.App
	Db  *mongo.Client
}

func NewSever(cfg *configs.Config, db *mongo.Client) *server {
	fiberConfigs := configs.NewFiberConfig(cfg.App)
	return &server{
		Cfg: cfg,
		App: fiber.New(fiberConfigs),
		Db:  db,
	}
}

func (s *server) Start() {
	if err := s.mapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	fiberConnURL := fmt.Sprintf("%s:%s", s.Cfg.App.Host, s.Cfg.App.Port)
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}

func (s *server) mapHandlers() error {
	middlewares.NewCorsFiberHandler(s.App)

	// Group a version
	v1 := s.App.Group("/v1")
	_ = v1

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		log.Println("error, endpoint is not found")
		return c.Status(fiber.StatusNotFound).JSON(entities.ErrResponse{
			Status:     fiber.ErrNotFound.Message,
			StatusCode: fiber.StatusNotFound,
			Message:    "error, endpoint is not found",
		})
	})
	return nil
}
