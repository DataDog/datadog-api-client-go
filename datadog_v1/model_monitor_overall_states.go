/*
 * (C) Datadog, Inc. 2019
 * All rights reserved
 * Licensed under a 3-clause BSD style license (see LICENSE)
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog_v1

import (
	"bytes"
	"encoding/json"
)

// MonitorOverallStates the model 'MonitorOverallStates'
type MonitorOverallStates string

// List of MonitorOverallStates
const (
	ALERT   MonitorOverallStates = "Alert"
	IGNORED MonitorOverallStates = "Ignored"
	NO_DATA MonitorOverallStates = "No Data"
	OK      MonitorOverallStates = "OK"
	SKIPPED MonitorOverallStates = "Skipped"
	UNKNOWN MonitorOverallStates = "Unknown"
	WARN    MonitorOverallStates = "Warn"
)

type NullableMonitorOverallStates struct {
	Value        MonitorOverallStates
	ExplicitNull bool
}

func (v NullableMonitorOverallStates) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull && v.Value != "":
		return nil, ErrInvalidNullable
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMonitorOverallStates) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
