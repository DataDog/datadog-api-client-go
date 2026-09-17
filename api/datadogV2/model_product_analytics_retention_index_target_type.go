// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionIndexTargetType The discriminator identifying a target selected by index.
type ProductAnalyticsRetentionIndexTargetType string

// List of ProductAnalyticsRetentionIndexTargetType.
const (
	PRODUCTANALYTICSRETENTIONINDEXTARGETTYPE_INDEX ProductAnalyticsRetentionIndexTargetType = "index"
)

var allowedProductAnalyticsRetentionIndexTargetTypeEnumValues = []ProductAnalyticsRetentionIndexTargetType{
	PRODUCTANALYTICSRETENTIONINDEXTARGETTYPE_INDEX,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionIndexTargetType) GetAllowedValues() []ProductAnalyticsRetentionIndexTargetType {
	return allowedProductAnalyticsRetentionIndexTargetTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionIndexTargetType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionIndexTargetType(value)
	return nil
}

// NewProductAnalyticsRetentionIndexTargetTypeFromValue returns a pointer to a valid ProductAnalyticsRetentionIndexTargetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionIndexTargetTypeFromValue(v string) (*ProductAnalyticsRetentionIndexTargetType, error) {
	ev := ProductAnalyticsRetentionIndexTargetType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionIndexTargetType: valid values are %v", v, allowedProductAnalyticsRetentionIndexTargetTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionIndexTargetType) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionIndexTargetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionIndexTargetType value.
func (v ProductAnalyticsRetentionIndexTargetType) Ptr() *ProductAnalyticsRetentionIndexTargetType {
	return &v
}
