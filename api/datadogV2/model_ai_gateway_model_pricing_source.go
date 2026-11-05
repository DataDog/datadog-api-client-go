// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingSource Source information for a model's pricing.
type AiGatewayModelPricingSource struct {
	// Pricing catalog provider identifier.
	ProviderId *string `json:"provider_id,omitempty"`
	// Source type for the pricing configuration.
	Type AiGatewayModelPricingSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingSource instantiates a new AiGatewayModelPricingSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingSource(typeVar AiGatewayModelPricingSourceType) *AiGatewayModelPricingSource {
	this := AiGatewayModelPricingSource{}
	this.Type = typeVar
	return &this
}

// NewAiGatewayModelPricingSourceWithDefaults instantiates a new AiGatewayModelPricingSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingSourceWithDefaults() *AiGatewayModelPricingSource {
	this := AiGatewayModelPricingSource{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *AiGatewayModelPricingSource) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingSource) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *AiGatewayModelPricingSource) HasProviderId() bool {
	return o != nil && o.ProviderId != nil
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *AiGatewayModelPricingSource) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetType returns the Type field value.
func (o *AiGatewayModelPricingSource) GetType() AiGatewayModelPricingSourceType {
	if o == nil {
		var ret AiGatewayModelPricingSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingSource) GetTypeOk() (*AiGatewayModelPricingSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayModelPricingSource) SetType(v AiGatewayModelPricingSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ProviderId != nil {
		toSerialize["provider_id"] = o.ProviderId
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProviderId *string                          `json:"provider_id,omitempty"`
		Type       *AiGatewayModelPricingSourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"provider_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ProviderId = all.ProviderId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
