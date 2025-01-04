package utils

import (
	"fmt"
	"my-realm/internal/models"
	"sort"
	"strings"
)

const neutral = "#171717"
const white = "white"
const gray = "#E5E5E5"

func GenerateLanguagesSVG(languageCount map[string]int, totalRepos int, username, color, background string) string {
	themeColor := ColorSchemes[color]
	if themeColor == "" {
		themeColor = ColorSchemes["red"]
	}

	bgColor := BackgroundSchemes[background]
	if bgColor == "" {
		bgColor = BackgroundSchemes["black"]
	}

	type langData struct {
		Name       string
		Count      int
		Percentage float64
	}

	var languages []langData
	for lang, count := range languageCount {
		percentage := (float64(count) / float64(totalRepos)) * 100
		languages = append(languages, langData{
			Name:       lang,
			Count:      count,
			Percentage: float64(int(percentage*100)) / 100,
		})
	}

	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Percentage > languages[j].Percentage
	})

	svgTemplate := `<?xml version="1.0" encoding="UTF-8"?>
    <svg width="500" height="%d" xmlns="http://www.w3.org/2000/svg">
        <style>
            .title { 
                font: 600 18px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
            }
            .lang-text { 
                font: 400 14px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
                opacity: 0.9;
            }
            .percentage-text { 
                font: 500 14px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
                opacity: 0.9;
            }
            .bar-bg { 
                fill: %s;
            }
            .bar { 
                fill: %s; 
                opacity: 0.8;
            }
        </style>

        <rect 
            x="0" 
            y="0" 
            width="500" 
            height="%d" 
            fill="%s"
            rx="12" 
            ry="12"
            stroke="%s" 
            stroke-width="3"
            stroke-opacity="0.7"
        />
        
        <g transform="translate(30, 35)">
            <text x="0" y="0" class="title">@%s's Languages</text>
            <g transform="translate(0, 30)">
                %s
            </g>
        </g>
    </svg>`

	var languageBars strings.Builder
	barBgColor := neutral
	if background == white {
		barBgColor = gray
	}

	const baseHeight = 100
	height := baseHeight + (len(languages) * 40)

	for i, lang := range languages {
		languageBars.WriteString(fmt.Sprintf(`
            <g transform="translate(0, %d)">
                <text x="0" y="0" class="lang-text">%s</text>
                <text x="440" y="0" class="percentage-text" text-anchor="end">%.1f%%</text>
                <g transform="translate(0, 10)">
                    <rect 
                        x="0" 
                        y="0" 
                        width="440" 
                        height="8" 
                        rx="4" 
                        class="bar-bg"
                    />
                    <rect 
                        x="0" 
                        y="0" 
                        width="%.1f" 
                        height="8" 
                        rx="4" 
                        class="bar"
                    />
                </g>
            </g>
        `, i*40, lang.Name, lang.Percentage, 440*(lang.Percentage/100)))
	}

	return fmt.Sprintf(svgTemplate,
		height,
		themeColor,
		themeColor,
		themeColor,
		barBgColor,
		themeColor,
		height,
		bgColor,
		themeColor,
		username,
		languageBars.String())
}

func GenerateStatsSVG(stats models.ProfileStats, username string, color string, background string) string {
	if color == "" {
		color = "red"
	}
	if background == "" {
		background = "black"
	}

	themeColor := ColorSchemes[color]
	if themeColor == "" {
		themeColor = ColorSchemes["red"]
	}

	bgColor := BackgroundSchemes[background]
	if bgColor == "" {
		bgColor = BackgroundSchemes["black"]
	}

	barBgColor := neutral
	if background == white {
		barBgColor = gray
	}

	maxContributions := 0
	for _, day := range stats.ContributionsByDay {
		if day.ContributionCount > maxContributions {
			maxContributions = day.ContributionCount
		}
	}

	svgTemplate := `<?xml version="1.0" encoding="UTF-8"?>
    <svg width="500" height="400" xmlns="http://www.w3.org/2000/svg">
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
            .contribution-bar { 
                fill: %s; 
                opacity: 0.8;
            }
            .contribution-bar-bg { 
                fill: %s;
            }
            .day-text { 
                font: 400 12px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
                opacity: 0.7;
            }
            .count-text { 
                font: 500 12px 'Inter', 'Segoe UI', Ubuntu, Sans-Serif; 
                fill: %s; 
                opacity: 0.9;
            }
        </style>

        <rect 
            x="0" 
            y="0" 
            width="500" 
            height="400" 
            fill="%s"
            rx="12" 
            ry="12"
            stroke="%s" 
            stroke-width="3"
            stroke-opacity="0.7"
        />
        
        <g transform="translate(30, 35)">
            <text x="0" y="0" class="title">@%s</text>

            <g transform="translate(0, 40)">
                <text x="0" y="0" class="stat-title">Total Contributions</text>
                <text x="160" y="0" class="stat">%d</text>

                <text x="0" y="30" class="stat-title">Total Commits</text>
                <text x="160" y="30" class="stat">%d</text>

                <text x="0" y="60" class="stat-title">Pull Requests</text>
                <text x="160" y="60" class="stat">%d</text>

                <text x="0" y="90" class="stat-title">Issues</text>
                <text x="160" y="90" class="stat">%d</text>
            </g>

            <g transform="translate(0, 170)">
                <text x="0" y="0" class="stat-title">Last 7 Days</text>
                <g transform="translate(0, 20)">
                    %s
                </g>
            </g>
        </g>
    </svg>`

	var contributionBars strings.Builder
	lastSevenDays := stats.ContributionsByDay
	if len(lastSevenDays) > 7 {
		lastSevenDays = lastSevenDays[len(lastSevenDays)-7:]
	}

	for i, day := range lastSevenDays {
		percentage := float64(day.ContributionCount) / float64(maxContributions)
		if maxContributions == 0 {
			percentage = 0
		}
		barHeight := percentage * 100
		if barHeight == 0 {
			barHeight = 1
		}

		contributionBars.WriteString(fmt.Sprintf(`
            <g transform="translate(%d, 0)">
                <rect 
                    x="0" 
                    y="0" 
                    width="50" 
                    height="100" 
                    class="contribution-bar-bg"
                    rx="6"
                    ry="6"
                />
                <rect 
                    x="0" 
                    y="%f" 
                    width="50" 
                    height="%f" 
                    class="contribution-bar"
                    rx="6"
                    ry="6"
                />
                <text 
                    x="25" 
                    y="125" 
                    class="count-text" 
                    text-anchor="middle"
                >%d</text>
                <text 
                    x="25" 
                    y="145" 
                    class="day-text" 
                    text-anchor="middle"
                >%s</text>
            </g>
        `, i*60,
			100-barHeight,
			barHeight,
			day.ContributionCount,
			getDayName(day.Weekday)))
	}

	return fmt.Sprintf(svgTemplate,
		themeColor,
		themeColor,
		themeColor,
		themeColor,
		barBgColor,
		themeColor,
		themeColor,
		bgColor,
		themeColor,
		username,
		stats.TotalContributions,
		stats.TotalCommits,
		stats.TotalPRs,
		stats.TotalIssues,
		contributionBars.String())
}

func getDayName(weekday int) string {
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	return days[weekday]
}
