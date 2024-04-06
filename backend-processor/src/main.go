package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestObject struct {
	UserId      string `json:"user_id"`
	MetricId    string `json:"metric_id"`
	MetricValue int64  `json:"metric_value"`
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
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
