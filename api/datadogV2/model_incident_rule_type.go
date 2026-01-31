// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleType Incident rule resource type.
type IncidentRuleType string

// List of IncidentRuleType.
const (
	INCIDENTRULETYPE_INCIDENT_RULE IncidentRuleType = "incident_rule"
)

var allowedIncidentRuleTypeEnumValues = []IncidentRuleType{
	INCIDENTRULETYPE_INCIDENT_RULE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRuleType) GetAllowedValues() []IncidentRuleType {
	return allowedIncidentRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRuleType(value)
	return nil
}

// NewIncidentRuleTypeFromValue returns a pointer to a valid IncidentRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRuleTypeFromValue(v string) (*IncidentRuleType, error) {
	ev := IncidentRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRuleType: valid values are %v", v, allowedIncidentRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRuleType) IsValid() bool {
	for _, existing := range allowedIncidentRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRuleType value.
func (v IncidentRuleType) Ptr() *IncidentRuleType {
	return &v
}
