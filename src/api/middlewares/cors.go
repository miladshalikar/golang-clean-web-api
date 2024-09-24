package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/miladshalikar/golang-clean-web-api/src/config"
)

func Cors(cfg *config.Config) fiber.Handler {
	return func(ctx fiber.Ctx) error {

		ctx.Set("Access-Control-Allow-Origin", cfg.Cors.AllowOrigins)
		ctx.Set("Access-Control-Allow-Credentials", "true")
		ctx.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Set("Access-Control-Max-Age", "21600")
		ctx.Set("content-type", "application/json")

		if ctx.Method() == fiber.MethodOptions {
			return ctx.SendStatus(204)
		}

		return ctx.Next()
	}
}
