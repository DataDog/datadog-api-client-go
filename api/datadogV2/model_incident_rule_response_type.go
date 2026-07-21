// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleResponseType Incident rule response resource type.
type IncidentRuleResponseType string

// List of IncidentRuleResponseType.
const (
	INCIDENTRULERESPONSETYPE_INCIDENTS_RULES IncidentRuleResponseType = "incidents_rules"
)

var allowedIncidentRuleResponseTypeEnumValues = []IncidentRuleResponseType{
	INCIDENTRULERESPONSETYPE_INCIDENTS_RULES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRuleResponseType) GetAllowedValues() []IncidentRuleResponseType {
	return allowedIncidentRuleResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRuleResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRuleResponseType(value)
	return nil
}

// NewIncidentRuleResponseTypeFromValue returns a pointer to a valid IncidentRuleResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRuleResponseTypeFromValue(v string) (*IncidentRuleResponseType, error) {
	ev := IncidentRuleResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRuleResponseType: valid values are %v", v, allowedIncidentRuleResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRuleResponseType) IsValid() bool {
	for _, existing := range allowedIncidentRuleResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRuleResponseType value.
func (v IncidentRuleResponseType) Ptr() *IncidentRuleResponseType {
	return &v
}
