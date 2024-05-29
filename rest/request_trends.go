package rest

import (
	"encoding/json"
	"fmt"
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

func RequestNewsTrends() []TrendArticle {
	amountTrends := config.GetConfiguration().AmountNewsTrends
	return requestTrends("news", amountTrends)
}

func RequestEventsTrends() []TrendArticle {
	amountTrends := config.GetConfiguration().AmountEventsTrends
	return requestTrends("events", amountTrends)
}

func requestTrends(requestType string, amountTrends int) []TrendArticle {

	var requestUrl string = getRestUrl(requestType, amountTrends)

	response, err := http.Get(requestUrl)
	utils.CheckError(err)

	responseBody, err := io.ReadAll(response.Body)
	utils.CheckError(err)

	stringBody := string(responseBody)

	articles := parseJson(stringBody)
	logging.Debug("Trends", zap.String("Trends", stringBody))

	return articles
}

func parseJson(stringBody string) []TrendArticle {
	data := []TrendArticle{}
	err := json.Unmarshal([]byte(stringBody), &data)
	utils.CheckError(err)
	return data
}

func getRestUrl(requestType string, amountTrends int) string {
	configuration := config.GetConfiguration()
	amountDaysOfTrends := configuration.AmountDaysTrends
	remoteUrl := configuration.RemoteUrl
	return remoteUrl + "/api/trends/?filter=" + requestType + "&amount=" + fmt.Sprint(amountTrends) + "&days=" + fmt.Sprint(amountDaysOfTrends)
}
