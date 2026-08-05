// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ControlNotificationSettingsUpdateAttributes The attributes of a governance control's notification settings that can be updated.
type ControlNotificationSettingsUpdateAttributes struct {
	// The notification settings for each supported event type on the control.
	EventSettings []ControlNotificationEventSetting `json:"event_settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewControlNotificationSettingsUpdateAttributes instantiates a new ControlNotificationSettingsUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewControlNotificationSettingsUpdateAttributes() *ControlNotificationSettingsUpdateAttributes {
	this := ControlNotificationSettingsUpdateAttributes{}
	return &this
}

// NewControlNotificationSettingsUpdateAttributesWithDefaults instantiates a new ControlNotificationSettingsUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewControlNotificationSettingsUpdateAttributesWithDefaults() *ControlNotificationSettingsUpdateAttributes {
	this := ControlNotificationSettingsUpdateAttributes{}
	return &this
}

// GetEventSettings returns the EventSettings field value if set, zero value otherwise.
func (o *ControlNotificationSettingsUpdateAttributes) GetEventSettings() []ControlNotificationEventSetting {
	if o == nil || o.EventSettings == nil {
		var ret []ControlNotificationEventSetting
		return ret
	}
	return o.EventSettings
}

// GetEventSettingsOk returns a tuple with the EventSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlNotificationSettingsUpdateAttributes) GetEventSettingsOk() (*[]ControlNotificationEventSetting, bool) {
	if o == nil || o.EventSettings == nil {
		return nil, false
	}
	return &o.EventSettings, true
}

// HasEventSettings returns a boolean if a field has been set.
func (o *ControlNotificationSettingsUpdateAttributes) HasEventSettings() bool {
	return o != nil && o.EventSettings != nil
}

// SetEventSettings gets a reference to the given []ControlNotificationEventSetting and assigns it to the EventSettings field.
func (o *ControlNotificationSettingsUpdateAttributes) SetEventSettings(v []ControlNotificationEventSetting) {
	o.EventSettings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ControlNotificationSettingsUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EventSettings != nil {
		toSerialize["event_settings"] = o.EventSettings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ControlNotificationSettingsUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventSettings []ControlNotificationEventSetting `json:"event_settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_settings"})
	} else {
		return err
	}
	o.EventSettings = all.EventSettings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
