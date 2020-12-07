package games

import (
	"time"
	"strings"
)

func Handler(choice string, cTime string) {
	t := time.Now()
	t.Format("Mon Jan _2 15:04:05 2006")

	// Define time
	weekday := strings.ToLower(t.Weekday().String())

	switch choice {
	case "daily":
		daily := new(DAILY_INFO)
		daily.GetTime(weekday)
		daily.Run(cTime)
	}
}
