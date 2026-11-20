// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type LogsStreamRequestResponseFormat string

// List of LogsStreamRequestResponseFormat.
const (
	LOGSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST LogsStreamRequestResponseFormat = "event_list"
)

var allowedLogsStreamRequestResponseFormatEnumValues = []LogsStreamRequestResponseFormat{
	LOGSSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsStreamRequestResponseFormat) GetAllowedValues() []LogsStreamRequestResponseFormat {
	return allowedLogsStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsStreamRequestResponseFormat(value)
	return nil
}

// NewLogsStreamRequestResponseFormatFromValue returns a pointer to a valid LogsStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsStreamRequestResponseFormatFromValue(v string) (*LogsStreamRequestResponseFormat, error) {
	ev := LogsStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsStreamRequestResponseFormat: valid values are %v", v, allowedLogsStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedLogsStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsStreamRequestResponseFormat value.
func (v LogsStreamRequestResponseFormat) Ptr() *LogsStreamRequestResponseFormat {
	return &v
}
