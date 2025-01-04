package api

import (
	"my-realm/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/languages", controllers.GetMostUsedLanguages)
	app.Get("/api/languages/svg", controllers.GetLanguagesAsSVG)
	app.Get("/api/stats", controllers.GetProfileStats)
	app.Get("/api/stats/svg", controllers.GetStatsAsSVG)

	app.Get("/api/leetcode", controllers.GetLeetCodeStats)
	app.Get("/api/leetcode/svg", controllers.GetLeetCodeStatsAsSVG)
}
