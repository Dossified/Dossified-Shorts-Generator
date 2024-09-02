// Helper to generate video information like title, description, ...
package video

import (
	"fmt"
	"time"

	"github.com/Dossified/Dossified-Shorts-Generator/logging"
)

// Generates video title based on type of video & current date
func GetVideoTitle(videoMode string) string {
	var title string
	switch videoMode {
	case "events":
		title = "Crypto Events Week " +
			getDateString(0) +
			" #airdrops #events #crypto #blockchain #Dossified #shorts"
		break
	case "news":
		title = "Crypto News Trends " +
			getDateString(0) +
			" #blockchain #crypto #Dossified #shorts"
	}
	logging.Debug("Video title: " + title)
	return title
}

// Generates video description based on type of video & current date
func GetVideoDescription(videoMode string) string {
	var description string
	switch videoMode {
	case "events":
		description = "Crypto events week " + getDateString(0) + " - " + getDateString(7) + " #airdrops #events #crypto #blockchain #Dossified #shorts"
		break
	case "news":
		description = "Crypto news trends " + getDateString(0) + " - " + getDateString(7) + " #blockchain #crypto #Dossified #shorts"
		break
	}
	logging.Debug("Video description: " + description)
	return description
}

// Generates a date string, which is the current day + month + year
func getDateString(daysToAdd int) string {
	day := time.Now().AddDate(0, 0, daysToAdd).Day()
	dayString := fmt.Sprint(day)
	switch day {
	case 1, 21, 31:
		dayString += "st"
		break
	case 2, 22:
		dayString += "nd"
		break
	case 3, 23:
		dayString += "rd"
		break
	default:
		dayString += "th"
		break
	}
	return dayString + " " + fmt.Sprint(time.Now().AddDate(0, 0, daysToAdd).Month()) + " " + fmt.Sprint(time.Now().AddDate(0, 0, daysToAdd).Year())
}
