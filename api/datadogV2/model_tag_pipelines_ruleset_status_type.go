// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPipelinesRulesetStatusType Tag pipeline ruleset status resource type.
type TagPipelinesRulesetStatusType string

// List of TagPipelinesRulesetStatusType.
const (
	TAGPIPELINESRULESETSTATUSTYPE_RULESET_STATUS TagPipelinesRulesetStatusType = "ruleset_status"
)

var allowedTagPipelinesRulesetStatusTypeEnumValues = []TagPipelinesRulesetStatusType{
	TAGPIPELINESRULESETSTATUSTYPE_RULESET_STATUS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagPipelinesRulesetStatusType) GetAllowedValues() []TagPipelinesRulesetStatusType {
	return allowedTagPipelinesRulesetStatusTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagPipelinesRulesetStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagPipelinesRulesetStatusType(value)
	return nil
}

// NewTagPipelinesRulesetStatusTypeFromValue returns a pointer to a valid TagPipelinesRulesetStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagPipelinesRulesetStatusTypeFromValue(v string) (*TagPipelinesRulesetStatusType, error) {
	ev := TagPipelinesRulesetStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagPipelinesRulesetStatusType: valid values are %v", v, allowedTagPipelinesRulesetStatusTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagPipelinesRulesetStatusType) IsValid() bool {
	for _, existing := range allowedTagPipelinesRulesetStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagPipelinesRulesetStatusType value.
func (v TagPipelinesRulesetStatusType) Ptr() *TagPipelinesRulesetStatusType {
	return &v
}
