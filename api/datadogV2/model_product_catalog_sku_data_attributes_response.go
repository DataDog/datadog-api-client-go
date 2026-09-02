// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUDataAttributesResponse The pricing and allotment metadata of a SKU.
type ProductCatalogSKUDataAttributesResponse struct {
	// The allotments the SKU provides to other SKUs. Every entry carries the code of this
	// SKU as its `parent_sku_code`. Empty when the SKU provides no allotments.
	Allotments []ProductCatalogSKUAllotment `json:"allotments"`
	// The identifier of the billing dimension the SKU is billed on, as used by the usage
	// metering endpoints. Several SKUs can share one billing dimension, so this value does
	// not identify a SKU.
	BillingDimension string `json:"billing_dimension"`
	// The billable usage unit the SKU is priced per. `null` for SKUs that are not priced
	// per unit of usage, such as those whose `pricing_type` is `percent`.
	BillingUnits datadog.NullableString `json:"billing_units"`
	// The ISO-4217 code of the currency the prices are expressed in.
	Currency string `json:"currency"`
	// The billing frequency applied to on-demand usage of the SKU by default.
	DefaultOnDemandOption ProductCatalogSKUOnDemandOption `json:"default_on_demand_option"`
	// The number of billable usage units included in the price. `0` for SKUs that are not
	// priced per unit of usage, such as those whose `pricing_type` is `percent`.
	NumberOfUnitsIncludedInPrice int64 `json:"number_of_units_included_in_price"`
	// The public list price of on-demand usage of the SKU, as a decimal string. The number
	// of decimal places is not normalized, so values such as `0`, `0.9`, and `30000.00`
	// all occur. `null` when the SKU is priced with tiers, in which case the prices are in
	// `on_demand_tiered`.
	OnDemandListPrice datadog.NullableString `json:"on_demand_list_price"`
	// The tiered pricing applied to on-demand usage of the SKU. `null` when the SKU is priced
	// with a single list price instead.
	OnDemandTiered NullableProductCatalogSKUTieredPricing `json:"on_demand_tiered"`
	// How the SKU is priced. `usage` prices each billable usage unit, and `percent` prices a
	// percentage; percent-priced SKUs have no `billing_units`.
	PricingType ProductCatalogSKUPricingType `json:"pricing_type"`
	// The human-readable name of the SKU.
	SkuName string `json:"sku_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductCatalogSKUDataAttributesResponse instantiates a new ProductCatalogSKUDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductCatalogSKUDataAttributesResponse(allotments []ProductCatalogSKUAllotment, billingDimension string, billingUnits datadog.NullableString, currency string, defaultOnDemandOption ProductCatalogSKUOnDemandOption, numberOfUnitsIncludedInPrice int64, onDemandListPrice datadog.NullableString, onDemandTiered NullableProductCatalogSKUTieredPricing, pricingType ProductCatalogSKUPricingType, skuName string) *ProductCatalogSKUDataAttributesResponse {
	this := ProductCatalogSKUDataAttributesResponse{}
	this.Allotments = allotments
	this.BillingDimension = billingDimension
	this.BillingUnits = billingUnits
	this.Currency = currency
	this.DefaultOnDemandOption = defaultOnDemandOption
	this.NumberOfUnitsIncludedInPrice = numberOfUnitsIncludedInPrice
	this.OnDemandListPrice = onDemandListPrice
	this.OnDemandTiered = onDemandTiered
	this.PricingType = pricingType
	this.SkuName = skuName
	return &this
}

// NewProductCatalogSKUDataAttributesResponseWithDefaults instantiates a new ProductCatalogSKUDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductCatalogSKUDataAttributesResponseWithDefaults() *ProductCatalogSKUDataAttributesResponse {
	this := ProductCatalogSKUDataAttributesResponse{}
	return &this
}

// GetAllotments returns the Allotments field value.
func (o *ProductCatalogSKUDataAttributesResponse) GetAllotments() []ProductCatalogSKUAllotment {
	if o == nil {
		var ret []ProductCatalogSKUAllotment
		return ret
	}
	return o.Allotments
}

// GetAllotmentsOk returns a tuple with the Allotments field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUDataAttributesResponse) GetAllotmentsOk() (*[]ProductCatalogSKUAllotment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allotments, true
}

// SetAllotments sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetAllotments(v []ProductCatalogSKUAllotment) {
	o.Allotments = v
}

// GetBillingDimension returns the BillingDimension field value.
func (o *ProductCatalogSKUDataAttributesResponse) GetBillingDimension() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BillingDimension
}

// GetBillingDimensionOk returns a tuple with the BillingDimension field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUDataAttributesResponse) GetBillingDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingDimension, true
}

// SetBillingDimension sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetBillingDimension(v string) {
	o.BillingDimension = v
}

// GetBillingUnits returns the BillingUnits field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ProductCatalogSKUDataAttributesResponse) GetBillingUnits() string {
	if o == nil || o.BillingUnits.Get() == nil {
		var ret string
		return ret
	}
	return *o.BillingUnits.Get()
}

// GetBillingUnitsOk returns a tuple with the BillingUnits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductCatalogSKUDataAttributesResponse) GetBillingUnitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingUnits.Get(), o.BillingUnits.IsSet()
}

// SetBillingUnits sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetBillingUnits(v string) {
	o.BillingUnits.Set(&v)
}

// GetCurrency returns the Currency field value.
func (o *ProductCatalogSKUDataAttributesResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUDataAttributesResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetDefaultOnDemandOption returns the DefaultOnDemandOption field value.
func (o *ProductCatalogSKUDataAttributesResponse) GetDefaultOnDemandOption() ProductCatalogSKUOnDemandOption {
	if o == nil {
		var ret ProductCatalogSKUOnDemandOption
		return ret
	}
	return o.DefaultOnDemandOption
}

// GetDefaultOnDemandOptionOk returns a tuple with the DefaultOnDemandOption field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUDataAttributesResponse) GetDefaultOnDemandOptionOk() (*ProductCatalogSKUOnDemandOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultOnDemandOption, true
}

// SetDefaultOnDemandOption sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetDefaultOnDemandOption(v ProductCatalogSKUOnDemandOption) {
	o.DefaultOnDemandOption = v
}

// GetNumberOfUnitsIncludedInPrice returns the NumberOfUnitsIncludedInPrice field value.
func (o *ProductCatalogSKUDataAttributesResponse) GetNumberOfUnitsIncludedInPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NumberOfUnitsIncludedInPrice
}

// GetNumberOfUnitsIncludedInPriceOk returns a tuple with the NumberOfUnitsIncludedInPrice field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUDataAttributesResponse) GetNumberOfUnitsIncludedInPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfUnitsIncludedInPrice, true
}

// SetNumberOfUnitsIncludedInPrice sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetNumberOfUnitsIncludedInPrice(v int64) {
	o.NumberOfUnitsIncludedInPrice = v
}

// GetOnDemandListPrice returns the OnDemandListPrice field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ProductCatalogSKUDataAttributesResponse) GetOnDemandListPrice() string {
	if o == nil || o.OnDemandListPrice.Get() == nil {
		var ret string
		return ret
	}
	return *o.OnDemandListPrice.Get()
}

// GetOnDemandListPriceOk returns a tuple with the OnDemandListPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductCatalogSKUDataAttributesResponse) GetOnDemandListPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnDemandListPrice.Get(), o.OnDemandListPrice.IsSet()
}

// SetOnDemandListPrice sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetOnDemandListPrice(v string) {
	o.OnDemandListPrice.Set(&v)
}

// GetOnDemandTiered returns the OnDemandTiered field value.
// If the value is explicit nil, the zero value for ProductCatalogSKUTieredPricing will be returned.
func (o *ProductCatalogSKUDataAttributesResponse) GetOnDemandTiered() ProductCatalogSKUTieredPricing {
	if o == nil || o.OnDemandTiered.Get() == nil {
		var ret ProductCatalogSKUTieredPricing
		return ret
	}
	return *o.OnDemandTiered.Get()
}

// GetOnDemandTieredOk returns a tuple with the OnDemandTiered field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductCatalogSKUDataAttributesResponse) GetOnDemandTieredOk() (*ProductCatalogSKUTieredPricing, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnDemandTiered.Get(), o.OnDemandTiered.IsSet()
}

// SetOnDemandTiered sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetOnDemandTiered(v ProductCatalogSKUTieredPricing) {
	o.OnDemandTiered.Set(&v)
}

// GetPricingType returns the PricingType field value.
func (o *ProductCatalogSKUDataAttributesResponse) GetPricingType() ProductCatalogSKUPricingType {
	if o == nil {
		var ret ProductCatalogSKUPricingType
		return ret
	}
	return o.PricingType
}

// GetPricingTypeOk returns a tuple with the PricingType field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUDataAttributesResponse) GetPricingTypeOk() (*ProductCatalogSKUPricingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingType, true
}

// SetPricingType sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetPricingType(v ProductCatalogSKUPricingType) {
	o.PricingType = v
}

// GetSkuName returns the SkuName field value.
func (o *ProductCatalogSKUDataAttributesResponse) GetSkuName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SkuName
}

// GetSkuNameOk returns a tuple with the SkuName field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUDataAttributesResponse) GetSkuNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuName, true
}

// SetSkuName sets field value.
func (o *ProductCatalogSKUDataAttributesResponse) SetSkuName(v string) {
	o.SkuName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductCatalogSKUDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["allotments"] = o.Allotments
	toSerialize["billing_dimension"] = o.BillingDimension
	toSerialize["billing_units"] = o.BillingUnits.Get()
	toSerialize["currency"] = o.Currency
	toSerialize["default_on_demand_option"] = o.DefaultOnDemandOption
	toSerialize["number_of_units_included_in_price"] = o.NumberOfUnitsIncludedInPrice
	toSerialize["on_demand_list_price"] = o.OnDemandListPrice.Get()
	toSerialize["on_demand_tiered"] = o.OnDemandTiered.Get()
	toSerialize["pricing_type"] = o.PricingType
	toSerialize["sku_name"] = o.SkuName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductCatalogSKUDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Allotments                   *[]ProductCatalogSKUAllotment          `json:"allotments"`
		BillingDimension             *string                                `json:"billing_dimension"`
		BillingUnits                 datadog.NullableString                 `json:"billing_units"`
		Currency                     *string                                `json:"currency"`
		DefaultOnDemandOption        *ProductCatalogSKUOnDemandOption       `json:"default_on_demand_option"`
		NumberOfUnitsIncludedInPrice *int64                                 `json:"number_of_units_included_in_price"`
		OnDemandListPrice            datadog.NullableString                 `json:"on_demand_list_price"`
		OnDemandTiered               NullableProductCatalogSKUTieredPricing `json:"on_demand_tiered"`
		PricingType                  *ProductCatalogSKUPricingType          `json:"pricing_type"`
		SkuName                      *string                                `json:"sku_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Allotments == nil {
		return fmt.Errorf("required field allotments missing")
	}
	if all.BillingDimension == nil {
		return fmt.Errorf("required field billing_dimension missing")
	}
	if !all.BillingUnits.IsSet() {
		return fmt.Errorf("required field billing_units missing")
	}
	if all.Currency == nil {
		return fmt.Errorf("required field currency missing")
	}
	if all.DefaultOnDemandOption == nil {
		return fmt.Errorf("required field default_on_demand_option missing")
	}
	if all.NumberOfUnitsIncludedInPrice == nil {
		return fmt.Errorf("required field number_of_units_included_in_price missing")
	}
	if !all.OnDemandListPrice.IsSet() {
		return fmt.Errorf("required field on_demand_list_price missing")
	}
	if !all.OnDemandTiered.IsSet() {
		return fmt.Errorf("required field on_demand_tiered missing")
	}
	if all.PricingType == nil {
		return fmt.Errorf("required field pricing_type missing")
	}
	if all.SkuName == nil {
		return fmt.Errorf("required field sku_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allotments", "billing_dimension", "billing_units", "currency", "default_on_demand_option", "number_of_units_included_in_price", "on_demand_list_price", "on_demand_tiered", "pricing_type", "sku_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Allotments = *all.Allotments
	o.BillingDimension = *all.BillingDimension
	o.BillingUnits = all.BillingUnits
	o.Currency = *all.Currency
	if !all.DefaultOnDemandOption.IsValid() {
		hasInvalidField = true
	} else {
		o.DefaultOnDemandOption = *all.DefaultOnDemandOption
	}
	o.NumberOfUnitsIncludedInPrice = *all.NumberOfUnitsIncludedInPrice
	o.OnDemandListPrice = all.OnDemandListPrice
	o.OnDemandTiered = all.OnDemandTiered
	if !all.PricingType.IsValid() {
		hasInvalidField = true
	} else {
		o.PricingType = *all.PricingType
	}
	o.SkuName = *all.SkuName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
