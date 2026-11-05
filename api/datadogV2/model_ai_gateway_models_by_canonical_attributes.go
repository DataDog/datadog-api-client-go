// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelsByCanonicalAttributes Provider entries grouped under one canonical model ID.
type AiGatewayModelsByCanonicalAttributes struct {
	// Provider-independent model identifier.
	ModelIdCanonical string `json:"model_id_canonical"`
	//
	Providers []AiGatewayProviderModelEntry `json:"providers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelsByCanonicalAttributes instantiates a new AiGatewayModelsByCanonicalAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelsByCanonicalAttributes(modelIdCanonical string, providers []AiGatewayProviderModelEntry) *AiGatewayModelsByCanonicalAttributes {
	this := AiGatewayModelsByCanonicalAttributes{}
	this.ModelIdCanonical = modelIdCanonical
	this.Providers = providers
	return &this
}

// NewAiGatewayModelsByCanonicalAttributesWithDefaults instantiates a new AiGatewayModelsByCanonicalAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelsByCanonicalAttributesWithDefaults() *AiGatewayModelsByCanonicalAttributes {
	this := AiGatewayModelsByCanonicalAttributes{}
	return &this
}

// GetModelIdCanonical returns the ModelIdCanonical field value.
func (o *AiGatewayModelsByCanonicalAttributes) GetModelIdCanonical() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelIdCanonical
}

// GetModelIdCanonicalOk returns a tuple with the ModelIdCanonical field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelsByCanonicalAttributes) GetModelIdCanonicalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelIdCanonical, true
}

// SetModelIdCanonical sets field value.
func (o *AiGatewayModelsByCanonicalAttributes) SetModelIdCanonical(v string) {
	o.ModelIdCanonical = v
}

// GetProviders returns the Providers field value.
func (o *AiGatewayModelsByCanonicalAttributes) GetProviders() []AiGatewayProviderModelEntry {
	if o == nil {
		var ret []AiGatewayProviderModelEntry
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelsByCanonicalAttributes) GetProvidersOk() (*[]AiGatewayProviderModelEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Providers, true
}

// SetProviders sets field value.
func (o *AiGatewayModelsByCanonicalAttributes) SetProviders(v []AiGatewayProviderModelEntry) {
	o.Providers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelsByCanonicalAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["model_id_canonical"] = o.ModelIdCanonical
	toSerialize["providers"] = o.Providers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelsByCanonicalAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModelIdCanonical *string                        `json:"model_id_canonical"`
		Providers        *[]AiGatewayProviderModelEntry `json:"providers"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ModelIdCanonical == nil {
		return fmt.Errorf("required field model_id_canonical missing")
	}
	if all.Providers == nil {
		return fmt.Errorf("required field providers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"model_id_canonical", "providers"})
	} else {
		return err
	}
	o.ModelIdCanonical = *all.ModelIdCanonical
	o.Providers = *all.Providers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
