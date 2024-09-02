// Base package for the short video generator
//
// 1. Calls the website screenshot REST api for the images
// 2. Calls ffmpeg to create the video
// 3. Calls automatic video upload functions

package main

import (
	"os"

	"github.com/Dossified/Dossified-Shorts-Generator/config"
	"github.com/Dossified/Dossified-Shorts-Generator/logging"
	"github.com/Dossified/Dossified-Shorts-Generator/rest"
	"github.com/Dossified/Dossified-Shorts-Generator/screenshot"
	"github.com/Dossified/Dossified-Shorts-Generator/upload/instagram"
	"github.com/Dossified/Dossified-Shorts-Generator/upload/youtube"
	"github.com/Dossified/Dossified-Shorts-Generator/video"
)

func main() {
    // Setting up logger
	logging.InitLogger()
	logging.Info("Dossified Shorts Generator v0.1")

    // Req. argument video mode.
    // e.g. `go run . events` or `go run . news`
	if len(os.Args) <= 1 {
		logging.Error("No video mode defined!")
		os.Exit(1)
	}
	videoMode := os.Args[1]

	switch videoMode {
        case "news":
            // Requesting events from REST api
            trendingNews := rest.RequestNewsTrends()
            // Call Gowitness REST api to take screenshots
            screenshot.ScreenshotTrends(trendingNews, "news")
            break
        case "events":
            // Requesting events from REST api
            trendingEvents := rest.RequestUpcomingEvents()
            // Call Gowitness REST api to take screenshots
            screenshot.ScreenshotEvents(trendingEvents)
            break
        default:
            logging.Error("Unknown video mode defined!")
            os.Exit(2)
	}

    // Cutting & rendering video
	videoPath, videoPathYT := video.CreateVideo(videoMode)

    // YouTube upload
	if config.GetConfiguration().UploadToYouTube {
		youtube.UploadVideo(videoPathYT, videoMode)
	}

    // Instagram upload
	if config.GetConfiguration().UploadToInstagram {
		instagram.UploadToInstagram(videoPath, videoMode)
	}
}
