// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomAllocationRuleStatusType Custom allocation rule status resource type.
type CustomAllocationRuleStatusType string

// List of CustomAllocationRuleStatusType.
const (
	CUSTOMALLOCATIONRULESTATUSTYPE_ARBITRARY_RULE_STATUS CustomAllocationRuleStatusType = "arbitrary_rule_status"
)

var allowedCustomAllocationRuleStatusTypeEnumValues = []CustomAllocationRuleStatusType{
	CUSTOMALLOCATIONRULESTATUSTYPE_ARBITRARY_RULE_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomAllocationRuleStatusType) GetAllowedValues() []CustomAllocationRuleStatusType {
	return allowedCustomAllocationRuleStatusTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomAllocationRuleStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomAllocationRuleStatusType(value)
	return nil
}

// NewCustomAllocationRuleStatusTypeFromValue returns a pointer to a valid CustomAllocationRuleStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomAllocationRuleStatusTypeFromValue(v string) (*CustomAllocationRuleStatusType, error) {
	ev := CustomAllocationRuleStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomAllocationRuleStatusType: valid values are %v", v, allowedCustomAllocationRuleStatusTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomAllocationRuleStatusType) IsValid() bool {
	for _, existing := range allowedCustomAllocationRuleStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomAllocationRuleStatusType value.
func (v CustomAllocationRuleStatusType) Ptr() *CustomAllocationRuleStatusType {
	return &v
}
