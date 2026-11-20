// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsEventOccurrenceQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ProductAnalyticsEventOccurrenceQueryDataSource string

// List of ProductAnalyticsEventOccurrenceQueryDataSource.
const (
	PRODUCTANALYTICSEVENTOCCURRENCEQUERYDATASOURCE_PRODUCT_ANALYTICS_OCCURRENCE ProductAnalyticsEventOccurrenceQueryDataSource = "product_analytics_occurrence"
)

var allowedProductAnalyticsEventOccurrenceQueryDataSourceEnumValues = []ProductAnalyticsEventOccurrenceQueryDataSource{
	PRODUCTANALYTICSEVENTOCCURRENCEQUERYDATASOURCE_PRODUCT_ANALYTICS_OCCURRENCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsEventOccurrenceQueryDataSource) GetAllowedValues() []ProductAnalyticsEventOccurrenceQueryDataSource {
	return allowedProductAnalyticsEventOccurrenceQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsEventOccurrenceQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsEventOccurrenceQueryDataSource(value)
	return nil
}

// NewProductAnalyticsEventOccurrenceQueryDataSourceFromValue returns a pointer to a valid ProductAnalyticsEventOccurrenceQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsEventOccurrenceQueryDataSourceFromValue(v string) (*ProductAnalyticsEventOccurrenceQueryDataSource, error) {
	ev := ProductAnalyticsEventOccurrenceQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsEventOccurrenceQueryDataSource: valid values are %v", v, allowedProductAnalyticsEventOccurrenceQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsEventOccurrenceQueryDataSource) IsValid() bool {
	for _, existing := range allowedProductAnalyticsEventOccurrenceQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsEventOccurrenceQueryDataSource value.
func (v ProductAnalyticsEventOccurrenceQueryDataSource) Ptr() *ProductAnalyticsEventOccurrenceQueryDataSource {
	return &v
}
