package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/ministry-of-programming-tasks/config"
	"github.com/sachinsmc/ministry-of-programming-tasks/routes"
	"github.com/spf13/viper"
)

func Run() {

	config.Init()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":" + viper.GetString("server.port"))

}
