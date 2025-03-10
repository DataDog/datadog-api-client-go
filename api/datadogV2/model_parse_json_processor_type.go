// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ParseJSONProcessorType The type of processor.
type ParseJSONProcessorType string

// List of ParseJSONProcessorType.
const (
	PARSEJSONPROCESSORTYPE_PARSE_JSON ParseJSONProcessorType = "parse_json"
)

var allowedParseJSONProcessorTypeEnumValues = []ParseJSONProcessorType{
	PARSEJSONPROCESSORTYPE_PARSE_JSON,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ParseJSONProcessorType) GetAllowedValues() []ParseJSONProcessorType {
	return allowedParseJSONProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ParseJSONProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ParseJSONProcessorType(value)
	return nil
}

// NewParseJSONProcessorTypeFromValue returns a pointer to a valid ParseJSONProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewParseJSONProcessorTypeFromValue(v string) (*ParseJSONProcessorType, error) {
	ev := ParseJSONProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ParseJSONProcessorType: valid values are %v", v, allowedParseJSONProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ParseJSONProcessorType) IsValid() bool {
	for _, existing := range allowedParseJSONProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ParseJSONProcessorType value.
func (v ParseJSONProcessorType) Ptr() *ParseJSONProcessorType {
	return &v
}
