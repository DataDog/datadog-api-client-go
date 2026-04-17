// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlagSuggestionProperty The flag property being changed.
type FlagSuggestionProperty string

// List of FlagSuggestionProperty.
const (
	FLAGSUGGESTIONPROPERTY_FLAG                 FlagSuggestionProperty = "FLAG"
	FLAGSUGGESTIONPROPERTY_FLAG_NAME            FlagSuggestionProperty = "FLAG_NAME"
	FLAGSUGGESTIONPROPERTY_FLAG_DESCRIPTION     FlagSuggestionProperty = "FLAG_DESCRIPTION"
	FLAGSUGGESTIONPROPERTY_JSON_SCHEMA          FlagSuggestionProperty = "JSON_SCHEMA"
	FLAGSUGGESTIONPROPERTY_DISTRIBUTION_CHANNEL FlagSuggestionProperty = "DISTRIBUTION_CHANNEL"
	FLAGSUGGESTIONPROPERTY_VARIANT              FlagSuggestionProperty = "VARIANT"
	FLAGSUGGESTIONPROPERTY_VARIANT_NAME         FlagSuggestionProperty = "VARIANT_NAME"
	FLAGSUGGESTIONPROPERTY_VARIANT_VALUE        FlagSuggestionProperty = "VARIANT_VALUE"
	FLAGSUGGESTIONPROPERTY_ALLOCATIONS          FlagSuggestionProperty = "ALLOCATIONS"
	FLAGSUGGESTIONPROPERTY_ROLLOUT              FlagSuggestionProperty = "ROLLOUT"
	FLAGSUGGESTIONPROPERTY_ENVIRONMENT_STATUS   FlagSuggestionProperty = "ENVIRONMENT_STATUS"
	FLAGSUGGESTIONPROPERTY_DEFAULT_VARIANT      FlagSuggestionProperty = "DEFAULT_VARIANT"
	FLAGSUGGESTIONPROPERTY_OVERRIDE_VARIANT     FlagSuggestionProperty = "OVERRIDE_VARIANT"
)

var allowedFlagSuggestionPropertyEnumValues = []FlagSuggestionProperty{
	FLAGSUGGESTIONPROPERTY_FLAG,
	FLAGSUGGESTIONPROPERTY_FLAG_NAME,
	FLAGSUGGESTIONPROPERTY_FLAG_DESCRIPTION,
	FLAGSUGGESTIONPROPERTY_JSON_SCHEMA,
	FLAGSUGGESTIONPROPERTY_DISTRIBUTION_CHANNEL,
	FLAGSUGGESTIONPROPERTY_VARIANT,
	FLAGSUGGESTIONPROPERTY_VARIANT_NAME,
	FLAGSUGGESTIONPROPERTY_VARIANT_VALUE,
	FLAGSUGGESTIONPROPERTY_ALLOCATIONS,
	FLAGSUGGESTIONPROPERTY_ROLLOUT,
	FLAGSUGGESTIONPROPERTY_ENVIRONMENT_STATUS,
	FLAGSUGGESTIONPROPERTY_DEFAULT_VARIANT,
	FLAGSUGGESTIONPROPERTY_OVERRIDE_VARIANT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FlagSuggestionProperty) GetAllowedValues() []FlagSuggestionProperty {
	return allowedFlagSuggestionPropertyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FlagSuggestionProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FlagSuggestionProperty(value)
	return nil
}

// NewFlagSuggestionPropertyFromValue returns a pointer to a valid FlagSuggestionProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFlagSuggestionPropertyFromValue(v string) (*FlagSuggestionProperty, error) {
	ev := FlagSuggestionProperty(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FlagSuggestionProperty: valid values are %v", v, allowedFlagSuggestionPropertyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FlagSuggestionProperty) IsValid() bool {
	for _, existing := range allowedFlagSuggestionPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlagSuggestionProperty value.
func (v FlagSuggestionProperty) Ptr() *FlagSuggestionProperty {
	return &v
}
