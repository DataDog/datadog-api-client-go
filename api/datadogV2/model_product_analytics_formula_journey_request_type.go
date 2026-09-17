// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFormulaJourneyRequestType The resource type identifier for a journey timeseries or scalar request.
type ProductAnalyticsFormulaJourneyRequestType string

// List of ProductAnalyticsFormulaJourneyRequestType.
const (
	PRODUCTANALYTICSFORMULAJOURNEYREQUESTTYPE_FORMULA_JOURNEY_REQUEST ProductAnalyticsFormulaJourneyRequestType = "formula_journey_request"
)

var allowedProductAnalyticsFormulaJourneyRequestTypeEnumValues = []ProductAnalyticsFormulaJourneyRequestType{
	PRODUCTANALYTICSFORMULAJOURNEYREQUESTTYPE_FORMULA_JOURNEY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFormulaJourneyRequestType) GetAllowedValues() []ProductAnalyticsFormulaJourneyRequestType {
	return allowedProductAnalyticsFormulaJourneyRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFormulaJourneyRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFormulaJourneyRequestType(value)
	return nil
}

// NewProductAnalyticsFormulaJourneyRequestTypeFromValue returns a pointer to a valid ProductAnalyticsFormulaJourneyRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFormulaJourneyRequestTypeFromValue(v string) (*ProductAnalyticsFormulaJourneyRequestType, error) {
	ev := ProductAnalyticsFormulaJourneyRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFormulaJourneyRequestType: valid values are %v", v, allowedProductAnalyticsFormulaJourneyRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFormulaJourneyRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFormulaJourneyRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFormulaJourneyRequestType value.
func (v ProductAnalyticsFormulaJourneyRequestType) Ptr() *ProductAnalyticsFormulaJourneyRequestType {
	return &v
}
