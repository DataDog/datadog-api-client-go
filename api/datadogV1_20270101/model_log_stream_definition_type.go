// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogStreamDefinitionType Test-only generated schema for the dashboard versioning proof of concept.
type LogStreamDefinitionType string

// List of LogStreamDefinitionType.
const (
	LOGSTREAMDEFINITIONTYPE_LOG_STREAM LogStreamDefinitionType = "log_stream"
)

var allowedLogStreamDefinitionTypeEnumValues = []LogStreamDefinitionType{
	LOGSTREAMDEFINITIONTYPE_LOG_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *LogStreamDefinitionType) GetAllowedValues() []LogStreamDefinitionType {
	return allowedLogStreamDefinitionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *LogStreamDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = LogStreamDefinitionType(value)
	return nil
}

// NewLogStreamDefinitionTypeFromValue returns a pointer to a valid LogStreamDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewLogStreamDefinitionTypeFromValue(v string) (*LogStreamDefinitionType, error) {
	ev := LogStreamDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LogStreamDefinitionType: valid values are %v", v, allowedLogStreamDefinitionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v LogStreamDefinitionType) IsValid() bool {
	for _, existing := range allowedLogStreamDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogStreamDefinitionType value.
func (v LogStreamDefinitionType) Ptr() *LogStreamDefinitionType {
	return &v
}
