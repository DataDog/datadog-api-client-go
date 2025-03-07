// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FilterProcessorType The type of processor.
type FilterProcessorType string

// List of FilterProcessorType.
const (
	FILTERPROCESSORTYPE_FILTER FilterProcessorType = "filter"
)

var allowedFilterProcessorTypeEnumValues = []FilterProcessorType{
	FILTERPROCESSORTYPE_FILTER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FilterProcessorType) GetAllowedValues() []FilterProcessorType {
	return allowedFilterProcessorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FilterProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FilterProcessorType(value)
	return nil
}

// NewFilterProcessorTypeFromValue returns a pointer to a valid FilterProcessorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFilterProcessorTypeFromValue(v string) (*FilterProcessorType, error) {
	ev := FilterProcessorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FilterProcessorType: valid values are %v", v, allowedFilterProcessorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FilterProcessorType) IsValid() bool {
	for _, existing := range allowedFilterProcessorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterProcessorType value.
func (v FilterProcessorType) Ptr() *FilterProcessorType {
	return &v
}
