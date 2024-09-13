package api

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/miladshalikar/golang-clean-web-api/src/api/routers"
	"github.com/miladshalikar/golang-clean-web-api/src/config"
)

func InitServer() {

	cfg := config.GetConfig()

	r := fiber.New()
	r.Use(logger.New())
	r.Use(recover.New())

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	r.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
