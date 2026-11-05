// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelUpdateAttributes Attributes used to update a model catalog entry.
type AiGatewayModelUpdateAttributes struct {
	// Model identifier sent to the provider.
	ModelIdWire *string `json:"model_id_wire,omitempty"`
	// Operations supported by the model.
	SupportedOperations []AiGatewaySupportedOperation `json:"supported_operations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelUpdateAttributes instantiates a new AiGatewayModelUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelUpdateAttributes() *AiGatewayModelUpdateAttributes {
	this := AiGatewayModelUpdateAttributes{}
	return &this
}

// NewAiGatewayModelUpdateAttributesWithDefaults instantiates a new AiGatewayModelUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelUpdateAttributesWithDefaults() *AiGatewayModelUpdateAttributes {
	this := AiGatewayModelUpdateAttributes{}
	return &this
}

// GetModelIdWire returns the ModelIdWire field value if set, zero value otherwise.
func (o *AiGatewayModelUpdateAttributes) GetModelIdWire() string {
	if o == nil || o.ModelIdWire == nil {
		var ret string
		return ret
	}
	return *o.ModelIdWire
}

// GetModelIdWireOk returns a tuple with the ModelIdWire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelUpdateAttributes) GetModelIdWireOk() (*string, bool) {
	if o == nil || o.ModelIdWire == nil {
		return nil, false
	}
	return o.ModelIdWire, true
}

// HasModelIdWire returns a boolean if a field has been set.
func (o *AiGatewayModelUpdateAttributes) HasModelIdWire() bool {
	return o != nil && o.ModelIdWire != nil
}

// SetModelIdWire gets a reference to the given string and assigns it to the ModelIdWire field.
func (o *AiGatewayModelUpdateAttributes) SetModelIdWire(v string) {
	o.ModelIdWire = &v
}

// GetSupportedOperations returns the SupportedOperations field value if set, zero value otherwise.
func (o *AiGatewayModelUpdateAttributes) GetSupportedOperations() []AiGatewaySupportedOperation {
	if o == nil || o.SupportedOperations == nil {
		var ret []AiGatewaySupportedOperation
		return ret
	}
	return o.SupportedOperations
}

// GetSupportedOperationsOk returns a tuple with the SupportedOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelUpdateAttributes) GetSupportedOperationsOk() (*[]AiGatewaySupportedOperation, bool) {
	if o == nil || o.SupportedOperations == nil {
		return nil, false
	}
	return &o.SupportedOperations, true
}

// HasSupportedOperations returns a boolean if a field has been set.
func (o *AiGatewayModelUpdateAttributes) HasSupportedOperations() bool {
	return o != nil && o.SupportedOperations != nil
}

// SetSupportedOperations gets a reference to the given []AiGatewaySupportedOperation and assigns it to the SupportedOperations field.
func (o *AiGatewayModelUpdateAttributes) SetSupportedOperations(v []AiGatewaySupportedOperation) {
	o.SupportedOperations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ModelIdWire != nil {
		toSerialize["model_id_wire"] = o.ModelIdWire
	}
	if o.SupportedOperations != nil {
		toSerialize["supported_operations"] = o.SupportedOperations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModelIdWire         *string                       `json:"model_id_wire,omitempty"`
		SupportedOperations []AiGatewaySupportedOperation `json:"supported_operations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"model_id_wire", "supported_operations"})
	} else {
		return err
	}
	o.ModelIdWire = all.ModelIdWire
	o.SupportedOperations = all.SupportedOperations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
