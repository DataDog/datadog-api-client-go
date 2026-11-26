// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsTransactionStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type LogsTransactionStreamRequestResponseFormat string

// List of LogsTransactionStreamRequestResponseFormat.
const (
	LOGSTRANSACTIONSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST LogsTransactionStreamRequestResponseFormat = "event_list"
)

var allowedLogsTransactionStreamRequestResponseFormatEnumValues = []LogsTransactionStreamRequestResponseFormat{
	LOGSTRANSACTIONSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsTransactionStreamRequestResponseFormat) GetAllowedValues() []LogsTransactionStreamRequestResponseFormat {
	return allowedLogsTransactionStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsTransactionStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsTransactionStreamRequestResponseFormat(value)
	return nil
}

// NewLogsTransactionStreamRequestResponseFormatFromValue returns a pointer to a valid LogsTransactionStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsTransactionStreamRequestResponseFormatFromValue(v string) (*LogsTransactionStreamRequestResponseFormat, error) {
	ev := LogsTransactionStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsTransactionStreamRequestResponseFormat: valid values are %v", v, allowedLogsTransactionStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsTransactionStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedLogsTransactionStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsTransactionStreamRequestResponseFormat value.
func (v LogsTransactionStreamRequestResponseFormat) Ptr() *LogsTransactionStreamRequestResponseFormat {
	return &v
}
