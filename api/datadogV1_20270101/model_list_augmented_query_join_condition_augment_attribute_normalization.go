// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAugmentedQueryJoinConditionAugmentAttributeNormalization Test-only generated schema for the dashboard versioning proof of concept.
type ListAugmentedQueryJoinConditionAugmentAttributeNormalization string

// List of ListAugmentedQueryJoinConditionAugmentAttributeNormalization.
const (
	LISTAUGMENTEDQUERYJOINCONDITIONAUGMENTATTRIBUTENORMALIZATION_METRIC_TAG ListAugmentedQueryJoinConditionAugmentAttributeNormalization = "metric_tag"
)

var allowedListAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues = []ListAugmentedQueryJoinConditionAugmentAttributeNormalization{
	LISTAUGMENTEDQUERYJOINCONDITIONAUGMENTATTRIBUTENORMALIZATION_METRIC_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ListAugmentedQueryJoinConditionAugmentAttributeNormalization) GetAllowedValues() []ListAugmentedQueryJoinConditionAugmentAttributeNormalization {
	return allowedListAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ListAugmentedQueryJoinConditionAugmentAttributeNormalization) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ListAugmentedQueryJoinConditionAugmentAttributeNormalization(value)
	return nil
}

// NewListAugmentedQueryJoinConditionAugmentAttributeNormalizationFromValue returns a pointer to a valid ListAugmentedQueryJoinConditionAugmentAttributeNormalization
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewListAugmentedQueryJoinConditionAugmentAttributeNormalizationFromValue(v string) (*ListAugmentedQueryJoinConditionAugmentAttributeNormalization, error) {
	ev := ListAugmentedQueryJoinConditionAugmentAttributeNormalization(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ListAugmentedQueryJoinConditionAugmentAttributeNormalization: valid values are %v", v, allowedListAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ListAugmentedQueryJoinConditionAugmentAttributeNormalization) IsValid() bool {
	for _, existing := range allowedListAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListAugmentedQueryJoinConditionAugmentAttributeNormalization value.
func (v ListAugmentedQueryJoinConditionAugmentAttributeNormalization) Ptr() *ListAugmentedQueryJoinConditionAugmentAttributeNormalization {
	return &v
}
