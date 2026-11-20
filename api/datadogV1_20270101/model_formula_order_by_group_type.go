// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaOrderByGroupType Test-only generated schema for the dashboard versioning proof of concept.
type FormulaOrderByGroupType string

// List of FormulaOrderByGroupType.
const (
	FORMULAORDERBYGROUPTYPE_GROUP FormulaOrderByGroupType = "group"
)

var allowedFormulaOrderByGroupTypeEnumValues = []FormulaOrderByGroupType{
	FORMULAORDERBYGROUPTYPE_GROUP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaOrderByGroupType) GetAllowedValues() []FormulaOrderByGroupType {
	return allowedFormulaOrderByGroupTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaOrderByGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaOrderByGroupType(value)
	return nil
}

// NewFormulaOrderByGroupTypeFromValue returns a pointer to a valid FormulaOrderByGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaOrderByGroupTypeFromValue(v string) (*FormulaOrderByGroupType, error) {
	ev := FormulaOrderByGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaOrderByGroupType: valid values are %v", v, allowedFormulaOrderByGroupTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaOrderByGroupType) IsValid() bool {
	for _, existing := range allowedFormulaOrderByGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaOrderByGroupType value.
func (v FormulaOrderByGroupType) Ptr() *FormulaOrderByGroupType {
	return &v
}
