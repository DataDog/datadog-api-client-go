// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyDropOffRequestType The resource type identifier for a journey drop-off analysis request.
type ProductAnalyticsJourneyDropOffRequestType string

// List of ProductAnalyticsJourneyDropOffRequestType.
const (
	PRODUCTANALYTICSJOURNEYDROPOFFREQUESTTYPE_JOURNEY_DROP_OFF_ANALYSIS_REQUEST ProductAnalyticsJourneyDropOffRequestType = "journey_drop_off_analysis_request"
)

var allowedProductAnalyticsJourneyDropOffRequestTypeEnumValues = []ProductAnalyticsJourneyDropOffRequestType{
	PRODUCTANALYTICSJOURNEYDROPOFFREQUESTTYPE_JOURNEY_DROP_OFF_ANALYSIS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyDropOffRequestType) GetAllowedValues() []ProductAnalyticsJourneyDropOffRequestType {
	return allowedProductAnalyticsJourneyDropOffRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyDropOffRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyDropOffRequestType(value)
	return nil
}

// NewProductAnalyticsJourneyDropOffRequestTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyDropOffRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyDropOffRequestTypeFromValue(v string) (*ProductAnalyticsJourneyDropOffRequestType, error) {
	ev := ProductAnalyticsJourneyDropOffRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyDropOffRequestType: valid values are %v", v, allowedProductAnalyticsJourneyDropOffRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyDropOffRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyDropOffRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyDropOffRequestType value.
func (v ProductAnalyticsJourneyDropOffRequestType) Ptr() *ProductAnalyticsJourneyDropOffRequestType {
	return &v
}
