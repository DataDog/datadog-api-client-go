// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlackTrigger Trigger a workflow from Slack. The workflow must be published.
type SlackTrigger struct {
	// Slack emoji reactions that trigger the workflow.
	ReactionTriggers []SlackReactionConfig `json:"reactionTriggers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSlackTrigger instantiates a new SlackTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSlackTrigger() *SlackTrigger {
	this := SlackTrigger{}
	return &this
}

// NewSlackTriggerWithDefaults instantiates a new SlackTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSlackTriggerWithDefaults() *SlackTrigger {
	this := SlackTrigger{}
	return &this
}

// GetReactionTriggers returns the ReactionTriggers field value if set, zero value otherwise.
func (o *SlackTrigger) GetReactionTriggers() []SlackReactionConfig {
	if o == nil || o.ReactionTriggers == nil {
		var ret []SlackReactionConfig
		return ret
	}
	return o.ReactionTriggers
}

// GetReactionTriggersOk returns a tuple with the ReactionTriggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackTrigger) GetReactionTriggersOk() (*[]SlackReactionConfig, bool) {
	if o == nil || o.ReactionTriggers == nil {
		return nil, false
	}
	return &o.ReactionTriggers, true
}

// HasReactionTriggers returns a boolean if a field has been set.
func (o *SlackTrigger) HasReactionTriggers() bool {
	return o != nil && o.ReactionTriggers != nil
}

// SetReactionTriggers gets a reference to the given []SlackReactionConfig and assigns it to the ReactionTriggers field.
func (o *SlackTrigger) SetReactionTriggers(v []SlackReactionConfig) {
	o.ReactionTriggers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SlackTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ReactionTriggers != nil {
		toSerialize["reactionTriggers"] = o.ReactionTriggers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SlackTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReactionTriggers []SlackReactionConfig `json:"reactionTriggers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reactionTriggers"})
	} else {
		return err
	}
	o.ReactionTriggers = all.ReactionTriggers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
