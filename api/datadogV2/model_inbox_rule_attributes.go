// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InboxRuleAttributes Attributes of the inbox rule
type InboxRuleAttributes struct {
	// Action of the inbox rule
	Action ActionInbox `json:"action"`
	// Date as Unix timestamp in milliseconds
	CreatedAt int64 `json:"created_at"`
	// User creating or modifying a rule
	CreatedBy RuleUser `json:"created_by"`
	// Field used to enable or disable the rule.
	Enabled bool `json:"enabled"`
	// Date as Unix timestamp in milliseconds
	ModifiedAt int64 `json:"modified_at"`
	// User creating or modifying a rule
	ModifiedBy RuleUser `json:"modified_by"`
	// Name of the pipeline rule
	Name string `json:"name"`
	// The definition of an automation pipeline rule scope.
	// A rule can act on specific issue types, security rule types, security rule IDs, rule severities, or a query.
	// The query can be used to filter resources on tags and attributes.
	// The issue type and rule types fields are required.
	Rule Rule `json:"rule"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInboxRuleAttributes instantiates a new InboxRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInboxRuleAttributes(action ActionInbox, createdAt int64, createdBy RuleUser, enabled bool, modifiedAt int64, modifiedBy RuleUser, name string, rule Rule) *InboxRuleAttributes {
	this := InboxRuleAttributes{}
	this.Action = action
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.Rule = rule
	return &this
}

// NewInboxRuleAttributesWithDefaults instantiates a new InboxRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInboxRuleAttributesWithDefaults() *InboxRuleAttributes {
	this := InboxRuleAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *InboxRuleAttributes) GetAction() ActionInbox {
	if o == nil {
		var ret ActionInbox
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetActionOk() (*ActionInbox, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *InboxRuleAttributes) SetAction(v ActionInbox) {
	o.Action = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *InboxRuleAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *InboxRuleAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *InboxRuleAttributes) GetCreatedBy() RuleUser {
	if o == nil {
		var ret RuleUser
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetCreatedByOk() (*RuleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *InboxRuleAttributes) SetCreatedBy(v RuleUser) {
	o.CreatedBy = v
}

// GetEnabled returns the Enabled field value.
func (o *InboxRuleAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *InboxRuleAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *InboxRuleAttributes) GetModifiedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetModifiedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *InboxRuleAttributes) SetModifiedAt(v int64) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *InboxRuleAttributes) GetModifiedBy() RuleUser {
	if o == nil {
		var ret RuleUser
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetModifiedByOk() (*RuleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *InboxRuleAttributes) SetModifiedBy(v RuleUser) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *InboxRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InboxRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetRule returns the Rule field value.
func (o *InboxRuleAttributes) GetRule() Rule {
	if o == nil {
		var ret Rule
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *InboxRuleAttributes) GetRuleOk() (*Rule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *InboxRuleAttributes) SetRule(v Rule) {
	o.Rule = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InboxRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["enabled"] = o.Enabled
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["rule"] = o.Rule

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InboxRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action     *ActionInbox `json:"action"`
		CreatedAt  *int64       `json:"created_at"`
		CreatedBy  *RuleUser    `json:"created_by"`
		Enabled    *bool        `json:"enabled"`
		ModifiedAt *int64       `json:"modified_at"`
		ModifiedBy *RuleUser    `json:"modified_by"`
		Name       *string      `json:"name"`
		Rule       *Rule        `json:"rule"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "created_at", "created_by", "enabled", "modified_at", "modified_by", "name", "rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Action = *all.Action
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	o.Enabled = *all.Enabled
	o.ModifiedAt = *all.ModifiedAt
	if all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = *all.ModifiedBy
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
