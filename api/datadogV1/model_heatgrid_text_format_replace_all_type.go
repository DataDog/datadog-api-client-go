// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTextFormatReplaceAllType Replace the entire matching group label.
type HeatgridTextFormatReplaceAllType string

// List of HeatgridTextFormatReplaceAllType.
const (
	HEATGRIDTEXTFORMATREPLACEALLTYPE_ALL HeatgridTextFormatReplaceAllType = "all"
)

var allowedHeatgridTextFormatReplaceAllTypeEnumValues = []HeatgridTextFormatReplaceAllType{
	HEATGRIDTEXTFORMATREPLACEALLTYPE_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridTextFormatReplaceAllType) GetAllowedValues() []HeatgridTextFormatReplaceAllType {
	return allowedHeatgridTextFormatReplaceAllTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridTextFormatReplaceAllType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridTextFormatReplaceAllType(value)
	return nil
}

// NewHeatgridTextFormatReplaceAllTypeFromValue returns a pointer to a valid HeatgridTextFormatReplaceAllType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridTextFormatReplaceAllTypeFromValue(v string) (*HeatgridTextFormatReplaceAllType, error) {
	ev := HeatgridTextFormatReplaceAllType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridTextFormatReplaceAllType: valid values are %v", v, allowedHeatgridTextFormatReplaceAllTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridTextFormatReplaceAllType) IsValid() bool {
	for _, existing := range allowedHeatgridTextFormatReplaceAllTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridTextFormatReplaceAllType value.
func (v HeatgridTextFormatReplaceAllType) Ptr() *HeatgridTextFormatReplaceAllType {
	return &v
}
