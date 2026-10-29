// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExecutionPolicySlackScopeRule A rule restricting a Slack scope to specific private channel names.
type ExecutionPolicySlackScopeRule struct {
	// The private channel name selectors this rule applies to. Selectors are lowercase private channel names
	// without the leading `#`, and can contain `*` wildcards. Use `*` alone to match any private channel name.
	ChannelNameSelectors []string `json:"channel_name_selectors"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExecutionPolicySlackScopeRule instantiates a new ExecutionPolicySlackScopeRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExecutionPolicySlackScopeRule(channelNameSelectors []string) *ExecutionPolicySlackScopeRule {
	this := ExecutionPolicySlackScopeRule{}
	this.ChannelNameSelectors = channelNameSelectors
	return &this
}

// NewExecutionPolicySlackScopeRuleWithDefaults instantiates a new ExecutionPolicySlackScopeRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExecutionPolicySlackScopeRuleWithDefaults() *ExecutionPolicySlackScopeRule {
	this := ExecutionPolicySlackScopeRule{}
	return &this
}

// GetChannelNameSelectors returns the ChannelNameSelectors field value.
func (o *ExecutionPolicySlackScopeRule) GetChannelNameSelectors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ChannelNameSelectors
}

// GetChannelNameSelectorsOk returns a tuple with the ChannelNameSelectors field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicySlackScopeRule) GetChannelNameSelectorsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelNameSelectors, true
}

// SetChannelNameSelectors sets field value.
func (o *ExecutionPolicySlackScopeRule) SetChannelNameSelectors(v []string) {
	o.ChannelNameSelectors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExecutionPolicySlackScopeRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["channel_name_selectors"] = o.ChannelNameSelectors

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExecutionPolicySlackScopeRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelNameSelectors *[]string `json:"channel_name_selectors"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChannelNameSelectors == nil {
		return fmt.Errorf("required field channel_name_selectors missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel_name_selectors"})
	} else {
		return err
	}
	o.ChannelNameSelectors = *all.ChannelNameSelectors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
