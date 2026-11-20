// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigType The type of the resource. The value should always be `rum_sdk_config`.
type RUMSdkConfigType string

// List of RUMSdkConfigType.
const (
	RUMSDKCONFIGTYPE_RUM_SDK_CONFIG RUMSdkConfigType = "rum_sdk_config"
)

var allowedRUMSdkConfigTypeEnumValues = []RUMSdkConfigType{
	RUMSDKCONFIGTYPE_RUM_SDK_CONFIG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMSdkConfigType) GetAllowedValues() []RUMSdkConfigType {
	return allowedRUMSdkConfigTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMSdkConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMSdkConfigType(value)
	return nil
}

// NewRUMSdkConfigTypeFromValue returns a pointer to a valid RUMSdkConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMSdkConfigTypeFromValue(v string) (*RUMSdkConfigType, error) {
	ev := RUMSdkConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMSdkConfigType: valid values are %v", v, allowedRUMSdkConfigTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMSdkConfigType) IsValid() bool {
	for _, existing := range allowedRUMSdkConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMSdkConfigType value.
func (v RUMSdkConfigType) Ptr() *RUMSdkConfigType {
	return &v
}
