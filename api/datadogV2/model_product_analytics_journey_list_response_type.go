// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyListResponseType The resource type identifier for a journey list response.
type ProductAnalyticsJourneyListResponseType string

// List of ProductAnalyticsJourneyListResponseType.
const (
	PRODUCTANALYTICSJOURNEYLISTRESPONSETYPE_JOURNEY_LIST_RESPONSE ProductAnalyticsJourneyListResponseType = "journey_list_response"
)

var allowedProductAnalyticsJourneyListResponseTypeEnumValues = []ProductAnalyticsJourneyListResponseType{
	PRODUCTANALYTICSJOURNEYLISTRESPONSETYPE_JOURNEY_LIST_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyListResponseType) GetAllowedValues() []ProductAnalyticsJourneyListResponseType {
	return allowedProductAnalyticsJourneyListResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyListResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyListResponseType(value)
	return nil
}

// NewProductAnalyticsJourneyListResponseTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyListResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyListResponseTypeFromValue(v string) (*ProductAnalyticsJourneyListResponseType, error) {
	ev := ProductAnalyticsJourneyListResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyListResponseType: valid values are %v", v, allowedProductAnalyticsJourneyListResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyListResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyListResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyListResponseType value.
func (v ProductAnalyticsJourneyListResponseType) Ptr() *ProductAnalyticsJourneyListResponseType {
	return &v
}
