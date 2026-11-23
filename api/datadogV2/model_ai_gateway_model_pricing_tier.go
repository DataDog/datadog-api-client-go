// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingTier Custom price for requests at or above `context_min_tokens` input tokens, in `service_tier` when it is set.
// Each tier needs a `service_tier`, a positive `context_min_tokens`, or both. Rates it omits keep the base custom rate.
type AiGatewayModelPricingTier struct {
	// Minimum input tokens for the price to apply.
	ContextMinTokens *int64 `json:"context_min_tokens,omitempty"`
	// Pricing rates for a model.
	Rates AiGatewayModelPricingRates `json:"rates"`
	// Provider service tier the price applies to. Lowercase letters, digits, `_` and `-`.
	ServiceTier *string `json:"service_tier,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingTier instantiates a new AiGatewayModelPricingTier object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingTier(rates AiGatewayModelPricingRates) *AiGatewayModelPricingTier {
	this := AiGatewayModelPricingTier{}
	this.Rates = rates
	return &this
}

// NewAiGatewayModelPricingTierWithDefaults instantiates a new AiGatewayModelPricingTier object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingTierWithDefaults() *AiGatewayModelPricingTier {
	this := AiGatewayModelPricingTier{}
	return &this
}

// GetContextMinTokens returns the ContextMinTokens field value if set, zero value otherwise.
func (o *AiGatewayModelPricingTier) GetContextMinTokens() int64 {
	if o == nil || o.ContextMinTokens == nil {
		var ret int64
		return ret
	}
	return *o.ContextMinTokens
}

// GetContextMinTokensOk returns a tuple with the ContextMinTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingTier) GetContextMinTokensOk() (*int64, bool) {
	if o == nil || o.ContextMinTokens == nil {
		return nil, false
	}
	return o.ContextMinTokens, true
}

// HasContextMinTokens returns a boolean if a field has been set.
func (o *AiGatewayModelPricingTier) HasContextMinTokens() bool {
	return o != nil && o.ContextMinTokens != nil
}

// SetContextMinTokens gets a reference to the given int64 and assigns it to the ContextMinTokens field.
func (o *AiGatewayModelPricingTier) SetContextMinTokens(v int64) {
	o.ContextMinTokens = &v
}

// GetRates returns the Rates field value.
func (o *AiGatewayModelPricingTier) GetRates() AiGatewayModelPricingRates {
	if o == nil {
		var ret AiGatewayModelPricingRates
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingTier) GetRatesOk() (*AiGatewayModelPricingRates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rates, true
}

// SetRates sets field value.
func (o *AiGatewayModelPricingTier) SetRates(v AiGatewayModelPricingRates) {
	o.Rates = v
}

// GetServiceTier returns the ServiceTier field value if set, zero value otherwise.
func (o *AiGatewayModelPricingTier) GetServiceTier() string {
	if o == nil || o.ServiceTier == nil {
		var ret string
		return ret
	}
	return *o.ServiceTier
}

// GetServiceTierOk returns a tuple with the ServiceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingTier) GetServiceTierOk() (*string, bool) {
	if o == nil || o.ServiceTier == nil {
		return nil, false
	}
	return o.ServiceTier, true
}

// HasServiceTier returns a boolean if a field has been set.
func (o *AiGatewayModelPricingTier) HasServiceTier() bool {
	return o != nil && o.ServiceTier != nil
}

// SetServiceTier gets a reference to the given string and assigns it to the ServiceTier field.
func (o *AiGatewayModelPricingTier) SetServiceTier(v string) {
	o.ServiceTier = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ContextMinTokens != nil {
		toSerialize["context_min_tokens"] = o.ContextMinTokens
	}
	toSerialize["rates"] = o.Rates
	if o.ServiceTier != nil {
		toSerialize["service_tier"] = o.ServiceTier
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingTier) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContextMinTokens *int64                      `json:"context_min_tokens,omitempty"`
		Rates            *AiGatewayModelPricingRates `json:"rates"`
		ServiceTier      *string                     `json:"service_tier,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rates == nil {
		return fmt.Errorf("required field rates missing")
	}

	hasInvalidField := false
	o.ContextMinTokens = all.ContextMinTokens
	if all.Rates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rates = *all.Rates
	o.ServiceTier = all.ServiceTier

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
