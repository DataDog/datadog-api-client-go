// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TraceStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type TraceStreamQueryDataSource string

// List of TraceStreamQueryDataSource.
const (
	TRACESTREAMQUERYDATASOURCE_TRACE_STREAM TraceStreamQueryDataSource = "trace_stream"
)

var allowedTraceStreamQueryDataSourceEnumValues = []TraceStreamQueryDataSource{
	TRACESTREAMQUERYDATASOURCE_TRACE_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TraceStreamQueryDataSource) GetAllowedValues() []TraceStreamQueryDataSource {
	return allowedTraceStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TraceStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TraceStreamQueryDataSource(value)
	return nil
}

// NewTraceStreamQueryDataSourceFromValue returns a pointer to a valid TraceStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTraceStreamQueryDataSourceFromValue(v string) (*TraceStreamQueryDataSource, error) {
	ev := TraceStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TraceStreamQueryDataSource: valid values are %v", v, allowedTraceStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TraceStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedTraceStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TraceStreamQueryDataSource value.
func (v TraceStreamQueryDataSource) Ptr() *TraceStreamQueryDataSource {
	return &v
}
