// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyRequestType The resource type identifier for a journey funnel request.
type ProductAnalyticsJourneyRequestType string

// List of ProductAnalyticsJourneyRequestType.
const (
	PRODUCTANALYTICSJOURNEYREQUESTTYPE_JOURNEY_REQUEST ProductAnalyticsJourneyRequestType = "journey_request"
)

var allowedProductAnalyticsJourneyRequestTypeEnumValues = []ProductAnalyticsJourneyRequestType{
	PRODUCTANALYTICSJOURNEYREQUESTTYPE_JOURNEY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyRequestType) GetAllowedValues() []ProductAnalyticsJourneyRequestType {
	return allowedProductAnalyticsJourneyRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyRequestType(value)
	return nil
}

// NewProductAnalyticsJourneyRequestTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyRequestTypeFromValue(v string) (*ProductAnalyticsJourneyRequestType, error) {
	ev := ProductAnalyticsJourneyRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyRequestType: valid values are %v", v, allowedProductAnalyticsJourneyRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyRequestType value.
func (v ProductAnalyticsJourneyRequestType) Ptr() *ProductAnalyticsJourneyRequestType {
	return &v
}
