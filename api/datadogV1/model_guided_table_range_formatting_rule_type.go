// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRangeFormattingRuleType
type GuidedTableRangeFormattingRuleType string

// List of GuidedTableRangeFormattingRuleType.
const (
	GUIDEDTABLERANGEFORMATTINGRULETYPE_RANGE GuidedTableRangeFormattingRuleType = "range"
)

var allowedGuidedTableRangeFormattingRuleTypeEnumValues = []GuidedTableRangeFormattingRuleType{
	GUIDEDTABLERANGEFORMATTINGRULETYPE_RANGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableRangeFormattingRuleType) GetAllowedValues() []GuidedTableRangeFormattingRuleType {
	return allowedGuidedTableRangeFormattingRuleTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableRangeFormattingRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableRangeFormattingRuleType(value)
	return nil
}

// NewGuidedTableRangeFormattingRuleTypeFromValue returns a pointer to a valid GuidedTableRangeFormattingRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableRangeFormattingRuleTypeFromValue(v string) (*GuidedTableRangeFormattingRuleType, error) {
	ev := GuidedTableRangeFormattingRuleType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableRangeFormattingRuleType: valid values are %v", v, allowedGuidedTableRangeFormattingRuleTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableRangeFormattingRuleType) IsValid() bool {
	for _, existing := range allowedGuidedTableRangeFormattingRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableRangeFormattingRuleType value.
func (v GuidedTableRangeFormattingRuleType) Ptr() *GuidedTableRangeFormattingRuleType {
	return &v
}
