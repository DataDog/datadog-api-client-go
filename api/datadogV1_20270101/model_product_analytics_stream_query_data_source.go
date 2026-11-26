// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsStreamQueryDataSource Test-only generated schema for the dashboard versioning proof of concept.
type ProductAnalyticsStreamQueryDataSource string

// List of ProductAnalyticsStreamQueryDataSource.
const (
	PRODUCTANALYTICSSTREAMQUERYDATASOURCE_PRODUCT_ANALYTICS_STREAM ProductAnalyticsStreamQueryDataSource = "product_analytics_stream"
)

var allowedProductAnalyticsStreamQueryDataSourceEnumValues = []ProductAnalyticsStreamQueryDataSource{
	PRODUCTANALYTICSSTREAMQUERYDATASOURCE_PRODUCT_ANALYTICS_STREAM,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsStreamQueryDataSource) GetAllowedValues() []ProductAnalyticsStreamQueryDataSource {
	return allowedProductAnalyticsStreamQueryDataSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsStreamQueryDataSource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsStreamQueryDataSource(value)
	return nil
}

// NewProductAnalyticsStreamQueryDataSourceFromValue returns a pointer to a valid ProductAnalyticsStreamQueryDataSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsStreamQueryDataSourceFromValue(v string) (*ProductAnalyticsStreamQueryDataSource, error) {
	ev := ProductAnalyticsStreamQueryDataSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsStreamQueryDataSource: valid values are %v", v, allowedProductAnalyticsStreamQueryDataSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsStreamQueryDataSource) IsValid() bool {
	for _, existing := range allowedProductAnalyticsStreamQueryDataSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsStreamQueryDataSource value.
func (v ProductAnalyticsStreamQueryDataSource) Ptr() *ProductAnalyticsStreamQueryDataSource {
	return &v
}
