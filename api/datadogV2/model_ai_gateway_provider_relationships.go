// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderRelationships Provider relationships.
type AiGatewayProviderRelationships struct {
	// Model catalog entries associated with a provider.
	Models AiGatewayProviderModelsRelationship `json:"models"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayProviderRelationships instantiates a new AiGatewayProviderRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayProviderRelationships(models AiGatewayProviderModelsRelationship) *AiGatewayProviderRelationships {
	this := AiGatewayProviderRelationships{}
	this.Models = models
	return &this
}

// NewAiGatewayProviderRelationshipsWithDefaults instantiates a new AiGatewayProviderRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayProviderRelationshipsWithDefaults() *AiGatewayProviderRelationships {
	this := AiGatewayProviderRelationships{}
	return &this
}

// GetModels returns the Models field value.
func (o *AiGatewayProviderRelationships) GetModels() AiGatewayProviderModelsRelationship {
	if o == nil {
		var ret AiGatewayProviderModelsRelationship
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderRelationships) GetModelsOk() (*AiGatewayProviderModelsRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Models, true
}

// SetModels sets field value.
func (o *AiGatewayProviderRelationships) SetModels(v AiGatewayProviderModelsRelationship) {
	o.Models = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayProviderRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["models"] = o.Models

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayProviderRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Models *AiGatewayProviderModelsRelationship `json:"models"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Models == nil {
		return fmt.Errorf("required field models missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"models"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Models.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Models = *all.Models

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
