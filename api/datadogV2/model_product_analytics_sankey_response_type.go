// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyResponseType
type ProductAnalyticsSankeyResponseType string

// List of ProductAnalyticsSankeyResponseType.
const (
	PRODUCTANALYTICSSANKEYRESPONSETYPE_SANKEY_RESPONSE ProductAnalyticsSankeyResponseType = "sankey_response"
)

var allowedProductAnalyticsSankeyResponseTypeEnumValues = []ProductAnalyticsSankeyResponseType{
	PRODUCTANALYTICSSANKEYRESPONSETYPE_SANKEY_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsSankeyResponseType) GetAllowedValues() []ProductAnalyticsSankeyResponseType {
	return allowedProductAnalyticsSankeyResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsSankeyResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsSankeyResponseType(value)
	return nil
}

// NewProductAnalyticsSankeyResponseTypeFromValue returns a pointer to a valid ProductAnalyticsSankeyResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsSankeyResponseTypeFromValue(v string) (*ProductAnalyticsSankeyResponseType, error) {
	ev := ProductAnalyticsSankeyResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsSankeyResponseType: valid values are %v", v, allowedProductAnalyticsSankeyResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsSankeyResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsSankeyResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsSankeyResponseType value.
func (v ProductAnalyticsSankeyResponseType) Ptr() *ProductAnalyticsSankeyResponseType {
	return &v
}
