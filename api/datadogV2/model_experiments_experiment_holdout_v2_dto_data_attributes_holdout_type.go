// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType How Datadog serves a MANAGED holdout. GLOBAL holdouts apply to every eligible experiment of a subject type. SELECTIVE holdouts apply only to the experiments attached to them. Null for EXTERNAL holdouts.
type ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType string

// List of ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType.
const (
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESHOLDOUTTYPE_GLOBAL    ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType = "GLOBAL"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESHOLDOUTTYPE_SELECTIVE ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType = "SELECTIVE"
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESHOLDOUTTYPE_UNKNOWN   ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType = "UNKNOWN"
)

var allowedExperimentsExperimentHoldoutV2DTODataAttributesHoldoutTypeEnumValues = []ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType{
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESHOLDOUTTYPE_GLOBAL,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESHOLDOUTTYPE_SELECTIVE,
	EXPERIMENTSEXPERIMENTHOLDOUTV2DTODATAATTRIBUTESHOLDOUTTYPE_UNKNOWN,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) GetAllowedValues() []ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType {
	return allowedExperimentsExperimentHoldoutV2DTODataAttributesHoldoutTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType(value)
	return nil
}

// NewExperimentsExperimentHoldoutV2DTODataAttributesHoldoutTypeFromValue returns a pointer to a valid ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewExperimentsExperimentHoldoutV2DTODataAttributesHoldoutTypeFromValue(v string) (*ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType, error) {
	ev := ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType: valid values are %v", v, allowedExperimentsExperimentHoldoutV2DTODataAttributesHoldoutTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) IsValid() bool {
	for _, existing := range allowedExperimentsExperimentHoldoutV2DTODataAttributesHoldoutTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType value.
func (v ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) Ptr() *ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType {
	return &v
}

// NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType handles when a null is used for ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType.
type NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType struct {
	value *ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType
	isSet bool
}

// Get returns the associated value.
func (v NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) Get() *ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) Set(val *ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType initializes the struct as if Set has been called.
func NewNullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType(val *ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) *NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType {
	return &NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return datadog.Unmarshal(src, &v.value)
}
