// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionEntity The entity whose retention is measured.
type ProductAnalyticsRetentionEntity string

// List of ProductAnalyticsRetentionEntity.
const (
	PRODUCTANALYTICSRETENTIONENTITY_USER_ID    ProductAnalyticsRetentionEntity = "@usr.id"
	PRODUCTANALYTICSRETENTIONENTITY_ACCOUNT_ID ProductAnalyticsRetentionEntity = "@account.id"
)

var allowedProductAnalyticsRetentionEntityEnumValues = []ProductAnalyticsRetentionEntity{
	PRODUCTANALYTICSRETENTIONENTITY_USER_ID,
	PRODUCTANALYTICSRETENTIONENTITY_ACCOUNT_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductAnalyticsRetentionEntity) GetAllowedValues() []ProductAnalyticsRetentionEntity {
	return allowedProductAnalyticsRetentionEntityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductAnalyticsRetentionEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductAnalyticsRetentionEntity(value)
	return nil
}

// NewProductAnalyticsRetentionEntityFromValue returns a pointer to a valid ProductAnalyticsRetentionEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductAnalyticsRetentionEntityFromValue(v string) (*ProductAnalyticsRetentionEntity, error) {
	ev := ProductAnalyticsRetentionEntity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductAnalyticsRetentionEntity: valid values are %v", v, allowedProductAnalyticsRetentionEntityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductAnalyticsRetentionEntity) IsValid() bool {
	for _, existing := range allowedProductAnalyticsRetentionEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductAnalyticsRetentionEntity value.
func (v ProductAnalyticsRetentionEntity) Ptr() *ProductAnalyticsRetentionEntity {
	return &v
}
