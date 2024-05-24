package utils

import (
    "fmt"
    "time"
)

func GetVideoTitle() string {
	return "Crypto News week " + fmt.Sprint(getCurrentWeekNumber()) + " #shorts"
}

func GetVideoDescription() string {
	return "News provided by Dossified.com for week " + fmt.Sprint(
		getCurrentWeekNumber(),
	) + " " + fmt.Sprint(time.Now().Year()) + " #shorts"
}

func getCurrentWeekNumber() int {
	timeNow := time.Now()
	_, weekNumber := timeNow.ISOWeek()
	return weekNumber
}
