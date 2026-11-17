// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelSuggestionsResponseType Type of a feature suggestions response.
type ProductAnalyticsFunnelSuggestionsResponseType string

// List of ProductAnalyticsFunnelSuggestionsResponseType.
const (
	PRODUCTANALYTICSFUNNELSUGGESTIONSRESPONSETYPE_FUNNEL_SUGGESTIONS_RESPONSE ProductAnalyticsFunnelSuggestionsResponseType = "funnel_suggestions_response"
)

var allowedProductAnalyticsFunnelSuggestionsResponseTypeEnumValues = []ProductAnalyticsFunnelSuggestionsResponseType{
	PRODUCTANALYTICSFUNNELSUGGESTIONSRESPONSETYPE_FUNNEL_SUGGESTIONS_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsFunnelSuggestionsResponseType) GetAllowedValues() []ProductAnalyticsFunnelSuggestionsResponseType {
	return allowedProductAnalyticsFunnelSuggestionsResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsFunnelSuggestionsResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsFunnelSuggestionsResponseType(value)
	return nil
}

// NewProductAnalyticsFunnelSuggestionsResponseTypeFromValue returns a pointer to a valid ProductAnalyticsFunnelSuggestionsResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsFunnelSuggestionsResponseTypeFromValue(v string) (*ProductAnalyticsFunnelSuggestionsResponseType, error) {
	ev := ProductAnalyticsFunnelSuggestionsResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsFunnelSuggestionsResponseType: valid values are %v", v, allowedProductAnalyticsFunnelSuggestionsResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsFunnelSuggestionsResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsFunnelSuggestionsResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsFunnelSuggestionsResponseType value.
func (v ProductAnalyticsFunnelSuggestionsResponseType) Ptr() *ProductAnalyticsFunnelSuggestionsResponseType {
	return &v
}
