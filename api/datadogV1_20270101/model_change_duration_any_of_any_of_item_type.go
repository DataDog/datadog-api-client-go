// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeDurationAnyOfAnyOfItemType Test-only generated schema for the dashboard versioning proof of concept.
type ChangeDurationAnyOfAnyOfItemType string

// List of ChangeDurationAnyOfAnyOfItemType.
const (
	CHANGEDURATIONANYOFANYOFITEMTYPE_CUSTOM_TIMEFRAME ChangeDurationAnyOfAnyOfItemType = "custom_timeframe"
)

var allowedChangeDurationAnyOfAnyOfItemTypeEnumValues = []ChangeDurationAnyOfAnyOfItemType{
	CHANGEDURATIONANYOFANYOFITEMTYPE_CUSTOM_TIMEFRAME,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ChangeDurationAnyOfAnyOfItemType) GetAllowedValues() []ChangeDurationAnyOfAnyOfItemType {
	return allowedChangeDurationAnyOfAnyOfItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ChangeDurationAnyOfAnyOfItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ChangeDurationAnyOfAnyOfItemType(value)
	return nil
}

// NewChangeDurationAnyOfAnyOfItemTypeFromValue returns a pointer to a valid ChangeDurationAnyOfAnyOfItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewChangeDurationAnyOfAnyOfItemTypeFromValue(v string) (*ChangeDurationAnyOfAnyOfItemType, error) {
	ev := ChangeDurationAnyOfAnyOfItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ChangeDurationAnyOfAnyOfItemType: valid values are %v", v, allowedChangeDurationAnyOfAnyOfItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ChangeDurationAnyOfAnyOfItemType) IsValid() bool {
	for _, existing := range allowedChangeDurationAnyOfAnyOfItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChangeDurationAnyOfAnyOfItemType value.
func (v ChangeDurationAnyOfAnyOfItemType) Ptr() *ChangeDurationAnyOfAnyOfItemType {
	return &v
}
