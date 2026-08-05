// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ControlNotificationSettingsAttributes The attributes of a governance control's notification settings.
type ControlNotificationSettingsAttributes struct {
	// The notification settings for each supported event type on the control.
	EventSettings []ControlNotificationEventSetting `json:"event_settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewControlNotificationSettingsAttributes instantiates a new ControlNotificationSettingsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewControlNotificationSettingsAttributes(eventSettings []ControlNotificationEventSetting) *ControlNotificationSettingsAttributes {
	this := ControlNotificationSettingsAttributes{}
	this.EventSettings = eventSettings
	return &this
}

// NewControlNotificationSettingsAttributesWithDefaults instantiates a new ControlNotificationSettingsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewControlNotificationSettingsAttributesWithDefaults() *ControlNotificationSettingsAttributes {
	this := ControlNotificationSettingsAttributes{}
	return &this
}

// GetEventSettings returns the EventSettings field value.
func (o *ControlNotificationSettingsAttributes) GetEventSettings() []ControlNotificationEventSetting {
	if o == nil {
		var ret []ControlNotificationEventSetting
		return ret
	}
	return o.EventSettings
}

// GetEventSettingsOk returns a tuple with the EventSettings field value
// and a boolean to check if the value has been set.
func (o *ControlNotificationSettingsAttributes) GetEventSettingsOk() (*[]ControlNotificationEventSetting, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventSettings, true
}

// SetEventSettings sets field value.
func (o *ControlNotificationSettingsAttributes) SetEventSettings(v []ControlNotificationEventSetting) {
	o.EventSettings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ControlNotificationSettingsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event_settings"] = o.EventSettings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ControlNotificationSettingsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventSettings *[]ControlNotificationEventSetting `json:"event_settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EventSettings == nil {
		return fmt.Errorf("required field event_settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event_settings"})
	} else {
		return err
	}
	o.EventSettings = *all.EventSettings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
