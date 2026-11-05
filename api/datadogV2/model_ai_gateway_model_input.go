// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelInput Manually managed model definition supplied when creating a provider or catalog entry.
type AiGatewayModelInput struct {
	// Provider-independent model identifier.
	ModelIdCanonical string `json:"model_id_canonical"`
	// Model identifier sent to the provider.
	ModelIdWire string `json:"model_id_wire"`
	// Operations supported by the model.
	SupportedOperations []AiGatewaySupportedOperation `json:"supported_operations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelInput instantiates a new AiGatewayModelInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelInput(modelIdCanonical string, modelIdWire string) *AiGatewayModelInput {
	this := AiGatewayModelInput{}
	this.ModelIdCanonical = modelIdCanonical
	this.ModelIdWire = modelIdWire
	return &this
}

// NewAiGatewayModelInputWithDefaults instantiates a new AiGatewayModelInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelInputWithDefaults() *AiGatewayModelInput {
	this := AiGatewayModelInput{}
	return &this
}

// GetModelIdCanonical returns the ModelIdCanonical field value.
func (o *AiGatewayModelInput) GetModelIdCanonical() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelIdCanonical
}

// GetModelIdCanonicalOk returns a tuple with the ModelIdCanonical field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelInput) GetModelIdCanonicalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelIdCanonical, true
}

// SetModelIdCanonical sets field value.
func (o *AiGatewayModelInput) SetModelIdCanonical(v string) {
	o.ModelIdCanonical = v
}

// GetModelIdWire returns the ModelIdWire field value.
func (o *AiGatewayModelInput) GetModelIdWire() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelIdWire
}

// GetModelIdWireOk returns a tuple with the ModelIdWire field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelInput) GetModelIdWireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelIdWire, true
}

// SetModelIdWire sets field value.
func (o *AiGatewayModelInput) SetModelIdWire(v string) {
	o.ModelIdWire = v
}

// GetSupportedOperations returns the SupportedOperations field value if set, zero value otherwise.
func (o *AiGatewayModelInput) GetSupportedOperations() []AiGatewaySupportedOperation {
	if o == nil || o.SupportedOperations == nil {
		var ret []AiGatewaySupportedOperation
		return ret
	}
	return o.SupportedOperations
}

// GetSupportedOperationsOk returns a tuple with the SupportedOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelInput) GetSupportedOperationsOk() (*[]AiGatewaySupportedOperation, bool) {
	if o == nil || o.SupportedOperations == nil {
		return nil, false
	}
	return &o.SupportedOperations, true
}

// HasSupportedOperations returns a boolean if a field has been set.
func (o *AiGatewayModelInput) HasSupportedOperations() bool {
	return o != nil && o.SupportedOperations != nil
}

// SetSupportedOperations gets a reference to the given []AiGatewaySupportedOperation and assigns it to the SupportedOperations field.
func (o *AiGatewayModelInput) SetSupportedOperations(v []AiGatewaySupportedOperation) {
	o.SupportedOperations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["model_id_canonical"] = o.ModelIdCanonical
	toSerialize["model_id_wire"] = o.ModelIdWire
	if o.SupportedOperations != nil {
		toSerialize["supported_operations"] = o.SupportedOperations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModelIdCanonical    *string                       `json:"model_id_canonical"`
		ModelIdWire         *string                       `json:"model_id_wire"`
		SupportedOperations []AiGatewaySupportedOperation `json:"supported_operations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ModelIdCanonical == nil {
		return fmt.Errorf("required field model_id_canonical missing")
	}
	if all.ModelIdWire == nil {
		return fmt.Errorf("required field model_id_wire missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"model_id_canonical", "model_id_wire", "supported_operations"})
	} else {
		return err
	}
	o.ModelIdCanonical = *all.ModelIdCanonical
	o.ModelIdWire = *all.ModelIdWire
	o.SupportedOperations = all.SupportedOperations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
