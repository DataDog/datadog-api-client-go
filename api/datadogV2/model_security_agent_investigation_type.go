// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationType Type of the investigation resource.
type SecurityAgentInvestigationType string

// List of SecurityAgentInvestigationType.
const (
	SECURITYAGENTINVESTIGATIONTYPE_INVESTIGATIONS SecurityAgentInvestigationType = "investigations"
)

var allowedSecurityAgentInvestigationTypeEnumValues = []SecurityAgentInvestigationType{
	SECURITYAGENTINVESTIGATIONTYPE_INVESTIGATIONS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentInvestigationType) GetAllowedValues() []SecurityAgentInvestigationType {
	return allowedSecurityAgentInvestigationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentInvestigationType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentInvestigationType(value)
	return nil
}

// NewSecurityAgentInvestigationTypeFromValue returns a pointer to a valid SecurityAgentInvestigationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentInvestigationTypeFromValue(v string) (*SecurityAgentInvestigationType, error) {
	ev := SecurityAgentInvestigationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentInvestigationType: valid values are %v", v, allowedSecurityAgentInvestigationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentInvestigationType) IsValid() bool {
	for _, existing := range allowedSecurityAgentInvestigationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentInvestigationType value.
func (v SecurityAgentInvestigationType) Ptr() *SecurityAgentInvestigationType {
	return &v
}
