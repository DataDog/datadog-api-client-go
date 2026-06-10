// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlagSuggestionAttributes Attributes of a flag suggestion.
type FlagSuggestionAttributes struct {
	// The type of change action for a suggestion.
	Action FlagSuggestionAction `json:"action"`
	// The flag history version this suggestion was based on.
	BaseFlagHistoryId *uuid.UUID `json:"base_flag_history_id,omitempty"`
	// Optional comment from the requester.
	Comment datadog.NullableString `json:"comment,omitempty"`
	// When the suggestion was created.
	CreatedAt time.Time `json:"created_at"`
	// UUID of the user who created the suggestion.
	CreatedBy uuid.UUID `json:"created_by"`
	// The status of a flag suggestion.
	CurrentStatus FlagSuggestionStatus `json:"current_status"`
	// The current value before the suggested change (empty string for flag-level actions like archive).
	CurrentValue *string `json:"current_value,omitempty"`
	// When the suggestion was soft-deleted.
	DeletedAt datadog.NullableTime `json:"deleted_at,omitempty"`
	// UUID of the user who deleted the suggestion.
	DeletedBy datadog.NullableString `json:"deleted_by,omitempty"`
	// The environment ID for environment-scoped suggestions. Null for flag-level changes.
	EnvironmentId datadog.NullableString `json:"environment_id,omitempty"`
	// The ID of the feature flag this suggestion applies to.
	FeatureFlagId uuid.UUID `json:"feature_flag_id"`
	// Human-readable message about the suggestion (populated on auto-created suggestions).
	Message *string `json:"message,omitempty"`
	// The flag property being changed.
	Property FlagSuggestionProperty `json:"property"`
	// The suggested new value (JSON-encoded for complex properties, empty string for flag-level actions like archive).
	Suggestion *string `json:"suggestion,omitempty"`
	// Optional metadata for a suggestion.
	SuggestionMetadata *SuggestionMetadata `json:"suggestion_metadata,omitempty"`
	// When the suggestion was last updated.
	UpdatedAt datadog.NullableTime `json:"updated_at,omitempty"`
	// UUID of the user who last updated the suggestion.
	UpdatedBy datadog.NullableString `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlagSuggestionAttributes instantiates a new FlagSuggestionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlagSuggestionAttributes(action FlagSuggestionAction, createdAt time.Time, createdBy uuid.UUID, currentStatus FlagSuggestionStatus, featureFlagId uuid.UUID, property FlagSuggestionProperty) *FlagSuggestionAttributes {
	this := FlagSuggestionAttributes{}
	this.Action = action
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.CurrentStatus = currentStatus
	this.FeatureFlagId = featureFlagId
	this.Property = property
	return &this
}

// NewFlagSuggestionAttributesWithDefaults instantiates a new FlagSuggestionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlagSuggestionAttributesWithDefaults() *FlagSuggestionAttributes {
	this := FlagSuggestionAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *FlagSuggestionAttributes) GetAction() FlagSuggestionAction {
	if o == nil {
		var ret FlagSuggestionAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetActionOk() (*FlagSuggestionAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *FlagSuggestionAttributes) SetAction(v FlagSuggestionAction) {
	o.Action = v
}

// GetBaseFlagHistoryId returns the BaseFlagHistoryId field value if set, zero value otherwise.
func (o *FlagSuggestionAttributes) GetBaseFlagHistoryId() uuid.UUID {
	if o == nil || o.BaseFlagHistoryId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.BaseFlagHistoryId
}

// GetBaseFlagHistoryIdOk returns a tuple with the BaseFlagHistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetBaseFlagHistoryIdOk() (*uuid.UUID, bool) {
	if o == nil || o.BaseFlagHistoryId == nil {
		return nil, false
	}
	return o.BaseFlagHistoryId, true
}

// HasBaseFlagHistoryId returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasBaseFlagHistoryId() bool {
	return o != nil && o.BaseFlagHistoryId != nil
}

// SetBaseFlagHistoryId gets a reference to the given uuid.UUID and assigns it to the BaseFlagHistoryId field.
func (o *FlagSuggestionAttributes) SetBaseFlagHistoryId(v uuid.UUID) {
	o.BaseFlagHistoryId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagSuggestionAttributes) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlagSuggestionAttributes) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasComment() bool {
	return o != nil && o.Comment.IsSet()
}

// SetComment gets a reference to the given datadog.NullableString and assigns it to the Comment field.
func (o *FlagSuggestionAttributes) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil.
func (o *FlagSuggestionAttributes) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil.
func (o *FlagSuggestionAttributes) UnsetComment() {
	o.Comment.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FlagSuggestionAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FlagSuggestionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *FlagSuggestionAttributes) GetCreatedBy() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *FlagSuggestionAttributes) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = v
}

// GetCurrentStatus returns the CurrentStatus field value.
func (o *FlagSuggestionAttributes) GetCurrentStatus() FlagSuggestionStatus {
	if o == nil {
		var ret FlagSuggestionStatus
		return ret
	}
	return o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetCurrentStatusOk() (*FlagSuggestionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentStatus, true
}

// SetCurrentStatus sets field value.
func (o *FlagSuggestionAttributes) SetCurrentStatus(v FlagSuggestionStatus) {
	o.CurrentStatus = v
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise.
func (o *FlagSuggestionAttributes) GetCurrentValue() string {
	if o == nil || o.CurrentValue == nil {
		var ret string
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetCurrentValueOk() (*string, bool) {
	if o == nil || o.CurrentValue == nil {
		return nil, false
	}
	return o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasCurrentValue() bool {
	return o != nil && o.CurrentValue != nil
}

// SetCurrentValue gets a reference to the given string and assigns it to the CurrentValue field.
func (o *FlagSuggestionAttributes) SetCurrentValue(v string) {
	o.CurrentValue = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagSuggestionAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlagSuggestionAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt.IsSet()
}

// SetDeletedAt gets a reference to the given datadog.NullableTime and assigns it to the DeletedAt field.
func (o *FlagSuggestionAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil.
func (o *FlagSuggestionAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil.
func (o *FlagSuggestionAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagSuggestionAttributes) GetDeletedBy() string {
	if o == nil || o.DeletedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeletedBy.Get()
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlagSuggestionAttributes) GetDeletedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedBy.Get(), o.DeletedBy.IsSet()
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasDeletedBy() bool {
	return o != nil && o.DeletedBy.IsSet()
}

// SetDeletedBy gets a reference to the given datadog.NullableString and assigns it to the DeletedBy field.
func (o *FlagSuggestionAttributes) SetDeletedBy(v string) {
	o.DeletedBy.Set(&v)
}

// SetDeletedByNil sets the value for DeletedBy to be an explicit nil.
func (o *FlagSuggestionAttributes) SetDeletedByNil() {
	o.DeletedBy.Set(nil)
}

// UnsetDeletedBy ensures that no value is present for DeletedBy, not even an explicit nil.
func (o *FlagSuggestionAttributes) UnsetDeletedBy() {
	o.DeletedBy.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagSuggestionAttributes) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlagSuggestionAttributes) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId.IsSet()
}

// SetEnvironmentId gets a reference to the given datadog.NullableString and assigns it to the EnvironmentId field.
func (o *FlagSuggestionAttributes) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}

// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil.
func (o *FlagSuggestionAttributes) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil.
func (o *FlagSuggestionAttributes) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetFeatureFlagId returns the FeatureFlagId field value.
func (o *FlagSuggestionAttributes) GetFeatureFlagId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagId
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetFeatureFlagIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagId, true
}

// SetFeatureFlagId sets field value.
func (o *FlagSuggestionAttributes) SetFeatureFlagId(v uuid.UUID) {
	o.FeatureFlagId = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FlagSuggestionAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FlagSuggestionAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetProperty returns the Property field value.
func (o *FlagSuggestionAttributes) GetProperty() FlagSuggestionProperty {
	if o == nil {
		var ret FlagSuggestionProperty
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetPropertyOk() (*FlagSuggestionProperty, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *FlagSuggestionAttributes) SetProperty(v FlagSuggestionProperty) {
	o.Property = v
}

// GetSuggestion returns the Suggestion field value if set, zero value otherwise.
func (o *FlagSuggestionAttributes) GetSuggestion() string {
	if o == nil || o.Suggestion == nil {
		var ret string
		return ret
	}
	return *o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetSuggestionOk() (*string, bool) {
	if o == nil || o.Suggestion == nil {
		return nil, false
	}
	return o.Suggestion, true
}

// HasSuggestion returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasSuggestion() bool {
	return o != nil && o.Suggestion != nil
}

// SetSuggestion gets a reference to the given string and assigns it to the Suggestion field.
func (o *FlagSuggestionAttributes) SetSuggestion(v string) {
	o.Suggestion = &v
}

// GetSuggestionMetadata returns the SuggestionMetadata field value if set, zero value otherwise.
func (o *FlagSuggestionAttributes) GetSuggestionMetadata() SuggestionMetadata {
	if o == nil || o.SuggestionMetadata == nil {
		var ret SuggestionMetadata
		return ret
	}
	return *o.SuggestionMetadata
}

// GetSuggestionMetadataOk returns a tuple with the SuggestionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagSuggestionAttributes) GetSuggestionMetadataOk() (*SuggestionMetadata, bool) {
	if o == nil || o.SuggestionMetadata == nil {
		return nil, false
	}
	return o.SuggestionMetadata, true
}

// HasSuggestionMetadata returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasSuggestionMetadata() bool {
	return o != nil && o.SuggestionMetadata != nil
}

// SetSuggestionMetadata gets a reference to the given SuggestionMetadata and assigns it to the SuggestionMetadata field.
func (o *FlagSuggestionAttributes) SetSuggestionMetadata(v SuggestionMetadata) {
	o.SuggestionMetadata = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagSuggestionAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlagSuggestionAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given datadog.NullableTime and assigns it to the UpdatedAt field.
func (o *FlagSuggestionAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *FlagSuggestionAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *FlagSuggestionAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagSuggestionAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlagSuggestionAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FlagSuggestionAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy.IsSet()
}

// SetUpdatedBy gets a reference to the given datadog.NullableString and assigns it to the UpdatedBy field.
func (o *FlagSuggestionAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil.
func (o *FlagSuggestionAttributes) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil.
func (o *FlagSuggestionAttributes) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FlagSuggestionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	if o.BaseFlagHistoryId != nil {
		toSerialize["base_flag_history_id"] = o.BaseFlagHistoryId
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["current_status"] = o.CurrentStatus
	if o.CurrentValue != nil {
		toSerialize["current_value"] = o.CurrentValue
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.DeletedBy.IsSet() {
		toSerialize["deleted_by"] = o.DeletedBy.Get()
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	toSerialize["feature_flag_id"] = o.FeatureFlagId
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	toSerialize["property"] = o.Property
	if o.Suggestion != nil {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.SuggestionMetadata != nil {
		toSerialize["suggestion_metadata"] = o.SuggestionMetadata
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updated_by"] = o.UpdatedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlagSuggestionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action             *FlagSuggestionAction   `json:"action"`
		BaseFlagHistoryId  *uuid.UUID              `json:"base_flag_history_id,omitempty"`
		Comment            datadog.NullableString  `json:"comment,omitempty"`
		CreatedAt          *time.Time              `json:"created_at"`
		CreatedBy          *uuid.UUID              `json:"created_by"`
		CurrentStatus      *FlagSuggestionStatus   `json:"current_status"`
		CurrentValue       *string                 `json:"current_value,omitempty"`
		DeletedAt          datadog.NullableTime    `json:"deleted_at,omitempty"`
		DeletedBy          datadog.NullableString  `json:"deleted_by,omitempty"`
		EnvironmentId      datadog.NullableString  `json:"environment_id,omitempty"`
		FeatureFlagId      *uuid.UUID              `json:"feature_flag_id"`
		Message            *string                 `json:"message,omitempty"`
		Property           *FlagSuggestionProperty `json:"property"`
		Suggestion         *string                 `json:"suggestion,omitempty"`
		SuggestionMetadata *SuggestionMetadata     `json:"suggestion_metadata,omitempty"`
		UpdatedAt          datadog.NullableTime    `json:"updated_at,omitempty"`
		UpdatedBy          datadog.NullableString  `json:"updated_by,omitempty"`
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
	if all.CurrentStatus == nil {
		return fmt.Errorf("required field current_status missing")
	}
	if all.FeatureFlagId == nil {
		return fmt.Errorf("required field feature_flag_id missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "base_flag_history_id", "comment", "created_at", "created_by", "current_status", "current_value", "deleted_at", "deleted_by", "environment_id", "feature_flag_id", "message", "property", "suggestion", "suggestion_metadata", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.BaseFlagHistoryId = all.BaseFlagHistoryId
	o.Comment = all.Comment
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if !all.CurrentStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.CurrentStatus = *all.CurrentStatus
	}
	o.CurrentValue = all.CurrentValue
	o.DeletedAt = all.DeletedAt
	o.DeletedBy = all.DeletedBy
	o.EnvironmentId = all.EnvironmentId
	o.FeatureFlagId = *all.FeatureFlagId
	o.Message = all.Message
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
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
