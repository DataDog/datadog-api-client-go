// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelStepSuggestionsRequestType Type of a step suggestions request.
type ProductAnalyticsFunnelStepSuggestionsRequestType string

// List of ProductAnalyticsFunnelStepSuggestionsRequestType.
const (
	PRODUCTANALYTICSFUNNELSTEPSUGGESTIONSREQUESTTYPE_FUNNEL_SUGGESTIONS_STEPS_REQUEST ProductAnalyticsFunnelStepSuggestionsRequestType = "funnel_suggestions_steps_request"
)

var allowedProductAnalyticsFunnelStepSuggestionsRequestTypeEnumValues = []ProductAnalyticsFunnelStepSuggestionsRequestType{
	PRODUCTANALYTICSFUNNELSTEPSUGGESTIONSREQUESTTYPE_FUNNEL_SUGGESTIONS_STEPS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFunnelStepSuggestionsRequestType) GetAllowedValues() []ProductAnalyticsFunnelStepSuggestionsRequestType {
	return allowedProductAnalyticsFunnelStepSuggestionsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFunnelStepSuggestionsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFunnelStepSuggestionsRequestType(value)
	return nil
}

// NewProductAnalyticsFunnelStepSuggestionsRequestTypeFromValue returns a pointer to a valid ProductAnalyticsFunnelStepSuggestionsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFunnelStepSuggestionsRequestTypeFromValue(v string) (*ProductAnalyticsFunnelStepSuggestionsRequestType, error) {
	ev := ProductAnalyticsFunnelStepSuggestionsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFunnelStepSuggestionsRequestType: valid values are %v", v, allowedProductAnalyticsFunnelStepSuggestionsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFunnelStepSuggestionsRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFunnelStepSuggestionsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFunnelStepSuggestionsRequestType value.
func (v ProductAnalyticsFunnelStepSuggestionsRequestType) Ptr() *ProductAnalyticsFunnelStepSuggestionsRequestType {
	return &v
}
