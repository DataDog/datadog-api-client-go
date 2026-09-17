// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyScalarResponseType The resource type identifier for a journey scalar response.
type ProductAnalyticsJourneyScalarResponseType string

// List of ProductAnalyticsJourneyScalarResponseType.
const (
	PRODUCTANALYTICSJOURNEYSCALARRESPONSETYPE_JOURNEY_SCALAR_RESPONSE ProductAnalyticsJourneyScalarResponseType = "journey_scalar_response"
)

var allowedProductAnalyticsJourneyScalarResponseTypeEnumValues = []ProductAnalyticsJourneyScalarResponseType{
	PRODUCTANALYTICSJOURNEYSCALARRESPONSETYPE_JOURNEY_SCALAR_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyScalarResponseType) GetAllowedValues() []ProductAnalyticsJourneyScalarResponseType {
	return allowedProductAnalyticsJourneyScalarResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyScalarResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyScalarResponseType(value)
	return nil
}

// NewProductAnalyticsJourneyScalarResponseTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyScalarResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyScalarResponseTypeFromValue(v string) (*ProductAnalyticsJourneyScalarResponseType, error) {
	ev := ProductAnalyticsJourneyScalarResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyScalarResponseType: valid values are %v", v, allowedProductAnalyticsJourneyScalarResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyScalarResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyScalarResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyScalarResponseType value.
func (v ProductAnalyticsJourneyScalarResponseType) Ptr() *ProductAnalyticsJourneyScalarResponseType {
	return &v
}
