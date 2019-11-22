package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/influxdata/influxdb-client-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var bg = context.Background

// IncidentMet ...
type IncidentMet struct {
	Result string `json:"result" flux:"result"`
	Table  string `json:"table" flux:"table"`
	// _time
	// Start time.Time `json:"_start"  flux:"_start"`
	// Stop  time.Time `json:"_stop" flux:"_stop"`
	// Time  time.Time `json:"_time" flux:"_time"`
	Start time.Time `json:"start"  flux:"start"`
	Stop  time.Time `json:"stop" flux:"stop"`
	Time  time.Time `json:"time" flux:"time"`

	// _fields
	// Value interface{} `json:"_value" flux:"_value"`
	// Field string      `json:"_field" flux:"_field"`
	Value interface{} `json:"value" flux:"value"`
	Field string      `json:"field" flux:"field"`

	// _measurement
	// Measurement string `json:"_measurement" flux:"_measurement"`
	Measurement string `json:"measurement" flux:"measurement"`

	// _tags
	IncidentID    string `json:"incident_id" flux:"incident_id"`
	OrgnizationID string `json:"orgnization_id" flux:"orgnization_id"`
	ServiceID     string `json:"service_id" flux:"service_id"`
	AlertSourceID string `json:"alert_source_id" flux:"alert_source_id"`
	Status        string `json:"status,omitempty" flux:"status"`
	EventType     string `json:"event_type,omitempty" flux:"event_type"`
}

// IncidentMetric ...
type IncidentMetric struct {
	Result string `json:"result" flux:"result"`
	Table  string `json:"table" flux:"table"`
	// _time
	Start time.Time `json:"_start"  flux:"_start"`
	Stop  time.Time `json:"_stop" flux:"_stop"`
	Time  time.Time `json:"_time" flux:"_time"`

	// _fields
	Value interface{} `json:"_value" flux:"_value"`
	Field string      `json:"_field" flux:"_field"`

	// _measurement
	Measurement string `json:"_measurement" flux:"_measurement"`

	// _tags
	IncidentID    string `json:"incident_id" flux:"incident_id"`
	OrgnizationID string `json:"orgnization_id" flux:"orgnization_id"`
	ServiceID     string `json:"service_id" flux:"service_id"`
	AlertSourceID string `json:"alert_source_id" flux:"alert_source_id"`
	Status        string `json:"status,omitempty" flux:"status"`
	EventType     string `json:"event_type,omitempty" flux:"event_type"`
}

// map[ alert_source_id:8 event_type:9 incident_id:10 orgnization_id:11  service_id:12 status:13 ]

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

	// dump data
	// writeData(influx)

	// read data
	fluxQuery := `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => r._measurement == "incident-metrics")
	|> filter(fn: (r) => r._field == "tta" or r._field == "ttr")
	|> filter(fn: (r) => r.alert_source_id == "5dd524fc8c61caf1309984ce")
	|> filter(fn: (r) => r.event_type == "acknowledged")`

	var extern IncidentMet
	res, err := influx.QueryCSV(bg(), fluxQuery, influxOrgName, extern)
	if err != nil {
		log.Fatal("Read Err: ", err)
	}
	if res.Err != nil {
		log.Fatal("Read Err: ", res.Err)
	}

	type QueryCSVResult struct {
		io.ReadCloser
		csvReader   *csv.Reader
		Row         []string
		ColNames    []string
		colNamesMap map[string]int
		dataTypes   []string
		group       []bool
		defaultVals []string
		Err         error
	}

	for res.Next() {
		fmt.Println("Row: ", res.Row)
		fmt.Println("ColNames: ", res.ColNames)

		var data IncidentMetric
		if err := res.Unmarshal(&data); err != nil {
			log.Fatal("Unmarshal Err: ", err)
		}
		fmt.Printf("%+v\n", data)
		break
	}

	if err := influx.Close(); err != nil {
		log.Fatal(err)
	}
}

func writeData(client *influxdb.Client) {
	primitive.NewObjectID()
	measurement := "incident-metrics"
	organizationID := primitive.NewObjectID().Hex()
	serviceID := primitive.NewObjectID().Hex()
	alertSourceID := primitive.NewObjectID().Hex()

	var allIncidentMerics []influxdb.Metric
	now := time.Now()
	fmt.Println(now)
	fmt.Println(strings.Repeat("-", 80))

	for i := 10; i >= 1; i-- {
		incidentID := primitive.NewObjectID().Hex()

		triggeredAt := now.Add(-(time.Duration(i+3) * time.Minute))
		acknowledgedAt := triggeredAt.Add((time.Duration(i) * time.Second))
		resolvedAt := triggeredAt.Add((time.Duration(i) * time.Minute))
		tta := acknowledgedAt.Sub(triggeredAt).Seconds()
		ttr := resolvedAt.Sub(triggeredAt).Seconds()
		fmt.Println("triggeredAt: ", triggeredAt)
		fmt.Println("acknowledgedAt: ", acknowledgedAt)
		fmt.Println("resolvedAt: ", resolvedAt)
		fmt.Println("TTA: ", tta)
		fmt.Println("TTR: ", ttr)
		fmt.Println(strings.Repeat("~", 80))

		incidentMerics := []influxdb.Metric{
			influxdb.NewRowMetric(
				map[string]interface{}{"init": "triggered"}, // _fields
				measurement,
				map[string]string{ // tags
					"incident_id":     incidentID,
					"orgnization_id":  organizationID,
					"service_id":      serviceID,
					"alert_source_id": alertSourceID,
					"status":          "triggered",
					"event_type":      "triggered",
				},
				triggeredAt,
			),
			influxdb.NewRowMetric(
				map[string]interface{}{"tta": tta}, // _fields
				measurement,
				map[string]string{ // tags
					"incident_id":     incidentID,
					"orgnization_id":  organizationID,
					"service_id":      serviceID,
					"alert_source_id": alertSourceID,
					"status":          "acknowledged",
					"event_type":      "acknowledged",
				},
				acknowledgedAt,
			),
			influxdb.NewRowMetric(
				map[string]interface{}{"ttr": ttr}, // _fields
				measurement,
				map[string]string{ // tags
					"incident_id":     incidentID,
					"orgnization_id":  organizationID,
					"service_id":      serviceID,
					"alert_source_id": alertSourceID,
					"status":          "resolved",
					"event_type":      "resolved",
				},
				resolvedAt,
			),
		}

		allIncidentMerics = append(allIncidentMerics, incidentMerics...)
	}

	if _, err := client.Write(bg(), bucketName, influxOrgName, allIncidentMerics...); err != nil {
		log.Fatal("write err: ", err)
	}
}
