// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448 Test-only generated schema for the dashboard versioning proof of concept.
type ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448 string

// List of ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448.
const (
	TOPLISTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE4453028448_FLAT ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448 = "flat"
)

var allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448EnumValues = []ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448{
	TOPLISTBACKENDDEFINITIONSTYLEDISPLAYANYOFANYOFITEMTYPE4453028448_FLAT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448) GetAllowedValues() []ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448 {
	return allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448(value)
	return nil
}

// NewToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448FromValue returns a pointer to a valid ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448FromValue(v string) (*ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448, error) {
	ev := ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448: valid values are %v", v, allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448) IsValid() bool {
	for _, existing := range allowedToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448 value.
func (v ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448) Ptr() *ToplistBackendDefinitionStyleDisplayAnyOfAnyOfItemType4453028448 {
	return &v
}
