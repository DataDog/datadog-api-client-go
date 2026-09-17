// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUPricingType How the SKU is priced. `usage` prices each billable usage unit, and `percent` prices a
// percentage; percent-priced SKUs have no `billing_units`.
type ProductCatalogSKUPricingType string

// List of ProductCatalogSKUPricingType.
const (
	PRODUCTCATALOGSKUPRICINGTYPE_USAGE   ProductCatalogSKUPricingType = "usage"
	PRODUCTCATALOGSKUPRICINGTYPE_PERCENT ProductCatalogSKUPricingType = "percent"
)

var allowedProductCatalogSKUPricingTypeEnumValues = []ProductCatalogSKUPricingType{
	PRODUCTCATALOGSKUPRICINGTYPE_USAGE,
	PRODUCTCATALOGSKUPRICINGTYPE_PERCENT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductCatalogSKUPricingType) GetAllowedValues() []ProductCatalogSKUPricingType {
	return allowedProductCatalogSKUPricingTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductCatalogSKUPricingType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductCatalogSKUPricingType(value)
	return nil
}

// NewProductCatalogSKUPricingTypeFromValue returns a pointer to a valid ProductCatalogSKUPricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductCatalogSKUPricingTypeFromValue(v string) (*ProductCatalogSKUPricingType, error) {
	ev := ProductCatalogSKUPricingType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductCatalogSKUPricingType: valid values are %v", v, allowedProductCatalogSKUPricingTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductCatalogSKUPricingType) IsValid() bool {
	for _, existing := range allowedProductCatalogSKUPricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductCatalogSKUPricingType value.
func (v ProductCatalogSKUPricingType) Ptr() *ProductCatalogSKUPricingType {
	return &v
}
