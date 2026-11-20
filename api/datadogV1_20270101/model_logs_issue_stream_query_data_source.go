// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsIssueStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type LogsIssueStreamQueryDataSource string

// List of LogsIssueStreamQueryDataSource.
const (
	LOGSISSUESTREAMQUERYDATASOURCE_LOGS_ISSUE_STREAM LogsIssueStreamQueryDataSource = "logs_issue_stream"
)

var allowedLogsIssueStreamQueryDataSourceEnumValues = []LogsIssueStreamQueryDataSource{
	LOGSISSUESTREAMQUERYDATASOURCE_LOGS_ISSUE_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsIssueStreamQueryDataSource) GetAllowedValues() []LogsIssueStreamQueryDataSource {
	return allowedLogsIssueStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsIssueStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsIssueStreamQueryDataSource(value)
	return nil
}

// NewLogsIssueStreamQueryDataSourceFromValue returns a pointer to a valid LogsIssueStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsIssueStreamQueryDataSourceFromValue(v string) (*LogsIssueStreamQueryDataSource, error) {
	ev := LogsIssueStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsIssueStreamQueryDataSource: valid values are %v", v, allowedLogsIssueStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsIssueStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedLogsIssueStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsIssueStreamQueryDataSource value.
func (v LogsIssueStreamQueryDataSource) Ptr() *LogsIssueStreamQueryDataSource {
	return &v
}
