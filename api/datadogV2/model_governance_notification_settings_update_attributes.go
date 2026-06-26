// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceNotificationSettingsUpdateAttributes The attributes of the governance notification settings that can be updated. Only the attributes present in the request are modified.
type GovernanceNotificationSettingsUpdateAttributes struct {
	// Whether notifications are sent to users when detections are assigned to them.
	AssignmentNotificationsEnabled *bool `json:"assignment_notifications_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceNotificationSettingsUpdateAttributes instantiates a new GovernanceNotificationSettingsUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceNotificationSettingsUpdateAttributes() *GovernanceNotificationSettingsUpdateAttributes {
	this := GovernanceNotificationSettingsUpdateAttributes{}
	return &this
}

// NewGovernanceNotificationSettingsUpdateAttributesWithDefaults instantiates a new GovernanceNotificationSettingsUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceNotificationSettingsUpdateAttributesWithDefaults() *GovernanceNotificationSettingsUpdateAttributes {
	this := GovernanceNotificationSettingsUpdateAttributes{}
	return &this
}

// GetAssignmentNotificationsEnabled returns the AssignmentNotificationsEnabled field value if set, zero value otherwise.
func (o *GovernanceNotificationSettingsUpdateAttributes) GetAssignmentNotificationsEnabled() bool {
	if o == nil || o.AssignmentNotificationsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AssignmentNotificationsEnabled
}

// GetAssignmentNotificationsEnabledOk returns a tuple with the AssignmentNotificationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceNotificationSettingsUpdateAttributes) GetAssignmentNotificationsEnabledOk() (*bool, bool) {
	if o == nil || o.AssignmentNotificationsEnabled == nil {
		return nil, false
	}
	return o.AssignmentNotificationsEnabled, true
}

// HasAssignmentNotificationsEnabled returns a boolean if a field has been set.
func (o *GovernanceNotificationSettingsUpdateAttributes) HasAssignmentNotificationsEnabled() bool {
	return o != nil && o.AssignmentNotificationsEnabled != nil
}

// SetAssignmentNotificationsEnabled gets a reference to the given bool and assigns it to the AssignmentNotificationsEnabled field.
func (o *GovernanceNotificationSettingsUpdateAttributes) SetAssignmentNotificationsEnabled(v bool) {
	o.AssignmentNotificationsEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceNotificationSettingsUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignmentNotificationsEnabled != nil {
		toSerialize["assignment_notifications_enabled"] = o.AssignmentNotificationsEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceNotificationSettingsUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentNotificationsEnabled *bool `json:"assignment_notifications_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_notifications_enabled"})
	} else {
		return err
	}
	o.AssignmentNotificationsEnabled = all.AssignmentNotificationsEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
