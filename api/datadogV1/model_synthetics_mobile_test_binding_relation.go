// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestBindingRelation The type of relation for the binding.
type SyntheticsMobileTestBindingRelation string

// List of SyntheticsMobileTestBindingRelation.
const (
	SYNTHETICSMOBILETESTBINDINGRELATION_EDITOR SyntheticsMobileTestBindingRelation = "editor"
	SYNTHETICSMOBILETESTBINDINGRELATION_VIEWER SyntheticsMobileTestBindingRelation = "viewer"
)

var allowedSyntheticsMobileTestBindingRelationEnumValues = []SyntheticsMobileTestBindingRelation{
	SYNTHETICSMOBILETESTBINDINGRELATION_EDITOR,
	SYNTHETICSMOBILETESTBINDINGRELATION_VIEWER,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SyntheticsMobileTestBindingRelation) GetAllowedValues() []SyntheticsMobileTestBindingRelation {
	return allowedSyntheticsMobileTestBindingRelationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsMobileTestBindingRelation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsMobileTestBindingRelation(value)
	return nil
}

// NewSyntheticsMobileTestBindingRelationFromValue returns a pointer to a valid SyntheticsMobileTestBindingRelation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsMobileTestBindingRelationFromValue(v string) (*SyntheticsMobileTestBindingRelation, error) {
	ev := SyntheticsMobileTestBindingRelation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsMobileTestBindingRelation: valid values are %v", v, allowedSyntheticsMobileTestBindingRelationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsMobileTestBindingRelation) IsValid() bool {
	for _, existing := range allowedSyntheticsMobileTestBindingRelationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsMobileTestBindingRelation value.
func (v SyntheticsMobileTestBindingRelation) Ptr() *SyntheticsMobileTestBindingRelation {
	return &v
}
