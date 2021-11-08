package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(&Response{Message: "List of all products"})
}
