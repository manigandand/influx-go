package main

import "time"

type mtta struct {
	Result string `json:"result,omitempty" flux:"result"`
	Table  string `json:"table" flux:"table"`
	// _time
	Start time.Time `json:"start"  flux:"_start"`
	Stop  time.Time `json:"stop" flux:"_stop"`

	// _fields
	Value interface{} `json:"value" flux:"_value"`
}

// IncidentMetric ...
type IncidentMetric struct {
	Result string `json:"result,omitempty" flux:"result"`
	Table  string `json:"table" flux:"table"`
	// _time
	Start time.Time `json:"start"  flux:"_start"`
	Stop  time.Time `json:"stop" flux:"_stop"`
	Time  time.Time `json:"time" flux:"_time"`

	// _fields
	Value interface{} `json:"value" flux:"_value"`
	Field string      `json:"field" flux:"_field"`

	// _measurement
	Measurement string `json:"measurement" flux:"_measurement"`

	// _tags
	IncidentID    string `json:"incident_id" flux:"incident_id"`
	OrgnizationID string `json:"orgnization_id" flux:"orgnization_id"`
	ServiceID     string `json:"service_id" flux:"service_id"`
	AlertSourceID string `json:"alert_source_id" flux:"alert_source_id"`
	Status        string `json:"status,omitempty" flux:"status"`
	EventType     string `json:"event_type,omitempty" flux:"event_type"`
}

// Extern ...
type Extern struct {
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
