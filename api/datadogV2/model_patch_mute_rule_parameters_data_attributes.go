// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchMuteRuleParametersDataAttributes Attributes of the mute rule patch request: the rule name, the rule details, the associated action, and the enabled field.
type PatchMuteRuleParametersDataAttributes struct {
	// Action of the mute rule
	Action *ActionMute `json:"action,omitempty"`
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

// NewPatchMuteRuleParametersDataAttributes instantiates a new PatchMuteRuleParametersDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchMuteRuleParametersDataAttributes() *PatchMuteRuleParametersDataAttributes {
	this := PatchMuteRuleParametersDataAttributes{}
	return &this
}

// NewPatchMuteRuleParametersDataAttributesWithDefaults instantiates a new PatchMuteRuleParametersDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchMuteRuleParametersDataAttributesWithDefaults() *PatchMuteRuleParametersDataAttributes {
	this := PatchMuteRuleParametersDataAttributes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PatchMuteRuleParametersDataAttributes) GetAction() ActionMute {
	if o == nil || o.Action == nil {
		var ret ActionMute
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMuteRuleParametersDataAttributes) GetActionOk() (*ActionMute, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PatchMuteRuleParametersDataAttributes) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given ActionMute and assigns it to the Action field.
func (o *PatchMuteRuleParametersDataAttributes) SetAction(v ActionMute) {
	o.Action = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchMuteRuleParametersDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMuteRuleParametersDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchMuteRuleParametersDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchMuteRuleParametersDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchMuteRuleParametersDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMuteRuleParametersDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchMuteRuleParametersDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchMuteRuleParametersDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *PatchMuteRuleParametersDataAttributes) GetRule() Rule {
	if o == nil || o.Rule == nil {
		var ret Rule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchMuteRuleParametersDataAttributes) GetRuleOk() (*Rule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *PatchMuteRuleParametersDataAttributes) HasRule() bool {
	return o != nil && o.Rule != nil
}

// SetRule gets a reference to the given Rule and assigns it to the Rule field.
func (o *PatchMuteRuleParametersDataAttributes) SetRule(v Rule) {
	o.Rule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchMuteRuleParametersDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *PatchMuteRuleParametersDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action  *ActionMute `json:"action,omitempty"`
		Enabled *bool       `json:"enabled,omitempty"`
		Name    *string     `json:"name,omitempty"`
		Rule    *Rule       `json:"rule,omitempty"`
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
