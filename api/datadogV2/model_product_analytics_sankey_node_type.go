// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyNodeType Node type.
type ProductAnalyticsSankeyNodeType string

// List of ProductAnalyticsSankeyNodeType.
const (
	PRODUCTANALYTICSSANKEYNODETYPE_REGULAR ProductAnalyticsSankeyNodeType = "regular"
	PRODUCTANALYTICSSANKEYNODETYPE_OTHER   ProductAnalyticsSankeyNodeType = "other"
	PRODUCTANALYTICSSANKEYNODETYPE_DROPOFF ProductAnalyticsSankeyNodeType = "dropoff"
)

var allowedProductAnalyticsSankeyNodeTypeEnumValues = []ProductAnalyticsSankeyNodeType{
	PRODUCTANALYTICSSANKEYNODETYPE_REGULAR,
	PRODUCTANALYTICSSANKEYNODETYPE_OTHER,
	PRODUCTANALYTICSSANKEYNODETYPE_DROPOFF,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsSankeyNodeType) GetAllowedValues() []ProductAnalyticsSankeyNodeType {
	return allowedProductAnalyticsSankeyNodeTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsSankeyNodeType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsSankeyNodeType(value)
	return nil
}

// NewProductAnalyticsSankeyNodeTypeFromValue returns a pointer to a valid ProductAnalyticsSankeyNodeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsSankeyNodeTypeFromValue(v string) (*ProductAnalyticsSankeyNodeType, error) {
	ev := ProductAnalyticsSankeyNodeType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsSankeyNodeType: valid values are %v", v, allowedProductAnalyticsSankeyNodeTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsSankeyNodeType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsSankeyNodeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsSankeyNodeType value.
func (v ProductAnalyticsSankeyNodeType) Ptr() *ProductAnalyticsSankeyNodeType {
	return &v
}
