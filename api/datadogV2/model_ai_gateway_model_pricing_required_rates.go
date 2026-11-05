// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingRequiredRates Complete set of rates required for custom pricing.
type AiGatewayModelPricingRequiredRates struct {
	// Cost per million tokens for cache reads.
	CacheReadCostPerMtok float64 `json:"cache_read_cost_per_mtok"`
	// Cost per million tokens for cache writes.
	CacheWriteCostPerMtok float64 `json:"cache_write_cost_per_mtok"`
	// Cost per million tokens for non-cached input.
	NonCachedInputCostPerMtok float64 `json:"non_cached_input_cost_per_mtok"`
	// Cost per million tokens for output.
	OutputCostPerMtok float64 `json:"output_cost_per_mtok"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingRequiredRates instantiates a new AiGatewayModelPricingRequiredRates object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingRequiredRates(cacheReadCostPerMtok float64, cacheWriteCostPerMtok float64, nonCachedInputCostPerMtok float64, outputCostPerMtok float64) *AiGatewayModelPricingRequiredRates {
	this := AiGatewayModelPricingRequiredRates{}
	this.CacheReadCostPerMtok = cacheReadCostPerMtok
	this.CacheWriteCostPerMtok = cacheWriteCostPerMtok
	this.NonCachedInputCostPerMtok = nonCachedInputCostPerMtok
	this.OutputCostPerMtok = outputCostPerMtok
	return &this
}

// NewAiGatewayModelPricingRequiredRatesWithDefaults instantiates a new AiGatewayModelPricingRequiredRates object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingRequiredRatesWithDefaults() *AiGatewayModelPricingRequiredRates {
	this := AiGatewayModelPricingRequiredRates{}
	return &this
}

// GetCacheReadCostPerMtok returns the CacheReadCostPerMtok field value.
func (o *AiGatewayModelPricingRequiredRates) GetCacheReadCostPerMtok() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CacheReadCostPerMtok
}

// GetCacheReadCostPerMtokOk returns a tuple with the CacheReadCostPerMtok field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRequiredRates) GetCacheReadCostPerMtokOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheReadCostPerMtok, true
}

// SetCacheReadCostPerMtok sets field value.
func (o *AiGatewayModelPricingRequiredRates) SetCacheReadCostPerMtok(v float64) {
	o.CacheReadCostPerMtok = v
}

// GetCacheWriteCostPerMtok returns the CacheWriteCostPerMtok field value.
func (o *AiGatewayModelPricingRequiredRates) GetCacheWriteCostPerMtok() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CacheWriteCostPerMtok
}

// GetCacheWriteCostPerMtokOk returns a tuple with the CacheWriteCostPerMtok field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRequiredRates) GetCacheWriteCostPerMtokOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheWriteCostPerMtok, true
}

// SetCacheWriteCostPerMtok sets field value.
func (o *AiGatewayModelPricingRequiredRates) SetCacheWriteCostPerMtok(v float64) {
	o.CacheWriteCostPerMtok = v
}

// GetNonCachedInputCostPerMtok returns the NonCachedInputCostPerMtok field value.
func (o *AiGatewayModelPricingRequiredRates) GetNonCachedInputCostPerMtok() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.NonCachedInputCostPerMtok
}

// GetNonCachedInputCostPerMtokOk returns a tuple with the NonCachedInputCostPerMtok field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRequiredRates) GetNonCachedInputCostPerMtokOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonCachedInputCostPerMtok, true
}

// SetNonCachedInputCostPerMtok sets field value.
func (o *AiGatewayModelPricingRequiredRates) SetNonCachedInputCostPerMtok(v float64) {
	o.NonCachedInputCostPerMtok = v
}

// GetOutputCostPerMtok returns the OutputCostPerMtok field value.
func (o *AiGatewayModelPricingRequiredRates) GetOutputCostPerMtok() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.OutputCostPerMtok
}

// GetOutputCostPerMtokOk returns a tuple with the OutputCostPerMtok field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingRequiredRates) GetOutputCostPerMtokOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputCostPerMtok, true
}

// SetOutputCostPerMtok sets field value.
func (o *AiGatewayModelPricingRequiredRates) SetOutputCostPerMtok(v float64) {
	o.OutputCostPerMtok = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingRequiredRates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cache_read_cost_per_mtok"] = o.CacheReadCostPerMtok
	toSerialize["cache_write_cost_per_mtok"] = o.CacheWriteCostPerMtok
	toSerialize["non_cached_input_cost_per_mtok"] = o.NonCachedInputCostPerMtok
	toSerialize["output_cost_per_mtok"] = o.OutputCostPerMtok
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingRequiredRates) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CacheReadCostPerMtok      *float64 `json:"cache_read_cost_per_mtok"`
		CacheWriteCostPerMtok     *float64 `json:"cache_write_cost_per_mtok"`
		NonCachedInputCostPerMtok *float64 `json:"non_cached_input_cost_per_mtok"`
		OutputCostPerMtok         *float64 `json:"output_cost_per_mtok"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CacheReadCostPerMtok == nil {
		return fmt.Errorf("required field cache_read_cost_per_mtok missing")
	}
	if all.CacheWriteCostPerMtok == nil {
		return fmt.Errorf("required field cache_write_cost_per_mtok missing")
	}
	if all.NonCachedInputCostPerMtok == nil {
		return fmt.Errorf("required field non_cached_input_cost_per_mtok missing")
	}
	if all.OutputCostPerMtok == nil {
		return fmt.Errorf("required field output_cost_per_mtok missing")
	}
	o.CacheReadCostPerMtok = *all.CacheReadCostPerMtok
	o.CacheWriteCostPerMtok = *all.CacheWriteCostPerMtok
	o.NonCachedInputCostPerMtok = *all.NonCachedInputCostPerMtok
	o.OutputCostPerMtok = *all.OutputCostPerMtok

	return nil
}
