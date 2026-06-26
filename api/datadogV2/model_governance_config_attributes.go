// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceConfigAttributes The attributes of a Governance Console configuration.
type GovernanceConfigAttributes struct {
	// Whether notifications are sent to users when detections are assigned to them.
	AssignmentNotificationsEnabled bool `json:"assignment_notifications_enabled"`
	// Whether the Governance Console is enabled for the organization.
	Enabled bool `json:"enabled"`
	// Whether usage attribution is configured for the organization.
	UsageAttributionConfigured bool `json:"usage_attribution_configured"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceConfigAttributes instantiates a new GovernanceConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceConfigAttributes(assignmentNotificationsEnabled bool, enabled bool, usageAttributionConfigured bool) *GovernanceConfigAttributes {
	this := GovernanceConfigAttributes{}
	this.AssignmentNotificationsEnabled = assignmentNotificationsEnabled
	this.Enabled = enabled
	this.UsageAttributionConfigured = usageAttributionConfigured
	return &this
}

// NewGovernanceConfigAttributesWithDefaults instantiates a new GovernanceConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceConfigAttributesWithDefaults() *GovernanceConfigAttributes {
	this := GovernanceConfigAttributes{}
	return &this
}

// GetAssignmentNotificationsEnabled returns the AssignmentNotificationsEnabled field value.
func (o *GovernanceConfigAttributes) GetAssignmentNotificationsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AssignmentNotificationsEnabled
}

// GetAssignmentNotificationsEnabledOk returns a tuple with the AssignmentNotificationsEnabled field value
// and a boolean to check if the value has been set.
func (o *GovernanceConfigAttributes) GetAssignmentNotificationsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentNotificationsEnabled, true
}

// SetAssignmentNotificationsEnabled sets field value.
func (o *GovernanceConfigAttributes) SetAssignmentNotificationsEnabled(v bool) {
	o.AssignmentNotificationsEnabled = v
}

// GetEnabled returns the Enabled field value.
func (o *GovernanceConfigAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GovernanceConfigAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *GovernanceConfigAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetUsageAttributionConfigured returns the UsageAttributionConfigured field value.
func (o *GovernanceConfigAttributes) GetUsageAttributionConfigured() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.UsageAttributionConfigured
}

// GetUsageAttributionConfiguredOk returns a tuple with the UsageAttributionConfigured field value
// and a boolean to check if the value has been set.
func (o *GovernanceConfigAttributes) GetUsageAttributionConfiguredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageAttributionConfigured, true
}

// SetUsageAttributionConfigured sets field value.
func (o *GovernanceConfigAttributes) SetUsageAttributionConfigured(v bool) {
	o.UsageAttributionConfigured = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["assignment_notifications_enabled"] = o.AssignmentNotificationsEnabled
	toSerialize["enabled"] = o.Enabled
	toSerialize["usage_attribution_configured"] = o.UsageAttributionConfigured

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentNotificationsEnabled *bool `json:"assignment_notifications_enabled"`
		Enabled                        *bool `json:"enabled"`
		UsageAttributionConfigured     *bool `json:"usage_attribution_configured"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AssignmentNotificationsEnabled == nil {
		return fmt.Errorf("required field assignment_notifications_enabled missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.UsageAttributionConfigured == nil {
		return fmt.Errorf("required field usage_attribution_configured missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_notifications_enabled", "enabled", "usage_attribution_configured"})
	} else {
		return err
	}
	o.AssignmentNotificationsEnabled = *all.AssignmentNotificationsEnabled
	o.Enabled = *all.Enabled
	o.UsageAttributionConfigured = *all.UsageAttributionConfigured

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
