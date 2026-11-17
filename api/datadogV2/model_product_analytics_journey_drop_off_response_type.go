// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyDropOffResponseType The resource type identifier for a journey drop-off analysis response.
type ProductAnalyticsJourneyDropOffResponseType string

// List of ProductAnalyticsJourneyDropOffResponseType.
const (
	PRODUCTANALYTICSJOURNEYDROPOFFRESPONSETYPE_JOURNEY_DROP_OFF_ANALYSIS_RESPONSE ProductAnalyticsJourneyDropOffResponseType = "journey_drop_off_analysis_response"
)

var allowedProductAnalyticsJourneyDropOffResponseTypeEnumValues = []ProductAnalyticsJourneyDropOffResponseType{
	PRODUCTANALYTICSJOURNEYDROPOFFRESPONSETYPE_JOURNEY_DROP_OFF_ANALYSIS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyDropOffResponseType) GetAllowedValues() []ProductAnalyticsJourneyDropOffResponseType {
	return allowedProductAnalyticsJourneyDropOffResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyDropOffResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyDropOffResponseType(value)
	return nil
}

// NewProductAnalyticsJourneyDropOffResponseTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyDropOffResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyDropOffResponseTypeFromValue(v string) (*ProductAnalyticsJourneyDropOffResponseType, error) {
	ev := ProductAnalyticsJourneyDropOffResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyDropOffResponseType: valid values are %v", v, allowedProductAnalyticsJourneyDropOffResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyDropOffResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyDropOffResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyDropOffResponseType value.
func (v ProductAnalyticsJourneyDropOffResponseType) Ptr() *ProductAnalyticsJourneyDropOffResponseType {
	return &v
}
