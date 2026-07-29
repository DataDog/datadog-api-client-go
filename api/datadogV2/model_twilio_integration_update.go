// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationUpdate Partial Twilio integration configuration for updates.
type TwilioIntegrationUpdate struct {
	// Partial Twilio interface (source-type) configuration for updates.
	Interface *TwilioInterfaceUpdate `json:"interface,omitempty"`
	// Integration discriminator for Twilio.
	Type TwilioIntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationUpdate instantiates a new TwilioIntegrationUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationUpdate(typeVar TwilioIntegrationType) *TwilioIntegrationUpdate {
	this := TwilioIntegrationUpdate{}
	this.Type = typeVar
	return &this
}

// NewTwilioIntegrationUpdateWithDefaults instantiates a new TwilioIntegrationUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationUpdateWithDefaults() *TwilioIntegrationUpdate {
	this := TwilioIntegrationUpdate{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *TwilioIntegrationUpdate) GetInterface() TwilioInterfaceUpdate {
	if o == nil || o.Interface == nil {
		var ret TwilioInterfaceUpdate
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationUpdate) GetInterfaceOk() (*TwilioInterfaceUpdate, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *TwilioIntegrationUpdate) HasInterface() bool {
	return o != nil && o.Interface != nil
}

// SetInterface gets a reference to the given TwilioInterfaceUpdate and assigns it to the Interface field.
func (o *TwilioIntegrationUpdate) SetInterface(v TwilioInterfaceUpdate) {
	o.Interface = &v
}

// GetType returns the Type field value.
func (o *TwilioIntegrationUpdate) GetType() TwilioIntegrationType {
	if o == nil {
		var ret TwilioIntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationUpdate) GetTypeOk() (*TwilioIntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TwilioIntegrationUpdate) SetType(v TwilioIntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interface *TwilioInterfaceUpdate `json:"interface,omitempty"`
		Type      *TwilioIntegrationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.Interface != nil && all.Interface.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Interface = all.Interface
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
