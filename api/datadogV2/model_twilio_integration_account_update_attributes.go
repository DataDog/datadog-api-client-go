// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioIntegrationAccountUpdateAttributes Updatable attributes of a Twilio integration account. Every field is optional; only the fields provided are changed.
type TwilioIntegrationAccountUpdateAttributes struct {
	// Partial Twilio interface (source-type) configuration for updates.
	Interface *TwilioInterfaceUpdate `json:"interface,omitempty"`
	// Human-readable name of the account.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioIntegrationAccountUpdateAttributes instantiates a new TwilioIntegrationAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioIntegrationAccountUpdateAttributes() *TwilioIntegrationAccountUpdateAttributes {
	this := TwilioIntegrationAccountUpdateAttributes{}
	return &this
}

// NewTwilioIntegrationAccountUpdateAttributesWithDefaults instantiates a new TwilioIntegrationAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioIntegrationAccountUpdateAttributesWithDefaults() *TwilioIntegrationAccountUpdateAttributes {
	this := TwilioIntegrationAccountUpdateAttributes{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *TwilioIntegrationAccountUpdateAttributes) GetInterface() TwilioInterfaceUpdate {
	if o == nil || o.Interface == nil {
		var ret TwilioInterfaceUpdate
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) GetInterfaceOk() (*TwilioInterfaceUpdate, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) HasInterface() bool {
	return o != nil && o.Interface != nil
}

// SetInterface gets a reference to the given TwilioInterfaceUpdate and assigns it to the Interface field.
func (o *TwilioIntegrationAccountUpdateAttributes) SetInterface(v TwilioInterfaceUpdate) {
	o.Interface = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TwilioIntegrationAccountUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TwilioIntegrationAccountUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TwilioIntegrationAccountUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioIntegrationAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioIntegrationAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interface *TwilioInterfaceUpdate `json:"interface,omitempty"`
		Name      *string                `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interface", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Interface != nil && all.Interface.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Interface = all.Interface
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
