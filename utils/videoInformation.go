package utils

import (
	"fmt"
	"time"
)

func GetVideoTitle(videoMode string) string {
    var title string
    switch videoMode {
        case "events":
            title = "Crypto Events week " + fmt.Sprint(getCurrentWeekNumber()) + " #shorts #blockchain #crypto #events"
            break;
        case "news":
            title = "Crypto News Trends week " + fmt.Sprint(getCurrentWeekNumber()) + " #shorts #blockchain #crypto #events"
    }
	return title
}

func GetVideoDescription(videoMode string) string {
    var description string
    switch videoMode {
        case "events":
            description = "Upcoming events provided by Dossified.com for week " + fmt.Sprint(
                getCurrentWeekNumber(),
            ) + " " + fmt.Sprint(time.Now().Year()) + " #shorts"
            break;
        case "news":
            description = "Trending news provided by Dossified.com for week " + fmt.Sprint(
                getCurrentWeekNumber(),
            ) + " " + fmt.Sprint(time.Now().Year()) + " #shorts"
            break;
    }
	return description
}

func getCurrentWeekNumber() int {
	timeNow := time.Now()
	_, weekNumber := timeNow.ISOWeek()
	return weekNumber
}
