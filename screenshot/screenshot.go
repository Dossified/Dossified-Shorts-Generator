package screenshot

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"go.uber.org/zap"

	"github.com/Dossified/Dossified-Shorts-Generator/config"
	"github.com/Dossified/Dossified-Shorts-Generator/logging"
	"github.com/Dossified/Dossified-Shorts-Generator/rest"
	"github.com/Dossified/Dossified-Shorts-Generator/utils"
)

func ScreenshotEvents(events []rest.EventItem) {
	logging.Info("Taking screenshots")
	configuration := config.GetConfiguration()
	gowitnessHost := configuration.GowitnessHost
	restApiHost := configuration.RemoteUrl

	screenshotPostUrl := gowitnessHost + "/api/screenshot"
	screenshotPath := createScreenshotDir("events")

	for _, event := range events {
		requestUrl := fmt.Sprintf(
			"%s/vid_gen/?item_id=%d&obj_type=%s",
			restApiHost,
			event.EventId,
			"events",
		)
		logging.Debug("URL: ", zap.String("URL", requestUrl))
		requestBody := getRequestBody(requestUrl)

		screenshotData := screenshotHttpPostRequest(screenshotPostUrl, requestBody)
		saveScreenshotToFile(getScreenshotPath(screenshotPath, event.EventId), screenshotData)
		defer screenshotData.Close()
		logging.Debug("Successfully saved event " + fmt.Sprint(event.EventId))
	}
	logging.Info("Screenshots taken")
}

func ScreenshotTrends(trends []rest.TrendArticle, subFolder string) {
	logging.Info("Taking screenshots")
	configuration := config.GetConfiguration()
	gowitnessHost := configuration.GowitnessHost
	restApiHost := configuration.RemoteUrl

	screenshotPostUrl := gowitnessHost + "/api/screenshot"
	screenshotPath := createScreenshotDir(subFolder)

	for _, article := range trends {
		requestUrl := fmt.Sprintf(
			"%s/vid_gen/?item_id=%d&obj_type=%s",
			restApiHost,
			article.ArticleId,
			article.ArticleType,
		)
		logging.Debug("url", zap.String("url", requestUrl))
		requestBody := getRequestBody(requestUrl)

		screenshotData := screenshotHttpPostRequest(screenshotPostUrl, requestBody)
		saveScreenshotToFile(getScreenshotPath(screenshotPath, article.ArticleId), screenshotData)
		logging.Debug("Successfully taken screenshot of trend " + article.ArticleType + " " + fmt.Sprint(article.ArticleId))
	}
	logging.Info("Screenshots taken")
}

func getRequestBody(requestUrl string) []byte {
	body := []byte(`{
        "url": "` + requestUrl + `",
        "oneshot": "true"
    }`)
	return body
}

func createScreenshotDir(subFolder string) string {
	path := filepath.Join(".", "output/screenshots/"+subFolder)
	err := os.MkdirAll(path, os.ModePerm)
	utils.CheckError(err)
	logging.Debug("Screenshot path", zap.String("path", path))
	return path
}

func screenshotHttpPostRequest(url string, body []byte) io.ReadCloser {
	screenshotRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	utils.CheckError(err)
	screenshotRequest.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(screenshotRequest)
	utils.CheckError(err)
	logging.Debug("HTTP Response " + fmt.Sprint(res.Body))
	return res.Body
}

func saveScreenshotToFile(filePath string, screenshotData io.ReadCloser) {
	file, err := os.Create(filePath)
	utils.CheckError(err)
	defer file.Close()

	_, err = io.Copy(file, screenshotData)
	utils.CheckError(err)
}

func getScreenshotPath(screenshotPath string, itemId int) string {
	return screenshotPath + "/" + fmt.Sprint(itemId) + ".png"
}
