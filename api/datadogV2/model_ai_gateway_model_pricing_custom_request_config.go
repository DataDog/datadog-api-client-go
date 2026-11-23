// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingCustomRequestConfig Custom pricing configuration. All four rates are required.
type AiGatewayModelPricingCustomRequestConfig struct {
	// Complete set of rates required for custom pricing.
	Rates AiGatewayModelPricingRequiredRates `json:"rates"`
	// Source for a custom pricing configuration.
	Source AiGatewayModelPricingCustomRequestSource `json:"source"`
	// Custom prices that replace the base rates for the requests they select. Only valid for custom pricing.
	// At most 16 tiers, and no two may share the same `service_tier` and `context_min_tokens`.
	Tiers []AiGatewayModelPricingTier `json:"tiers,omitempty"`
	// Pricing configuration version. The only supported version is `1`.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingCustomRequestConfig instantiates a new AiGatewayModelPricingCustomRequestConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingCustomRequestConfig(rates AiGatewayModelPricingRequiredRates, source AiGatewayModelPricingCustomRequestSource, version int64) *AiGatewayModelPricingCustomRequestConfig {
	this := AiGatewayModelPricingCustomRequestConfig{}
	this.Rates = rates
	this.Source = source
	this.Version = version
	return &this
}

// NewAiGatewayModelPricingCustomRequestConfigWithDefaults instantiates a new AiGatewayModelPricingCustomRequestConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingCustomRequestConfigWithDefaults() *AiGatewayModelPricingCustomRequestConfig {
	this := AiGatewayModelPricingCustomRequestConfig{}
	return &this
}

// GetRates returns the Rates field value.
func (o *AiGatewayModelPricingCustomRequestConfig) GetRates() AiGatewayModelPricingRequiredRates {
	if o == nil {
		var ret AiGatewayModelPricingRequiredRates
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCustomRequestConfig) GetRatesOk() (*AiGatewayModelPricingRequiredRates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rates, true
}

// SetRates sets field value.
func (o *AiGatewayModelPricingCustomRequestConfig) SetRates(v AiGatewayModelPricingRequiredRates) {
	o.Rates = v
}

// GetSource returns the Source field value.
func (o *AiGatewayModelPricingCustomRequestConfig) GetSource() AiGatewayModelPricingCustomRequestSource {
	if o == nil {
		var ret AiGatewayModelPricingCustomRequestSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCustomRequestConfig) GetSourceOk() (*AiGatewayModelPricingCustomRequestSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *AiGatewayModelPricingCustomRequestConfig) SetSource(v AiGatewayModelPricingCustomRequestSource) {
	o.Source = v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *AiGatewayModelPricingCustomRequestConfig) GetTiers() []AiGatewayModelPricingTier {
	if o == nil || o.Tiers == nil {
		var ret []AiGatewayModelPricingTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCustomRequestConfig) GetTiersOk() (*[]AiGatewayModelPricingTier, bool) {
	if o == nil || o.Tiers == nil {
		return nil, false
	}
	return &o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *AiGatewayModelPricingCustomRequestConfig) HasTiers() bool {
	return o != nil && o.Tiers != nil
}

// SetTiers gets a reference to the given []AiGatewayModelPricingTier and assigns it to the Tiers field.
func (o *AiGatewayModelPricingCustomRequestConfig) SetTiers(v []AiGatewayModelPricingTier) {
	o.Tiers = v
}

// GetVersion returns the Version field value.
func (o *AiGatewayModelPricingCustomRequestConfig) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCustomRequestConfig) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *AiGatewayModelPricingCustomRequestConfig) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingCustomRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rates"] = o.Rates
	toSerialize["source"] = o.Source
	if o.Tiers != nil {
		toSerialize["tiers"] = o.Tiers
	}
	toSerialize["version"] = o.Version
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingCustomRequestConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rates   *AiGatewayModelPricingRequiredRates       `json:"rates"`
		Source  *AiGatewayModelPricingCustomRequestSource `json:"source"`
		Tiers   []AiGatewayModelPricingTier               `json:"tiers,omitempty"`
		Version *int64                                    `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rates == nil {
		return fmt.Errorf("required field rates missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}

	hasInvalidField := false
	if all.Rates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rates = *all.Rates
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	o.Tiers = all.Tiers
	o.Version = *all.Version

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
