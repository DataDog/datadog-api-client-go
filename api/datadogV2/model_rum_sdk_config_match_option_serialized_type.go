// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigMatchOptionSerializedType The type of match pattern, either a literal string or a regex.
type RUMSdkConfigMatchOptionSerializedType string

// List of RUMSdkConfigMatchOptionSerializedType.
const (
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING RUMSdkConfigMatchOptionSerializedType = "string"
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_REGEX  RUMSdkConfigMatchOptionSerializedType = "regex"
)

var allowedRUMSdkConfigMatchOptionSerializedTypeEnumValues = []RUMSdkConfigMatchOptionSerializedType{
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_STRING,
	RUMSDKCONFIGMATCHOPTIONSERIALIZEDTYPE_REGEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMSdkConfigMatchOptionSerializedType) GetAllowedValues() []RUMSdkConfigMatchOptionSerializedType {
	return allowedRUMSdkConfigMatchOptionSerializedTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMSdkConfigMatchOptionSerializedType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMSdkConfigMatchOptionSerializedType(value)
	return nil
}

// NewRUMSdkConfigMatchOptionSerializedTypeFromValue returns a pointer to a valid RUMSdkConfigMatchOptionSerializedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMSdkConfigMatchOptionSerializedTypeFromValue(v string) (*RUMSdkConfigMatchOptionSerializedType, error) {
	ev := RUMSdkConfigMatchOptionSerializedType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMSdkConfigMatchOptionSerializedType: valid values are %v", v, allowedRUMSdkConfigMatchOptionSerializedTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMSdkConfigMatchOptionSerializedType) IsValid() bool {
	for _, existing := range allowedRUMSdkConfigMatchOptionSerializedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMSdkConfigMatchOptionSerializedType value.
func (v RUMSdkConfigMatchOptionSerializedType) Ptr() *RUMSdkConfigMatchOptionSerializedType {
	return &v
}
