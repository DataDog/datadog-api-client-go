// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentOutreachPolicyAttributes Attributes of the organization's outreach policy.
type SecurityAgentOutreachPolicyAttributes struct {
	// Email domains user outreach can be sent to, at most 20. `null` when none are allowed.
	AllowedDomains datadog.NullableList[string] `json:"allowed_domains,omitempty"`
	// The organization's outreach mode. `off` disables user outreach. `manual` only sends outreach an analyst
	// triggers. `automated` auto-sends automated interviews for rules with outreach enabled, to users whose email
	// domain is allowed.
	Mode SecurityAgentOutreachMode `json:"mode"`
	// The stored per-rule automation state, including rules that are explicitly disabled. `null` when the
	// organization has no stored rule state.
	Rules []SecurityAgentOutreachRule `json:"rules,omitempty"`
	// Domains the caller could allow next: their own email domain when it is not allowed yet and is not a
	// consumer mailbox provider. Read-only, never stored. `null` when there is no suggestion.
	SuggestedDomains datadog.NullableList[string] `json:"suggested_domains,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentOutreachPolicyAttributes instantiates a new SecurityAgentOutreachPolicyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentOutreachPolicyAttributes(mode SecurityAgentOutreachMode) *SecurityAgentOutreachPolicyAttributes {
	this := SecurityAgentOutreachPolicyAttributes{}
	this.Mode = mode
	return &this
}

// NewSecurityAgentOutreachPolicyAttributesWithDefaults instantiates a new SecurityAgentOutreachPolicyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentOutreachPolicyAttributesWithDefaults() *SecurityAgentOutreachPolicyAttributes {
	this := SecurityAgentOutreachPolicyAttributes{}
	return &this
}

// GetAllowedDomains returns the AllowedDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityAgentOutreachPolicyAttributes) GetAllowedDomains() []string {
	if o == nil || o.AllowedDomains.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AllowedDomains.Get()
}

// GetAllowedDomainsOk returns a tuple with the AllowedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityAgentOutreachPolicyAttributes) GetAllowedDomainsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedDomains.Get(), o.AllowedDomains.IsSet()
}

// HasAllowedDomains returns a boolean if a field has been set.
func (o *SecurityAgentOutreachPolicyAttributes) HasAllowedDomains() bool {
	return o != nil && o.AllowedDomains.IsSet()
}

// SetAllowedDomains gets a reference to the given datadog.NullableList[string] and assigns it to the AllowedDomains field.
func (o *SecurityAgentOutreachPolicyAttributes) SetAllowedDomains(v []string) {
	o.AllowedDomains.Set(&v)
}

// SetAllowedDomainsNil sets the value for AllowedDomains to be an explicit nil.
func (o *SecurityAgentOutreachPolicyAttributes) SetAllowedDomainsNil() {
	o.AllowedDomains.Set(nil)
}

// UnsetAllowedDomains ensures that no value is present for AllowedDomains, not even an explicit nil.
func (o *SecurityAgentOutreachPolicyAttributes) UnsetAllowedDomains() {
	o.AllowedDomains.Unset()
}

// GetMode returns the Mode field value.
func (o *SecurityAgentOutreachPolicyAttributes) GetMode() SecurityAgentOutreachMode {
	if o == nil {
		var ret SecurityAgentOutreachMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentOutreachPolicyAttributes) GetModeOk() (*SecurityAgentOutreachMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *SecurityAgentOutreachPolicyAttributes) SetMode(v SecurityAgentOutreachMode) {
	o.Mode = v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityAgentOutreachPolicyAttributes) GetRules() []SecurityAgentOutreachRule {
	if o == nil {
		var ret []SecurityAgentOutreachRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityAgentOutreachPolicyAttributes) GetRulesOk() (*[]SecurityAgentOutreachRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SecurityAgentOutreachPolicyAttributes) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []SecurityAgentOutreachRule and assigns it to the Rules field.
func (o *SecurityAgentOutreachPolicyAttributes) SetRules(v []SecurityAgentOutreachRule) {
	o.Rules = v
}

// GetSuggestedDomains returns the SuggestedDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityAgentOutreachPolicyAttributes) GetSuggestedDomains() []string {
	if o == nil || o.SuggestedDomains.Get() == nil {
		var ret []string
		return ret
	}
	return *o.SuggestedDomains.Get()
}

// GetSuggestedDomainsOk returns a tuple with the SuggestedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityAgentOutreachPolicyAttributes) GetSuggestedDomainsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuggestedDomains.Get(), o.SuggestedDomains.IsSet()
}

// HasSuggestedDomains returns a boolean if a field has been set.
func (o *SecurityAgentOutreachPolicyAttributes) HasSuggestedDomains() bool {
	return o != nil && o.SuggestedDomains.IsSet()
}

// SetSuggestedDomains gets a reference to the given datadog.NullableList[string] and assigns it to the SuggestedDomains field.
func (o *SecurityAgentOutreachPolicyAttributes) SetSuggestedDomains(v []string) {
	o.SuggestedDomains.Set(&v)
}

// SetSuggestedDomainsNil sets the value for SuggestedDomains to be an explicit nil.
func (o *SecurityAgentOutreachPolicyAttributes) SetSuggestedDomainsNil() {
	o.SuggestedDomains.Set(nil)
}

// UnsetSuggestedDomains ensures that no value is present for SuggestedDomains, not even an explicit nil.
func (o *SecurityAgentOutreachPolicyAttributes) UnsetSuggestedDomains() {
	o.SuggestedDomains.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentOutreachPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowedDomains.IsSet() {
		toSerialize["allowed_domains"] = o.AllowedDomains.Get()
	}
	toSerialize["mode"] = o.Mode
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.SuggestedDomains.IsSet() {
		toSerialize["suggested_domains"] = o.SuggestedDomains.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentOutreachPolicyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowedDomains   datadog.NullableList[string] `json:"allowed_domains,omitempty"`
		Mode             *SecurityAgentOutreachMode   `json:"mode"`
		Rules            []SecurityAgentOutreachRule  `json:"rules,omitempty"`
		SuggestedDomains datadog.NullableList[string] `json:"suggested_domains,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowed_domains", "mode", "rules", "suggested_domains"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowedDomains = all.AllowedDomains
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.Rules = all.Rules
	o.SuggestedDomains = all.SuggestedDomains

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
