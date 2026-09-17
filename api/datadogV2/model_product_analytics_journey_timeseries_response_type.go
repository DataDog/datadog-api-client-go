// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyTimeseriesResponseType The resource type identifier for a journey timeseries response.
type ProductAnalyticsJourneyTimeseriesResponseType string

// List of ProductAnalyticsJourneyTimeseriesResponseType.
const (
	PRODUCTANALYTICSJOURNEYTIMESERIESRESPONSETYPE_JOURNEY_TIMESERIES_RESPONSE ProductAnalyticsJourneyTimeseriesResponseType = "journey_timeseries_response"
)

var allowedProductAnalyticsJourneyTimeseriesResponseTypeEnumValues = []ProductAnalyticsJourneyTimeseriesResponseType{
	PRODUCTANALYTICSJOURNEYTIMESERIESRESPONSETYPE_JOURNEY_TIMESERIES_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyTimeseriesResponseType) GetAllowedValues() []ProductAnalyticsJourneyTimeseriesResponseType {
	return allowedProductAnalyticsJourneyTimeseriesResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyTimeseriesResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyTimeseriesResponseType(value)
	return nil
}

// NewProductAnalyticsJourneyTimeseriesResponseTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyTimeseriesResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyTimeseriesResponseTypeFromValue(v string) (*ProductAnalyticsJourneyTimeseriesResponseType, error) {
	ev := ProductAnalyticsJourneyTimeseriesResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyTimeseriesResponseType: valid values are %v", v, allowedProductAnalyticsJourneyTimeseriesResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyTimeseriesResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyTimeseriesResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyTimeseriesResponseType value.
func (v ProductAnalyticsJourneyTimeseriesResponseType) Ptr() *ProductAnalyticsJourneyTimeseriesResponseType {
	return &v
}
