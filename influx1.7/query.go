package main

const (
	q  = `SELECT * FROM "test_incident"."autogen"."incident-metrics"`
	q1 = `SHOW TAG KEYS ON "test_incident" FROM "incident-metrics"`
	q2 = `SHOW TAG VALUES ON "test_incident" FROM "incident-metrics" WITH KEY = "incident_id"`
	q3 = `SHOW FIELD KEY CARDINALITY ON "test_incident"`
	q4 = `SHOW SERIES CARDINALITY ON "test_incident"`
	q5 = `SHOW TAG KEY CARDINALITY ON "test_incident"`

	meanTTA = `
	SELECT mean("tta") AS "mean_tta" 
	FROM "test_incident"."autogen"."incident-metrics" 
	WHERE time > now() - 12h
	AND "orgnization_id"='5dee1330c745915bbe752c8e'`

	sumTTA = `SELECT sum("tta") AS "mean_tta" 
	FROM "test_incident"."autogen"."incident-metrics" 
	WHERE time > now() - 12h 
	AND "orgnization_id"='5dee1330c745915bbe752c8e'`

	maxTTA = `SELECT max("tta") AS "max_tta" 
	FROM "test_incident"."autogen"."incident-metrics" 
	WHERE time > now() - 12h 
	AND "orgnization_id"='5dee1330c745915bbe752c8e'`

	allData = `SELECT * FROM "test_incident"."autogen"."incident-metrics" 
	WHERE time > now() - 12h 
	AND "orgnization_id"='5dee1330c745915bbe752c8e'`

	allDataByService = `SELECT * AS "mean_tta" 
	FROM "test_incident"."autogen"."incident-metrics" 
	WHERE time > now() - 12h 
	AND "orgnization_id"='5dee01af528a60a1d65f8741' 
	AND "service_id"='5dee01af528a60a1d65f8742'`
)
