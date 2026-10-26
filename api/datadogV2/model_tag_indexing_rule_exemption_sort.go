// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleExemptionSort Sort field for tag indexing rule exemptions.
type TagIndexingRuleExemptionSort string

// List of TagIndexingRuleExemptionSort.
const (
	TAGINDEXINGRULEEXEMPTIONSORT_NAME                                  TagIndexingRuleExemptionSort = "name"
	TAGINDEXINGRULEEXEMPTIONSORT_NAME_DESCENDING                       TagIndexingRuleExemptionSort = "-name"
	TAGINDEXINGRULEEXEMPTIONSORT_ESTIMATED_CHANGE_IN_VOLUME            TagIndexingRuleExemptionSort = "estimated_change_in_volume"
	TAGINDEXINGRULEEXEMPTIONSORT_ESTIMATED_CHANGE_IN_VOLUME_DESCENDING TagIndexingRuleExemptionSort = "-estimated_change_in_volume"
)

var allowedTagIndexingRuleExemptionSortEnumValues = []TagIndexingRuleExemptionSort{
	TAGINDEXINGRULEEXEMPTIONSORT_NAME,
	TAGINDEXINGRULEEXEMPTIONSORT_NAME_DESCENDING,
	TAGINDEXINGRULEEXEMPTIONSORT_ESTIMATED_CHANGE_IN_VOLUME,
	TAGINDEXINGRULEEXEMPTIONSORT_ESTIMATED_CHANGE_IN_VOLUME_DESCENDING,
}

// GetAllowedValues reeturns the list of possible values.
func (v *TagIndexingRuleExemptionSort) GetAllowedValues() []TagIndexingRuleExemptionSort {
	return allowedTagIndexingRuleExemptionSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TagIndexingRuleExemptionSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TagIndexingRuleExemptionSort(value)
	return nil
}

// NewTagIndexingRuleExemptionSortFromValue returns a pointer to a valid TagIndexingRuleExemptionSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTagIndexingRuleExemptionSortFromValue(v string) (*TagIndexingRuleExemptionSort, error) {
	ev := TagIndexingRuleExemptionSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TagIndexingRuleExemptionSort: valid values are %v", v, allowedTagIndexingRuleExemptionSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TagIndexingRuleExemptionSort) IsValid() bool {
	for _, existing := range allowedTagIndexingRuleExemptionSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagIndexingRuleExemptionSort value.
func (v TagIndexingRuleExemptionSort) Ptr() *TagIndexingRuleExemptionSort {
	return &v
}
