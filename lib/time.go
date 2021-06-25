package lib

import "time"

const (
	day = 24 * time.Hour
)

func LastSunday() time.Time {
	today := time.Now().Truncate(day)
	daysSinceSun := time.Duration(today.Weekday()) * day
	return today.Add(-daysSinceSun)
}
