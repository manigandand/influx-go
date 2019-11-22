package main

import (
	"context"
	"log"
	"net/http"

	"github.com/influxdata/influxdb-client-go"
)

var bg = context.Background

const (
	influxOrgName       = "GopherHut"
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

	// dump data
	// writeData(influx)

	// read data
	// readData(influx, fluxQueryAck)
	readData(influx, fluxQueryByIncID)

	if err := influx.Close(); err != nil {
		log.Fatal(err)
	}
}
