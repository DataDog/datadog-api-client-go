// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelReferencePrice Informational source pricing, not gateway billing configuration. Prices can differ by provider, mode, and input length.
type AiGatewayModelReferencePrice struct {
	// Exact decimal amount per unit, zero or greater, represented as a string to preserve precision. Supports fixed decimal and exponent notation.
	Amount string `json:"amount"`
	// Token category priced, such as input, output, cache_read, or cache_write.
	Component string `json:"component"`
	// Largest input length covered by this price. Omitted when there is no upper bound.
	ContextMaxTokens *int64 `json:"context_max_tokens,omitempty"`
	// Input length above which this price applies. Zero identifies the base tier.
	ContextMinTokens int64 `json:"context_min_tokens"`
	// Currency of the price. The only supported value is USD.
	Currency string `json:"currency"`
	// Source mode name. Standard is the base mode; other values, such as fast, are source-defined.
	Mode string `json:"mode"`
	// Provider service tier explicitly associated with this mode by the source. Omitted when unknown.
	ServiceTier *string `json:"service_tier,omitempty"`
	// Unit of the quoted amount. The only supported value is million_tokens.
	Unit string `json:"unit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelReferencePrice instantiates a new AiGatewayModelReferencePrice object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelReferencePrice(amount string, component string, contextMinTokens int64, currency string, mode string, unit string) *AiGatewayModelReferencePrice {
	this := AiGatewayModelReferencePrice{}
	this.Amount = amount
	this.Component = component
	this.ContextMinTokens = contextMinTokens
	this.Currency = currency
	this.Mode = mode
	this.Unit = unit
	return &this
}

// NewAiGatewayModelReferencePriceWithDefaults instantiates a new AiGatewayModelReferencePrice object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelReferencePriceWithDefaults() *AiGatewayModelReferencePrice {
	this := AiGatewayModelReferencePrice{}
	return &this
}

// GetAmount returns the Amount field value.
func (o *AiGatewayModelReferencePrice) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value.
func (o *AiGatewayModelReferencePrice) SetAmount(v string) {
	o.Amount = v
}

// GetComponent returns the Component field value.
func (o *AiGatewayModelReferencePrice) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *AiGatewayModelReferencePrice) SetComponent(v string) {
	o.Component = v
}

// GetContextMaxTokens returns the ContextMaxTokens field value if set, zero value otherwise.
func (o *AiGatewayModelReferencePrice) GetContextMaxTokens() int64 {
	if o == nil || o.ContextMaxTokens == nil {
		var ret int64
		return ret
	}
	return *o.ContextMaxTokens
}

// GetContextMaxTokensOk returns a tuple with the ContextMaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetContextMaxTokensOk() (*int64, bool) {
	if o == nil || o.ContextMaxTokens == nil {
		return nil, false
	}
	return o.ContextMaxTokens, true
}

// HasContextMaxTokens returns a boolean if a field has been set.
func (o *AiGatewayModelReferencePrice) HasContextMaxTokens() bool {
	return o != nil && o.ContextMaxTokens != nil
}

// SetContextMaxTokens gets a reference to the given int64 and assigns it to the ContextMaxTokens field.
func (o *AiGatewayModelReferencePrice) SetContextMaxTokens(v int64) {
	o.ContextMaxTokens = &v
}

// GetContextMinTokens returns the ContextMinTokens field value.
func (o *AiGatewayModelReferencePrice) GetContextMinTokens() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ContextMinTokens
}

// GetContextMinTokensOk returns a tuple with the ContextMinTokens field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetContextMinTokensOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextMinTokens, true
}

// SetContextMinTokens sets field value.
func (o *AiGatewayModelReferencePrice) SetContextMinTokens(v int64) {
	o.ContextMinTokens = v
}

// GetCurrency returns the Currency field value.
func (o *AiGatewayModelReferencePrice) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value.
func (o *AiGatewayModelReferencePrice) SetCurrency(v string) {
	o.Currency = v
}

// GetMode returns the Mode field value.
func (o *AiGatewayModelReferencePrice) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *AiGatewayModelReferencePrice) SetMode(v string) {
	o.Mode = v
}

// GetServiceTier returns the ServiceTier field value if set, zero value otherwise.
func (o *AiGatewayModelReferencePrice) GetServiceTier() string {
	if o == nil || o.ServiceTier == nil {
		var ret string
		return ret
	}
	return *o.ServiceTier
}

// GetServiceTierOk returns a tuple with the ServiceTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetServiceTierOk() (*string, bool) {
	if o == nil || o.ServiceTier == nil {
		return nil, false
	}
	return o.ServiceTier, true
}

// HasServiceTier returns a boolean if a field has been set.
func (o *AiGatewayModelReferencePrice) HasServiceTier() bool {
	return o != nil && o.ServiceTier != nil
}

// SetServiceTier gets a reference to the given string and assigns it to the ServiceTier field.
func (o *AiGatewayModelReferencePrice) SetServiceTier(v string) {
	o.ServiceTier = &v
}

// GetUnit returns the Unit field value.
func (o *AiGatewayModelReferencePrice) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelReferencePrice) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *AiGatewayModelReferencePrice) SetUnit(v string) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelReferencePrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["amount"] = o.Amount
	toSerialize["component"] = o.Component
	if o.ContextMaxTokens != nil {
		toSerialize["context_max_tokens"] = o.ContextMaxTokens
	}
	toSerialize["context_min_tokens"] = o.ContextMinTokens
	toSerialize["currency"] = o.Currency
	toSerialize["mode"] = o.Mode
	if o.ServiceTier != nil {
		toSerialize["service_tier"] = o.ServiceTier
	}
	toSerialize["unit"] = o.Unit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelReferencePrice) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Amount           *string `json:"amount"`
		Component        *string `json:"component"`
		ContextMaxTokens *int64  `json:"context_max_tokens,omitempty"`
		ContextMinTokens *int64  `json:"context_min_tokens"`
		Currency         *string `json:"currency"`
		Mode             *string `json:"mode"`
		ServiceTier      *string `json:"service_tier,omitempty"`
		Unit             *string `json:"unit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Amount == nil {
		return fmt.Errorf("required field amount missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.ContextMinTokens == nil {
		return fmt.Errorf("required field context_min_tokens missing")
	}
	if all.Currency == nil {
		return fmt.Errorf("required field currency missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"amount", "component", "context_max_tokens", "context_min_tokens", "currency", "mode", "service_tier", "unit"})
	} else {
		return err
	}
	o.Amount = *all.Amount
	o.Component = *all.Component
	o.ContextMaxTokens = all.ContextMaxTokens
	o.ContextMinTokens = *all.ContextMinTokens
	o.Currency = *all.Currency
	o.Mode = *all.Mode
	o.ServiceTier = all.ServiceTier
	o.Unit = *all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
