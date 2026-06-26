// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlDetectionResourceType Governance control detection resource type.
type GovernanceControlDetectionResourceType string

// List of GovernanceControlDetectionResourceType.
const (
	GOVERNANCECONTROLDETECTIONRESOURCETYPE_GOVERNANCE_CONTROL_DETECTION GovernanceControlDetectionResourceType = "governance_control_detection"
)

var allowedGovernanceControlDetectionResourceTypeEnumValues = []GovernanceControlDetectionResourceType{
	GOVERNANCECONTROLDETECTIONRESOURCETYPE_GOVERNANCE_CONTROL_DETECTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GovernanceControlDetectionResourceType) GetAllowedValues() []GovernanceControlDetectionResourceType {
	return allowedGovernanceControlDetectionResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GovernanceControlDetectionResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GovernanceControlDetectionResourceType(value)
	return nil
}

// NewGovernanceControlDetectionResourceTypeFromValue returns a pointer to a valid GovernanceControlDetectionResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGovernanceControlDetectionResourceTypeFromValue(v string) (*GovernanceControlDetectionResourceType, error) {
	ev := GovernanceControlDetectionResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GovernanceControlDetectionResourceType: valid values are %v", v, allowedGovernanceControlDetectionResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GovernanceControlDetectionResourceType) IsValid() bool {
	for _, existing := range allowedGovernanceControlDetectionResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GovernanceControlDetectionResourceType value.
func (v GovernanceControlDetectionResourceType) Ptr() *GovernanceControlDetectionResourceType {
	return &v
}
