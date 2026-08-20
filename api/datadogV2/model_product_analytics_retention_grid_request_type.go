// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridRequestType The resource type identifier for a retention grid request.
type ProductAnalyticsRetentionGridRequestType string

// List of ProductAnalyticsRetentionGridRequestType.
const (
	PRODUCTANALYTICSRETENTIONGRIDREQUESTTYPE_RETENTION_GRID_REQUEST ProductAnalyticsRetentionGridRequestType = "retention_grid_request"
)

var allowedProductAnalyticsRetentionGridRequestTypeEnumValues = []ProductAnalyticsRetentionGridRequestType{
	PRODUCTANALYTICSRETENTIONGRIDREQUESTTYPE_RETENTION_GRID_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionGridRequestType) GetAllowedValues() []ProductAnalyticsRetentionGridRequestType {
	return allowedProductAnalyticsRetentionGridRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionGridRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionGridRequestType(value)
	return nil
}

// NewProductAnalyticsRetentionGridRequestTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionGridRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionGridRequestTypeFromValue(v string) (*ProductAnalyticsRetentionGridRequestType, error) {
	ev := ProductAnalyticsRetentionGridRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionGridRequestType: valid values are %v", v, allowedProductAnalyticsRetentionGridRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionGridRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionGridRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionGridRequestType value.
func (v ProductAnalyticsRetentionGridRequestType) Ptr() *ProductAnalyticsRetentionGridRequestType {
	return &v
}
