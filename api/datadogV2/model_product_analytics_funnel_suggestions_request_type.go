// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelSuggestionsRequestType Type of a feature suggestions request.
type ProductAnalyticsFunnelSuggestionsRequestType string

// List of ProductAnalyticsFunnelSuggestionsRequestType.
const (
	PRODUCTANALYTICSFUNNELSUGGESTIONSREQUESTTYPE_FUNNEL_SUGGESTIONS_REQUEST ProductAnalyticsFunnelSuggestionsRequestType = "funnel_suggestions_request"
)

var allowedProductAnalyticsFunnelSuggestionsRequestTypeEnumValues = []ProductAnalyticsFunnelSuggestionsRequestType{
	PRODUCTANALYTICSFUNNELSUGGESTIONSREQUESTTYPE_FUNNEL_SUGGESTIONS_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFunnelSuggestionsRequestType) GetAllowedValues() []ProductAnalyticsFunnelSuggestionsRequestType {
	return allowedProductAnalyticsFunnelSuggestionsRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFunnelSuggestionsRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFunnelSuggestionsRequestType(value)
	return nil
}

// NewProductAnalyticsFunnelSuggestionsRequestTypeFromValue returns a pointer to a valid ProductAnalyticsFunnelSuggestionsRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFunnelSuggestionsRequestTypeFromValue(v string) (*ProductAnalyticsFunnelSuggestionsRequestType, error) {
	ev := ProductAnalyticsFunnelSuggestionsRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFunnelSuggestionsRequestType: valid values are %v", v, allowedProductAnalyticsFunnelSuggestionsRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFunnelSuggestionsRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFunnelSuggestionsRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFunnelSuggestionsRequestType value.
func (v ProductAnalyticsFunnelSuggestionsRequestType) Ptr() *ProductAnalyticsFunnelSuggestionsRequestType {
	return &v
}
