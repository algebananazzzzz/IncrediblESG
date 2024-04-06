package main

import (
	"fmt"
	"math"
)

func calculateDeviation(data int64, metricData MetricAverage) int {
	if data == metricData.Average {
		return 3
	} else if data < metricData.Average && metricData.Positive {
		return int(math.Round(3 + 7*(math.Abs(float64(data-metricData.Average))/math.Abs(float64(metricData.Min)-float64(metricData.Average)))))
	} else if data > metricData.Average && !metricData.Positive {
		return int(math.Round(3 + 7*(math.Abs(float64(data-metricData.Average))/math.Abs(float64(metricData.Max)-float64(metricData.Average)))))
	} else if metricData.Positive {
		fmt.Print("here")
		return int(math.Round(3 - 2*(math.Abs(float64(data-metricData.Average))/math.Abs(float64(metricData.Min)-float64(metricData.Average)))))
	} else {
		return int(math.Round(3 - 2*(math.Abs(float64(data-metricData.Average))/math.Abs(float64(metricData.Max)-float64(metricData.Average)))))
	}
}

func updateAverage(data int64, metricData MetricAverage) MetricAverage {
	metricData.Average = (metricData.Average*metricData.NumberOfRecords + data) / (metricData.NumberOfRecords + 1)
	metricData.NumberOfRecords += 1

	if data > metricData.Max {
		metricData.Max = data
	}

	if data < metricData.Min {
		metricData.Min = data
	}
	return metricData
}
