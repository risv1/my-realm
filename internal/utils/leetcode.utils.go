package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"my-realm/internal/models"
	"net/http"
	"time"
)

func FetchLeetCodeStats(username string) (*models.LeetCodeStats, error) {
	query := `
    query userSessionProgress($username: String!) {
        allQuestionsCount {
            difficulty
            count
        }
        matchedUser(username: $username) {
            profile {
                ranking
            }
            submitStats {
                acSubmissionNum {
                    difficulty
                    count
                    submissions
                }
                totalSubmissionNum {
                    difficulty
                    count
                    submissions
                }
            }
        }
    }`

	requestBody := map[string]interface{}{
		"query": query,
		"variables": map[string]string{
			"username": username,
		},
		"operationName": "userSessionProgress",
	}

	jsonValue, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data struct {
			AllQuestionsCount []struct {
				Difficulty string `json:"difficulty"`
				Count      int    `json:"count"`
			} `json:"allQuestionsCount"`
			MatchedUser struct {
				Profile struct {
					Ranking int `json:"ranking"`
				} `json:"profile"`
				SubmitStats struct {
					AcSubmissionNum []struct {
						Difficulty  string `json:"difficulty"`
						Count       int    `json:"count"`
						Submissions int    `json:"submissions"`
					} `json:"acSubmissionNum"`
					TotalSubmissionNum []struct {
						Difficulty  string `json:"difficulty"`
						Count       int    `json:"count"`
						Submissions int    `json:"submissions"`
					} `json:"totalSubmissionNum"`
				} `json:"submitStats"`
			} `json:"matchedUser"`
		} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("leetcode API error: %s", result.Errors[0].Message)
	}

	stats := &models.LeetCodeStats{}

	for _, qCount := range result.Data.AllQuestionsCount {
		stats.TotalQuestions += qCount.Count
	}

	var totalSubmissions, acceptedSubmissions int
	for _, submission := range result.Data.MatchedUser.SubmitStats.AcSubmissionNum {
		switch submission.Difficulty {
		case "Easy":
			stats.EasySolved = submission.Count
		case "Medium":
			stats.MediumSolved = submission.Count
		case "Hard":
			stats.HardSolved = submission.Count
		}
		acceptedSubmissions += submission.Submissions
	}

	for _, submission := range result.Data.MatchedUser.SubmitStats.TotalSubmissionNum {
		totalSubmissions += submission.Submissions
	}

	stats.TotalSolved = stats.EasySolved + stats.MediumSolved + stats.HardSolved
	stats.Ranking = result.Data.MatchedUser.Profile.Ranking

	if totalSubmissions > 0 {
		stats.AcceptanceRate = float64(acceptedSubmissions) / float64(totalSubmissions) * 100
	}

	return stats, nil
}

func GenerateLeetCodeStatsSVG(stats *models.LeetCodeStats, username, color, background string) string {
	themeColor := ColorSchemes[color]
	if themeColor == "" {
		themeColor = ColorSchemes["red"]
	}

	bgColor := BackgroundSchemes[background]
	if bgColor == "" {
		bgColor = BackgroundSchemes["black"]
	}

	easyPercent := float64(stats.EasySolved) / float64(stats.TotalQuestions) * 100
	mediumPercent := float64(stats.MediumSolved) / float64(stats.TotalQuestions) * 100
	hardPercent := float64(stats.HardSolved) / float64(stats.TotalQuestions) * 100
	totalPercent := float64(stats.TotalSolved) / float64(stats.TotalQuestions) * 100

	svgTemplate := `<?xml version="1.0" encoding="UTF-8"?>
    <svg width="500" height="450" xmlns="http://www.w3.org/2000/svg">
        <style>
            .title { 
                font: 600 18px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
            }
            .stat { 
                font: 500 14px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
                opacity: 0.9;
            }
            .stat-title { 
                font: 400 14px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
                opacity: 0.8;
            }
            .rank { 
                font: 700 24px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
                opacity: 0.9;
            }
            .progress-bar-bg { 
                fill: %s;
                opacity: 0.2;
            }
            .progress-bar { 
                fill: %s; 
                opacity: 0.8;
            }
            .easy { fill: rgb(0, 184, 163); }
            .medium { fill: rgb(255, 192, 30); }
            .hard { fill: rgb(255, 55, 95); }
        </style>

        <rect 
            x="0" 
            y="0" 
            width="500" 
            height="450" 
            fill="%s"
            rx="12" 
            ry="12"
            stroke="%s" 
            stroke-width="3"
            stroke-opacity="0.7"
        />
        
        <g transform="translate(25, 35)">
            <text x="0" y="0" class="title">@%s's LeetCode Stats</text>

            <g transform="translate(0, 55)">
                <text class="stat-title">Ranking</text>
                <text x="0" y="25" class="rank">#%d</text>
            </g>

            <g transform="translate(0, 120)">
                <text class="stat-title">Total Progress</text>
                <text x="430" y="0" class="stat" text-anchor="end">%d / %d</text>
                <rect x="0" y="10" width="440" height="10" rx="5" class="progress-bar-bg"/>
                <rect x="0" y="10" width="%.1f" height="10" rx="5" class="progress-bar"/>
            </g>

            <g transform="translate(0, 175)">
                <text class="stat-title">Problems Solved</text>
            
                <g transform="translate(0, 30)">
                    <text class="stat-title" style="fill: rgb(0, 184, 163)">Easy</text>
                    <text x="430" y="0" class="stat" text-anchor="end">%d</text>
                    <rect x="0" y="10" width="440" height="8" rx="4" class="progress-bar-bg"/>
                    <rect x="0" y="10" width="%.1f" height="8" rx="4" style="fill: rgb(0, 184, 163); opacity: 0.8;"/>
                </g>

                <g transform="translate(0, 75)">
                    <text class="stat-title" style="fill: rgb(255, 192, 30)">Medium</text>
                    <text x="430" y="0" class="stat" text-anchor="end">%d</text>
                    <rect x="0" y="10" width="440" height="8" rx="4" class="progress-bar-bg"/>
                    <rect x="0" y="10" width="%.1f" height="8" rx="4" style="fill: rgb(255, 192, 30); opacity: 0.8;"/>
                </g>

                <g transform="translate(0, 120)">
                    <text class="stat-title" style="fill: rgb(255, 55, 95)">Hard</text>
                    <text x="430" y="0" class="stat" text-anchor="end">%d</text>
                    <rect x="0" y="10" width="440" height="8" rx="4" class="progress-bar-bg"/>
                    <rect x="0" y="10" width="%.1f" height="8" rx="4" style="fill: rgb(255, 55, 95); opacity: 0.8;"/>
                </g>
            </g>

            <g transform="translate(0, 350)">
                <text class="stat-title">Acceptance Rate</text>
                <text x="430" y="0" class="stat" text-anchor="end">%.1f%%</text>
            </g>
        </g>
    </svg>`

	barBgColor := "#171717"
	if background == "white" {
		barBgColor = "#E5E5E5"
	}

	return fmt.Sprintf(svgTemplate,
		themeColor,
		themeColor,
		themeColor,
		themeColor,
		barBgColor,
		themeColor,
		bgColor,
		themeColor,
		username,
		stats.Ranking,
		stats.TotalSolved,
		stats.TotalQuestions,
		440*totalPercent/100,
		stats.EasySolved,
		440*easyPercent/100,
		stats.MediumSolved,
		440*mediumPercent/100,
		stats.HardSolved,
		440*hardPercent/100,
		stats.AcceptanceRate)
}
