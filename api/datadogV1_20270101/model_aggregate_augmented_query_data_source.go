// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateAugmentedQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type AggregateAugmentedQueryDataSource string

// List of AggregateAugmentedQueryDataSource.
const (
	AGGREGATEAUGMENTEDQUERYDATASOURCE_AGGREGATE_AUGMENTED_QUERY AggregateAugmentedQueryDataSource = "aggregate_augmented_query"
)

var allowedAggregateAugmentedQueryDataSourceEnumValues = []AggregateAugmentedQueryDataSource{
	AGGREGATEAUGMENTEDQUERYDATASOURCE_AGGREGATE_AUGMENTED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregateAugmentedQueryDataSource) GetAllowedValues() []AggregateAugmentedQueryDataSource {
	return allowedAggregateAugmentedQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregateAugmentedQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregateAugmentedQueryDataSource(value)
	return nil
}

// NewAggregateAugmentedQueryDataSourceFromValue returns a pointer to a valid AggregateAugmentedQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregateAugmentedQueryDataSourceFromValue(v string) (*AggregateAugmentedQueryDataSource, error) {
	ev := AggregateAugmentedQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregateAugmentedQueryDataSource: valid values are %v", v, allowedAggregateAugmentedQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregateAugmentedQueryDataSource) IsValid() bool {
	for _, existing := range allowedAggregateAugmentedQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregateAugmentedQueryDataSource value.
func (v AggregateAugmentedQueryDataSource) Ptr() *AggregateAugmentedQueryDataSource {
	return &v
}
