// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReturnPeriodScopeTargetType Test-only generated schema for the dashboard versioning proof of concept.
type ReturnPeriodScopeTargetType string

// List of ReturnPeriodScopeTargetType.
const (
	RETURNPERIODSCOPETARGETTYPE_INDEX ReturnPeriodScopeTargetType = "index"
)

var allowedReturnPeriodScopeTargetTypeEnumValues = []ReturnPeriodScopeTargetType{
	RETURNPERIODSCOPETARGETTYPE_INDEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ReturnPeriodScopeTargetType) GetAllowedValues() []ReturnPeriodScopeTargetType {
	return allowedReturnPeriodScopeTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ReturnPeriodScopeTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ReturnPeriodScopeTargetType(value)
	return nil
}

// NewReturnPeriodScopeTargetTypeFromValue returns a pointer to a valid ReturnPeriodScopeTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewReturnPeriodScopeTargetTypeFromValue(v string) (*ReturnPeriodScopeTargetType, error) {
	ev := ReturnPeriodScopeTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ReturnPeriodScopeTargetType: valid values are %v", v, allowedReturnPeriodScopeTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ReturnPeriodScopeTargetType) IsValid() bool {
	for _, existing := range allowedReturnPeriodScopeTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnPeriodScopeTargetType value.
func (v ReturnPeriodScopeTargetType) Ptr() *ReturnPeriodScopeTargetType {
	return &v
}
