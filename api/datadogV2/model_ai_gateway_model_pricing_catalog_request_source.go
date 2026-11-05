// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelPricingCatalogRequestSource Source for a catalog pricing configuration.
type AiGatewayModelPricingCatalogRequestSource struct {
	// Pricing catalog provider identifier.
	ProviderId string `json:"provider_id"`
	// Catalog pricing source type.
	Type AiGatewayModelPricingCatalogRequestSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAiGatewayModelPricingCatalogRequestSource instantiates a new AiGatewayModelPricingCatalogRequestSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelPricingCatalogRequestSource(providerId string, typeVar AiGatewayModelPricingCatalogRequestSourceType) *AiGatewayModelPricingCatalogRequestSource {
	this := AiGatewayModelPricingCatalogRequestSource{}
	this.ProviderId = providerId
	this.Type = typeVar
	return &this
}

// NewAiGatewayModelPricingCatalogRequestSourceWithDefaults instantiates a new AiGatewayModelPricingCatalogRequestSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelPricingCatalogRequestSourceWithDefaults() *AiGatewayModelPricingCatalogRequestSource {
	this := AiGatewayModelPricingCatalogRequestSource{}
	return &this
}

// GetProviderId returns the ProviderId field value.
func (o *AiGatewayModelPricingCatalogRequestSource) GetProviderId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCatalogRequestSource) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value.
func (o *AiGatewayModelPricingCatalogRequestSource) SetProviderId(v string) {
	o.ProviderId = v
}

// GetType returns the Type field value.
func (o *AiGatewayModelPricingCatalogRequestSource) GetType() AiGatewayModelPricingCatalogRequestSourceType {
	if o == nil {
		var ret AiGatewayModelPricingCatalogRequestSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelPricingCatalogRequestSource) GetTypeOk() (*AiGatewayModelPricingCatalogRequestSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiGatewayModelPricingCatalogRequestSource) SetType(v AiGatewayModelPricingCatalogRequestSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelPricingCatalogRequestSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["provider_id"] = o.ProviderId
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelPricingCatalogRequestSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProviderId *string                                        `json:"provider_id"`
		Type       *AiGatewayModelPricingCatalogRequestSourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProviderId == nil {
		return fmt.Errorf("required field provider_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.ProviderId = *all.ProviderId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
