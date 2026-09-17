// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRuleAttributesCreate Attributes for creating or updating a severity modifier rule.
type SeverityModifierRuleAttributesCreate struct {
	// The action to take when a severity modifier rule matches a finding. This is a discriminated union on `type`: `set` assigns a fixed severity, while `shift` moves the severity up or down by one rank.
	//
	// A severity modifier rule's `rule.query` must not filter on `@severity` or on the `@severity_details.user_adjusted.*` namespace.
	//
	// Use `@severity_details.adjusted.value` instead, which reflects the severity before user-defined adjustments.
	Action SeverityModifierRuleAction `json:"action"`
	// Whether the severity modifier rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the severity modifier rule.
	Name string `json:"name"`
	// Defines the scope of findings to which the automation rule applies.
	Rule AutomationRuleScope `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityModifierRuleAttributesCreate instantiates a new SeverityModifierRuleAttributesCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityModifierRuleAttributesCreate(action SeverityModifierRuleAction, name string, rule AutomationRuleScope) *SeverityModifierRuleAttributesCreate {
	this := SeverityModifierRuleAttributesCreate{}
	this.Action = action
	this.Name = name
	this.Rule = rule
	return &this
}

// NewSeverityModifierRuleAttributesCreateWithDefaults instantiates a new SeverityModifierRuleAttributesCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityModifierRuleAttributesCreateWithDefaults() *SeverityModifierRuleAttributesCreate {
	this := SeverityModifierRuleAttributesCreate{}
	return &this
}

// GetAction returns the Action field value.
func (o *SeverityModifierRuleAttributesCreate) GetAction() SeverityModifierRuleAction {
	if o == nil {
		var ret SeverityModifierRuleAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleAttributesCreate) GetActionOk() (*SeverityModifierRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *SeverityModifierRuleAttributesCreate) SetAction(v SeverityModifierRuleAction) {
	o.Action = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SeverityModifierRuleAttributesCreate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleAttributesCreate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SeverityModifierRuleAttributesCreate) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SeverityModifierRuleAttributesCreate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value.
func (o *SeverityModifierRuleAttributesCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleAttributesCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SeverityModifierRuleAttributesCreate) SetName(v string) {
	o.Name = v
}

// GetRule returns the Rule field value.
func (o *SeverityModifierRuleAttributesCreate) GetRule() AutomationRuleScope {
	if o == nil {
		var ret AutomationRuleScope
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRuleAttributesCreate) GetRuleOk() (*AutomationRuleScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *SeverityModifierRuleAttributesCreate) SetRule(v AutomationRuleScope) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityModifierRuleAttributesCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["name"] = o.Name
	toSerialize["rule"] = o.Rule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityModifierRuleAttributesCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *SeverityModifierRuleAction `json:"action"`
		Enabled *bool                       `json:"enabled,omitempty"`
		Name    *string                     `json:"name"`
		Rule    *AutomationRuleScope        `json:"rule"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "enabled", "name", "rule"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Action = *all.Action
	o.Enabled = all.Enabled
	o.Name = *all.Name
	if all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = *all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
