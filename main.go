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

type PostRequestBody struct {
	Url     string   `json:"url"`
	Oneshot bool     `json:"oneshot"`
	Headers []string `json:"headers"`
}

func main() {
	logging.InitLogger()
	logging.Info("Dossified Shorts Generator v0.1")
	if len(os.Args) <= 1 {
		logging.Error("No video mode defined!")
		os.Exit(1)
	}
	videoMode := os.Args[1]
	switch videoMode {
	case "news":
		trendingNews := rest.RequestNewsTrends()
		screenshot.ScreenshotTrends(trendingNews, "news")
		break
	case "events":
		trendingEvents := rest.RequestUpcomingEvents()
		screenshot.ScreenshotEvents(trendingEvents)
		break
	case "coins":
		// ToDo
		break
	default:
		logging.Error("Unknown video mode defined!")
		os.Exit(2)
	}
	videoPath, videoPathYT := video.CreateVideo(videoMode)

	if config.GetConfiguration().UploadToYouTube {
		youtube.UploadVideo(videoPathYT, videoMode)
	}
	if config.GetConfiguration().UploadToInstagram {
		instagram.UploadToInstagram(videoPath, videoMode)
	}
}
