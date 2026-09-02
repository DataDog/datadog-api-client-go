// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUPricingTier A usage range and the price that applies to usage falling inside it.
type ProductCatalogSKUPricingTier struct {
	// The exclusive upper bound of the usage range the tier prices. `null` on the final
	// tier, which is unbounded.
	MaxUsageQuantity datadog.NullableInt64 `json:"max_usage_quantity"`
	// The inclusive lower bound of the usage range the tier prices.
	MinUsageQuantity int64 `json:"min_usage_quantity"`
	// The price applied to usage in the tier, as a decimal string. The number of decimal
	// places is not normalized, so free tiers appear as either `0` or `0.00`.
	Price string `json:"price"`
	// Whether the tier's price applies per unit of usage or to a block of usage.
	PricingUnitType ProductCatalogSKUPricingUnitType `json:"pricing_unit_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductCatalogSKUPricingTier instantiates a new ProductCatalogSKUPricingTier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductCatalogSKUPricingTier(maxUsageQuantity datadog.NullableInt64, minUsageQuantity int64, price string, pricingUnitType ProductCatalogSKUPricingUnitType) *ProductCatalogSKUPricingTier {
	this := ProductCatalogSKUPricingTier{}
	this.MaxUsageQuantity = maxUsageQuantity
	this.MinUsageQuantity = minUsageQuantity
	this.Price = price
	this.PricingUnitType = pricingUnitType
	return &this
}

// NewProductCatalogSKUPricingTierWithDefaults instantiates a new ProductCatalogSKUPricingTier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductCatalogSKUPricingTierWithDefaults() *ProductCatalogSKUPricingTier {
	this := ProductCatalogSKUPricingTier{}
	return &this
}

// GetMaxUsageQuantity returns the MaxUsageQuantity field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ProductCatalogSKUPricingTier) GetMaxUsageQuantity() int64 {
	if o == nil || o.MaxUsageQuantity.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MaxUsageQuantity.Get()
}

// GetMaxUsageQuantityOk returns a tuple with the MaxUsageQuantity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductCatalogSKUPricingTier) GetMaxUsageQuantityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxUsageQuantity.Get(), o.MaxUsageQuantity.IsSet()
}

// SetMaxUsageQuantity sets field value.
func (o *ProductCatalogSKUPricingTier) SetMaxUsageQuantity(v int64) {
	o.MaxUsageQuantity.Set(&v)
}

// GetMinUsageQuantity returns the MinUsageQuantity field value.
func (o *ProductCatalogSKUPricingTier) GetMinUsageQuantity() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MinUsageQuantity
}

// GetMinUsageQuantityOk returns a tuple with the MinUsageQuantity field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUPricingTier) GetMinUsageQuantityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinUsageQuantity, true
}

// SetMinUsageQuantity sets field value.
func (o *ProductCatalogSKUPricingTier) SetMinUsageQuantity(v int64) {
	o.MinUsageQuantity = v
}

// GetPrice returns the Price field value.
func (o *ProductCatalogSKUPricingTier) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUPricingTier) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value.
func (o *ProductCatalogSKUPricingTier) SetPrice(v string) {
	o.Price = v
}

// GetPricingUnitType returns the PricingUnitType field value.
func (o *ProductCatalogSKUPricingTier) GetPricingUnitType() ProductCatalogSKUPricingUnitType {
	if o == nil {
		var ret ProductCatalogSKUPricingUnitType
		return ret
	}
	return o.PricingUnitType
}

// GetPricingUnitTypeOk returns a tuple with the PricingUnitType field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUPricingTier) GetPricingUnitTypeOk() (*ProductCatalogSKUPricingUnitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingUnitType, true
}

// SetPricingUnitType sets field value.
func (o *ProductCatalogSKUPricingTier) SetPricingUnitType(v ProductCatalogSKUPricingUnitType) {
	o.PricingUnitType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductCatalogSKUPricingTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["max_usage_quantity"] = o.MaxUsageQuantity.Get()
	toSerialize["min_usage_quantity"] = o.MinUsageQuantity
	toSerialize["price"] = o.Price
	toSerialize["pricing_unit_type"] = o.PricingUnitType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductCatalogSKUPricingTier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxUsageQuantity datadog.NullableInt64             `json:"max_usage_quantity"`
		MinUsageQuantity *int64                            `json:"min_usage_quantity"`
		Price            *string                           `json:"price"`
		PricingUnitType  *ProductCatalogSKUPricingUnitType `json:"pricing_unit_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.MaxUsageQuantity.IsSet() {
		return fmt.Errorf("required field max_usage_quantity missing")
	}
	if all.MinUsageQuantity == nil {
		return fmt.Errorf("required field min_usage_quantity missing")
	}
	if all.Price == nil {
		return fmt.Errorf("required field price missing")
	}
	if all.PricingUnitType == nil {
		return fmt.Errorf("required field pricing_unit_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_usage_quantity", "min_usage_quantity", "price", "pricing_unit_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MaxUsageQuantity = all.MaxUsageQuantity
	o.MinUsageQuantity = *all.MinUsageQuantity
	o.Price = *all.Price
	if !all.PricingUnitType.IsValid() {
		hasInvalidField = true
	} else {
		o.PricingUnitType = *all.PricingUnitType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
