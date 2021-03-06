package main

import (
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func writeData(client *influxdb.Client) {
	measurement := "incident-metrics"
	organizationID := primitive.NewObjectID().Hex()
	serviceID := primitive.NewObjectID().Hex()
	alertSourceID := primitive.NewObjectID().Hex()

	var allIncidentMerics []influxdb.Metric
	now := time.Now()

	for i := 10; i >= 1; i-- {
		incidentID := primitive.NewObjectID().Hex()

		triggeredAt := now.Add(-(time.Duration(i+3) * time.Minute))
		acknowledgedAt := triggeredAt.Add((time.Duration(i) * time.Second))
		resolvedAt := triggeredAt.Add((time.Duration(i) * time.Minute))
		tta := acknowledgedAt.Sub(triggeredAt).Seconds()
		ttr := resolvedAt.Sub(triggeredAt).Seconds()
		// fmt.Println("triggeredAt: ", triggeredAt)
		// fmt.Println("acknowledgedAt: ", acknowledgedAt)
		// fmt.Println("resolvedAt: ", resolvedAt)
		// fmt.Println("TTA: ", tta)
		// fmt.Println("TTR: ", ttr)
		// fmt.Println(strings.Repeat("~", 80))

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
