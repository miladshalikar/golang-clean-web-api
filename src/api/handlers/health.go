package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/miladshalikar/golang-clean-web-api/src/api/helper"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helper.GenerateBaseResponse("working!", true, 0))
}
