package main

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return
	}
	defer c.Close()

	d, s, err := c.Ping(1 * time.Second)
	if err != nil {
		fmt.Println("Ping Error: ", err.Error())
		return
	}
	fmt.Println("Ping: ", d, s)

	// write data
	// writeData(c)

	// read data
	readDataValue(c, meanTTA)
	// readData(c, allData)

}
