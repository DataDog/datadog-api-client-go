// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUsAPIVersion The version of the product catalog contract to return.
type ProductCatalogSKUsAPIVersion string

// List of ProductCatalogSKUsAPIVersion.
const (
	PRODUCTCATALOGSKUSAPIVERSION_V1 ProductCatalogSKUsAPIVersion = "v1"
)

var allowedProductCatalogSKUsAPIVersionEnumValues = []ProductCatalogSKUsAPIVersion{
	PRODUCTCATALOGSKUSAPIVERSION_V1,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductCatalogSKUsAPIVersion) GetAllowedValues() []ProductCatalogSKUsAPIVersion {
	return allowedProductCatalogSKUsAPIVersionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductCatalogSKUsAPIVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductCatalogSKUsAPIVersion(value)
	return nil
}

// NewProductCatalogSKUsAPIVersionFromValue returns a pointer to a valid ProductCatalogSKUsAPIVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductCatalogSKUsAPIVersionFromValue(v string) (*ProductCatalogSKUsAPIVersion, error) {
	ev := ProductCatalogSKUsAPIVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductCatalogSKUsAPIVersion: valid values are %v", v, allowedProductCatalogSKUsAPIVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductCatalogSKUsAPIVersion) IsValid() bool {
	for _, existing := range allowedProductCatalogSKUsAPIVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductCatalogSKUsAPIVersion value.
func (v ProductCatalogSKUsAPIVersion) Ptr() *ProductCatalogSKUsAPIVersion {
	return &v
}
