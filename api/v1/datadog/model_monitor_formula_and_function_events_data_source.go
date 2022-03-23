/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// MonitorFormulaAndFunctionEventsDataSource Data source for event platform-based queries.
type MonitorFormulaAndFunctionEventsDataSource string

// List of MonitorFormulaAndFunctionEventsDataSource
const (
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_RUM          MonitorFormulaAndFunctionEventsDataSource = "rum"
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES MonitorFormulaAndFunctionEventsDataSource = "ci_pipelines"
)

var allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues = []MonitorFormulaAndFunctionEventsDataSource{
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_RUM,
	MONITORFORMULAANDFUNCTIONEVENTSDATASOURCE_CI_PIPELINES,
}

func (w *MonitorFormulaAndFunctionEventsDataSource) GetAllowedValues() []MonitorFormulaAndFunctionEventsDataSource {
	return allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues
}

func (v *MonitorFormulaAndFunctionEventsDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonitorFormulaAndFunctionEventsDataSource(value)
	return nil
}

// NewMonitorFormulaAndFunctionEventsDataSourceFromValue returns a pointer to a valid MonitorFormulaAndFunctionEventsDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorFormulaAndFunctionEventsDataSourceFromValue(v string) (*MonitorFormulaAndFunctionEventsDataSource, error) {
	ev := MonitorFormulaAndFunctionEventsDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorFormulaAndFunctionEventsDataSource: valid values are %v", v, allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorFormulaAndFunctionEventsDataSource) IsValid() bool {
	for _, existing := range allowedMonitorFormulaAndFunctionEventsDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorFormulaAndFunctionEventsDataSource value
func (v MonitorFormulaAndFunctionEventsDataSource) Ptr() *MonitorFormulaAndFunctionEventsDataSource {
	return &v
}

type NullableMonitorFormulaAndFunctionEventsDataSource struct {
	value *MonitorFormulaAndFunctionEventsDataSource
	isSet bool
}

func (v NullableMonitorFormulaAndFunctionEventsDataSource) Get() *MonitorFormulaAndFunctionEventsDataSource {
	return v.value
}

func (v *NullableMonitorFormulaAndFunctionEventsDataSource) Set(val *MonitorFormulaAndFunctionEventsDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorFormulaAndFunctionEventsDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorFormulaAndFunctionEventsDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorFormulaAndFunctionEventsDataSource(val *MonitorFormulaAndFunctionEventsDataSource) *NullableMonitorFormulaAndFunctionEventsDataSource {
	return &NullableMonitorFormulaAndFunctionEventsDataSource{value: val, isSet: true}
}

func (v NullableMonitorFormulaAndFunctionEventsDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorFormulaAndFunctionEventsDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
