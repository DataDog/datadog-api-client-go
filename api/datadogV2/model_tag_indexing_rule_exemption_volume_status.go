// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleExemptionVolumeStatus Indicates whether `estimated_change_in_volume` is available and, if not, why. `ok` indicates a fresh estimate.
// `unindexed_tags` indicates that the matching rule uses tags that are not indexed for the metric.
// `unavailable` indicates that the estimate is stale, was not computed, or could not be retrieved.
type TagIndexingRuleExemptionVolumeStatus string

// List of TagIndexingRuleExemptionVolumeStatus.
const (
	TAGINDEXINGRULEEXEMPTIONVOLUMESTATUS_OK             TagIndexingRuleExemptionVolumeStatus = "ok"
	TAGINDEXINGRULEEXEMPTIONVOLUMESTATUS_UNINDEXED_TAGS TagIndexingRuleExemptionVolumeStatus = "unindexed_tags"
	TAGINDEXINGRULEEXEMPTIONVOLUMESTATUS_UNAVAILABLE    TagIndexingRuleExemptionVolumeStatus = "unavailable"
)

var allowedTagIndexingRuleExemptionVolumeStatusEnumValues = []TagIndexingRuleExemptionVolumeStatus{
	TAGINDEXINGRULEEXEMPTIONVOLUMESTATUS_OK,
	TAGINDEXINGRULEEXEMPTIONVOLUMESTATUS_UNINDEXED_TAGS,
	TAGINDEXINGRULEEXEMPTIONVOLUMESTATUS_UNAVAILABLE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagIndexingRuleExemptionVolumeStatus) GetAllowedValues() []TagIndexingRuleExemptionVolumeStatus {
	return allowedTagIndexingRuleExemptionVolumeStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagIndexingRuleExemptionVolumeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagIndexingRuleExemptionVolumeStatus(value)
	return nil
}

// NewTagIndexingRuleExemptionVolumeStatusFromValue returns a pointer to a valid TagIndexingRuleExemptionVolumeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagIndexingRuleExemptionVolumeStatusFromValue(v string) (*TagIndexingRuleExemptionVolumeStatus, error) {
	ev := TagIndexingRuleExemptionVolumeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagIndexingRuleExemptionVolumeStatus: valid values are %v", v, allowedTagIndexingRuleExemptionVolumeStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagIndexingRuleExemptionVolumeStatus) IsValid() bool {
	for _, existing := range allowedTagIndexingRuleExemptionVolumeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagIndexingRuleExemptionVolumeStatus value.
func (v TagIndexingRuleExemptionVolumeStatus) Ptr() *TagIndexingRuleExemptionVolumeStatus {
	return &v
}
