// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTextFormatReplaceSubstringType Replace a substring of a matching group label.
type HeatgridTextFormatReplaceSubstringType string

// List of HeatgridTextFormatReplaceSubstringType.
const (
	HEATGRIDTEXTFORMATREPLACESUBSTRINGTYPE_SUBSTRING HeatgridTextFormatReplaceSubstringType = "substring"
)

var allowedHeatgridTextFormatReplaceSubstringTypeEnumValues = []HeatgridTextFormatReplaceSubstringType{
	HEATGRIDTEXTFORMATREPLACESUBSTRINGTYPE_SUBSTRING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridTextFormatReplaceSubstringType) GetAllowedValues() []HeatgridTextFormatReplaceSubstringType {
	return allowedHeatgridTextFormatReplaceSubstringTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridTextFormatReplaceSubstringType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridTextFormatReplaceSubstringType(value)
	return nil
}

// NewHeatgridTextFormatReplaceSubstringTypeFromValue returns a pointer to a valid HeatgridTextFormatReplaceSubstringType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridTextFormatReplaceSubstringTypeFromValue(v string) (*HeatgridTextFormatReplaceSubstringType, error) {
	ev := HeatgridTextFormatReplaceSubstringType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridTextFormatReplaceSubstringType: valid values are %v", v, allowedHeatgridTextFormatReplaceSubstringTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridTextFormatReplaceSubstringType) IsValid() bool {
	for _, existing := range allowedHeatgridTextFormatReplaceSubstringTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridTextFormatReplaceSubstringType value.
func (v HeatgridTextFormatReplaceSubstringType) Ptr() *HeatgridTextFormatReplaceSubstringType {
	return &v
}
