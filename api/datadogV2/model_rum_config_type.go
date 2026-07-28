// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumConfigType The type of the resource. The value should always be `rum_config`.
type RumConfigType string

// List of RumConfigType.
const (
	RUMCONFIGTYPE_RUM_CONFIG RumConfigType = "rum_config"
)

var allowedRumConfigTypeEnumValues = []RumConfigType{
	RUMCONFIGTYPE_RUM_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RumConfigType) GetAllowedValues() []RumConfigType {
	return allowedRumConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RumConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RumConfigType(value)
	return nil
}

// NewRumConfigTypeFromValue returns a pointer to a valid RumConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRumConfigTypeFromValue(v string) (*RumConfigType, error) {
	ev := RumConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RumConfigType: valid values are %v", v, allowedRumConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RumConfigType) IsValid() bool {
	for _, existing := range allowedRumConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RumConfigType value.
func (v RumConfigType) Ptr() *RumConfigType {
	return &v
}
