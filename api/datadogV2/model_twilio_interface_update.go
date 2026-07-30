// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TwilioInterfaceUpdate Partial Twilio interface (source-type) configuration for updates.
type TwilioInterfaceUpdate struct {
	// Authentication methods supported by the Twilio interface. Exactly one is set, selected by its `type`.
	Authentication *TwilioAuthentication `json:"authentication,omitempty"`
	// Dataflows for the Twilio interface.
	Dataflows []TwilioDataflow `json:"dataflows,omitempty"`
	// Partial Twilio interface settings for updates.
	Settings *TwilioSettingsUpdate `json:"settings,omitempty"`
	// Interface discriminator for Twilio.
	Type TwilioInterfaceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTwilioInterfaceUpdate instantiates a new TwilioInterfaceUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTwilioInterfaceUpdate(typeVar TwilioInterfaceType) *TwilioInterfaceUpdate {
	this := TwilioInterfaceUpdate{}
	this.Type = typeVar
	return &this
}

// NewTwilioInterfaceUpdateWithDefaults instantiates a new TwilioInterfaceUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTwilioInterfaceUpdateWithDefaults() *TwilioInterfaceUpdate {
	this := TwilioInterfaceUpdate{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *TwilioInterfaceUpdate) GetAuthentication() TwilioAuthentication {
	if o == nil || o.Authentication == nil {
		var ret TwilioAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioInterfaceUpdate) GetAuthenticationOk() (*TwilioAuthentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *TwilioInterfaceUpdate) HasAuthentication() bool {
	return o != nil && o.Authentication != nil
}

// SetAuthentication gets a reference to the given TwilioAuthentication and assigns it to the Authentication field.
func (o *TwilioInterfaceUpdate) SetAuthentication(v TwilioAuthentication) {
	o.Authentication = &v
}

// GetDataflows returns the Dataflows field value if set, zero value otherwise.
func (o *TwilioInterfaceUpdate) GetDataflows() []TwilioDataflow {
	if o == nil || o.Dataflows == nil {
		var ret []TwilioDataflow
		return ret
	}
	return o.Dataflows
}

// GetDataflowsOk returns a tuple with the Dataflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioInterfaceUpdate) GetDataflowsOk() (*[]TwilioDataflow, bool) {
	if o == nil || o.Dataflows == nil {
		return nil, false
	}
	return &o.Dataflows, true
}

// HasDataflows returns a boolean if a field has been set.
func (o *TwilioInterfaceUpdate) HasDataflows() bool {
	return o != nil && o.Dataflows != nil
}

// SetDataflows gets a reference to the given []TwilioDataflow and assigns it to the Dataflows field.
func (o *TwilioInterfaceUpdate) SetDataflows(v []TwilioDataflow) {
	o.Dataflows = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *TwilioInterfaceUpdate) GetSettings() TwilioSettingsUpdate {
	if o == nil || o.Settings == nil {
		var ret TwilioSettingsUpdate
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TwilioInterfaceUpdate) GetSettingsOk() (*TwilioSettingsUpdate, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *TwilioInterfaceUpdate) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given TwilioSettingsUpdate and assigns it to the Settings field.
func (o *TwilioInterfaceUpdate) SetSettings(v TwilioSettingsUpdate) {
	o.Settings = &v
}

// GetType returns the Type field value.
func (o *TwilioInterfaceUpdate) GetType() TwilioInterfaceType {
	if o == nil {
		var ret TwilioInterfaceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TwilioInterfaceUpdate) GetTypeOk() (*TwilioInterfaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TwilioInterfaceUpdate) SetType(v TwilioInterfaceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TwilioInterfaceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if o.Dataflows != nil {
		toSerialize["dataflows"] = o.Dataflows
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TwilioInterfaceUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Authentication *TwilioAuthentication `json:"authentication,omitempty"`
		Dataflows      []TwilioDataflow      `json:"dataflows,omitempty"`
		Settings       *TwilioSettingsUpdate `json:"settings,omitempty"`
		Type           *TwilioInterfaceType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"authentication", "dataflows", "settings", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Authentication = all.Authentication
	o.Dataflows = all.Dataflows
	if all.Settings != nil && all.Settings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Settings = all.Settings
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
