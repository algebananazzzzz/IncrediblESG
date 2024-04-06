package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

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

type UserMetricData struct {
	MetricId  string `json:"metric_id"`
	Value     int64  `json:"value"`
	Timestamp int64  `json:"timestamp"`
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

func getMetricIds() ([]string, error) {
	var metricIds []string
	redisData, err := redisClient.HKeys(ctx, os.Getenv("REDIS_KEY")).Result()
	if err != nil {
		return nil, err
	}
	for _, v := range redisData {
		metricId, found := strings.CutPrefix(v, "metric:")
		if found {
			metricIds = append(metricIds, metricId)
		}
	}

	return metricIds, nil
}

func dumpAverage(data MetricAverage) error {
	if data.NumberOfRecords == 1 {
		data.Max = data.Average
		data.Min = data.Average
	}
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

func getUserData(userId string) (map[string]UserMetricData, error) {

	userData, err := redisClient.HGet(ctx, os.Getenv("REDIS_KEY"), fmt.Sprintf("user:%v", userId)).Result()
	if err == redis.Nil {
		return map[string]UserMetricData{}, nil
	}
	if err != nil {
		return nil, err
	}

	var data map[string]UserMetricData
	if err := json.Unmarshal([]byte(userData), &data); err != nil {
		return nil, err
	}

	return data, nil
}

func dumpUserData(userId string, metricId string, value int64) error {
	data, err := getUserData(userId)
	if err != nil {
		return err
	}

	data[metricId] = UserMetricData{
		Value:     value,
		MetricId:  metricId,
		Timestamp: time.Now().Unix(),
	}
	redisData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := redisClient.HSet(ctx, os.Getenv("REDIS_KEY"), fmt.Sprintf("user:%v", userId), redisData).Err(); err != nil && err != redis.Nil {
		return err
	}
	return nil
}

func deleteAllData() error {
	if err := redisClient.Del(ctx, os.Getenv("REDIS_KEY")).Err(); err != nil {
		return err
	}
	return nil
}
