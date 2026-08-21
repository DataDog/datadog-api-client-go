// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsGraphQueryGroupBySource Audience dimension to group by, instead of an event facet.
type ProductAnalyticsGraphQueryGroupBySource string

// List of ProductAnalyticsGraphQueryGroupBySource.
const (
	PRODUCTANALYTICSGRAPHQUERYGROUPBYSOURCE_USERS    ProductAnalyticsGraphQueryGroupBySource = "product_analytics_audience_filters.users"
	PRODUCTANALYTICSGRAPHQUERYGROUPBYSOURCE_ACCOUNTS ProductAnalyticsGraphQueryGroupBySource = "product_analytics_audience_filters.accounts"
)

var allowedProductAnalyticsGraphQueryGroupBySourceEnumValues = []ProductAnalyticsGraphQueryGroupBySource{
	PRODUCTANALYTICSGRAPHQUERYGROUPBYSOURCE_USERS,
	PRODUCTANALYTICSGRAPHQUERYGROUPBYSOURCE_ACCOUNTS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsGraphQueryGroupBySource) GetAllowedValues() []ProductAnalyticsGraphQueryGroupBySource {
	return allowedProductAnalyticsGraphQueryGroupBySourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsGraphQueryGroupBySource) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsGraphQueryGroupBySource(value)
	return nil
}

// NewProductAnalyticsGraphQueryGroupBySourceFromValue returns a pointer to a valid ProductAnalyticsGraphQueryGroupBySource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsGraphQueryGroupBySourceFromValue(v string) (*ProductAnalyticsGraphQueryGroupBySource, error) {
	ev := ProductAnalyticsGraphQueryGroupBySource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsGraphQueryGroupBySource: valid values are %v", v, allowedProductAnalyticsGraphQueryGroupBySourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsGraphQueryGroupBySource) IsValid() bool {
	for _, existing := range allowedProductAnalyticsGraphQueryGroupBySourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsGraphQueryGroupBySource value.
func (v ProductAnalyticsGraphQueryGroupBySource) Ptr() *ProductAnalyticsGraphQueryGroupBySource {
	return &v
}
