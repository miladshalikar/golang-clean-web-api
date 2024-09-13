package routers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/miladshalikar/golang-clean-web-api/src/api/handlers"
)

func Health(r fiber.Router) {
	handler := handlers.NewHealthHandler()

	r.Get("/", handler.Health)

}
