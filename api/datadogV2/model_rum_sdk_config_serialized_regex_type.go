// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigSerializedRegexType The type identifier for a serialized regex. Always `regex`.
type RUMSdkConfigSerializedRegexType string

// List of RUMSdkConfigSerializedRegexType.
const (
	RUMSDKCONFIGSERIALIZEDREGEXTYPE_REGEX RUMSdkConfigSerializedRegexType = "regex"
)

var allowedRUMSdkConfigSerializedRegexTypeEnumValues = []RUMSdkConfigSerializedRegexType{
	RUMSDKCONFIGSERIALIZEDREGEXTYPE_REGEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMSdkConfigSerializedRegexType) GetAllowedValues() []RUMSdkConfigSerializedRegexType {
	return allowedRUMSdkConfigSerializedRegexTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMSdkConfigSerializedRegexType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMSdkConfigSerializedRegexType(value)
	return nil
}

// NewRUMSdkConfigSerializedRegexTypeFromValue returns a pointer to a valid RUMSdkConfigSerializedRegexType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMSdkConfigSerializedRegexTypeFromValue(v string) (*RUMSdkConfigSerializedRegexType, error) {
	ev := RUMSdkConfigSerializedRegexType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMSdkConfigSerializedRegexType: valid values are %v", v, allowedRUMSdkConfigSerializedRegexTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMSdkConfigSerializedRegexType) IsValid() bool {
	for _, existing := range allowedRUMSdkConfigSerializedRegexTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMSdkConfigSerializedRegexType value.
func (v RUMSdkConfigSerializedRegexType) Ptr() *RUMSdkConfigSerializedRegexType {
	return &v
}
