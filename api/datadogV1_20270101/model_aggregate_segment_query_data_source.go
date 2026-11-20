// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateSegmentQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type AggregateSegmentQueryDataSource string

// List of AggregateSegmentQueryDataSource.
const (
	AGGREGATESEGMENTQUERYDATASOURCE_AGGREGATE_SEGMENT_QUERY AggregateSegmentQueryDataSource = "aggregate_segment_query"
)

var allowedAggregateSegmentQueryDataSourceEnumValues = []AggregateSegmentQueryDataSource{
	AGGREGATESEGMENTQUERYDATASOURCE_AGGREGATE_SEGMENT_QUERY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AggregateSegmentQueryDataSource) GetAllowedValues() []AggregateSegmentQueryDataSource {
	return allowedAggregateSegmentQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AggregateSegmentQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AggregateSegmentQueryDataSource(value)
	return nil
}

// NewAggregateSegmentQueryDataSourceFromValue returns a pointer to a valid AggregateSegmentQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAggregateSegmentQueryDataSourceFromValue(v string) (*AggregateSegmentQueryDataSource, error) {
	ev := AggregateSegmentQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AggregateSegmentQueryDataSource: valid values are %v", v, allowedAggregateSegmentQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AggregateSegmentQueryDataSource) IsValid() bool {
	for _, existing := range allowedAggregateSegmentQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AggregateSegmentQueryDataSource value.
func (v AggregateSegmentQueryDataSource) Ptr() *AggregateSegmentQueryDataSource {
	return &v
}
