// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsTransactionStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type LogsTransactionStreamQueryDataSource string

// List of LogsTransactionStreamQueryDataSource.
const (
	LOGSTRANSACTIONSTREAMQUERYDATASOURCE_LOGS_TRANSACTION_STREAM LogsTransactionStreamQueryDataSource = "logs_transaction_stream"
)

var allowedLogsTransactionStreamQueryDataSourceEnumValues = []LogsTransactionStreamQueryDataSource{
	LOGSTRANSACTIONSTREAMQUERYDATASOURCE_LOGS_TRANSACTION_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsTransactionStreamQueryDataSource) GetAllowedValues() []LogsTransactionStreamQueryDataSource {
	return allowedLogsTransactionStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsTransactionStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsTransactionStreamQueryDataSource(value)
	return nil
}

// NewLogsTransactionStreamQueryDataSourceFromValue returns a pointer to a valid LogsTransactionStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsTransactionStreamQueryDataSourceFromValue(v string) (*LogsTransactionStreamQueryDataSource, error) {
	ev := LogsTransactionStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsTransactionStreamQueryDataSource: valid values are %v", v, allowedLogsTransactionStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsTransactionStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedLogsTransactionStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsTransactionStreamQueryDataSource value.
func (v LogsTransactionStreamQueryDataSource) Ptr() *LogsTransactionStreamQueryDataSource {
	return &v
}
