// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionListResponseType The resource type identifier for a retention list response.
type ProductAnalyticsRetentionListResponseType string

// List of ProductAnalyticsRetentionListResponseType.
const (
	PRODUCTANALYTICSRETENTIONLISTRESPONSETYPE_RETENTION_LIST_RESPONSE ProductAnalyticsRetentionListResponseType = "retention_list_response"
)

var allowedProductAnalyticsRetentionListResponseTypeEnumValues = []ProductAnalyticsRetentionListResponseType{
	PRODUCTANALYTICSRETENTIONLISTRESPONSETYPE_RETENTION_LIST_RESPONSE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionListResponseType) GetAllowedValues() []ProductAnalyticsRetentionListResponseType {
	return allowedProductAnalyticsRetentionListResponseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionListResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionListResponseType(value)
	return nil
}

// NewProductAnalyticsRetentionListResponseTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionListResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionListResponseTypeFromValue(v string) (*ProductAnalyticsRetentionListResponseType, error) {
	ev := ProductAnalyticsRetentionListResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionListResponseType: valid values are %v", v, allowedProductAnalyticsRetentionListResponseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionListResponseType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionListResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionListResponseType value.
func (v ProductAnalyticsRetentionListResponseType) Ptr() *ProductAnalyticsRetentionListResponseType {
	return &v
}
