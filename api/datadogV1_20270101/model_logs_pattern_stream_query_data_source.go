// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsPatternStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type LogsPatternStreamQueryDataSource string

// List of LogsPatternStreamQueryDataSource.
const (
	LOGSPATTERNSTREAMQUERYDATASOURCE_LOGS_PATTERN_STREAM LogsPatternStreamQueryDataSource = "logs_pattern_stream"
)

var allowedLogsPatternStreamQueryDataSourceEnumValues = []LogsPatternStreamQueryDataSource{
	LOGSPATTERNSTREAMQUERYDATASOURCE_LOGS_PATTERN_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsPatternStreamQueryDataSource) GetAllowedValues() []LogsPatternStreamQueryDataSource {
	return allowedLogsPatternStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsPatternStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsPatternStreamQueryDataSource(value)
	return nil
}

// NewLogsPatternStreamQueryDataSourceFromValue returns a pointer to a valid LogsPatternStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsPatternStreamQueryDataSourceFromValue(v string) (*LogsPatternStreamQueryDataSource, error) {
	ev := LogsPatternStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsPatternStreamQueryDataSource: valid values are %v", v, allowedLogsPatternStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsPatternStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedLogsPatternStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsPatternStreamQueryDataSource value.
func (v LogsPatternStreamQueryDataSource) Ptr() *LogsPatternStreamQueryDataSource {
	return &v
}
