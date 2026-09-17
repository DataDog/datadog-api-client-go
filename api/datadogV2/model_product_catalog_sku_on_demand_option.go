// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUOnDemandOption The billing frequency applied to on-demand usage of the SKU by default.
type ProductCatalogSKUOnDemandOption string

// List of ProductCatalogSKUOnDemandOption.
const (
	PRODUCTCATALOGSKUONDEMANDOPTION_HOURLY  ProductCatalogSKUOnDemandOption = "hourly"
	PRODUCTCATALOGSKUONDEMANDOPTION_MONTHLY ProductCatalogSKUOnDemandOption = "monthly"
)

var allowedProductCatalogSKUOnDemandOptionEnumValues = []ProductCatalogSKUOnDemandOption{
	PRODUCTCATALOGSKUONDEMANDOPTION_HOURLY,
	PRODUCTCATALOGSKUONDEMANDOPTION_MONTHLY,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ProductCatalogSKUOnDemandOption) GetAllowedValues() []ProductCatalogSKUOnDemandOption {
	return allowedProductCatalogSKUOnDemandOptionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ProductCatalogSKUOnDemandOption) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ProductCatalogSKUOnDemandOption(value)
	return nil
}

// NewProductCatalogSKUOnDemandOptionFromValue returns a pointer to a valid ProductCatalogSKUOnDemandOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewProductCatalogSKUOnDemandOptionFromValue(v string) (*ProductCatalogSKUOnDemandOption, error) {
	ev := ProductCatalogSKUOnDemandOption(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ProductCatalogSKUOnDemandOption: valid values are %v", v, allowedProductCatalogSKUOnDemandOptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ProductCatalogSKUOnDemandOption) IsValid() bool {
	for _, existing := range allowedProductCatalogSKUOnDemandOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductCatalogSKUOnDemandOption value.
func (v ProductCatalogSKUOnDemandOption) Ptr() *ProductCatalogSKUOnDemandOption {
	return &v
}
