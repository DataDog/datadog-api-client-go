// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OverviewItemDataType Overview items resource type.
type OverviewItemDataType string

// List of OverviewItemDataType.
const (
	OVERVIEWITEMDATATYPE_OVERVIEW_ITEMS OverviewItemDataType = "overview_items"
)

var allowedOverviewItemDataTypeEnumValues = []OverviewItemDataType{
	OVERVIEWITEMDATATYPE_OVERVIEW_ITEMS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OverviewItemDataType) GetAllowedValues() []OverviewItemDataType {
	return allowedOverviewItemDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OverviewItemDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OverviewItemDataType(value)
	return nil
}

// NewOverviewItemDataTypeFromValue returns a pointer to a valid OverviewItemDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOverviewItemDataTypeFromValue(v string) (*OverviewItemDataType, error) {
	ev := OverviewItemDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OverviewItemDataType: valid values are %v", v, allowedOverviewItemDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OverviewItemDataType) IsValid() bool {
	for _, existing := range allowedOverviewItemDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OverviewItemDataType value.
func (v OverviewItemDataType) Ptr() *OverviewItemDataType {
	return &v
}
