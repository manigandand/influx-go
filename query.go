package main

const fluxQueryAck = `from(bucket: "test_bucket")
	|> range(start: -350h, stop: now())
	|> filter(fn: (r) => r._measurement == "incident-metrics")
	|> filter(fn: (r) => r._field == "tta" or r._field == "ttr")
	|> filter(fn: (r) => r.alert_source_id == "5dd524fc8c61caf1309984ce")
	|> filter(fn: (r) => r.event_type == "acknowledged")`
