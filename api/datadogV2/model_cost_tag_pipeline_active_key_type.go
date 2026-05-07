// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostTagPipelineActiveKeyType Active tag key resource type.
type CostTagPipelineActiveKeyType string

// List of CostTagPipelineActiveKeyType.
const (
	COSTTAGPIPELINEACTIVEKEYTYPE_ACTIVE_TAG_KEY CostTagPipelineActiveKeyType = "active_tag_key"
)

var allowedCostTagPipelineActiveKeyTypeEnumValues = []CostTagPipelineActiveKeyType{
	COSTTAGPIPELINEACTIVEKEYTYPE_ACTIVE_TAG_KEY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostTagPipelineActiveKeyType) GetAllowedValues() []CostTagPipelineActiveKeyType {
	return allowedCostTagPipelineActiveKeyTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostTagPipelineActiveKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostTagPipelineActiveKeyType(value)
	return nil
}

// NewCostTagPipelineActiveKeyTypeFromValue returns a pointer to a valid CostTagPipelineActiveKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostTagPipelineActiveKeyTypeFromValue(v string) (*CostTagPipelineActiveKeyType, error) {
	ev := CostTagPipelineActiveKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostTagPipelineActiveKeyType: valid values are %v", v, allowedCostTagPipelineActiveKeyTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostTagPipelineActiveKeyType) IsValid() bool {
	for _, existing := range allowedCostTagPipelineActiveKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostTagPipelineActiveKeyType value.
func (v CostTagPipelineActiveKeyType) Ptr() *CostTagPipelineActiveKeyType {
	return &v
}
