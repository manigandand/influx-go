package main

import (
	"log"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func writeData(con client.Client) {
	measurement := "incident-metrics"
	organizationID := primitive.NewObjectID().Hex()
	serviceID := primitive.NewObjectID().Hex()
	alertSourceID := primitive.NewObjectID().Hex()

	batchPoint, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: "test_incident",
	})
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	for i := 10; i >= 1; i-- {
		incidentID := primitive.NewObjectID().Hex()

		triggeredAt := now.Add(-(time.Duration(i+3) * time.Minute))
		acknowledgedAt := triggeredAt.Add((time.Duration(i) * time.Second))
		resolvedAt := triggeredAt.Add((time.Duration(i) * time.Minute))
		tta := acknowledgedAt.Sub(triggeredAt).Seconds()
		ttr := resolvedAt.Sub(triggeredAt).Seconds()

		triggeredPoint, err := client.NewPoint(
			measurement,
			map[string]string{ // tags
				"incident_id":     incidentID,
				"orgnization_id":  organizationID,
				"service_id":      serviceID,
				"alert_source_id": alertSourceID,
				"status":          "triggered",
				"event_type":      "triggered",
			},
			map[string]interface{}{"init": "triggered"}, // _fields
			triggeredAt,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		batchPoint.AddPoint(triggeredPoint)

		acknowledgedPoint, err := client.NewPoint(
			measurement,
			map[string]string{ // tags
				"incident_id":     incidentID,
				"orgnization_id":  organizationID,
				"service_id":      serviceID,
				"alert_source_id": alertSourceID,
				"status":          "acknowledged",
				"event_type":      "acknowledged",
			},
			map[string]interface{}{"tta": tta}, // _fields
			acknowledgedAt,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		batchPoint.AddPoint(acknowledgedPoint)

		resolvedPoint, err := client.NewPoint(
			measurement,
			map[string]string{ // tags
				"incident_id":     incidentID,
				"orgnization_id":  organizationID,
				"service_id":      serviceID,
				"alert_source_id": alertSourceID,
				"status":          "resolved",
				"event_type":      "resolved",
			},
			map[string]interface{}{"ttr": ttr}, // _fields
			resolvedAt,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		batchPoint.AddPoint(resolvedPoint)
	}

	if err := con.Write(batchPoint); err != nil {
		log.Fatal("write err: ", err)
	}
}
