// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyFunnelResponseType The resource type identifier for a journey funnel response.
type ProductAnalyticsJourneyFunnelResponseType string

// List of ProductAnalyticsJourneyFunnelResponseType.
const (
	PRODUCTANALYTICSJOURNEYFUNNELRESPONSETYPE_FUNNEL_RESPONSE ProductAnalyticsJourneyFunnelResponseType = "funnel_response"
)

var allowedProductAnalyticsJourneyFunnelResponseTypeEnumValues = []ProductAnalyticsJourneyFunnelResponseType{
	PRODUCTANALYTICSJOURNEYFUNNELRESPONSETYPE_FUNNEL_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyFunnelResponseType) GetAllowedValues() []ProductAnalyticsJourneyFunnelResponseType {
	return allowedProductAnalyticsJourneyFunnelResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyFunnelResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyFunnelResponseType(value)
	return nil
}

// NewProductAnalyticsJourneyFunnelResponseTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyFunnelResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyFunnelResponseTypeFromValue(v string) (*ProductAnalyticsJourneyFunnelResponseType, error) {
	ev := ProductAnalyticsJourneyFunnelResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyFunnelResponseType: valid values are %v", v, allowedProductAnalyticsJourneyFunnelResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyFunnelResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyFunnelResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyFunnelResponseType value.
func (v ProductAnalyticsJourneyFunnelResponseType) Ptr() *ProductAnalyticsJourneyFunnelResponseType {
	return &v
}
