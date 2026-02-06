// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationAssignmentType Integration assignment resource type.
type IntegrationAssignmentType string

// List of IntegrationAssignmentType.
const (
	INTEGRATIONASSIGNMENTTYPE_ISSUE_ASSIGNMENT IntegrationAssignmentType = "issue_assignment"
)

var allowedIntegrationAssignmentTypeEnumValues = []IntegrationAssignmentType{
	INTEGRATIONASSIGNMENTTYPE_ISSUE_ASSIGNMENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IntegrationAssignmentType) GetAllowedValues() []IntegrationAssignmentType {
	return allowedIntegrationAssignmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IntegrationAssignmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IntegrationAssignmentType(value)
	return nil
}

// NewIntegrationAssignmentTypeFromValue returns a pointer to a valid IntegrationAssignmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIntegrationAssignmentTypeFromValue(v string) (*IntegrationAssignmentType, error) {
	ev := IntegrationAssignmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IntegrationAssignmentType: valid values are %v", v, allowedIntegrationAssignmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IntegrationAssignmentType) IsValid() bool {
	for _, existing := range allowedIntegrationAssignmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationAssignmentType value.
func (v IntegrationAssignmentType) Ptr() *IntegrationAssignmentType {
	return &v
}
