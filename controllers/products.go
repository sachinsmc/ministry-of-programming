package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
}

func GetAllProducts(c *fiber.Ctx) error {
	fmt.Println("getting all products")
	return c.JSON(&Response{Message: "List of all products"})
}
