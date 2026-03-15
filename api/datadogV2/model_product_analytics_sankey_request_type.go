// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyRequestType
type ProductAnalyticsSankeyRequestType string

// List of ProductAnalyticsSankeyRequestType.
const (
	PRODUCTANALYTICSSANKEYREQUESTTYPE_SANKEY_REQUEST ProductAnalyticsSankeyRequestType = "sankey_request"
)

var allowedProductAnalyticsSankeyRequestTypeEnumValues = []ProductAnalyticsSankeyRequestType{
	PRODUCTANALYTICSSANKEYREQUESTTYPE_SANKEY_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsSankeyRequestType) GetAllowedValues() []ProductAnalyticsSankeyRequestType {
	return allowedProductAnalyticsSankeyRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsSankeyRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsSankeyRequestType(value)
	return nil
}

// NewProductAnalyticsSankeyRequestTypeFromValue returns a pointer to a valid ProductAnalyticsSankeyRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsSankeyRequestTypeFromValue(v string) (*ProductAnalyticsSankeyRequestType, error) {
	ev := ProductAnalyticsSankeyRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsSankeyRequestType: valid values are %v", v, allowedProductAnalyticsSankeyRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsSankeyRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsSankeyRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsSankeyRequestType value.
func (v ProductAnalyticsSankeyRequestType) Ptr() *ProductAnalyticsSankeyRequestType {
	return &v
}
