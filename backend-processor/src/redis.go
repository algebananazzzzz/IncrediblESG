package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type MetricAverage struct {
	Id              string `json:"id"`
	Average         int64  `json:"average"`
	Max             int64  `json:"max"`
	Min             int64  `json:"min"`
	NumberOfRecords int64  `json:"number"`
	Positive        bool   `json:"positive"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func getAverageData(metricId string) (*MetricAverage, error) {
	redisData, err := redisClient.HGet(ctx, os.Getenv("REDIS_KEY"), fmt.Sprintf("metric:%v", metricId)).Result()
	if err == redis.Nil {
		return &MetricAverage{Id: metricId}, nil
	}
	if err != nil {
		return nil, err
	}
	var data MetricAverage
	if err := json.Unmarshal([]byte(redisData), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func dumpAverage(data MetricAverage) error {
	redisData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = redisClient.HSet(ctx, os.Getenv("REDIS_KEY"), fmt.Sprintf("metric:%v", data.Id), redisData).Err()
	if err != nil && err != redis.Nil {
		return err
	}
	return nil
}
