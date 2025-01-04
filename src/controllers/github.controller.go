package controllers

import (
	"encoding/json"
	"fmt"
	"my-realm/internal/config"
	"my-realm/internal/utils"
	"my-realm/src/constants"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Repository struct {
	Language string `json:"language"`
}

func GetMostUsedLanguages(c *fiber.Ctx) error {
	username := c.Query("username", "risv1")
	githubAPIURL := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := http.Get(githubAPIURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}

	languageCount := make(map[string]int)
	totalRepos := 0

	for _, repo := range repos {
		if repo.Language != "" {
			languageCount[repo.Language]++
			totalRepos++
		}
	}

	languagePercentages := make(map[string]float64)
	for lang, count := range languageCount {
		percentage := (float64(count) / float64(totalRepos)) * 100
		languagePercentages[lang] = float64(int(percentage*100)) / 100
	}

	response := constants.Response{
		Message:       "OK",
		PrettyMessage: "Successfully retrieved language usage statistics",
		Status:        200,
		Data:          languagePercentages,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

var (
	svgCache    string
	cacheTTL    = 1 * time.Second
	lastUpdated time.Time
	cacheMutex  sync.RWMutex
)

func GetProfileStats(c *fiber.Ctx) error {
	env := config.LoadEnv()
	username := c.Query("username", "risv1")
	token := env.GithubToken

	stats, err := utils.FetchGitHubStats(username, token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}

	response := constants.Response{
		Message:       "OK",
		PrettyMessage: "Successfully retrieved profile statistics",
		Status:        200,
		Data:          stats,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func GetLanguagesAsSVG(c *fiber.Ctx) error {
	username := c.Query("username", "risv1")
	color := c.Query("color", "red")
	background := c.Query("background", "black")

	githubAPIURL := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	resp, err := http.Get(githubAPIURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}

	languageCount := make(map[string]int)
	totalRepos := 0

	for _, repo := range repos {
		if repo.Language != "" {
			languageCount[repo.Language]++
			totalRepos++
		}
	}

	svg := utils.GenerateLanguagesSVG(languageCount, totalRepos, username, color, background)

	c.Set("Content-Type", "image/svg+xml")
	return c.SendString(svg)
}

func GetStatsAsSVG(c *fiber.Ctx) error {
	cacheMutex.RLock()
	if time.Since(lastUpdated) < cacheTTL && svgCache != "" {
		defer cacheMutex.RUnlock()
		c.Set("Content-Type", "image/svg+xml")
		return c.SendString(svgCache)
	}
	cacheMutex.RUnlock()

	env := config.LoadEnv()
	username := c.Query("username", "risv1")
	token := env.GithubToken
	color := c.Query("color", "red")
	background := c.Query("background", "black")

	stats, err := utils.FetchGitHubStats(username, token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(constants.ErrorInternalServerError)
	}

	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	svg := utils.GenerateStatsSVG(stats, username, color, background)
	svgCache = svg
	lastUpdated = time.Now()

	c.Set("Content-Type", "image/svg+xml")
	return c.SendString(svg)
}
