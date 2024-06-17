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

type EventItem struct {
	EventId         int    `json:"pk"`
	EventTitle      string `json:"title"`
	EventSource     string `json:"source"`
	EventDatePublic string `json:"date_public"`
	EventDateStart  string `json:"date_start"`
	EventDateEnd    string `json:"date_end"`
	EventCoinId     int    `json:"coin_id"`
	EventTags       int    `json:"tags"`
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

func RequestUpcomingEvents() []EventItem {
	return requestEvents("events")
}

func requestEvents(requestType string) []EventItem {
	var requestUrl string = getRestUrl(requestType, 0, "events")

	response, err := http.Get(requestUrl)
	utils.CheckError(err)

	responseBody, err := io.ReadAll(response.Body)
	utils.CheckError(err)

	stringBody := string(responseBody)

	articles := parseEventsJson(stringBody)
	logging.Debug("Events", zap.String("Events", stringBody))

	return articles
}

func requestTrends(requestType string, amountTrends int) []TrendArticle {
	var requestUrl string = getRestUrl(requestType, amountTrends, "trends")

	response, err := http.Get(requestUrl)
	utils.CheckError(err)

	responseBody, err := io.ReadAll(response.Body)
	utils.CheckError(err)

	stringBody := string(responseBody)

	articles := parseTrendsJson(stringBody)
	logging.Debug("Trends", zap.String("Trends", stringBody))

	return articles
}

func parseEventsJson(stringBody string) []EventItem {
	data := []EventItem{}
	logging.Debug("Events JSON: " + fmt.Sprint(data))
	err := json.Unmarshal([]byte(stringBody), &data)
	utils.CheckError(err)
	return data
}

func parseTrendsJson(stringBody string) []TrendArticle {
	data := []TrendArticle{}
	err := json.Unmarshal([]byte(stringBody), &data)
	utils.CheckError(err)
	return data
}

func getRestUrl(requestType string, amountTrends int, endpointType string) string {
	configuration := config.GetConfiguration()
	amountDaysOfTrends := configuration.AmountDaysTrends
	remoteUrl := configuration.RemoteUrl
	fullRemoteUrl := remoteUrl + "/api/" + endpointType + "/?filter=" + requestType + "&days=" + fmt.Sprint(amountDaysOfTrends)
	if amountTrends != 0 {
		fullRemoteUrl = fullRemoteUrl + "&amount=" + fmt.Sprint(amountTrends)
	}
	logging.Debug("REST URL: " + fullRemoteUrl)
	return fullRemoteUrl
}
