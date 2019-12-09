package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/influxdata/influxdb1-client/models"
	"github.com/mitchellh/mapstructure"
	"github.com/squadcastHQ/auxpkg/types"

	client "github.com/influxdata/influxdb1-client/v2"
)

func readDataValue(con client.Client, fluxQuery string) {
	q := client.NewQuery(fluxQuery, `test_incident`, "")
	response, err := con.Query(q)
	if err != nil {
		log.Fatal("read err: ", err)
	}
	if response.Error() != nil {
		log.Fatal("read err2: ", response.Error())
	}
	strings.Repeat("[ ]", 80)
	// decode into struct/map
	for _, res := range response.Results {
		// fmt.Printf("Result: %#v\n", res)
		// strings.Repeat("[ ]", 80)
		for _, row := range res.Series {
			// fmt.Printf("ROW: %#v\n", row)
			// strings.Repeat("-", 80)

			var datas []*map[string]interface{}

			unmarshal(row, func() interface{} {
				d := &map[string]interface{}{}
				datas = append(datas, d)
				return d
			})

			if bytes, err := json.Marshal(datas); err == nil {
				fmt.Printf("%+v\n", string(bytes))
			}
		}
		for _, msg := range res.Messages {
			fmt.Printf("MSG: %#v\n", msg)
			// strings.Repeat("*", 80)
		}
		strings.Repeat("~~~", 80)
	}

}

func readData(con client.Client, fluxQuery string) {
	q := client.NewQuery(fluxQuery, `test_incident`, "")
	response, err := con.Query(q)
	if err != nil {
		log.Fatal("read err: ", err)
	}
	if response.Error() != nil {
		log.Fatal("read err2: ", response.Error())
	}

	strings.Repeat("na", 2)

	// decode into struct/map
	for _, res := range response.Results {
		// fmt.Printf("Result: %#v\n", res)
		for _, row := range res.Series {
			// fmt.Printf("ROW Name: %#v\n\n", row.Name)
			// fmt.Printf("ROW Tags: %#v\n\n", row.Tags)
			// strings.Repeat("na", 2)
			// fmt.Printf("ROW Columns: %#v\n\n", row.Columns)
			// strings.Repeat("na", 2)
			// fmt.Printf("ROW Values: %#v\n\n", row.Values)
			// strings.Repeat("na", 2)
			// fmt.Printf("ROW Partial: %#v\n", row.Partial)
			// fmt.Println("----------------------------------------------------")
			// strings.Repeat("na", 2)

			var datas []*map[string]interface{}

			unmarshal(row, func() interface{} {
				d := &map[string]interface{}{}
				datas = append(datas, d)
				return d
			})

			if bytes, err := json.Marshal(datas); err == nil {
				fmt.Printf("%+v\n", string(bytes))
			}
		}

		for _, msg := range res.Messages {
			fmt.Printf("MSG: %#v\n", msg)
			strings.Repeat("*", 80)
		}
		strings.Repeat("~~~", 80)
	}
}

func unmarshal(row models.Row, v func() interface{}) error {
	keys := make(map[int]string)
	for i := range row.Columns {
		keys[i] = row.Columns[i]
	}
	for _, r := range row.Values {
		ptr := v()

		tmp := make(types.JSON)
		for i := range r {
			tmp[keys[i]] = r[i]
		}

		fmt.Printf("ROW Partial: %#v\n", row.Partial)

		err := mapstructure.Decode(tmp, ptr)
		if err != nil {
			return err
		}
	}

	return nil
}
