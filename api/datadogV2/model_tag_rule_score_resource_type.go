// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleScoreResourceType JSON:API resource type for a tag rule compliance score.
type TagRuleScoreResourceType string

// List of TagRuleScoreResourceType.
const (
	TAGRULESCORERESOURCETYPE_TAG_RULE_SCORE TagRuleScoreResourceType = "tag_rule_score"
)

var allowedTagRuleScoreResourceTypeEnumValues = []TagRuleScoreResourceType{
	TAGRULESCORERESOURCETYPE_TAG_RULE_SCORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagRuleScoreResourceType) GetAllowedValues() []TagRuleScoreResourceType {
	return allowedTagRuleScoreResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagRuleScoreResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagRuleScoreResourceType(value)
	return nil
}

// NewTagRuleScoreResourceTypeFromValue returns a pointer to a valid TagRuleScoreResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagRuleScoreResourceTypeFromValue(v string) (*TagRuleScoreResourceType, error) {
	ev := TagRuleScoreResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagRuleScoreResourceType: valid values are %v", v, allowedTagRuleScoreResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagRuleScoreResourceType) IsValid() bool {
	for _, existing := range allowedTagRuleScoreResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagRuleScoreResourceType value.
func (v TagRuleScoreResourceType) Ptr() *TagRuleScoreResourceType {
	return &v
}
