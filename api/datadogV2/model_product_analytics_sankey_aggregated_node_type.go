// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyAggregatedNodeType
type ProductAnalyticsSankeyAggregatedNodeType string

// List of ProductAnalyticsSankeyAggregatedNodeType.
const (
	PRODUCTANALYTICSSANKEYAGGREGATEDNODETYPE_AGGREGATED ProductAnalyticsSankeyAggregatedNodeType = "aggregated"
)

var allowedProductAnalyticsSankeyAggregatedNodeTypeEnumValues = []ProductAnalyticsSankeyAggregatedNodeType{
	PRODUCTANALYTICSSANKEYAGGREGATEDNODETYPE_AGGREGATED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsSankeyAggregatedNodeType) GetAllowedValues() []ProductAnalyticsSankeyAggregatedNodeType {
	return allowedProductAnalyticsSankeyAggregatedNodeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsSankeyAggregatedNodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsSankeyAggregatedNodeType(value)
	return nil
}

// NewProductAnalyticsSankeyAggregatedNodeTypeFromValue returns a pointer to a valid ProductAnalyticsSankeyAggregatedNodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsSankeyAggregatedNodeTypeFromValue(v string) (*ProductAnalyticsSankeyAggregatedNodeType, error) {
	ev := ProductAnalyticsSankeyAggregatedNodeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsSankeyAggregatedNodeType: valid values are %v", v, allowedProductAnalyticsSankeyAggregatedNodeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsSankeyAggregatedNodeType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsSankeyAggregatedNodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsSankeyAggregatedNodeType value.
func (v ProductAnalyticsSankeyAggregatedNodeType) Ptr() *ProductAnalyticsSankeyAggregatedNodeType {
	return &v
}
