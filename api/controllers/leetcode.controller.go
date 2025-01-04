package controllers

import (
	"my-realm/api/constants"
	"my-realm/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func GetLeetCodeStats(c *fiber.Ctx) error {
	username := c.Query("username")
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username is required",
		})
	}

	stats, err := utils.FetchLeetCodeStats(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}

	response := constants.Response{
		Message:       "OK",
		PrettyMessage: "Successfully retrieved LeetCode statistics",
		Status:        200,
		Data:          stats,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func GetLeetCodeStatsAsSVG(c *fiber.Ctx) error {
	username := c.Query("username")
	color := c.Query("color", "red")
	background := c.Query("background", "black")

	stats, err := utils.FetchLeetCodeStats(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}

	svg := utils.GenerateLeetCodeStatsSVG(stats, username, color, background)

	c.Set("Content-Type", "image/svg+xml")
	return c.SendString(svg)
}
