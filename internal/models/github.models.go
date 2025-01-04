package models

type ProfileStats struct {
	TotalContributions int               `json:"total_contributions"`
	TotalCommits       int               `json:"total_commits"`
	TotalPRs           int               `json:"total_pull_requests"`
	TotalIssues        int               `json:"total_issues"`
	ContributionsByDay []DayContribution `json:"contributions_by_day"`
}

type DayContribution struct {
	Date              string `json:"date"`
	ContributionCount int    `json:"count"`
	Weekday           int    `json:"weekday"`
}

type GraphQLResponse struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions      int `json:"totalCommitContributions"`
				TotalPullRequestContributions int `json:"totalPullRequestContributions"`
				TotalIssueContributions       int `json:"totalIssueContributions"`
				ContributionCalendar          struct {
					TotalContributions int `json:"totalContributions"`
					Weeks              []struct {
						ContributionDays []struct {
							ContributionCount int    `json:"contributionCount"`
							Date              string `json:"date"`
							Weekday           int    `json:"weekday"`
						} `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}
