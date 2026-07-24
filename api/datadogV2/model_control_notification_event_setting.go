// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ControlNotificationEventSetting The notification settings for a single event type on a control.
type ControlNotificationEventSetting struct {
	// Whether notifications are enabled for this event type.
	Enabled bool `json:"enabled"`
	// The event type the notification settings apply to, such as `new_detection`.
	EventType string `json:"event_type"`
	// The destinations that receive notifications for an event type.
	Targets []ControlNotificationTarget `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewControlNotificationEventSetting instantiates a new ControlNotificationEventSetting object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewControlNotificationEventSetting(enabled bool, eventType string, targets []ControlNotificationTarget) *ControlNotificationEventSetting {
	this := ControlNotificationEventSetting{}
	this.Enabled = enabled
	this.EventType = eventType
	this.Targets = targets
	return &this
}

// NewControlNotificationEventSettingWithDefaults instantiates a new ControlNotificationEventSetting object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewControlNotificationEventSettingWithDefaults() *ControlNotificationEventSetting {
	this := ControlNotificationEventSetting{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *ControlNotificationEventSetting) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ControlNotificationEventSetting) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ControlNotificationEventSetting) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEventType returns the EventType field value.
func (o *ControlNotificationEventSetting) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ControlNotificationEventSetting) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *ControlNotificationEventSetting) SetEventType(v string) {
	o.EventType = v
}

// GetTargets returns the Targets field value.
func (o *ControlNotificationEventSetting) GetTargets() []ControlNotificationTarget {
	if o == nil {
		var ret []ControlNotificationTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *ControlNotificationEventSetting) GetTargetsOk() (*[]ControlNotificationTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *ControlNotificationEventSetting) SetTargets(v []ControlNotificationTarget) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ControlNotificationEventSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["event_type"] = o.EventType
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ControlNotificationEventSetting) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                        `json:"enabled"`
		EventType *string                      `json:"event_type"`
		Targets   *[]ControlNotificationTarget `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.EventType == nil {
		return fmt.Errorf("required field event_type missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "event_type", "targets"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.EventType = *all.EventType
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
