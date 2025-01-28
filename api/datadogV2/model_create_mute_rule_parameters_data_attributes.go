// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateMuteRuleParametersDataAttributes Attributes of the mute rule create request: the rule name, the rule details, the associated action, and the optional enabled field.
type CreateMuteRuleParametersDataAttributes struct {
	// Action of the mute rule
	Action ActionMute `json:"action"`
	// Field used to enable or disable the rule.
	Enabled *bool `json:"enabled,omitempty"`
	// Name of the pipeline rule
	Name string `json:"name"`
	// The definition of an automation pipeline rule scope.
	// A rule can act on specific issue types, security rule types, security rule IDs, rule severities, or a query.
	// The query can be used to filter resources on tags and attributes.
	// The issue type and rule types fields are required.
	Rule AutomationRule `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateMuteRuleParametersDataAttributes instantiates a new CreateMuteRuleParametersDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateMuteRuleParametersDataAttributes(action ActionMute, name string, rule AutomationRule) *CreateMuteRuleParametersDataAttributes {
	this := CreateMuteRuleParametersDataAttributes{}
	this.Action = action
	this.Name = name
	this.Rule = rule
	return &this
}

// NewCreateMuteRuleParametersDataAttributesWithDefaults instantiates a new CreateMuteRuleParametersDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateMuteRuleParametersDataAttributesWithDefaults() *CreateMuteRuleParametersDataAttributes {
	this := CreateMuteRuleParametersDataAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *CreateMuteRuleParametersDataAttributes) GetAction() ActionMute {
	if o == nil {
		var ret ActionMute
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CreateMuteRuleParametersDataAttributes) GetActionOk() (*ActionMute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *CreateMuteRuleParametersDataAttributes) SetAction(v ActionMute) {
	o.Action = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateMuteRuleParametersDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMuteRuleParametersDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateMuteRuleParametersDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateMuteRuleParametersDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value.
func (o *CreateMuteRuleParametersDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMuteRuleParametersDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateMuteRuleParametersDataAttributes) SetName(v string) {
	o.Name = v
}

// GetRule returns the Rule field value.
func (o *CreateMuteRuleParametersDataAttributes) GetRule() AutomationRule {
	if o == nil {
		var ret AutomationRule
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *CreateMuteRuleParametersDataAttributes) GetRuleOk() (*AutomationRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *CreateMuteRuleParametersDataAttributes) SetRule(v AutomationRule) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateMuteRuleParametersDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *CreateMuteRuleParametersDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *ActionMute     `json:"action"`
		Enabled *bool           `json:"enabled,omitempty"`
		Name    *string         `json:"name"`
		Rule    *AutomationRule `json:"rule"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "enabled", "name", "rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
