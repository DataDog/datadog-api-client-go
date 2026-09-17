// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionListRequestType The resource type identifier for a retention list request.
type ProductAnalyticsRetentionListRequestType string

// List of ProductAnalyticsRetentionListRequestType.
const (
	PRODUCTANALYTICSRETENTIONLISTREQUESTTYPE_RETENTION_LIST_REQUEST ProductAnalyticsRetentionListRequestType = "retention_list_request"
)

var allowedProductAnalyticsRetentionListRequestTypeEnumValues = []ProductAnalyticsRetentionListRequestType{
	PRODUCTANALYTICSRETENTIONLISTREQUESTTYPE_RETENTION_LIST_REQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionListRequestType) GetAllowedValues() []ProductAnalyticsRetentionListRequestType {
	return allowedProductAnalyticsRetentionListRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionListRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionListRequestType(value)
	return nil
}

// NewProductAnalyticsRetentionListRequestTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionListRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionListRequestTypeFromValue(v string) (*ProductAnalyticsRetentionListRequestType, error) {
	ev := ProductAnalyticsRetentionListRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionListRequestType: valid values are %v", v, allowedProductAnalyticsRetentionListRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionListRequestType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionListRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionListRequestType value.
func (v ProductAnalyticsRetentionListRequestType) Ptr() *ProductAnalyticsRetentionListRequestType {
	return &v
}
