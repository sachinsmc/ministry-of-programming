package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/ministry-of-programming-tasks/config"
	"github.com/sachinsmc/ministry-of-programming-tasks/routes"
	"github.com/spf13/viper"
)

func Run() {

	config.Init()

	app := fiber.New()

	fmt.Println("Server Port : ", viper.GetString("server.port"))

	routes.Setup(app)

	app.Listen(":3003")

}
