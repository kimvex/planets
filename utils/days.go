package utils

import (
	"time"
)

//YearsSinceToday function for get days of a ten years since now
func YearsSinceToday(years int) int {
	now := time.Now()
	TenYears := now.AddDate(years, 0, 0)
	countOfDays := TenYears.Sub(now).Hours() / 24
	return int(countOfDays)
}

//YearBaseDayOfPlanets function for get days of orbit planets
func YearBaseDayOfPlanets(years int) int {
	return years * 360
}
