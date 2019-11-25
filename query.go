package main

const (
	//count
	fluxQueryCount = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => 
		r._measurement == "incident-metrics" and
		r.incident_id == "5dd524fc8c61caf1309984d3"
	)
	|> count(column: "_value")`

	// sum
	fluxQuerySum = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => 
		r._measurement == "incident-metrics" and
		r._field == "tta" and
		r.orgnization_id == "5dd524fc8c61caf1309984cc"
	)
    |> group()
    |> distinct()
	|> sum(column: "_value")`

	fluxQueryMean = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => 
	  r._measurement == "incident-metrics" and
	  r._field == "tta" and
	  r.orgnization_id == "5dd524fc8c61caf1309984cc"
	)
    |> group()
    |> distinct()
    |> mean(column: "_value")`

	fluxQueryAck = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => r._measurement == "incident-metrics")
	|> filter(fn: (r) => r._field == "tta" or r._field == "ttr")
	|> filter(fn: (r) => r.alert_source_id == "5dd524fc8c61caf1309984ce")
	|> filter(fn: (r) => r.event_type == "acknowledged")`

	fluxQueryByIncID = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => r._measurement == "incident-metrics")
	|> filter(fn: (r) => r._field == "tta" or r._field == "ttr")
	|> filter(fn: (r) => r.incident_id == "5dd524fc8c61caf1309984d3")`

	fluxQueryRangeTTA = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => 
	  r._measurement == "incident-metrics" and
	  r._field == "tta" and
	  r.orgnization_id == "5dd524fc8c61caf1309984cc"
	)
    |> filter(fn: (r) => 
	  r._value > 3 and r._value < 8
	)
	|> set(key: "host", value: "prod-node-1")
    |> sort(columns: ["_value"], desc: false)
	|> group()`

	fluxQueryHigestMaxTTA = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => 
	  r._measurement == "incident-metrics" and
	  r._field == "tta" and
	  r.orgnization_id == "5dd524fc8c61caf1309984cc"
	)
    |> sort(columns: ["_value"], desc: false)
    |> highestMax(n:3, groupColumns: ["incident_id", "orgnization_id"])
	|> group()`

	fluxQueryLowestMinTTA = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => 
	  r._measurement == "incident-metrics" and
	  r._field == "tta" and
	  r.orgnization_id == "5dd524fc8c61caf1309984cc"
	)
    
    |> sort(columns: ["_value"], desc: false)
    |> lowestMin(n:3, groupColumns: ["incident_id", "orgnization_id"])
    |> group()`
)
