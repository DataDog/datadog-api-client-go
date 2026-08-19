// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagRuleSource The telemetry source that a tag rule applies to.
type TagRuleSource string

// List of TagRuleSource.
const (
	TAGRULESOURCE_LOGS    TagRuleSource = "logs"
	TAGRULESOURCE_SPANS   TagRuleSource = "spans"
	TAGRULESOURCE_METRICS TagRuleSource = "metrics"
)

var allowedTagRuleSourceEnumValues = []TagRuleSource{
	TAGRULESOURCE_LOGS,
	TAGRULESOURCE_SPANS,
	TAGRULESOURCE_METRICS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagRuleSource) GetAllowedValues() []TagRuleSource {
	return allowedTagRuleSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagRuleSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagRuleSource(value)
	return nil
}

// NewTagRuleSourceFromValue returns a pointer to a valid TagRuleSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagRuleSourceFromValue(v string) (*TagRuleSource, error) {
	ev := TagRuleSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagRuleSource: valid values are %v", v, allowedTagRuleSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagRuleSource) IsValid() bool {
	for _, existing := range allowedTagRuleSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagRuleSource value.
func (v TagRuleSource) Ptr() *TagRuleSource {
	return &v
}
