// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeType Test-only generated schema for the dashboard versioning proof of concept.
type ChangeType string

// List of ChangeType.
const (
	CHANGETYPE_BOTH     ChangeType = "both"
	CHANGETYPE_RELATIVE ChangeType = "relative"
	CHANGETYPE_ABSOLUTE ChangeType = "absolute"
)

var allowedChangeTypeEnumValues = []ChangeType{
	CHANGETYPE_BOTH,
	CHANGETYPE_RELATIVE,
	CHANGETYPE_ABSOLUTE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeType) GetAllowedValues() []ChangeType {
	return allowedChangeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeType(value)
	return nil
}

// NewChangeTypeFromValue returns a pointer to a valid ChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeTypeFromValue(v string) (*ChangeType, error) {
	ev := ChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeType: valid values are %v", v, allowedChangeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeType) IsValid() bool {
	for _, existing := range allowedChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeType value.
func (v ChangeType) Ptr() *ChangeType {
	return &v
}
