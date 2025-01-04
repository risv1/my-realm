package models

import "time"

type LeetCodeCache struct {
	Stats     *LeetCodeStats `json:"stats"`
	Timestamp time.Time      `json:"timestamp"`
}

type LeetCodeStats struct {
	TotalSolved        int     `json:"totalSolved"`
	TotalQuestions     int     `json:"totalQuestions"`
	EasySolved         int     `json:"easySolved"`
	MediumSolved       int     `json:"mediumSolved"`
	HardSolved         int     `json:"hardSolved"`
	AcceptanceRate     float64 `json:"acceptanceRate"`
	Ranking            int     `json:"ranking"`
	ContributionPoints int     `json:"contributionPoints"`
}

type LeetCodeSubmissionStats struct {
	LastSubmissions []LeetCodeSubmission `json:"lastSubmissions"`
}

type LeetCodeSubmission struct {
	Title      string `json:"title"`
	Timestamp  string `json:"timestamp"`
	Status     string `json:"status"`
	Language   string `json:"language"`
	Difficulty string `json:"difficulty"`
}

type LeetCodeLanguageStats struct {
	Languages        map[string]int `json:"languages"`
	TotalSubmissions int            `json:"totalSubmissions"`
}

type LeetCodeResponse struct {
	Data struct {
		MatchedUser struct {
			Username    string `json:"username"`
			SubmitStats struct {
				AcSubmissionNum []struct {
					Count      int    `json:"count"`
					Difficulty string `json:"difficulty"`
				} `json:"acSubmissionNum"`
				TotalSubmissionNum []struct {
					Count      int    `json:"count"`
					Difficulty string `json:"difficulty"`
				} `json:"totalSubmissionNum"`
			} `json:"submitStats"`
			Profile struct {
				Ranking int `json:"ranking"`
				Stars   int `json:"starRating"`
			} `json:"profile"`
			SubmissionCalendar string `json:"submissionCalendar"`
		} `json:"matchedUser"`
		AllQuestionsCount []struct {
			Difficulty string `json:"difficulty"`
			Count      int    `json:"count"`
		} `json:"allQuestionsCount"`
	} `json:"data"`
}
