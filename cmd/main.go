package main

import (
	"my-realm/api/constants"
	"my-realm/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.LoadEnv()

	app := fiber.New()

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(constants.Response{
			Data:          time.Now(),
			Message:       "OK",
			PrettyMessage: "The server is healthy.",
			Status:        200,
		})
	})

	if err := app.Listen(":" + envConfig.Port); err != nil {
		panic(err)
	}
}
