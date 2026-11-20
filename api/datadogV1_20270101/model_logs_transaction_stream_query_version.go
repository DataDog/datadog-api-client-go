// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsTransactionStreamQueryVersion Test-only generated schema for the dashboard versioning proof of concept.
type LogsTransactionStreamQueryVersion string

// List of LogsTransactionStreamQueryVersion.
const (
	LOGSTRANSACTIONSTREAMQUERYVERSION_SEQUENTIAL_QUERY LogsTransactionStreamQueryVersion = "sequential_query"
)

var allowedLogsTransactionStreamQueryVersionEnumValues = []LogsTransactionStreamQueryVersion{
	LOGSTRANSACTIONSTREAMQUERYVERSION_SEQUENTIAL_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsTransactionStreamQueryVersion) GetAllowedValues() []LogsTransactionStreamQueryVersion {
	return allowedLogsTransactionStreamQueryVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsTransactionStreamQueryVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsTransactionStreamQueryVersion(value)
	return nil
}

// NewLogsTransactionStreamQueryVersionFromValue returns a pointer to a valid LogsTransactionStreamQueryVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsTransactionStreamQueryVersionFromValue(v string) (*LogsTransactionStreamQueryVersion, error) {
	ev := LogsTransactionStreamQueryVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsTransactionStreamQueryVersion: valid values are %v", v, allowedLogsTransactionStreamQueryVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsTransactionStreamQueryVersion) IsValid() bool {
	for _, existing := range allowedLogsTransactionStreamQueryVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsTransactionStreamQueryVersion value.
func (v LogsTransactionStreamQueryVersion) Ptr() *LogsTransactionStreamQueryVersion {
	return &v
}
