// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleInclude A related resource to include alongside a tag rule in the response. Currently the only supported value is `score`.
type TagRuleInclude string

// List of TagRuleInclude.
const (
	TAGRULEINCLUDE_SCORE TagRuleInclude = "score"
)

var allowedTagRuleIncludeEnumValues = []TagRuleInclude{
	TAGRULEINCLUDE_SCORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagRuleInclude) GetAllowedValues() []TagRuleInclude {
	return allowedTagRuleIncludeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagRuleInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagRuleInclude(value)
	return nil
}

// NewTagRuleIncludeFromValue returns a pointer to a valid TagRuleInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagRuleIncludeFromValue(v string) (*TagRuleInclude, error) {
	ev := TagRuleInclude(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagRuleInclude: valid values are %v", v, allowedTagRuleIncludeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagRuleInclude) IsValid() bool {
	for _, existing := range allowedTagRuleIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagRuleInclude value.
func (v TagRuleInclude) Ptr() *TagRuleInclude {
	return &v
}
