package handler

import (
	"my-realm/src"
	"my-realm/src/constants"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Forwarded-For",
	})

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(constants.Response{
			Data:          time.Now(),
			Message:       "OK",
			PrettyMessage: "The server is healthy.",
			Status:        200,
		})
	})

	src.SetupRoutes(app)

	return adaptor.FiberApp(app)
}
