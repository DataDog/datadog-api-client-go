// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsPatternStreamRequestResponseFormat Test-only generated schema for the dashboard versioning proof of concept.
type LogsPatternStreamRequestResponseFormat string

// List of LogsPatternStreamRequestResponseFormat.
const (
	LOGSPATTERNSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST LogsPatternStreamRequestResponseFormat = "event_list"
)

var allowedLogsPatternStreamRequestResponseFormatEnumValues = []LogsPatternStreamRequestResponseFormat{
	LOGSPATTERNSTREAMREQUESTRESPONSEFORMAT_EVENT_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogsPatternStreamRequestResponseFormat) GetAllowedValues() []LogsPatternStreamRequestResponseFormat {
	return allowedLogsPatternStreamRequestResponseFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogsPatternStreamRequestResponseFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogsPatternStreamRequestResponseFormat(value)
	return nil
}

// NewLogsPatternStreamRequestResponseFormatFromValue returns a pointer to a valid LogsPatternStreamRequestResponseFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogsPatternStreamRequestResponseFormatFromValue(v string) (*LogsPatternStreamRequestResponseFormat, error) {
	ev := LogsPatternStreamRequestResponseFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogsPatternStreamRequestResponseFormat: valid values are %v", v, allowedLogsPatternStreamRequestResponseFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogsPatternStreamRequestResponseFormat) IsValid() bool {
	for _, existing := range allowedLogsPatternStreamRequestResponseFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogsPatternStreamRequestResponseFormat value.
func (v LogsPatternStreamRequestResponseFormat) Ptr() *LogsPatternStreamRequestResponseFormat {
	return &v
}
