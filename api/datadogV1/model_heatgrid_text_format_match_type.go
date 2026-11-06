// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridTextFormatMatchType Comparison used to match a group label.
type HeatgridTextFormatMatchType string

// List of HeatgridTextFormatMatchType.
const (
	HEATGRIDTEXTFORMATMATCHTYPE_IS               HeatgridTextFormatMatchType = "is"
	HEATGRIDTEXTFORMATMATCHTYPE_IS_NOT           HeatgridTextFormatMatchType = "is_not"
	HEATGRIDTEXTFORMATMATCHTYPE_CONTAINS         HeatgridTextFormatMatchType = "contains"
	HEATGRIDTEXTFORMATMATCHTYPE_DOES_NOT_CONTAIN HeatgridTextFormatMatchType = "does_not_contain"
	HEATGRIDTEXTFORMATMATCHTYPE_STARTS_WITH      HeatgridTextFormatMatchType = "starts_with"
	HEATGRIDTEXTFORMATMATCHTYPE_ENDS_WITH        HeatgridTextFormatMatchType = "ends_with"
)

var allowedHeatgridTextFormatMatchTypeEnumValues = []HeatgridTextFormatMatchType{
	HEATGRIDTEXTFORMATMATCHTYPE_IS,
	HEATGRIDTEXTFORMATMATCHTYPE_IS_NOT,
	HEATGRIDTEXTFORMATMATCHTYPE_CONTAINS,
	HEATGRIDTEXTFORMATMATCHTYPE_DOES_NOT_CONTAIN,
	HEATGRIDTEXTFORMATMATCHTYPE_STARTS_WITH,
	HEATGRIDTEXTFORMATMATCHTYPE_ENDS_WITH,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HeatgridTextFormatMatchType) GetAllowedValues() []HeatgridTextFormatMatchType {
	return allowedHeatgridTextFormatMatchTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HeatgridTextFormatMatchType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HeatgridTextFormatMatchType(value)
	return nil
}

// NewHeatgridTextFormatMatchTypeFromValue returns a pointer to a valid HeatgridTextFormatMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHeatgridTextFormatMatchTypeFromValue(v string) (*HeatgridTextFormatMatchType, error) {
	ev := HeatgridTextFormatMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HeatgridTextFormatMatchType: valid values are %v", v, allowedHeatgridTextFormatMatchTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HeatgridTextFormatMatchType) IsValid() bool {
	for _, existing := range allowedHeatgridTextFormatMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HeatgridTextFormatMatchType value.
func (v HeatgridTextFormatMatchType) Ptr() *HeatgridTextFormatMatchType {
	return &v
}
