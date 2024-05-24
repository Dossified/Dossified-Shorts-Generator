package main

import (
	"github.com/Dossified/Dossified-Shorts-Generator/logging"
	"github.com/Dossified/Dossified-Shorts-Generator/rest"
	"github.com/Dossified/Dossified-Shorts-Generator/screenshot"
	"github.com/Dossified/Dossified-Shorts-Generator/upload/youtube"
	"github.com/Dossified/Dossified-Shorts-Generator/upload/instagram"
	"github.com/Dossified/Dossified-Shorts-Generator/video"
	"github.com/Dossified/Dossified-Shorts-Generator/config"
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

    if config.GetConfiguration().UploadToYouTube {
	    youtube.UploadVideo(videoPath)
    }
    if config.GetConfiguration().UploadToInstagram {
        instagram.UploadToInstagram(videoPath)
    }
}
