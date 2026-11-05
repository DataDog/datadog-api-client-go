// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingRates Pricing rates for a model.
type AiGatewayModelPricingRates struct {
	// Cost per million tokens for cache reads.
	CacheReadCostPerMtok *float64 `json:"cache_read_cost_per_mtok,omitempty"`
	// Cost per million tokens for cache writes.
	CacheWriteCostPerMtok *float64 `json:"cache_write_cost_per_mtok,omitempty"`
	// Cost per million tokens for non-cached input.
	NonCachedInputCostPerMtok *float64 `json:"non_cached_input_cost_per_mtok,omitempty"`
	// Cost per million tokens for output.
	OutputCostPerMtok *float64 `json:"output_cost_per_mtok,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingRates instantiates a new AiGatewayModelPricingRates object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingRates() *AiGatewayModelPricingRates {
	this := AiGatewayModelPricingRates{}
	return &this
}

// NewAiGatewayModelPricingRatesWithDefaults instantiates a new AiGatewayModelPricingRates object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingRatesWithDefaults() *AiGatewayModelPricingRates {
	this := AiGatewayModelPricingRates{}
	return &this
}

// GetCacheReadCostPerMtok returns the CacheReadCostPerMtok field value if set, zero value otherwise.
func (o *AiGatewayModelPricingRates) GetCacheReadCostPerMtok() float64 {
	if o == nil || o.CacheReadCostPerMtok == nil {
		var ret float64
		return ret
	}
	return *o.CacheReadCostPerMtok
}

// GetCacheReadCostPerMtokOk returns a tuple with the CacheReadCostPerMtok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRates) GetCacheReadCostPerMtokOk() (*float64, bool) {
	if o == nil || o.CacheReadCostPerMtok == nil {
		return nil, false
	}
	return o.CacheReadCostPerMtok, true
}

// HasCacheReadCostPerMtok returns a boolean if a field has been set.
func (o *AiGatewayModelPricingRates) HasCacheReadCostPerMtok() bool {
	return o != nil && o.CacheReadCostPerMtok != nil
}

// SetCacheReadCostPerMtok gets a reference to the given float64 and assigns it to the CacheReadCostPerMtok field.
func (o *AiGatewayModelPricingRates) SetCacheReadCostPerMtok(v float64) {
	o.CacheReadCostPerMtok = &v
}

// GetCacheWriteCostPerMtok returns the CacheWriteCostPerMtok field value if set, zero value otherwise.
func (o *AiGatewayModelPricingRates) GetCacheWriteCostPerMtok() float64 {
	if o == nil || o.CacheWriteCostPerMtok == nil {
		var ret float64
		return ret
	}
	return *o.CacheWriteCostPerMtok
}

// GetCacheWriteCostPerMtokOk returns a tuple with the CacheWriteCostPerMtok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRates) GetCacheWriteCostPerMtokOk() (*float64, bool) {
	if o == nil || o.CacheWriteCostPerMtok == nil {
		return nil, false
	}
	return o.CacheWriteCostPerMtok, true
}

// HasCacheWriteCostPerMtok returns a boolean if a field has been set.
func (o *AiGatewayModelPricingRates) HasCacheWriteCostPerMtok() bool {
	return o != nil && o.CacheWriteCostPerMtok != nil
}

// SetCacheWriteCostPerMtok gets a reference to the given float64 and assigns it to the CacheWriteCostPerMtok field.
func (o *AiGatewayModelPricingRates) SetCacheWriteCostPerMtok(v float64) {
	o.CacheWriteCostPerMtok = &v
}

// GetNonCachedInputCostPerMtok returns the NonCachedInputCostPerMtok field value if set, zero value otherwise.
func (o *AiGatewayModelPricingRates) GetNonCachedInputCostPerMtok() float64 {
	if o == nil || o.NonCachedInputCostPerMtok == nil {
		var ret float64
		return ret
	}
	return *o.NonCachedInputCostPerMtok
}

// GetNonCachedInputCostPerMtokOk returns a tuple with the NonCachedInputCostPerMtok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRates) GetNonCachedInputCostPerMtokOk() (*float64, bool) {
	if o == nil || o.NonCachedInputCostPerMtok == nil {
		return nil, false
	}
	return o.NonCachedInputCostPerMtok, true
}

// HasNonCachedInputCostPerMtok returns a boolean if a field has been set.
func (o *AiGatewayModelPricingRates) HasNonCachedInputCostPerMtok() bool {
	return o != nil && o.NonCachedInputCostPerMtok != nil
}

// SetNonCachedInputCostPerMtok gets a reference to the given float64 and assigns it to the NonCachedInputCostPerMtok field.
func (o *AiGatewayModelPricingRates) SetNonCachedInputCostPerMtok(v float64) {
	o.NonCachedInputCostPerMtok = &v
}

// GetOutputCostPerMtok returns the OutputCostPerMtok field value if set, zero value otherwise.
func (o *AiGatewayModelPricingRates) GetOutputCostPerMtok() float64 {
	if o == nil || o.OutputCostPerMtok == nil {
		var ret float64
		return ret
	}
	return *o.OutputCostPerMtok
}

// GetOutputCostPerMtokOk returns a tuple with the OutputCostPerMtok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRates) GetOutputCostPerMtokOk() (*float64, bool) {
	if o == nil || o.OutputCostPerMtok == nil {
		return nil, false
	}
	return o.OutputCostPerMtok, true
}

// HasOutputCostPerMtok returns a boolean if a field has been set.
func (o *AiGatewayModelPricingRates) HasOutputCostPerMtok() bool {
	return o != nil && o.OutputCostPerMtok != nil
}

// SetOutputCostPerMtok gets a reference to the given float64 and assigns it to the OutputCostPerMtok field.
func (o *AiGatewayModelPricingRates) SetOutputCostPerMtok(v float64) {
	o.OutputCostPerMtok = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingRates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CacheReadCostPerMtok != nil {
		toSerialize["cache_read_cost_per_mtok"] = o.CacheReadCostPerMtok
	}
	if o.CacheWriteCostPerMtok != nil {
		toSerialize["cache_write_cost_per_mtok"] = o.CacheWriteCostPerMtok
	}
	if o.NonCachedInputCostPerMtok != nil {
		toSerialize["non_cached_input_cost_per_mtok"] = o.NonCachedInputCostPerMtok
	}
	if o.OutputCostPerMtok != nil {
		toSerialize["output_cost_per_mtok"] = o.OutputCostPerMtok
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingRates) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CacheReadCostPerMtok      *float64 `json:"cache_read_cost_per_mtok,omitempty"`
		CacheWriteCostPerMtok     *float64 `json:"cache_write_cost_per_mtok,omitempty"`
		NonCachedInputCostPerMtok *float64 `json:"non_cached_input_cost_per_mtok,omitempty"`
		OutputCostPerMtok         *float64 `json:"output_cost_per_mtok,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.CacheReadCostPerMtok = all.CacheReadCostPerMtok
	o.CacheWriteCostPerMtok = all.CacheWriteCostPerMtok
	o.NonCachedInputCostPerMtok = all.NonCachedInputCostPerMtok
	o.OutputCostPerMtok = all.OutputCostPerMtok

	return nil
}
