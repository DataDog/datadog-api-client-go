// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUPricingUnitType Whether the tier's price applies per unit of usage or to a block of usage.
type ProductCatalogSKUPricingUnitType string

// List of ProductCatalogSKUPricingUnitType.
const (
	PRODUCTCATALOGSKUPRICINGUNITTYPE_BLOCK ProductCatalogSKUPricingUnitType = "block"
	PRODUCTCATALOGSKUPRICINGUNITTYPE_UNIT  ProductCatalogSKUPricingUnitType = "unit"
)

var allowedProductCatalogSKUPricingUnitTypeEnumValues = []ProductCatalogSKUPricingUnitType{
	PRODUCTCATALOGSKUPRICINGUNITTYPE_BLOCK,
	PRODUCTCATALOGSKUPRICINGUNITTYPE_UNIT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductCatalogSKUPricingUnitType) GetAllowedValues() []ProductCatalogSKUPricingUnitType {
	return allowedProductCatalogSKUPricingUnitTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductCatalogSKUPricingUnitType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductCatalogSKUPricingUnitType(value)
	return nil
}

// NewProductCatalogSKUPricingUnitTypeFromValue returns a pointer to a valid ProductCatalogSKUPricingUnitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductCatalogSKUPricingUnitTypeFromValue(v string) (*ProductCatalogSKUPricingUnitType, error) {
	ev := ProductCatalogSKUPricingUnitType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductCatalogSKUPricingUnitType: valid values are %v", v, allowedProductCatalogSKUPricingUnitTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductCatalogSKUPricingUnitType) IsValid() bool {
	for _, existing := range allowedProductCatalogSKUPricingUnitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductCatalogSKUPricingUnitType value.
func (v ProductCatalogSKUPricingUnitType) Ptr() *ProductCatalogSKUPricingUnitType {
	return &v
}
