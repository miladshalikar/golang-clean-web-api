package middlewares

import "github.com/gofiber/fiber/v3"

func TestMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		apiKey := ctx.Get("x-api-key")
		if apiKey == "1" {
			return ctx.Next()
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"result": "Api key is required",
		})
	}
}
