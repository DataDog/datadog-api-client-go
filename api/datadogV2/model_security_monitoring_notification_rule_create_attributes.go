// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringNotificationRuleCreateAttributes Attributes of a notification rule to be created.
type SecurityMonitoringNotificationRuleCreateAttributes struct {
	// Whether the notification rule is enabled.
	Enabled bool `json:"enabled"`
	// The name of the notification rule.
	Name string `json:"name"`
	// Selectors describing the notification rule.
	Selectors SecurityMonitoringNotificationRuleSelectors `json:"selectors"`
	// Set of targets to notify.
	Targets []string `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewSecurityMonitoringNotificationRuleCreateAttributes instantiates a new SecurityMonitoringNotificationRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringNotificationRuleCreateAttributes(enabled bool, name string, selectors SecurityMonitoringNotificationRuleSelectors, targets []string) *SecurityMonitoringNotificationRuleCreateAttributes {
	this := SecurityMonitoringNotificationRuleCreateAttributes{}
	this.Enabled = enabled
	this.Name = name
	this.Selectors = selectors
	this.Targets = targets
	return &this
}

// NewSecurityMonitoringNotificationRuleCreateAttributesWithDefaults instantiates a new SecurityMonitoringNotificationRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringNotificationRuleCreateAttributesWithDefaults() *SecurityMonitoringNotificationRuleCreateAttributes {
	this := SecurityMonitoringNotificationRuleCreateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetName returns the Name field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSelectors returns the Selectors field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetSelectors() SecurityMonitoringNotificationRuleSelectors {
	if o == nil {
		var ret SecurityMonitoringNotificationRuleSelectors
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetSelectorsOk() (*SecurityMonitoringNotificationRuleSelectors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selectors, true
}

// SetSelectors sets field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) SetSelectors(v SecurityMonitoringNotificationRuleSelectors) {
	o.Selectors = v
}

// GetTargets returns the Targets field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) GetTargetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) SetTargets(v []string) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringNotificationRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["name"] = o.Name
	toSerialize["selectors"] = o.Selectors
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringNotificationRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                                        `json:"enabled"`
		Name      *string                                      `json:"name"`
		Selectors *SecurityMonitoringNotificationRuleSelectors `json:"selectors"`
		Targets   *[]string                                    `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Selectors == nil {
		return fmt.Errorf("required field selectors missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "name", "selectors", "targets"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = *all.Enabled
	o.Name = *all.Name
	if all.Selectors.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Selectors = *all.Selectors
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
