// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchInboxRulesParametersDataAttributes Attributes of the inbox rule patch request: the rule name, the rule details, the associated action, and the enabled field.
type PatchInboxRulesParametersDataAttributes struct {
	// Action of the inbox rule
	Action *ActionInbox `json:"action,omitempty"`
	// Field used to enable or disable the rule.
	Enabled *bool `json:"enabled,omitempty"`
	// Name of the pipeline rule
	Name *string `json:"name,omitempty"`
	// The definition of an automation pipeline rule scope.
	// A rule can act on specific issue types, security rule types, security rule IDs, rule severities, or a query.
	// The query can be used to filter resources on tags and attributes.
	// The issue type and rule types fields are required.
	Rule *Rule `json:"rule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchInboxRulesParametersDataAttributes instantiates a new PatchInboxRulesParametersDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchInboxRulesParametersDataAttributes() *PatchInboxRulesParametersDataAttributes {
	this := PatchInboxRulesParametersDataAttributes{}
	return &this
}

// NewPatchInboxRulesParametersDataAttributesWithDefaults instantiates a new PatchInboxRulesParametersDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchInboxRulesParametersDataAttributesWithDefaults() *PatchInboxRulesParametersDataAttributes {
	this := PatchInboxRulesParametersDataAttributes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PatchInboxRulesParametersDataAttributes) GetAction() ActionInbox {
	if o == nil || o.Action == nil {
		var ret ActionInbox
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchInboxRulesParametersDataAttributes) GetActionOk() (*ActionInbox, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PatchInboxRulesParametersDataAttributes) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given ActionInbox and assigns it to the Action field.
func (o *PatchInboxRulesParametersDataAttributes) SetAction(v ActionInbox) {
	o.Action = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchInboxRulesParametersDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchInboxRulesParametersDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchInboxRulesParametersDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchInboxRulesParametersDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchInboxRulesParametersDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchInboxRulesParametersDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchInboxRulesParametersDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchInboxRulesParametersDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *PatchInboxRulesParametersDataAttributes) GetRule() Rule {
	if o == nil || o.Rule == nil {
		var ret Rule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchInboxRulesParametersDataAttributes) GetRuleOk() (*Rule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *PatchInboxRulesParametersDataAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given Rule and assigns it to the Rule field.
func (o *PatchInboxRulesParametersDataAttributes) SetRule(v Rule) {
	o.Rule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchInboxRulesParametersDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchInboxRulesParametersDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *ActionInbox `json:"action,omitempty"`
		Enabled *bool        `json:"enabled,omitempty"`
		Name    *string      `json:"name,omitempty"`
		Rule    *Rule        `json:"rule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "enabled", "name", "rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action != nil && all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Action = all.Action
	o.Enabled = all.Enabled
	o.Name = all.Name
	if all.Rule != nil && all.Rule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rule = all.Rule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
