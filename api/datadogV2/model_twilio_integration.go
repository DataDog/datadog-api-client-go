// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegration Twilio integration configuration.
type TwilioIntegration struct {
	// Twilio interface (source-type) configuration.
	Interface TwilioInterface `json:"interface"`
	// Integration discriminator for Twilio.
	Type TwilioIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegration instantiates a new TwilioIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegration(interfaceVar TwilioInterface, typeVar TwilioIntegrationType) *TwilioIntegration {
	this := TwilioIntegration{}
	this.Interface = interfaceVar
	this.Type = typeVar
	return &this
}

// NewTwilioIntegrationWithDefaults instantiates a new TwilioIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationWithDefaults() *TwilioIntegration {
	this := TwilioIntegration{}
	return &this
}

// GetInterface returns the Interface field value.
func (o *TwilioIntegration) GetInterface() TwilioInterface {
	if o == nil {
		var ret TwilioInterface
		return ret
	}
	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegration) GetInterfaceOk() (*TwilioInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value.
func (o *TwilioIntegration) SetInterface(v TwilioInterface) {
	o.Interface = v
}

// GetType returns the Type field value.
func (o *TwilioIntegration) GetType() TwilioIntegrationType {
	if o == nil {
		var ret TwilioIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegration) GetTypeOk() (*TwilioIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TwilioIntegration) SetType(v TwilioIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["interface"] = o.Interface
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interface *TwilioInterface       `json:"interface"`
		Type      *TwilioIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Interface == nil {
		return fmt.Errorf("required field interface missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interface", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Interface.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Interface = *all.Interface
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
