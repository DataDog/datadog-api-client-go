// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateFlagSuggestionAttributes Attributes for creating a flag suggestion.
type CreateFlagSuggestionAttributes struct {
	// The type of change action for a suggestion.
	Action FlagSuggestionAction `json:"action"`
	// Optional comment explaining the change.
	Comment *string `json:"comment,omitempty"`
	// The environment ID for environment-scoped changes.
	EnvironmentId *uuid.UUID `json:"environment_id,omitempty"`
	// Notification handles (without @ prefix) to receive approval or rejection notifications.
	// For example, an email address or Slack channel name.
	NotificationRuleTargets []string `json:"notification_rule_targets"`
	// The flag property being changed.
	Property FlagSuggestionProperty `json:"property"`
	// The suggested new value (empty string for flag-level actions like archive, JSON-encoded for complex properties like allocations).
	Suggestion *string `json:"suggestion,omitempty"`
	// Optional metadata for a suggestion.
	SuggestionMetadata *SuggestionMetadata `json:"suggestion_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateFlagSuggestionAttributes instantiates a new CreateFlagSuggestionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateFlagSuggestionAttributes(action FlagSuggestionAction, notificationRuleTargets []string, property FlagSuggestionProperty) *CreateFlagSuggestionAttributes {
	this := CreateFlagSuggestionAttributes{}
	this.Action = action
	this.NotificationRuleTargets = notificationRuleTargets
	this.Property = property
	return &this
}

// NewCreateFlagSuggestionAttributesWithDefaults instantiates a new CreateFlagSuggestionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateFlagSuggestionAttributesWithDefaults() *CreateFlagSuggestionAttributes {
	this := CreateFlagSuggestionAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *CreateFlagSuggestionAttributes) GetAction() FlagSuggestionAction {
	if o == nil {
		var ret FlagSuggestionAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *CreateFlagSuggestionAttributes) GetActionOk() (*FlagSuggestionAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *CreateFlagSuggestionAttributes) SetAction(v FlagSuggestionAction) {
	o.Action = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateFlagSuggestionAttributes) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlagSuggestionAttributes) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateFlagSuggestionAttributes) HasComment() bool {
	return o != nil && o.Comment != nil
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateFlagSuggestionAttributes) SetComment(v string) {
	o.Comment = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *CreateFlagSuggestionAttributes) GetEnvironmentId() uuid.UUID {
	if o == nil || o.EnvironmentId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlagSuggestionAttributes) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *CreateFlagSuggestionAttributes) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given uuid.UUID and assigns it to the EnvironmentId field.
func (o *CreateFlagSuggestionAttributes) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = &v
}

// GetNotificationRuleTargets returns the NotificationRuleTargets field value.
func (o *CreateFlagSuggestionAttributes) GetNotificationRuleTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NotificationRuleTargets
}

// GetNotificationRuleTargetsOk returns a tuple with the NotificationRuleTargets field value
// and a boolean to check if the value has been set.
func (o *CreateFlagSuggestionAttributes) GetNotificationRuleTargetsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationRuleTargets, true
}

// SetNotificationRuleTargets sets field value.
func (o *CreateFlagSuggestionAttributes) SetNotificationRuleTargets(v []string) {
	o.NotificationRuleTargets = v
}

// GetProperty returns the Property field value.
func (o *CreateFlagSuggestionAttributes) GetProperty() FlagSuggestionProperty {
	if o == nil {
		var ret FlagSuggestionProperty
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *CreateFlagSuggestionAttributes) GetPropertyOk() (*FlagSuggestionProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *CreateFlagSuggestionAttributes) SetProperty(v FlagSuggestionProperty) {
	o.Property = v
}

// GetSuggestion returns the Suggestion field value if set, zero value otherwise.
func (o *CreateFlagSuggestionAttributes) GetSuggestion() string {
	if o == nil || o.Suggestion == nil {
		var ret string
		return ret
	}
	return *o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlagSuggestionAttributes) GetSuggestionOk() (*string, bool) {
	if o == nil || o.Suggestion == nil {
		return nil, false
	}
	return o.Suggestion, true
}

// HasSuggestion returns a boolean if a field has been set.
func (o *CreateFlagSuggestionAttributes) HasSuggestion() bool {
	return o != nil && o.Suggestion != nil
}

// SetSuggestion gets a reference to the given string and assigns it to the Suggestion field.
func (o *CreateFlagSuggestionAttributes) SetSuggestion(v string) {
	o.Suggestion = &v
}

// GetSuggestionMetadata returns the SuggestionMetadata field value if set, zero value otherwise.
func (o *CreateFlagSuggestionAttributes) GetSuggestionMetadata() SuggestionMetadata {
	if o == nil || o.SuggestionMetadata == nil {
		var ret SuggestionMetadata
		return ret
	}
	return *o.SuggestionMetadata
}

// GetSuggestionMetadataOk returns a tuple with the SuggestionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlagSuggestionAttributes) GetSuggestionMetadataOk() (*SuggestionMetadata, bool) {
	if o == nil || o.SuggestionMetadata == nil {
		return nil, false
	}
	return o.SuggestionMetadata, true
}

// HasSuggestionMetadata returns a boolean if a field has been set.
func (o *CreateFlagSuggestionAttributes) HasSuggestionMetadata() bool {
	return o != nil && o.SuggestionMetadata != nil
}

// SetSuggestionMetadata gets a reference to the given SuggestionMetadata and assigns it to the SuggestionMetadata field.
func (o *CreateFlagSuggestionAttributes) SetSuggestionMetadata(v SuggestionMetadata) {
	o.SuggestionMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateFlagSuggestionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	toSerialize["notification_rule_targets"] = o.NotificationRuleTargets
	toSerialize["property"] = o.Property
	if o.Suggestion != nil {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.SuggestionMetadata != nil {
		toSerialize["suggestion_metadata"] = o.SuggestionMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateFlagSuggestionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action                  *FlagSuggestionAction   `json:"action"`
		Comment                 *string                 `json:"comment,omitempty"`
		EnvironmentId           *uuid.UUID              `json:"environment_id,omitempty"`
		NotificationRuleTargets *[]string               `json:"notification_rule_targets"`
		Property                *FlagSuggestionProperty `json:"property"`
		Suggestion              *string                 `json:"suggestion,omitempty"`
		SuggestionMetadata      *SuggestionMetadata     `json:"suggestion_metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.NotificationRuleTargets == nil {
		return fmt.Errorf("required field notification_rule_targets missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "comment", "environment_id", "notification_rule_targets", "property", "suggestion", "suggestion_metadata"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.Comment = all.Comment
	o.EnvironmentId = all.EnvironmentId
	o.NotificationRuleTargets = *all.NotificationRuleTargets
	if !all.Property.IsValid() {
		hasInvalidField = true
	} else {
		o.Property = *all.Property
	}
	o.Suggestion = all.Suggestion
	if all.SuggestionMetadata != nil && all.SuggestionMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SuggestionMetadata = all.SuggestionMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
