// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestRestrictionPolicyBindingRelation The type of relation for the binding.
type SyntheticsTestRestrictionPolicyBindingRelation string

// List of SyntheticsTestRestrictionPolicyBindingRelation.
const (
	SYNTHETICSTESTRESTRICTIONPOLICYBINDINGRELATION_EDITOR SyntheticsTestRestrictionPolicyBindingRelation = "editor"
	SYNTHETICSTESTRESTRICTIONPOLICYBINDINGRELATION_VIEWER SyntheticsTestRestrictionPolicyBindingRelation = "viewer"
)

var allowedSyntheticsTestRestrictionPolicyBindingRelationEnumValues = []SyntheticsTestRestrictionPolicyBindingRelation{
	SYNTHETICSTESTRESTRICTIONPOLICYBINDINGRELATION_EDITOR,
	SYNTHETICSTESTRESTRICTIONPOLICYBINDINGRELATION_VIEWER,
}

// GetAllowedValues returns the list of possible values.
func (v *SyntheticsTestRestrictionPolicyBindingRelation) GetAllowedValues() []SyntheticsTestRestrictionPolicyBindingRelation {
	return allowedSyntheticsTestRestrictionPolicyBindingRelationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SyntheticsTestRestrictionPolicyBindingRelation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SyntheticsTestRestrictionPolicyBindingRelation(value)
	return nil
}

// NewSyntheticsTestRestrictionPolicyBindingRelationFromValue returns a pointer to a valid SyntheticsTestRestrictionPolicyBindingRelation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSyntheticsTestRestrictionPolicyBindingRelationFromValue(v string) (*SyntheticsTestRestrictionPolicyBindingRelation, error) {
	ev := SyntheticsTestRestrictionPolicyBindingRelation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SyntheticsTestRestrictionPolicyBindingRelation: valid values are %v", v, allowedSyntheticsTestRestrictionPolicyBindingRelationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SyntheticsTestRestrictionPolicyBindingRelation) IsValid() bool {
	for _, existing := range allowedSyntheticsTestRestrictionPolicyBindingRelationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyntheticsTestRestrictionPolicyBindingRelation value.
func (v SyntheticsTestRestrictionPolicyBindingRelation) Ptr() *SyntheticsTestRestrictionPolicyBindingRelation {
	return &v
}
