// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUType The SKU resource type.
type ProductCatalogSKUType string

// List of ProductCatalogSKUType.
const (
	PRODUCTCATALOGSKUTYPE_SKU ProductCatalogSKUType = "Sku"
)

var allowedProductCatalogSKUTypeEnumValues = []ProductCatalogSKUType{
	PRODUCTCATALOGSKUTYPE_SKU,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductCatalogSKUType) GetAllowedValues() []ProductCatalogSKUType {
	return allowedProductCatalogSKUTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductCatalogSKUType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductCatalogSKUType(value)
	return nil
}

// NewProductCatalogSKUTypeFromValue returns a pointer to a valid ProductCatalogSKUType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductCatalogSKUTypeFromValue(v string) (*ProductCatalogSKUType, error) {
	ev := ProductCatalogSKUType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductCatalogSKUType: valid values are %v", v, allowedProductCatalogSKUTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductCatalogSKUType) IsValid() bool {
	for _, existing := range allowedProductCatalogSKUTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductCatalogSKUType value.
func (v ProductCatalogSKUType) Ptr() *ProductCatalogSKUType {
	return &v
}
