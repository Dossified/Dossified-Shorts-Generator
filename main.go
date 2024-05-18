package main

import (
	"github.com/Dominique-Roth/Dossified-Shorts-Generator/logging"
	"github.com/Dominique-Roth/Dossified-Shorts-Generator/rest"
	"github.com/Dominique-Roth/Dossified-Shorts-Generator/screenshot"
	"github.com/Dominique-Roth/Dossified-Shorts-Generator/upload/youtube"
	"github.com/Dominique-Roth/Dossified-Shorts-Generator/video"
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

	videoPath := video.CreateVideo()

	youtube.UploadVideo(videoPath)
}
