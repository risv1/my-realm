# My Realm

Get stats data from Github and Leetcode.

## Routes

- `/api/languages`: Query params are username
- `/api/languages/svg`: Query params are username, color, background
- `/api/stats`: Query params are username
- `/api/stats/svg`: Query params are username, color, background
- `/api/leetcode`: Query params are username
- `/api/leetcode/svg`: Query params are username, color, background

### Options

```go
var (
 ColorSchemes = map[string]string{
  "red":       "rgb(243, 69, 69)",
  "blue":      "rgb(59, 130, 246)",
  "lightBlue": "rgb(14, 165, 233)",
  "green":     "rgb(34, 197, 94)",
  "yellow":    "rgb(234, 179, 8)",
  "orange":    "rgb(249, 115, 22)",
  "purple":    "rgb(168, 85, 247)",
  "pink":      "rgb(236, 72, 153)",
  "white":     "rgb(255, 255, 255)",
  "black":     "rgb(0, 0, 0)",
 }

 BackgroundSchemes = map[string]string{
  "black":   "#0A0A0A",
  "neutral": "#171717",
  "white":   "#FFFFFF",
 }
)
```
