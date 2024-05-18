package main

import (
	"video_generator/logging"
	"video_generator/rest"
    "video_generator/screenshot"
)

type PostRequestBody struct {
	Url     string   `json:"url"`
	Oneshot bool     `json:"oneshot"`
	Headers []string `json:"headers"`
}

func main() {
	logging.InitLogger()
	logging.Info("Dossified Shorts Generator v0.1")
	logging.Debug("Test")

	trendingArticles := rest.RequestTrends(0)
    screenshot.ScreenshotTrends(trendingArticles)

}
