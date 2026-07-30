// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXIndicatorType The STIX object type for an indicator.
type STIXIndicatorType string

// List of STIXIndicatorType.
const (
	STIXINDICATORTYPE_INDICATOR STIXIndicatorType = "indicator"
)

var allowedSTIXIndicatorTypeEnumValues = []STIXIndicatorType{
	STIXINDICATORTYPE_INDICATOR,
}

// GetAllowedValues reeturns the list of possible values.
func (v *STIXIndicatorType) GetAllowedValues() []STIXIndicatorType {
	return allowedSTIXIndicatorTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *STIXIndicatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = STIXIndicatorType(value)
	return nil
}

// NewSTIXIndicatorTypeFromValue returns a pointer to a valid STIXIndicatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSTIXIndicatorTypeFromValue(v string) (*STIXIndicatorType, error) {
	ev := STIXIndicatorType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for STIXIndicatorType: valid values are %v", v, allowedSTIXIndicatorTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v STIXIndicatorType) IsValid() bool {
	for _, existing := range allowedSTIXIndicatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to STIXIndicatorType value.
func (v STIXIndicatorType) Ptr() *STIXIndicatorType {
	return &v
}
