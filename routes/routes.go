package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/sachinsmc/ministry-of-programming-tasks/controllers"
	"github.com/sachinsmc/ministry-of-programming-tasks/middlewares"
)

func Setup(app *fiber.App) {

	fmt.Println("config setup")
	middlewares.CORS(app)

	app.Get("/", monitor.New())

	app.Use(etag.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "Mon, 02 Jan 2006 15:04:05 MST",
		TimeZone:   "Asia/Dubai",
	}))
	app.Use(pprof.New())

	api := app.Group("/api")
	//api.Use(middlewares.JWT())

	v1 := api.Group("/v1")

	v1.Get("/products", controllers.GetAllProducts)

}
