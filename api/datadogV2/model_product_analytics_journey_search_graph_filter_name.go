// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneySearchGraphFilterName The journey-level metric the graph filter applies to.
type ProductAnalyticsJourneySearchGraphFilterName string

// List of ProductAnalyticsJourneySearchGraphFilterName.
const (
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTERNAME_TIME_TO_CONVERT ProductAnalyticsJourneySearchGraphFilterName = "__dd.time_to_convert"
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTERNAME_SESSION         ProductAnalyticsJourneySearchGraphFilterName = "__dd.session"
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTERNAME_DROPOFF_RATE    ProductAnalyticsJourneySearchGraphFilterName = "__dd.dropoff_rate"
)

var allowedProductAnalyticsJourneySearchGraphFilterNameEnumValues = []ProductAnalyticsJourneySearchGraphFilterName{
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTERNAME_TIME_TO_CONVERT,
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTERNAME_SESSION,
	PRODUCTANALYTICSJOURNEYSEARCHGRAPHFILTERNAME_DROPOFF_RATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneySearchGraphFilterName) GetAllowedValues() []ProductAnalyticsJourneySearchGraphFilterName {
	return allowedProductAnalyticsJourneySearchGraphFilterNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneySearchGraphFilterName) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneySearchGraphFilterName(value)
	return nil
}

// NewProductAnalyticsJourneySearchGraphFilterNameFromValue returns a pointer to a valid ProductAnalyticsJourneySearchGraphFilterName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneySearchGraphFilterNameFromValue(v string) (*ProductAnalyticsJourneySearchGraphFilterName, error) {
	ev := ProductAnalyticsJourneySearchGraphFilterName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneySearchGraphFilterName: valid values are %v", v, allowedProductAnalyticsJourneySearchGraphFilterNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneySearchGraphFilterName) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneySearchGraphFilterNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneySearchGraphFilterName value.
func (v ProductAnalyticsJourneySearchGraphFilterName) Ptr() *ProductAnalyticsJourneySearchGraphFilterName {
	return &v
}
