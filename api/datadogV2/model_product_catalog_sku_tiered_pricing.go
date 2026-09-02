// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUTieredPricing The tiered pricing applied to on-demand usage of the SKU. `null` when the SKU is priced
// with a single list price instead.
type ProductCatalogSKUTieredPricing struct {
	// The pricing tiers, ordered by ascending usage quantity.
	Tiers []ProductCatalogSKUPricingTier `json:"tiers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductCatalogSKUTieredPricing instantiates a new ProductCatalogSKUTieredPricing object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductCatalogSKUTieredPricing(tiers []ProductCatalogSKUPricingTier) *ProductCatalogSKUTieredPricing {
	this := ProductCatalogSKUTieredPricing{}
	this.Tiers = tiers
	return &this
}

// NewProductCatalogSKUTieredPricingWithDefaults instantiates a new ProductCatalogSKUTieredPricing object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductCatalogSKUTieredPricingWithDefaults() *ProductCatalogSKUTieredPricing {
	this := ProductCatalogSKUTieredPricing{}
	return &this
}

// GetTiers returns the Tiers field value.
func (o *ProductCatalogSKUTieredPricing) GetTiers() []ProductCatalogSKUPricingTier {
	if o == nil {
		var ret []ProductCatalogSKUPricingTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUTieredPricing) GetTiersOk() (*[]ProductCatalogSKUPricingTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tiers, true
}

// SetTiers sets field value.
func (o *ProductCatalogSKUTieredPricing) SetTiers(v []ProductCatalogSKUPricingTier) {
	o.Tiers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductCatalogSKUTieredPricing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["tiers"] = o.Tiers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductCatalogSKUTieredPricing) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tiers *[]ProductCatalogSKUPricingTier `json:"tiers"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Tiers == nil {
		return fmt.Errorf("required field tiers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tiers"})
	} else {
		return err
	}
	o.Tiers = *all.Tiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableProductCatalogSKUTieredPricing handles when a null is used for ProductCatalogSKUTieredPricing.
type NullableProductCatalogSKUTieredPricing struct {
	value *ProductCatalogSKUTieredPricing
	isSet bool
}

// Get returns the associated value.
func (v NullableProductCatalogSKUTieredPricing) Get() *ProductCatalogSKUTieredPricing {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableProductCatalogSKUTieredPricing) Set(val *ProductCatalogSKUTieredPricing) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableProductCatalogSKUTieredPricing) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableProductCatalogSKUTieredPricing) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableProductCatalogSKUTieredPricing initializes the struct as if Set has been called.
func NewNullableProductCatalogSKUTieredPricing(val *ProductCatalogSKUTieredPricing) *NullableProductCatalogSKUTieredPricing {
	return &NullableProductCatalogSKUTieredPricing{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableProductCatalogSKUTieredPricing) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableProductCatalogSKUTieredPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
