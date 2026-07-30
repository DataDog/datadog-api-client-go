// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleExecutionType The execution type of an incident rule.
type IncidentRuleExecutionType int64

// List of IncidentRuleExecutionType.
const (
	INCIDENTRULEEXECUTIONTYPE_SINGLE_EXECUTION IncidentRuleExecutionType = 1
	INCIDENTRULEEXECUTIONTYPE_MULTI_EXECUTION  IncidentRuleExecutionType = 2
)

var allowedIncidentRuleExecutionTypeEnumValues = []IncidentRuleExecutionType{
	INCIDENTRULEEXECUTIONTYPE_SINGLE_EXECUTION,
	INCIDENTRULEEXECUTIONTYPE_MULTI_EXECUTION,
}

// GetAllowedValues reeturns the list of possible values.
func (v *IncidentRuleExecutionType) GetAllowedValues() []IncidentRuleExecutionType {
	return allowedIncidentRuleExecutionTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IncidentRuleExecutionType) UnmarshalJSON(src []byte) error {
	var value int64
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IncidentRuleExecutionType(value)
	return nil
}

// NewIncidentRuleExecutionTypeFromValue returns a pointer to a valid IncidentRuleExecutionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIncidentRuleExecutionTypeFromValue(v int64) (*IncidentRuleExecutionType, error) {
	ev := IncidentRuleExecutionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IncidentRuleExecutionType: valid values are %v", v, allowedIncidentRuleExecutionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IncidentRuleExecutionType) IsValid() bool {
	for _, existing := range allowedIncidentRuleExecutionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncidentRuleExecutionType value.
func (v IncidentRuleExecutionType) Ptr() *IncidentRuleExecutionType {
	return &v
}
