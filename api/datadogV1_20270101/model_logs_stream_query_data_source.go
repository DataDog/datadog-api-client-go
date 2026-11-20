// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type LogsStreamQueryDataSource string

// List of LogsStreamQueryDataSource.
const (
	LOGSSTREAMQUERYDATASOURCE_LOGS_STREAM LogsStreamQueryDataSource = "logs_stream"
)

var allowedLogsStreamQueryDataSourceEnumValues = []LogsStreamQueryDataSource{
	LOGSSTREAMQUERYDATASOURCE_LOGS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsStreamQueryDataSource) GetAllowedValues() []LogsStreamQueryDataSource {
	return allowedLogsStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsStreamQueryDataSource(value)
	return nil
}

// NewLogsStreamQueryDataSourceFromValue returns a pointer to a valid LogsStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsStreamQueryDataSourceFromValue(v string) (*LogsStreamQueryDataSource, error) {
	ev := LogsStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsStreamQueryDataSource: valid values are %v", v, allowedLogsStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedLogsStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsStreamQueryDataSource value.
func (v LogsStreamQueryDataSource) Ptr() *LogsStreamQueryDataSource {
	return &v
}
