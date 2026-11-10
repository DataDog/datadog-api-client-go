// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleTagValueTransformType The kind of transform.
type ArbitraryRuleTagValueTransformType string

// List of ArbitraryRuleTagValueTransformType.
const (
	ARBITRARYRULETAGVALUETRANSFORMTYPE_STRING_REPLACE ArbitraryRuleTagValueTransformType = "string_replace"
	ARBITRARYRULETAGVALUETRANSFORMTYPE_REGEX_EXTRACT  ArbitraryRuleTagValueTransformType = "regex_extract"
	ARBITRARYRULETAGVALUETRANSFORMTYPE_REGEX_REPLACE  ArbitraryRuleTagValueTransformType = "regex_replace"
	ARBITRARYRULETAGVALUETRANSFORMTYPE_LOOKUP_MAP     ArbitraryRuleTagValueTransformType = "lookup_map"
)

var allowedArbitraryRuleTagValueTransformTypeEnumValues = []ArbitraryRuleTagValueTransformType{
	ARBITRARYRULETAGVALUETRANSFORMTYPE_STRING_REPLACE,
	ARBITRARYRULETAGVALUETRANSFORMTYPE_REGEX_EXTRACT,
	ARBITRARYRULETAGVALUETRANSFORMTYPE_REGEX_REPLACE,
	ARBITRARYRULETAGVALUETRANSFORMTYPE_LOOKUP_MAP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ArbitraryRuleTagValueTransformType) GetAllowedValues() []ArbitraryRuleTagValueTransformType {
	return allowedArbitraryRuleTagValueTransformTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ArbitraryRuleTagValueTransformType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ArbitraryRuleTagValueTransformType(value)
	return nil
}

// NewArbitraryRuleTagValueTransformTypeFromValue returns a pointer to a valid ArbitraryRuleTagValueTransformType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewArbitraryRuleTagValueTransformTypeFromValue(v string) (*ArbitraryRuleTagValueTransformType, error) {
	ev := ArbitraryRuleTagValueTransformType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ArbitraryRuleTagValueTransformType: valid values are %v", v, allowedArbitraryRuleTagValueTransformTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ArbitraryRuleTagValueTransformType) IsValid() bool {
	for _, existing := range allowedArbitraryRuleTagValueTransformTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ArbitraryRuleTagValueTransformType value.
func (v ArbitraryRuleTagValueTransformType) Ptr() *ArbitraryRuleTagValueTransformType {
	return &v
}
