// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeAudienceFiltersQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type CompositeAudienceFiltersQueryDataSource string

// List of CompositeAudienceFiltersQueryDataSource.
const (
	COMPOSITEAUDIENCEFILTERSQUERYDATASOURCE_PRODUCT_ANALYTICS_AUDIENCE_FILTERS CompositeAudienceFiltersQueryDataSource = "product_analytics_audience_filters"
)

var allowedCompositeAudienceFiltersQueryDataSourceEnumValues = []CompositeAudienceFiltersQueryDataSource{
	COMPOSITEAUDIENCEFILTERSQUERYDATASOURCE_PRODUCT_ANALYTICS_AUDIENCE_FILTERS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CompositeAudienceFiltersQueryDataSource) GetAllowedValues() []CompositeAudienceFiltersQueryDataSource {
	return allowedCompositeAudienceFiltersQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CompositeAudienceFiltersQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CompositeAudienceFiltersQueryDataSource(value)
	return nil
}

// NewCompositeAudienceFiltersQueryDataSourceFromValue returns a pointer to a valid CompositeAudienceFiltersQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCompositeAudienceFiltersQueryDataSourceFromValue(v string) (*CompositeAudienceFiltersQueryDataSource, error) {
	ev := CompositeAudienceFiltersQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CompositeAudienceFiltersQueryDataSource: valid values are %v", v, allowedCompositeAudienceFiltersQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CompositeAudienceFiltersQueryDataSource) IsValid() bool {
	for _, existing := range allowedCompositeAudienceFiltersQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompositeAudienceFiltersQueryDataSource value.
func (v CompositeAudienceFiltersQueryDataSource) Ptr() *CompositeAudienceFiltersQueryDataSource {
	return &v
}
