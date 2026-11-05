// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingCatalogRequestConfig Catalog pricing configuration with optional manual rate overrides.
type AiGatewayModelPricingCatalogRequestConfig struct {
	// Pricing rates for a model.
	Overrides *AiGatewayModelPricingRates `json:"overrides,omitempty"`
	// Source for a catalog pricing configuration.
	Source AiGatewayModelPricingCatalogRequestSource `json:"source"`
	// Pricing configuration version. The only supported version is `1`.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingCatalogRequestConfig instantiates a new AiGatewayModelPricingCatalogRequestConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingCatalogRequestConfig(source AiGatewayModelPricingCatalogRequestSource, version int64) *AiGatewayModelPricingCatalogRequestConfig {
	this := AiGatewayModelPricingCatalogRequestConfig{}
	this.Source = source
	this.Version = version
	return &this
}

// NewAiGatewayModelPricingCatalogRequestConfigWithDefaults instantiates a new AiGatewayModelPricingCatalogRequestConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingCatalogRequestConfigWithDefaults() *AiGatewayModelPricingCatalogRequestConfig {
	this := AiGatewayModelPricingCatalogRequestConfig{}
	return &this
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *AiGatewayModelPricingCatalogRequestConfig) GetOverrides() AiGatewayModelPricingRates {
	if o == nil || o.Overrides == nil {
		var ret AiGatewayModelPricingRates
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCatalogRequestConfig) GetOverridesOk() (*AiGatewayModelPricingRates, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *AiGatewayModelPricingCatalogRequestConfig) HasOverrides() bool {
	return o != nil && o.Overrides != nil
}

// SetOverrides gets a reference to the given AiGatewayModelPricingRates and assigns it to the Overrides field.
func (o *AiGatewayModelPricingCatalogRequestConfig) SetOverrides(v AiGatewayModelPricingRates) {
	o.Overrides = &v
}

// GetSource returns the Source field value.
func (o *AiGatewayModelPricingCatalogRequestConfig) GetSource() AiGatewayModelPricingCatalogRequestSource {
	if o == nil {
		var ret AiGatewayModelPricingCatalogRequestSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCatalogRequestConfig) GetSourceOk() (*AiGatewayModelPricingCatalogRequestSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *AiGatewayModelPricingCatalogRequestConfig) SetSource(v AiGatewayModelPricingCatalogRequestSource) {
	o.Source = v
}

// GetVersion returns the Version field value.
func (o *AiGatewayModelPricingCatalogRequestConfig) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCatalogRequestConfig) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *AiGatewayModelPricingCatalogRequestConfig) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingCatalogRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Overrides != nil {
		toSerialize["overrides"] = o.Overrides
	}
	toSerialize["source"] = o.Source
	toSerialize["version"] = o.Version
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingCatalogRequestConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Overrides *AiGatewayModelPricingRates                `json:"overrides,omitempty"`
		Source    *AiGatewayModelPricingCatalogRequestSource `json:"source"`
		Version   *int64                                     `json:"version"`
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

	hasInvalidField := false
	if all.Overrides != nil && all.Overrides.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Overrides = all.Overrides
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	o.Version = *all.Version

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
