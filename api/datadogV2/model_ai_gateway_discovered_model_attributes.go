// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayDiscoveredModelAttributes Model discovered from a provider integration account.
type AiGatewayDiscoveredModelAttributes struct {
	// Human-readable model name.
	DisplayName *string `json:"display_name,omitempty"`
	// Model identifier sent to the provider.
	ModelIdWire string `json:"model_id_wire"`
	// Operations supported by the model.
	SupportedOperations []AiGatewaySupportedOperation `json:"supported_operations"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayDiscoveredModelAttributes instantiates a new AiGatewayDiscoveredModelAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayDiscoveredModelAttributes(modelIdWire string, supportedOperations []AiGatewaySupportedOperation) *AiGatewayDiscoveredModelAttributes {
	this := AiGatewayDiscoveredModelAttributes{}
	this.ModelIdWire = modelIdWire
	this.SupportedOperations = supportedOperations
	return &this
}

// NewAiGatewayDiscoveredModelAttributesWithDefaults instantiates a new AiGatewayDiscoveredModelAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayDiscoveredModelAttributesWithDefaults() *AiGatewayDiscoveredModelAttributes {
	this := AiGatewayDiscoveredModelAttributes{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AiGatewayDiscoveredModelAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayDiscoveredModelAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AiGatewayDiscoveredModelAttributes) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AiGatewayDiscoveredModelAttributes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetModelIdWire returns the ModelIdWire field value.
func (o *AiGatewayDiscoveredModelAttributes) GetModelIdWire() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelIdWire
}

// GetModelIdWireOk returns a tuple with the ModelIdWire field value
// and a boolean to check if the value has been set.
func (o *AiGatewayDiscoveredModelAttributes) GetModelIdWireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelIdWire, true
}

// SetModelIdWire sets field value.
func (o *AiGatewayDiscoveredModelAttributes) SetModelIdWire(v string) {
	o.ModelIdWire = v
}

// GetSupportedOperations returns the SupportedOperations field value.
func (o *AiGatewayDiscoveredModelAttributes) GetSupportedOperations() []AiGatewaySupportedOperation {
	if o == nil {
		var ret []AiGatewaySupportedOperation
		return ret
	}
	return o.SupportedOperations
}

// GetSupportedOperationsOk returns a tuple with the SupportedOperations field value
// and a boolean to check if the value has been set.
func (o *AiGatewayDiscoveredModelAttributes) GetSupportedOperationsOk() (*[]AiGatewaySupportedOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedOperations, true
}

// SetSupportedOperations sets field value.
func (o *AiGatewayDiscoveredModelAttributes) SetSupportedOperations(v []AiGatewaySupportedOperation) {
	o.SupportedOperations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayDiscoveredModelAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["model_id_wire"] = o.ModelIdWire
	toSerialize["supported_operations"] = o.SupportedOperations

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayDiscoveredModelAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName         *string                        `json:"display_name,omitempty"`
		ModelIdWire         *string                        `json:"model_id_wire"`
		SupportedOperations *[]AiGatewaySupportedOperation `json:"supported_operations"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ModelIdWire == nil {
		return fmt.Errorf("required field model_id_wire missing")
	}
	if all.SupportedOperations == nil {
		return fmt.Errorf("required field supported_operations missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "model_id_wire", "supported_operations"})
	} else {
		return err
	}
	o.DisplayName = all.DisplayName
	o.ModelIdWire = *all.ModelIdWire
	o.SupportedOperations = *all.SupportedOperations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
