// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridResponseType The resource type identifier for a retention grid response.
type ProductAnalyticsRetentionGridResponseType string

// List of ProductAnalyticsRetentionGridResponseType.
const (
	PRODUCTANALYTICSRETENTIONGRIDRESPONSETYPE_RETENTION_GRID_RESPONSE ProductAnalyticsRetentionGridResponseType = "retention_grid_response"
)

var allowedProductAnalyticsRetentionGridResponseTypeEnumValues = []ProductAnalyticsRetentionGridResponseType{
	PRODUCTANALYTICSRETENTIONGRIDRESPONSETYPE_RETENTION_GRID_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionGridResponseType) GetAllowedValues() []ProductAnalyticsRetentionGridResponseType {
	return allowedProductAnalyticsRetentionGridResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionGridResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionGridResponseType(value)
	return nil
}

// NewProductAnalyticsRetentionGridResponseTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionGridResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionGridResponseTypeFromValue(v string) (*ProductAnalyticsRetentionGridResponseType, error) {
	ev := ProductAnalyticsRetentionGridResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionGridResponseType: valid values are %v", v, allowedProductAnalyticsRetentionGridResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionGridResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionGridResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionGridResponseType value.
func (v ProductAnalyticsRetentionGridResponseType) Ptr() *ProductAnalyticsRetentionGridResponseType {
	return &v
}
