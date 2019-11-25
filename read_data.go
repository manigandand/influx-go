package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/influxdata/influxdb-client-go"
)

func readDataValue(client *influxdb.Client, fluxQuery string) {
	res, err := client.QueryCSV(bg(), fluxQuery, influxOrgName)
	if err != nil {
		log.Fatal("Read Err: ", err)
	}
	if res.Err != nil {
		log.Fatal("Read Err: ", res.Err)
	}

	var mttaMetrics mtta
	for res.Next() {
		// fmt.Println("Row: ", res.Row)
		// fmt.Println("ColNames: ", res.ColNames)

		if err := res.Unmarshal(&mttaMetrics); err != nil {
			log.Fatal("Unmarshal Err: ", err)
		}
	}

	if bytes, err := json.Marshal(mttaMetrics); err == nil {
		fmt.Printf("%+v\n", string(bytes))
	}
}

func readData(client *influxdb.Client, fluxQuery string) {
	// var extern IncidentMet
	// extern := map[string]interface{}{
	// 	"result":          111,
	// 	"alert_source_id": "manigandan",
	// }
	res, err := client.QueryCSV(bg(), fluxQuery, influxOrgName)
	if err != nil {
		log.Fatal("Read Err: ", err)
	}
	if res.Err != nil {
		log.Fatal("Read Err: ", res.Err)
	}

	var incidentMetrics []*IncidentMetric
	for res.Next() {
		// fmt.Println("Row: ", res.Row)
		// fmt.Println("ColNames: ", res.ColNames)
		var data IncidentMetric
		if err := res.Unmarshal(&data); err != nil {
			log.Fatal("Unmarshal Err: ", err)
		}
		// fmt.Printf("%+v\n", data)
		incidentMetrics = append(incidentMetrics, &data)
	}

	if bytes, err := json.Marshal(incidentMetrics); err == nil {
		fmt.Printf("%+v\n", string(bytes))
	}
}
