// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventCreationSource Where the change was made: `UI` for the Datadog app, or `MCP` for a Model Context Protocol (MCP) client.
type FeatureFlagChangeEventCreationSource string

// List of FeatureFlagChangeEventCreationSource.
const (
	FEATUREFLAGCHANGEEVENTCREATIONSOURCE_UI  FeatureFlagChangeEventCreationSource = "UI"
	FEATUREFLAGCHANGEEVENTCREATIONSOURCE_MCP FeatureFlagChangeEventCreationSource = "MCP"
)

var allowedFeatureFlagChangeEventCreationSourceEnumValues = []FeatureFlagChangeEventCreationSource{
	FEATUREFLAGCHANGEEVENTCREATIONSOURCE_UI,
	FEATUREFLAGCHANGEEVENTCREATIONSOURCE_MCP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FeatureFlagChangeEventCreationSource) GetAllowedValues() []FeatureFlagChangeEventCreationSource {
	return allowedFeatureFlagChangeEventCreationSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeatureFlagChangeEventCreationSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeatureFlagChangeEventCreationSource(value)
	return nil
}

// NewFeatureFlagChangeEventCreationSourceFromValue returns a pointer to a valid FeatureFlagChangeEventCreationSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeatureFlagChangeEventCreationSourceFromValue(v string) (*FeatureFlagChangeEventCreationSource, error) {
	ev := FeatureFlagChangeEventCreationSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeatureFlagChangeEventCreationSource: valid values are %v", v, allowedFeatureFlagChangeEventCreationSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeatureFlagChangeEventCreationSource) IsValid() bool {
	for _, existing := range allowedFeatureFlagChangeEventCreationSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeatureFlagChangeEventCreationSource value.
func (v FeatureFlagChangeEventCreationSource) Ptr() *FeatureFlagChangeEventCreationSource {
	return &v
}

// NullableFeatureFlagChangeEventCreationSource handles when a null is used for FeatureFlagChangeEventCreationSource.
type NullableFeatureFlagChangeEventCreationSource struct {
	value *FeatureFlagChangeEventCreationSource
	isSet bool
}

// Get returns the associated value.
func (v NullableFeatureFlagChangeEventCreationSource) Get() *FeatureFlagChangeEventCreationSource {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableFeatureFlagChangeEventCreationSource) Set(val *FeatureFlagChangeEventCreationSource) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableFeatureFlagChangeEventCreationSource) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableFeatureFlagChangeEventCreationSource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFeatureFlagChangeEventCreationSource initializes the struct as if Set has been called.
func NewNullableFeatureFlagChangeEventCreationSource(val *FeatureFlagChangeEventCreationSource) *NullableFeatureFlagChangeEventCreationSource {
	return &NullableFeatureFlagChangeEventCreationSource{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableFeatureFlagChangeEventCreationSource) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableFeatureFlagChangeEventCreationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
