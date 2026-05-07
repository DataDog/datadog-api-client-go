// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAIPreferredTagsType Preferred tags resource type.
type CostAIPreferredTagsType string

// List of CostAIPreferredTagsType.
const (
	COSTAIPREFERREDTAGSTYPE_PREFERRED_TAGS CostAIPreferredTagsType = "preferred_tags"
)

var allowedCostAIPreferredTagsTypeEnumValues = []CostAIPreferredTagsType{
	COSTAIPREFERREDTAGSTYPE_PREFERRED_TAGS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CostAIPreferredTagsType) GetAllowedValues() []CostAIPreferredTagsType {
	return allowedCostAIPreferredTagsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CostAIPreferredTagsType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CostAIPreferredTagsType(value)
	return nil
}

// NewCostAIPreferredTagsTypeFromValue returns a pointer to a valid CostAIPreferredTagsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCostAIPreferredTagsTypeFromValue(v string) (*CostAIPreferredTagsType, error) {
	ev := CostAIPreferredTagsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CostAIPreferredTagsType: valid values are %v", v, allowedCostAIPreferredTagsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CostAIPreferredTagsType) IsValid() bool {
	for _, existing := range allowedCostAIPreferredTagsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CostAIPreferredTagsType value.
func (v CostAIPreferredTagsType) Ptr() *CostAIPreferredTagsType {
	return &v
}
