package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestObject struct {
	UserId      string `json:"user_id"`
	MetricId    string `json:"metric_id"`
	MetricValue int64  `json:"metric_value"`
}

type ResponseObject struct {
	UserId        string        `json:"user_id"`
	MetricId      string        `json:"metric_id"`
	MetricAverage MetricAverage `json:"metric_average"`
	Score         int           `json:"score"`
}

type SchedulerEvent struct {
	Command string `json:"command"`
	Payload string `json:"payload"`
}

type Event struct {
	events.APIGatewayProxyRequest
	SchedulerEvent
}

func Handler(ctx context.Context, event Event) (interface{}, error) {
	switch {
	case !reflect.DeepEqual(event.APIGatewayProxyRequest, events.APIGatewayProxyRequest{}):
		return ApiRequestHandler(ctx, event.APIGatewayProxyRequest)
	case !reflect.DeepEqual(event.SchedulerEvent, SchedulerEvent{}):
		return nil, SchedulerEventHandler(ctx, event.SchedulerEvent)
	default:
		return nil, fmt.Errorf("unknown event: %v", event)
	}
}

func ApiRequestHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestObject RequestObject

	if req.Body == "" {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: "Hello World"}, nil
	}

	if err := json.Unmarshal([]byte(req.Body), &requestObject); err != nil {
		log.Panic(err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	var metricAverageData MetricAverage

	if data, err := getAverageData(requestObject.MetricId); err != nil {
		log.Panic(err)
	} else {
		metricAverageData = updateAverage(requestObject.MetricValue, *data)
	}

	if err := dumpAverage(metricAverageData); err != nil {
		log.Panic(err)
	}

	score := calculateDeviation(requestObject.MetricValue, metricAverageData)

	if err := dumpUserData(requestObject.UserId, score); err != nil {
		log.Panic(err)
	}

	responseObject := ResponseObject{
		UserId:        requestObject.UserId,
		MetricId:      requestObject.MetricId,
		MetricAverage: metricAverageData,
		Score:         score,
	}

	responseJson, _ := json.Marshal(responseObject)

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(responseJson)}, nil
}

func SchedulerEventHandler(ctx context.Context, req SchedulerEvent) error {
	switch req.Command {
	case "clear_cache":
		deleteAllData()
	case "reset_cache":
		metricIds, err := getMetricIds()
		if err != nil {
			return err
		}
		for _, metricId := range metricIds {
			if data, err := getAverageData(metricId); err != nil {
				return err
			} else {
				data.NumberOfRecords = 1
				if err := dumpAverage(*data); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func main() {
	if os.Getenv("LAMBDA_TASK_ROOT") != "" {
		lambda.Start(Handler)
	} else {
		fmt.Print(calculateDeviation(0,
			MetricAverage{
				Average:  5,
				Max:      10,
				Min:      0,
				Positive: false,
			}))
	}
}
