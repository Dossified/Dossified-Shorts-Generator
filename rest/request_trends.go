package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"github.com/Dossified/Dossified-Shorts-Generator/config"
	"github.com/Dossified/Dossified-Shorts-Generator/logging"
	"github.com/Dossified/Dossified-Shorts-Generator/utils"
)

type RequestType int

type TrendArticle struct {
	ArticleType string `json:"type"`
	ArticleId   int    `json:"id"`
}

const (
	NEWS RequestType = iota
	EVENTS
	STAFF_PICKED
)

func RequestTrends(requestType int) []TrendArticle {

	var requestUrl string = getRestUrl(requestType)

	response, err := http.Get(requestUrl)
	utils.CheckError(err)

	responseBody, err := io.ReadAll(response.Body)
	utils.CheckError(err)

	stringBody := string(responseBody)

	articles := parseJson(stringBody)
	//logging.Debug("Trending articles", map[string][]TrendArticle{"articles": articles,})
	logging.Debug("Trends", zap.String("Trends", stringBody))

	return articles
}

func parseJson(stringBody string) []TrendArticle {
	data := []TrendArticle{}
	err := json.Unmarshal([]byte(stringBody), &data)
	utils.CheckError(err)
	return data
}

func getRestUrl(requestType int) string {
	// ToDo: Add filters to request & server to only provide certain types of trends e.g. news
	remoteUrl := config.GetConfiguration().RemoteUrl
	urlMap := map[int]string{
		0: remoteUrl + "/api/trends/",
	}
	return urlMap[requestType]
}
