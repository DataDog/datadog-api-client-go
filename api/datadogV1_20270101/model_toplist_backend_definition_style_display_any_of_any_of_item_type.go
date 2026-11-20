// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType string

// List of ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType.
const (
	TOPLISTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE_STACKED ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType = "stacked"
)

var allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues = []ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType{
	TOPLISTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE_STACKED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType) GetAllowedValues() []ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType {
	return allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType(value)
	return nil
}

// NewToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeFromValue returns a pointer to a valid ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeFromValue(v string) (*ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType, error) {
	ev := ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType: valid values are %v", v, allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType) IsValid() bool {
	for _, existing := range allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType value.
func (v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType) Ptr() *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType {
	return &v
}
