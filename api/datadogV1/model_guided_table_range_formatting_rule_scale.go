// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRangeFormattingRuleScale Scale function for mapping values to colors.
type GuidedTableRangeFormattingRuleScale string

// List of GuidedTableRangeFormattingRuleScale.
const (
	GUIDEDTABLERANGEFORMATTINGRULESCALE_LINEAR GuidedTableRangeFormattingRuleScale = "linear"
	GUIDEDTABLERANGEFORMATTINGRULESCALE_LOG    GuidedTableRangeFormattingRuleScale = "log"
	GUIDEDTABLERANGEFORMATTINGRULESCALE_POW    GuidedTableRangeFormattingRuleScale = "pow"
	GUIDEDTABLERANGEFORMATTINGRULESCALE_SQRT   GuidedTableRangeFormattingRuleScale = "sqrt"
)

var allowedGuidedTableRangeFormattingRuleScaleEnumValues = []GuidedTableRangeFormattingRuleScale{
	GUIDEDTABLERANGEFORMATTINGRULESCALE_LINEAR,
	GUIDEDTABLERANGEFORMATTINGRULESCALE_LOG,
	GUIDEDTABLERANGEFORMATTINGRULESCALE_POW,
	GUIDEDTABLERANGEFORMATTINGRULESCALE_SQRT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GuidedTableRangeFormattingRuleScale) GetAllowedValues() []GuidedTableRangeFormattingRuleScale {
	return allowedGuidedTableRangeFormattingRuleScaleEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GuidedTableRangeFormattingRuleScale) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GuidedTableRangeFormattingRuleScale(value)
	return nil
}

// NewGuidedTableRangeFormattingRuleScaleFromValue returns a pointer to a valid GuidedTableRangeFormattingRuleScale
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGuidedTableRangeFormattingRuleScaleFromValue(v string) (*GuidedTableRangeFormattingRuleScale, error) {
	ev := GuidedTableRangeFormattingRuleScale(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GuidedTableRangeFormattingRuleScale: valid values are %v", v, allowedGuidedTableRangeFormattingRuleScaleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GuidedTableRangeFormattingRuleScale) IsValid() bool {
	for _, existing := range allowedGuidedTableRangeFormattingRuleScaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GuidedTableRangeFormattingRuleScale value.
func (v GuidedTableRangeFormattingRuleScale) Ptr() *GuidedTableRangeFormattingRuleScale {
	return &v
}
