// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagChangeEventAttributes Attributes of a feature flag change event.
type FeatureFlagChangeEventAttributes struct {
	// The kind of change.
	Action FeatureFlagChangeEventAction `json:"action"`
	// Additional details about the change, as a JSON-encoded string.
	ChangeMetadata datadog.NullableString `json:"change_metadata,omitempty"`
	// Where the change was made: `UI` for the Datadog app, or `MCP` for a Model Context Protocol (MCP) client.
	CreationSource NullableFeatureFlagChangeEventCreationSource `json:"creation_source,omitempty"`
	// A readable sentence describing the change, as shown on the flag history page and in the Action column of the CSV export.
	Description *string `json:"description,omitempty"`
	// The environment the change applies to. Null for a change to the flag itself rather than to one environment.
	Environment NullableFeatureFlagChangeEventEnvironment `json:"environment,omitempty"`
	// The feature flag the change applies to.
	FeatureFlag FeatureFlagChangeEventFlag `json:"feature_flag"`
	// The value after the change, serialized as a string.
	NewValue datadog.NullableString `json:"new_value,omitempty"`
	// The value before the change, serialized as a string.
	PreviousValue datadog.NullableString `json:"previous_value,omitempty"`
	// The part of the feature flag the change applies to.
	Property FeatureFlagChangeEventProperty `json:"property"`
	// Review details of a change request that needed approval.
	// Present on review events, which `include_suggestion_events` adds.
	// Also present on a change that an approved change request applied, even when `include_suggestion_events` is `false`.
	SuggestionReview *FeatureFlagChangeEventSuggestionReview `json:"suggestion_review,omitempty"`
	// When the change was made.
	Timestamp time.Time `json:"timestamp"`
	// The ID of the user who made the change. Null when no user made the change.
	UserId datadog.NullableString `json:"user_id,omitempty"`
	// The version of the feature flag the change produced.
	VersionNumber int64 `json:"version_number"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagChangeEventAttributes instantiates a new FeatureFlagChangeEventAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagChangeEventAttributes(action FeatureFlagChangeEventAction, featureFlag FeatureFlagChangeEventFlag, property FeatureFlagChangeEventProperty, timestamp time.Time, versionNumber int64) *FeatureFlagChangeEventAttributes {
	this := FeatureFlagChangeEventAttributes{}
	this.Action = action
	this.FeatureFlag = featureFlag
	this.Property = property
	this.Timestamp = timestamp
	this.VersionNumber = versionNumber
	return &this
}

// NewFeatureFlagChangeEventAttributesWithDefaults instantiates a new FeatureFlagChangeEventAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagChangeEventAttributesWithDefaults() *FeatureFlagChangeEventAttributes {
	this := FeatureFlagChangeEventAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *FeatureFlagChangeEventAttributes) GetAction() FeatureFlagChangeEventAction {
	if o == nil {
		var ret FeatureFlagChangeEventAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventAttributes) GetActionOk() (*FeatureFlagChangeEventAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *FeatureFlagChangeEventAttributes) SetAction(v FeatureFlagChangeEventAction) {
	o.Action = v
}

// GetChangeMetadata returns the ChangeMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagChangeEventAttributes) GetChangeMetadata() string {
	if o == nil || o.ChangeMetadata.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChangeMetadata.Get()
}

// GetChangeMetadataOk returns a tuple with the ChangeMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagChangeEventAttributes) GetChangeMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangeMetadata.Get(), o.ChangeMetadata.IsSet()
}

// HasChangeMetadata returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasChangeMetadata() bool {
	return o != nil && o.ChangeMetadata.IsSet()
}

// SetChangeMetadata gets a reference to the given datadog.NullableString and assigns it to the ChangeMetadata field.
func (o *FeatureFlagChangeEventAttributes) SetChangeMetadata(v string) {
	o.ChangeMetadata.Set(&v)
}

// SetChangeMetadataNil sets the value for ChangeMetadata to be an explicit nil.
func (o *FeatureFlagChangeEventAttributes) SetChangeMetadataNil() {
	o.ChangeMetadata.Set(nil)
}

// UnsetChangeMetadata ensures that no value is present for ChangeMetadata, not even an explicit nil.
func (o *FeatureFlagChangeEventAttributes) UnsetChangeMetadata() {
	o.ChangeMetadata.Unset()
}

// GetCreationSource returns the CreationSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagChangeEventAttributes) GetCreationSource() FeatureFlagChangeEventCreationSource {
	if o == nil || o.CreationSource.Get() == nil {
		var ret FeatureFlagChangeEventCreationSource
		return ret
	}
	return *o.CreationSource.Get()
}

// GetCreationSourceOk returns a tuple with the CreationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagChangeEventAttributes) GetCreationSourceOk() (*FeatureFlagChangeEventCreationSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationSource.Get(), o.CreationSource.IsSet()
}

// HasCreationSource returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasCreationSource() bool {
	return o != nil && o.CreationSource.IsSet()
}

// SetCreationSource gets a reference to the given NullableFeatureFlagChangeEventCreationSource and assigns it to the CreationSource field.
func (o *FeatureFlagChangeEventAttributes) SetCreationSource(v FeatureFlagChangeEventCreationSource) {
	o.CreationSource.Set(&v)
}

// SetCreationSourceNil sets the value for CreationSource to be an explicit nil.
func (o *FeatureFlagChangeEventAttributes) SetCreationSourceNil() {
	o.CreationSource.Set(nil)
}

// UnsetCreationSource ensures that no value is present for CreationSource, not even an explicit nil.
func (o *FeatureFlagChangeEventAttributes) UnsetCreationSource() {
	o.CreationSource.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeatureFlagChangeEventAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeatureFlagChangeEventAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagChangeEventAttributes) GetEnvironment() FeatureFlagChangeEventEnvironment {
	if o == nil || o.Environment.Get() == nil {
		var ret FeatureFlagChangeEventEnvironment
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagChangeEventAttributes) GetEnvironmentOk() (*FeatureFlagChangeEventEnvironment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasEnvironment() bool {
	return o != nil && o.Environment.IsSet()
}

// SetEnvironment gets a reference to the given NullableFeatureFlagChangeEventEnvironment and assigns it to the Environment field.
func (o *FeatureFlagChangeEventAttributes) SetEnvironment(v FeatureFlagChangeEventEnvironment) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil.
func (o *FeatureFlagChangeEventAttributes) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil.
func (o *FeatureFlagChangeEventAttributes) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetFeatureFlag returns the FeatureFlag field value.
func (o *FeatureFlagChangeEventAttributes) GetFeatureFlag() FeatureFlagChangeEventFlag {
	if o == nil {
		var ret FeatureFlagChangeEventFlag
		return ret
	}
	return o.FeatureFlag
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventAttributes) GetFeatureFlagOk() (*FeatureFlagChangeEventFlag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlag, true
}

// SetFeatureFlag sets field value.
func (o *FeatureFlagChangeEventAttributes) SetFeatureFlag(v FeatureFlagChangeEventFlag) {
	o.FeatureFlag = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagChangeEventAttributes) GetNewValue() string {
	if o == nil || o.NewValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagChangeEventAttributes) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasNewValue() bool {
	return o != nil && o.NewValue.IsSet()
}

// SetNewValue gets a reference to the given datadog.NullableString and assigns it to the NewValue field.
func (o *FeatureFlagChangeEventAttributes) SetNewValue(v string) {
	o.NewValue.Set(&v)
}

// SetNewValueNil sets the value for NewValue to be an explicit nil.
func (o *FeatureFlagChangeEventAttributes) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil.
func (o *FeatureFlagChangeEventAttributes) UnsetNewValue() {
	o.NewValue.Unset()
}

// GetPreviousValue returns the PreviousValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagChangeEventAttributes) GetPreviousValue() string {
	if o == nil || o.PreviousValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousValue.Get()
}

// GetPreviousValueOk returns a tuple with the PreviousValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagChangeEventAttributes) GetPreviousValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousValue.Get(), o.PreviousValue.IsSet()
}

// HasPreviousValue returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasPreviousValue() bool {
	return o != nil && o.PreviousValue.IsSet()
}

// SetPreviousValue gets a reference to the given datadog.NullableString and assigns it to the PreviousValue field.
func (o *FeatureFlagChangeEventAttributes) SetPreviousValue(v string) {
	o.PreviousValue.Set(&v)
}

// SetPreviousValueNil sets the value for PreviousValue to be an explicit nil.
func (o *FeatureFlagChangeEventAttributes) SetPreviousValueNil() {
	o.PreviousValue.Set(nil)
}

// UnsetPreviousValue ensures that no value is present for PreviousValue, not even an explicit nil.
func (o *FeatureFlagChangeEventAttributes) UnsetPreviousValue() {
	o.PreviousValue.Unset()
}

// GetProperty returns the Property field value.
func (o *FeatureFlagChangeEventAttributes) GetProperty() FeatureFlagChangeEventProperty {
	if o == nil {
		var ret FeatureFlagChangeEventProperty
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventAttributes) GetPropertyOk() (*FeatureFlagChangeEventProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *FeatureFlagChangeEventAttributes) SetProperty(v FeatureFlagChangeEventProperty) {
	o.Property = v
}

// GetSuggestionReview returns the SuggestionReview field value if set, zero value otherwise.
func (o *FeatureFlagChangeEventAttributes) GetSuggestionReview() FeatureFlagChangeEventSuggestionReview {
	if o == nil || o.SuggestionReview == nil {
		var ret FeatureFlagChangeEventSuggestionReview
		return ret
	}
	return *o.SuggestionReview
}

// GetSuggestionReviewOk returns a tuple with the SuggestionReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventAttributes) GetSuggestionReviewOk() (*FeatureFlagChangeEventSuggestionReview, bool) {
	if o == nil || o.SuggestionReview == nil {
		return nil, false
	}
	return o.SuggestionReview, true
}

// HasSuggestionReview returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasSuggestionReview() bool {
	return o != nil && o.SuggestionReview != nil
}

// SetSuggestionReview gets a reference to the given FeatureFlagChangeEventSuggestionReview and assigns it to the SuggestionReview field.
func (o *FeatureFlagChangeEventAttributes) SetSuggestionReview(v FeatureFlagChangeEventSuggestionReview) {
	o.SuggestionReview = &v
}

// GetTimestamp returns the Timestamp field value.
func (o *FeatureFlagChangeEventAttributes) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *FeatureFlagChangeEventAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagChangeEventAttributes) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagChangeEventAttributes) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *FeatureFlagChangeEventAttributes) HasUserId() bool {
	return o != nil && o.UserId.IsSet()
}

// SetUserId gets a reference to the given datadog.NullableString and assigns it to the UserId field.
func (o *FeatureFlagChangeEventAttributes) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil.
func (o *FeatureFlagChangeEventAttributes) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil.
func (o *FeatureFlagChangeEventAttributes) UnsetUserId() {
	o.UserId.Unset()
}

// GetVersionNumber returns the VersionNumber field value.
func (o *FeatureFlagChangeEventAttributes) GetVersionNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagChangeEventAttributes) GetVersionNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionNumber, true
}

// SetVersionNumber sets field value.
func (o *FeatureFlagChangeEventAttributes) SetVersionNumber(v int64) {
	o.VersionNumber = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagChangeEventAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.ChangeMetadata.IsSet() {
		toSerialize["change_metadata"] = o.ChangeMetadata.Get()
	}
	if o.CreationSource.IsSet() {
		toSerialize["creation_source"] = o.CreationSource.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	toSerialize["feature_flag"] = o.FeatureFlag
	if o.NewValue.IsSet() {
		toSerialize["new_value"] = o.NewValue.Get()
	}
	if o.PreviousValue.IsSet() {
		toSerialize["previous_value"] = o.PreviousValue.Get()
	}
	toSerialize["property"] = o.Property
	if o.SuggestionReview != nil {
		toSerialize["suggestion_review"] = o.SuggestionReview
	}
	if o.Timestamp.Nanosecond() == 0 {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	toSerialize["version_number"] = o.VersionNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagChangeEventAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action           *FeatureFlagChangeEventAction                `json:"action"`
		ChangeMetadata   datadog.NullableString                       `json:"change_metadata,omitempty"`
		CreationSource   NullableFeatureFlagChangeEventCreationSource `json:"creation_source,omitempty"`
		Description      *string                                      `json:"description,omitempty"`
		Environment      NullableFeatureFlagChangeEventEnvironment    `json:"environment,omitempty"`
		FeatureFlag      *FeatureFlagChangeEventFlag                  `json:"feature_flag"`
		NewValue         datadog.NullableString                       `json:"new_value,omitempty"`
		PreviousValue    datadog.NullableString                       `json:"previous_value,omitempty"`
		Property         *FeatureFlagChangeEventProperty              `json:"property"`
		SuggestionReview *FeatureFlagChangeEventSuggestionReview      `json:"suggestion_review,omitempty"`
		Timestamp        *time.Time                                   `json:"timestamp"`
		UserId           datadog.NullableString                       `json:"user_id,omitempty"`
		VersionNumber    *int64                                       `json:"version_number"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.FeatureFlag == nil {
		return fmt.Errorf("required field feature_flag missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	if all.VersionNumber == nil {
		return fmt.Errorf("required field version_number missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "change_metadata", "creation_source", "description", "environment", "feature_flag", "new_value", "previous_value", "property", "suggestion_review", "timestamp", "user_id", "version_number"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.ChangeMetadata = all.ChangeMetadata
	if all.CreationSource.Get() != nil && !all.CreationSource.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.CreationSource = all.CreationSource
	}
	o.Description = all.Description
	o.Environment = all.Environment
	if all.FeatureFlag.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FeatureFlag = *all.FeatureFlag
	o.NewValue = all.NewValue
	o.PreviousValue = all.PreviousValue
	if !all.Property.IsValid() {
		hasInvalidField = true
	} else {
		o.Property = *all.Property
	}
	if all.SuggestionReview != nil && all.SuggestionReview.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SuggestionReview = all.SuggestionReview
	o.Timestamp = *all.Timestamp
	o.UserId = all.UserId
	o.VersionNumber = *all.VersionNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
