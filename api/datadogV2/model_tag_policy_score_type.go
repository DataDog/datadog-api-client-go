// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyScoreType The type of the resource. The value should always be `tag_policy_score`.
type TagPolicyScoreType string

// List of TagPolicyScoreType.
const (
	TAGPOLICYSCORETYPE_TAG_POLICY_SCORE TagPolicyScoreType = "tag_policy_score"
)

var allowedTagPolicyScoreTypeEnumValues = []TagPolicyScoreType{
	TAGPOLICYSCORETYPE_TAG_POLICY_SCORE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagPolicyScoreType) GetAllowedValues() []TagPolicyScoreType {
	return allowedTagPolicyScoreTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagPolicyScoreType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagPolicyScoreType(value)
	return nil
}

// NewTagPolicyScoreTypeFromValue returns a pointer to a valid TagPolicyScoreType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagPolicyScoreTypeFromValue(v string) (*TagPolicyScoreType, error) {
	ev := TagPolicyScoreType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagPolicyScoreType: valid values are %v", v, allowedTagPolicyScoreTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagPolicyScoreType) IsValid() bool {
	for _, existing := range allowedTagPolicyScoreTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagPolicyScoreType value.
func (v TagPolicyScoreType) Ptr() *TagPolicyScoreType {
	return &v
}
