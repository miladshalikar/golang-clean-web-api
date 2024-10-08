package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/miladshalikar/golang-clean-web-api/src/api/helper"
	"github.com/miladshalikar/golang-clean-web-api/src/api/validations"
	"net/http"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName    string `json:"first_name" validate:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" validate:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" validate:"required,mobile,min=11,max=11"`
}

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helper.GenerateBaseResponse("Test", true, 0))
}

func (h *TestHandler) Users(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(helper.GenerateBaseResponse("Users", true, 0))
}

func (h *TestHandler) UserById(c fiber.Ctx) error {

	id := c.Params("id")
	return c.Status(fiber.StatusOK).JSON(helper.GenerateBaseResponse(fiber.Map{
		"result": "UserById",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) UserByUsername(c fiber.Ctx) error {

	username := c.Params("username")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result":   "UserByUsername",
		"username": username,
	})
}

func (h *TestHandler) Accounts(c fiber.Ctx) error {

	id := c.Params("id")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "Accounts",
		"id":     id,
	})
}

func (h *TestHandler) AddUser(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "AddUser",
	})
}

func (h *TestHandler) HeaderBinder1(c fiber.Ctx) error {

	userId := c.Get("UserId")

	return c.Status(fiber.StatusOK).JSON(helper.GenerateBaseResponse(fiber.Map{
		"result": "HeaderBinder1",
		"userId": userId,
	}, true, 0))
}

func (h *TestHandler) HeaderBinder2(c fiber.Ctx) error {

	var headers header
	c.Bind().Header(&headers)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result":  "HeaderBinder1",
		"headers": headers,
	})

}

func (h *TestHandler) QueryBinder1(c fiber.Ctx) error {

	id := c.Query("id")
	name := c.Query("name")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "HeaderBinder1",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) QueryBinder2(c fiber.Ctx) error {

	args := c.Context().QueryArgs()

	var ids []string

	args.VisitAll(func(key, value []byte) {
		if string(key) == "id" {
			ids = append(ids, string(value))
		}
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "HeaderBinder1",
		"ids ":   ids,
	})
}

func (h *TestHandler) UriBinder(c fiber.Ctx) error {

	id := c.Params("id")
	name := c.Params("name")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}

func (h *TestHandler) BodyBinder(c fiber.Ctx) error {

	validate := validations.GetValidator()

	var p personData

	err := c.Bind().Body(&p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.GenerateBaseResponseWithError("Failed to parse JSON",
			false, -1, err))
	}

	err = validate.Struct(p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
	}

	return c.Status(fiber.StatusOK).JSON(helper.GenerateBaseResponse(fiber.Map{
		"result": "bodyBinder",
		"person": p,
	}, true, 0))
}

func (h *TestHandler) FormBinder(c fiber.Ctx) error {

	var p personData
	c.Bind().Body(&p)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "BodyBinder",
		"person": p,
	})
}

func (h *TestHandler) FileBinder(c fiber.Ctx) error {

	file, _ := c.FormFile("file")
	err := c.SaveFile(file, "file")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "FileBinder",
		"file":   file.Filename,
		"err":    err,
	})
}
