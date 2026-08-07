// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleResourceType JSON:API resource type for a tag rule.
type TagRuleResourceType string

// List of TagRuleResourceType.
const (
	TAGRULERESOURCETYPE_TAG_POLICY TagRuleResourceType = "tag_policy"
)

var allowedTagRuleResourceTypeEnumValues = []TagRuleResourceType{
	TAGRULERESOURCETYPE_TAG_POLICY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagRuleResourceType) GetAllowedValues() []TagRuleResourceType {
	return allowedTagRuleResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagRuleResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagRuleResourceType(value)
	return nil
}

// NewTagRuleResourceTypeFromValue returns a pointer to a valid TagRuleResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagRuleResourceTypeFromValue(v string) (*TagRuleResourceType, error) {
	ev := TagRuleResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagRuleResourceType: valid values are %v", v, allowedTagRuleResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagRuleResourceType) IsValid() bool {
	for _, existing := range allowedTagRuleResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagRuleResourceType value.
func (v TagRuleResourceType) Ptr() *TagRuleResourceType {
	return &v
}
