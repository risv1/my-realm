package main

import (
	"my-realm/src"
	"my-realm/src/constants"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(constants.Response{
			Data:          time.Now(),
			Message:       "OK",
			PrettyMessage: "The server is healthy.",
			Status:        200,
		})
	})

	src.SetupRoutes(app)

	if err := app.Listen(":8000"); err != nil {
		panic(err)
	}
}
