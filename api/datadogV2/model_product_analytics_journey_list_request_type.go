// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyListRequestType The resource type identifier for a journey list request.
type ProductAnalyticsJourneyListRequestType string

// List of ProductAnalyticsJourneyListRequestType.
const (
	PRODUCTANALYTICSJOURNEYLISTREQUESTTYPE_JOURNEY_LIST_REQUEST ProductAnalyticsJourneyListRequestType = "journey_list_request"
)

var allowedProductAnalyticsJourneyListRequestTypeEnumValues = []ProductAnalyticsJourneyListRequestType{
	PRODUCTANALYTICSJOURNEYLISTREQUESTTYPE_JOURNEY_LIST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsJourneyListRequestType) GetAllowedValues() []ProductAnalyticsJourneyListRequestType {
	return allowedProductAnalyticsJourneyListRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsJourneyListRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsJourneyListRequestType(value)
	return nil
}

// NewProductAnalyticsJourneyListRequestTypeFromValue returns a pointer to a valid ProductAnalyticsJourneyListRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsJourneyListRequestTypeFromValue(v string) (*ProductAnalyticsJourneyListRequestType, error) {
	ev := ProductAnalyticsJourneyListRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsJourneyListRequestType: valid values are %v", v, allowedProductAnalyticsJourneyListRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsJourneyListRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsJourneyListRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsJourneyListRequestType value.
func (v ProductAnalyticsJourneyListRequestType) Ptr() *ProductAnalyticsJourneyListRequestType {
	return &v
}
