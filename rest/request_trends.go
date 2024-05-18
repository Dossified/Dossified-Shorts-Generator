package rest

import (
    "fmt"
)

type RequestType int;

const (
    NEWS RequestType = iota
    EVENTS
    STAFF_PICKED
)

func RequestTrends(requestType int) {
    
}
