// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationCountsType Type of the investigation counts resource.
type SecurityAgentInvestigationCountsType string

// List of SecurityAgentInvestigationCountsType.
const (
	SECURITYAGENTINVESTIGATIONCOUNTSTYPE_INVESTIGATION_COUNTS SecurityAgentInvestigationCountsType = "investigation_counts"
)

var allowedSecurityAgentInvestigationCountsTypeEnumValues = []SecurityAgentInvestigationCountsType{
	SECURITYAGENTINVESTIGATIONCOUNTSTYPE_INVESTIGATION_COUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityAgentInvestigationCountsType) GetAllowedValues() []SecurityAgentInvestigationCountsType {
	return allowedSecurityAgentInvestigationCountsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityAgentInvestigationCountsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityAgentInvestigationCountsType(value)
	return nil
}

// NewSecurityAgentInvestigationCountsTypeFromValue returns a pointer to a valid SecurityAgentInvestigationCountsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityAgentInvestigationCountsTypeFromValue(v string) (*SecurityAgentInvestigationCountsType, error) {
	ev := SecurityAgentInvestigationCountsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityAgentInvestigationCountsType: valid values are %v", v, allowedSecurityAgentInvestigationCountsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityAgentInvestigationCountsType) IsValid() bool {
	for _, existing := range allowedSecurityAgentInvestigationCountsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityAgentInvestigationCountsType value.
func (v SecurityAgentInvestigationCountsType) Ptr() *SecurityAgentInvestigationCountsType {
	return &v
}
