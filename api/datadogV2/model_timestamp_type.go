// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimestampType The type of timestamp being overridden.
type TimestampType string

// List of TimestampType.
const (
	TIMESTAMPTYPE_CREATED  TimestampType = "created"
	TIMESTAMPTYPE_DETECTED TimestampType = "detected"
	TIMESTAMPTYPE_RESOLVED TimestampType = "resolved"
	TIMESTAMPTYPE_DECLARED TimestampType = "declared"
)

var allowedTimestampTypeEnumValues = []TimestampType{
	TIMESTAMPTYPE_CREATED,
	TIMESTAMPTYPE_DETECTED,
	TIMESTAMPTYPE_RESOLVED,
	TIMESTAMPTYPE_DECLARED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TimestampType) GetAllowedValues() []TimestampType {
	return allowedTimestampTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TimestampType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TimestampType(value)
	return nil
}

// NewTimestampTypeFromValue returns a pointer to a valid TimestampType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTimestampTypeFromValue(v string) (*TimestampType, error) {
	ev := TimestampType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TimestampType: valid values are %v", v, allowedTimestampTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TimestampType) IsValid() bool {
	for _, existing := range allowedTimestampTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimestampType value.
func (v TimestampType) Ptr() *TimestampType {
	return &v
}
