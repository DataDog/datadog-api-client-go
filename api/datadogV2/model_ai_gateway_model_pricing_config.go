// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingConfig Pricing configuration for a model.
type AiGatewayModelPricingConfig struct {
	// Pricing rates for a model.
	Overrides *AiGatewayModelPricingRates `json:"overrides,omitempty"`
	// Pricing rates for a model.
	Rates *AiGatewayModelPricingRates `json:"rates,omitempty"`
	// Source information for a model's pricing.
	Source AiGatewayModelPricingSource `json:"source"`
	// Pricing configuration version. The only supported version is `1`.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingConfig instantiates a new AiGatewayModelPricingConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingConfig(source AiGatewayModelPricingSource, version int64) *AiGatewayModelPricingConfig {
	this := AiGatewayModelPricingConfig{}
	this.Source = source
	this.Version = version
	return &this
}

// NewAiGatewayModelPricingConfigWithDefaults instantiates a new AiGatewayModelPricingConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingConfigWithDefaults() *AiGatewayModelPricingConfig {
	this := AiGatewayModelPricingConfig{}
	return &this
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *AiGatewayModelPricingConfig) GetOverrides() AiGatewayModelPricingRates {
	if o == nil || o.Overrides == nil {
		var ret AiGatewayModelPricingRates
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingConfig) GetOverridesOk() (*AiGatewayModelPricingRates, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *AiGatewayModelPricingConfig) HasOverrides() bool {
	return o != nil && o.Overrides != nil
}

// SetOverrides gets a reference to the given AiGatewayModelPricingRates and assigns it to the Overrides field.
func (o *AiGatewayModelPricingConfig) SetOverrides(v AiGatewayModelPricingRates) {
	o.Overrides = &v
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *AiGatewayModelPricingConfig) GetRates() AiGatewayModelPricingRates {
	if o == nil || o.Rates == nil {
		var ret AiGatewayModelPricingRates
		return ret
	}
	return *o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingConfig) GetRatesOk() (*AiGatewayModelPricingRates, bool) {
	if o == nil || o.Rates == nil {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *AiGatewayModelPricingConfig) HasRates() bool {
	return o != nil && o.Rates != nil
}

// SetRates gets a reference to the given AiGatewayModelPricingRates and assigns it to the Rates field.
func (o *AiGatewayModelPricingConfig) SetRates(v AiGatewayModelPricingRates) {
	o.Rates = &v
}

// GetSource returns the Source field value.
func (o *AiGatewayModelPricingConfig) GetSource() AiGatewayModelPricingSource {
	if o == nil {
		var ret AiGatewayModelPricingSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingConfig) GetSourceOk() (*AiGatewayModelPricingSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *AiGatewayModelPricingConfig) SetSource(v AiGatewayModelPricingSource) {
	o.Source = v
}

// GetVersion returns the Version field value.
func (o *AiGatewayModelPricingConfig) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingConfig) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *AiGatewayModelPricingConfig) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	if o.Rates != nil {
		toSerialize["rates"] = o.Rates
	}
	toSerialize["source"] = o.Source
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Overrides *AiGatewayModelPricingRates  `json:"overrides,omitempty"`
		Rates     *AiGatewayModelPricingRates  `json:"rates,omitempty"`
		Source    *AiGatewayModelPricingSource `json:"source"`
		Version   *int64                       `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"overrides", "rates", "source", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Overrides != nil && all.Overrides.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Overrides = all.Overrides
	if all.Rates != nil && all.Rates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rates = all.Rates
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
