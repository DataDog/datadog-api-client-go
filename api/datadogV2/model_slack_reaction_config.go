// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlackReactionConfig Configuration for a Slack emoji reaction trigger.
type SlackReactionConfig struct {
	// The Slack emoji reaction name.
	ReactionEmoji string `json:"reactionEmoji"`
	// The Slack workspace ID.
	TeamId string `json:"teamId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSlackReactionConfig instantiates a new SlackReactionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSlackReactionConfig(reactionEmoji string, teamId string) *SlackReactionConfig {
	this := SlackReactionConfig{}
	this.ReactionEmoji = reactionEmoji
	this.TeamId = teamId
	return &this
}

// NewSlackReactionConfigWithDefaults instantiates a new SlackReactionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSlackReactionConfigWithDefaults() *SlackReactionConfig {
	this := SlackReactionConfig{}
	return &this
}

// GetReactionEmoji returns the ReactionEmoji field value.
func (o *SlackReactionConfig) GetReactionEmoji() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReactionEmoji
}

// GetReactionEmojiOk returns a tuple with the ReactionEmoji field value
// and a boolean to check if the value has been set.
func (o *SlackReactionConfig) GetReactionEmojiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReactionEmoji, true
}

// SetReactionEmoji sets field value.
func (o *SlackReactionConfig) SetReactionEmoji(v string) {
	o.ReactionEmoji = v
}

// GetTeamId returns the TeamId field value.
func (o *SlackReactionConfig) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *SlackReactionConfig) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value.
func (o *SlackReactionConfig) SetTeamId(v string) {
	o.TeamId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SlackReactionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["reactionEmoji"] = o.ReactionEmoji
	toSerialize["teamId"] = o.TeamId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SlackReactionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReactionEmoji *string `json:"reactionEmoji"`
		TeamId        *string `json:"teamId"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ReactionEmoji == nil {
		return fmt.Errorf("required field reactionEmoji missing")
	}
	if all.TeamId == nil {
		return fmt.Errorf("required field teamId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reactionEmoji", "teamId"})
	} else {
		return err
	}
	o.ReactionEmoji = *all.ReactionEmoji
	o.TeamId = *all.TeamId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
