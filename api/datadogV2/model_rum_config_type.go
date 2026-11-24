// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMConfigType The type of the resource. The value should always be `rum_config`.
type RUMConfigType string

// List of RUMConfigType.
const (
	RUMCONFIGTYPE_RUM_CONFIG RUMConfigType = "rum_config"
)

var allowedRUMConfigTypeEnumValues = []RUMConfigType{
	RUMCONFIGTYPE_RUM_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMConfigType) GetAllowedValues() []RUMConfigType {
	return allowedRUMConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMConfigType(value)
	return nil
}

// NewRUMConfigTypeFromValue returns a pointer to a valid RUMConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMConfigTypeFromValue(v string) (*RUMConfigType, error) {
	ev := RUMConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMConfigType: valid values are %v", v, allowedRUMConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMConfigType) IsValid() bool {
	for _, existing := range allowedRUMConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMConfigType value.
func (v RUMConfigType) Ptr() *RUMConfigType {
	return &v
}
