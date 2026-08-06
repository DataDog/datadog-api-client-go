// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceNotificationSettingsAttributes The attributes of the organization-wide governance notification settings.
type GovernanceNotificationSettingsAttributes struct {
	// Whether notifications are sent to users when detections are assigned to them.
	AssignmentNotificationsEnabled bool `json:"assignment_notifications_enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceNotificationSettingsAttributes instantiates a new GovernanceNotificationSettingsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceNotificationSettingsAttributes(assignmentNotificationsEnabled bool) *GovernanceNotificationSettingsAttributes {
	this := GovernanceNotificationSettingsAttributes{}
	this.AssignmentNotificationsEnabled = assignmentNotificationsEnabled
	return &this
}

// NewGovernanceNotificationSettingsAttributesWithDefaults instantiates a new GovernanceNotificationSettingsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceNotificationSettingsAttributesWithDefaults() *GovernanceNotificationSettingsAttributes {
	this := GovernanceNotificationSettingsAttributes{}
	return &this
}

// GetAssignmentNotificationsEnabled returns the AssignmentNotificationsEnabled field value.
func (o *GovernanceNotificationSettingsAttributes) GetAssignmentNotificationsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AssignmentNotificationsEnabled
}

// GetAssignmentNotificationsEnabledOk returns a tuple with the AssignmentNotificationsEnabled field value
// and a boolean to check if the value has been set.
func (o *GovernanceNotificationSettingsAttributes) GetAssignmentNotificationsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentNotificationsEnabled, true
}

// SetAssignmentNotificationsEnabled sets field value.
func (o *GovernanceNotificationSettingsAttributes) SetAssignmentNotificationsEnabled(v bool) {
	o.AssignmentNotificationsEnabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceNotificationSettingsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["assignment_notifications_enabled"] = o.AssignmentNotificationsEnabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceNotificationSettingsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentNotificationsEnabled *bool `json:"assignment_notifications_enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AssignmentNotificationsEnabled == nil {
		return fmt.Errorf("required field assignment_notifications_enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_notifications_enabled"})
	} else {
		return err
	}
	o.AssignmentNotificationsEnabled = *all.AssignmentNotificationsEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
