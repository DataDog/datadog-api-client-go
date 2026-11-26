// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization Test-only generated schema for the dashboard versioning proof of concept.
type AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization string

// List of AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization.
const (
	AGGREGATEAUGMENTEDQUERYJOINCONDITIONAUGMENTATTRIBUTENORMALIZATION_METRIC_TAG AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization = "metric_tag"
)

var allowedAggregateAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues = []AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization{
	AGGREGATEAUGMENTEDQUERYJOINCONDITIONAUGMENTATTRIBUTENORMALIZATION_METRIC_TAG,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization) GetAllowedValues() []AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization {
	return allowedAggregateAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization(value)
	return nil
}

// NewAggregateAugmentedQueryJoinConditionAugmentAttributeNormalizationFromValue returns a pointer to a valid AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregateAugmentedQueryJoinConditionAugmentAttributeNormalizationFromValue(v string) (*AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization, error) {
	ev := AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization: valid values are %v", v, allowedAggregateAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization) IsValid() bool {
	for _, existing := range allowedAggregateAugmentedQueryJoinConditionAugmentAttributeNormalizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization value.
func (v AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization) Ptr() *AggregateAugmentedQueryJoinConditionAugmentAttributeNormalization {
	return &v
}
