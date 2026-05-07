// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostSettingType Cost setting resource type.
type CostSettingType string

// List of CostSettingType.
const (
	COSTSETTINGTYPE_SETTING CostSettingType = "setting"
)

var allowedCostSettingTypeEnumValues = []CostSettingType{
	COSTSETTINGTYPE_SETTING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostSettingType) GetAllowedValues() []CostSettingType {
	return allowedCostSettingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostSettingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostSettingType(value)
	return nil
}

// NewCostSettingTypeFromValue returns a pointer to a valid CostSettingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostSettingTypeFromValue(v string) (*CostSettingType, error) {
	ev := CostSettingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostSettingType: valid values are %v", v, allowedCostSettingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostSettingType) IsValid() bool {
	for _, existing := range allowedCostSettingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostSettingType value.
func (v CostSettingType) Ptr() *CostSettingType {
	return &v
}
