package main

import (
	"go-api/controller/api"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func apiPath(c *fiber.Ctx) error {
	return c.SendString("Hello, World! API")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api", apiPath)
	app.Post("/api/plus", api.GetPlus)
	app.Post("/api/minus", api.GetMinus)
	app.Post("/api/multiple", api.GetMultiple)
	app.Post("/api/divide", api.GetDivide)
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":3000")
}
