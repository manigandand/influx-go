package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

const (
	influxOrgName = "GopherHut"
	// bucketName          = "incident"
	bucketName          = "test_bucket"
	myHTTPInfluxAddress = "http://localhost:9999"
	myToken             = "RM2Ezx1ZYedbIqjs-_PHBGUN3MVB1RzIan_qwxZf7BTlyIT12RBqJVuPJIgMX91DPB9-YjB3T-rPvVxfrpUAdA=="
)

func main() {
	myHTTPClient := http.DefaultClient
	influx, err := influxdb.New(myHTTPInfluxAddress, myToken, influxdb.WithHTTPClient(myHTTPClient))
	if err != nil {
		panic(err)
	}

	myMetrics := []influxdb.Metric{
		influxdb.NewRowMetric(
			map[string]interface{}{"memory": 1000, "cpu": 0.93},
			"system-metrics",
			map[string]string{"hostname": "hal9000"},
			time.Date(2018, 3, 4, 5, 6, 7, 8, time.UTC)),
		influxdb.NewRowMetric(
			map[string]interface{}{"memory": 1000, "cpu": 0.93},
			"system-metrics",
			map[string]string{"hostname": "hal9000"},
			time.Date(2018, 3, 4, 5, 6, 7, 9, time.UTC)),
	}

	if _, err := influx.Write(context.Background(), bucketName, influxOrgName, myMetrics...); err != nil {
		log.Fatal(err)
	}

	influx.Close()
}
