// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaSemanticMode Test-only generated schema for the dashboard versioning proof of concept.
type FormulaSemanticMode string

// List of FormulaSemanticMode.
const (
	FORMULASEMANTICMODE_NATIVE   FormulaSemanticMode = "native"
	FORMULASEMANTICMODE_COMBINED FormulaSemanticMode = "combined"
)

var allowedFormulaSemanticModeEnumValues = []FormulaSemanticMode{
	FORMULASEMANTICMODE_NATIVE,
	FORMULASEMANTICMODE_COMBINED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *FormulaSemanticMode) GetAllowedValues() []FormulaSemanticMode {
	return allowedFormulaSemanticModeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FormulaSemanticMode) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FormulaSemanticMode(value)
	return nil
}

// NewFormulaSemanticModeFromValue returns a pointer to a valid FormulaSemanticMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFormulaSemanticModeFromValue(v string) (*FormulaSemanticMode, error) {
	ev := FormulaSemanticMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FormulaSemanticMode: valid values are %v", v, allowedFormulaSemanticModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FormulaSemanticMode) IsValid() bool {
	for _, existing := range allowedFormulaSemanticModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormulaSemanticMode value.
func (v FormulaSemanticMode) Ptr() *FormulaSemanticMode {
	return &v
}
