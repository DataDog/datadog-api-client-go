// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateFilteredQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type AggregateFilteredQueryDataSource string

// List of AggregateFilteredQueryDataSource.
const (
	AGGREGATEFILTEREDQUERYDATASOURCE_AGGREGATE_FILTERED_QUERY AggregateFilteredQueryDataSource = "aggregate_filtered_query"
)

var allowedAggregateFilteredQueryDataSourceEnumValues = []AggregateFilteredQueryDataSource{
	AGGREGATEFILTEREDQUERYDATASOURCE_AGGREGATE_FILTERED_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregateFilteredQueryDataSource) GetAllowedValues() []AggregateFilteredQueryDataSource {
	return allowedAggregateFilteredQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregateFilteredQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregateFilteredQueryDataSource(value)
	return nil
}

// NewAggregateFilteredQueryDataSourceFromValue returns a pointer to a valid AggregateFilteredQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregateFilteredQueryDataSourceFromValue(v string) (*AggregateFilteredQueryDataSource, error) {
	ev := AggregateFilteredQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregateFilteredQueryDataSource: valid values are %v", v, allowedAggregateFilteredQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregateFilteredQueryDataSource) IsValid() bool {
	for _, existing := range allowedAggregateFilteredQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregateFilteredQueryDataSource value.
func (v AggregateFilteredQueryDataSource) Ptr() *AggregateFilteredQueryDataSource {
	return &v
}
