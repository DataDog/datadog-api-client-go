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

// FeatureFlagSuggestionAttributes Attributes of a feature flag change suggestion.
type FeatureFlagSuggestionAttributes struct {
	// The action proposed by the suggestion.
	Action string `json:"action"`
	// The ID of the flag history entry on which the suggestion is based.
	BaseFlagHistoryId uuid.UUID `json:"base_flag_history_id"`
	// The comment supplied by the suggestion's creator.
	Comment *string `json:"comment,omitempty"`
	// The timestamp when the suggestion was created.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the user who created the suggestion.
	CreatedBy uuid.UUID `json:"created_by"`
	// The legacy review status. Use `lifecycle_status` to distinguish approval from application.
	CurrentStatus FeatureFlagSuggestionAttributesCurrentStatus `json:"current_status"`
	// The value before the proposed change, serialized as a string.
	CurrentValue string `json:"current_value"`
	// The timestamp when the suggestion was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The ID of the user who deleted the suggestion.
	DeletedBy *uuid.UUID `json:"deleted_by,omitempty"`
	// The ID of the environment targeted by the change, if the change is environment-specific.
	EnvironmentId *uuid.UUID `json:"environment_id,omitempty"`
	// The ID of the feature flag that the suggestion changes.
	FeatureFlagId uuid.UUID `json:"feature_flag_id"`
	// The lifecycle status of the suggestion.
	LifecycleStatus FeatureFlagSuggestionAttributesLifecycleStatus `json:"lifecycle_status"`
	// A message explaining why a suggestion was created.
	Message *string `json:"message,omitempty"`
	// The feature flag property to update.
	Property string `json:"property"`
	// The proposed value, serialized as a string. Structured values, such as variants and allocations,
	// are JSON-encoded strings.
	Suggestion string `json:"suggestion"`
	// Metadata associated with a suggestion. Fields other than `variant_id` are set by the server.
	SuggestionMetadata *FeatureFlagSuggestionMetadata `json:"suggestion_metadata,omitempty"`
	// The timestamp when the suggestion was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The ID of the user who last updated the suggestion.
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagSuggestionAttributes instantiates a new FeatureFlagSuggestionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagSuggestionAttributes(action string, baseFlagHistoryId uuid.UUID, createdAt time.Time, createdBy uuid.UUID, currentStatus FeatureFlagSuggestionAttributesCurrentStatus, currentValue string, featureFlagId uuid.UUID, lifecycleStatus FeatureFlagSuggestionAttributesLifecycleStatus, property string, suggestion string) *FeatureFlagSuggestionAttributes {
	this := FeatureFlagSuggestionAttributes{}
	this.Action = action
	this.BaseFlagHistoryId = baseFlagHistoryId
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.CurrentStatus = currentStatus
	this.CurrentValue = currentValue
	this.FeatureFlagId = featureFlagId
	this.LifecycleStatus = lifecycleStatus
	this.Property = property
	this.Suggestion = suggestion
	return &this
}

// NewFeatureFlagSuggestionAttributesWithDefaults instantiates a new FeatureFlagSuggestionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagSuggestionAttributesWithDefaults() *FeatureFlagSuggestionAttributes {
	this := FeatureFlagSuggestionAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *FeatureFlagSuggestionAttributes) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *FeatureFlagSuggestionAttributes) SetAction(v string) {
	o.Action = v
}

// GetBaseFlagHistoryId returns the BaseFlagHistoryId field value.
func (o *FeatureFlagSuggestionAttributes) GetBaseFlagHistoryId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.BaseFlagHistoryId
}

// GetBaseFlagHistoryIdOk returns a tuple with the BaseFlagHistoryId field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetBaseFlagHistoryIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseFlagHistoryId, true
}

// SetBaseFlagHistoryId sets field value.
func (o *FeatureFlagSuggestionAttributes) SetBaseFlagHistoryId(v uuid.UUID) {
	o.BaseFlagHistoryId = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasComment() bool {
	return o != nil && o.Comment != nil
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *FeatureFlagSuggestionAttributes) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FeatureFlagSuggestionAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FeatureFlagSuggestionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *FeatureFlagSuggestionAttributes) GetCreatedBy() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *FeatureFlagSuggestionAttributes) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = v
}

// GetCurrentStatus returns the CurrentStatus field value.
func (o *FeatureFlagSuggestionAttributes) GetCurrentStatus() FeatureFlagSuggestionAttributesCurrentStatus {
	if o == nil {
		var ret FeatureFlagSuggestionAttributesCurrentStatus
		return ret
	}
	return o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetCurrentStatusOk() (*FeatureFlagSuggestionAttributesCurrentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentStatus, true
}

// SetCurrentStatus sets field value.
func (o *FeatureFlagSuggestionAttributes) SetCurrentStatus(v FeatureFlagSuggestionAttributesCurrentStatus) {
	o.CurrentStatus = v
}

// GetCurrentValue returns the CurrentValue field value.
func (o *FeatureFlagSuggestionAttributes) GetCurrentValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetCurrentValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentValue, true
}

// SetCurrentValue sets field value.
func (o *FeatureFlagSuggestionAttributes) SetCurrentValue(v string) {
	o.CurrentValue = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasDeletedAt() bool {
	return o != nil && o.DeletedAt != nil
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *FeatureFlagSuggestionAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetDeletedBy() uuid.UUID {
	if o == nil || o.DeletedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetDeletedByOk() (*uuid.UUID, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasDeletedBy() bool {
	return o != nil && o.DeletedBy != nil
}

// SetDeletedBy gets a reference to the given uuid.UUID and assigns it to the DeletedBy field.
func (o *FeatureFlagSuggestionAttributes) SetDeletedBy(v uuid.UUID) {
	o.DeletedBy = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetEnvironmentId() uuid.UUID {
	if o == nil || o.EnvironmentId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given uuid.UUID and assigns it to the EnvironmentId field.
func (o *FeatureFlagSuggestionAttributes) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = &v
}

// GetFeatureFlagId returns the FeatureFlagId field value.
func (o *FeatureFlagSuggestionAttributes) GetFeatureFlagId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FeatureFlagId
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetFeatureFlagIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureFlagId, true
}

// SetFeatureFlagId sets field value.
func (o *FeatureFlagSuggestionAttributes) SetFeatureFlagId(v uuid.UUID) {
	o.FeatureFlagId = v
}

// GetLifecycleStatus returns the LifecycleStatus field value.
func (o *FeatureFlagSuggestionAttributes) GetLifecycleStatus() FeatureFlagSuggestionAttributesLifecycleStatus {
	if o == nil {
		var ret FeatureFlagSuggestionAttributesLifecycleStatus
		return ret
	}
	return o.LifecycleStatus
}

// GetLifecycleStatusOk returns a tuple with the LifecycleStatus field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetLifecycleStatusOk() (*FeatureFlagSuggestionAttributesLifecycleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifecycleStatus, true
}

// SetLifecycleStatus sets field value.
func (o *FeatureFlagSuggestionAttributes) SetLifecycleStatus(v FeatureFlagSuggestionAttributesLifecycleStatus) {
	o.LifecycleStatus = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FeatureFlagSuggestionAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetProperty returns the Property field value.
func (o *FeatureFlagSuggestionAttributes) GetProperty() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Property
}

// GetPropertyOk returns a tuple with the Property field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetPropertyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Property, true
}

// SetProperty sets field value.
func (o *FeatureFlagSuggestionAttributes) SetProperty(v string) {
	o.Property = v
}

// GetSuggestion returns the Suggestion field value.
func (o *FeatureFlagSuggestionAttributes) GetSuggestion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetSuggestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value.
func (o *FeatureFlagSuggestionAttributes) SetSuggestion(v string) {
	o.Suggestion = v
}

// GetSuggestionMetadata returns the SuggestionMetadata field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetSuggestionMetadata() FeatureFlagSuggestionMetadata {
	if o == nil || o.SuggestionMetadata == nil {
		var ret FeatureFlagSuggestionMetadata
		return ret
	}
	return *o.SuggestionMetadata
}

// GetSuggestionMetadataOk returns a tuple with the SuggestionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetSuggestionMetadataOk() (*FeatureFlagSuggestionMetadata, bool) {
	if o == nil || o.SuggestionMetadata == nil {
		return nil, false
	}
	return o.SuggestionMetadata, true
}

// HasSuggestionMetadata returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasSuggestionMetadata() bool {
	return o != nil && o.SuggestionMetadata != nil
}

// SetSuggestionMetadata gets a reference to the given FeatureFlagSuggestionMetadata and assigns it to the SuggestionMetadata field.
func (o *FeatureFlagSuggestionAttributes) SetSuggestionMetadata(v FeatureFlagSuggestionMetadata) {
	o.SuggestionMetadata = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FeatureFlagSuggestionAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *FeatureFlagSuggestionAttributes) GetUpdatedBy() uuid.UUID {
	if o == nil || o.UpdatedBy == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagSuggestionAttributes) GetUpdatedByOk() (*uuid.UUID, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *FeatureFlagSuggestionAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given uuid.UUID and assigns it to the UpdatedBy field.
func (o *FeatureFlagSuggestionAttributes) SetUpdatedBy(v uuid.UUID) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagSuggestionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["base_flag_history_id"] = o.BaseFlagHistoryId
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["current_status"] = o.CurrentStatus
	toSerialize["current_value"] = o.CurrentValue
	if o.DeletedAt != nil {
		if o.DeletedAt.Nanosecond() == 0 {
			toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DeletedBy != nil {
		toSerialize["deleted_by"] = o.DeletedBy
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	toSerialize["feature_flag_id"] = o.FeatureFlagId
	toSerialize["lifecycle_status"] = o.LifecycleStatus
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	toSerialize["property"] = o.Property
	toSerialize["suggestion"] = o.Suggestion
	if o.SuggestionMetadata != nil {
		toSerialize["suggestion_metadata"] = o.SuggestionMetadata
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagSuggestionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action             *string                                         `json:"action"`
		BaseFlagHistoryId  *uuid.UUID                                      `json:"base_flag_history_id"`
		Comment            *string                                         `json:"comment,omitempty"`
		CreatedAt          *time.Time                                      `json:"created_at"`
		CreatedBy          *uuid.UUID                                      `json:"created_by"`
		CurrentStatus      *FeatureFlagSuggestionAttributesCurrentStatus   `json:"current_status"`
		CurrentValue       *string                                         `json:"current_value"`
		DeletedAt          *time.Time                                      `json:"deleted_at,omitempty"`
		DeletedBy          *uuid.UUID                                      `json:"deleted_by,omitempty"`
		EnvironmentId      *uuid.UUID                                      `json:"environment_id,omitempty"`
		FeatureFlagId      *uuid.UUID                                      `json:"feature_flag_id"`
		LifecycleStatus    *FeatureFlagSuggestionAttributesLifecycleStatus `json:"lifecycle_status"`
		Message            *string                                         `json:"message,omitempty"`
		Property           *string                                         `json:"property"`
		Suggestion         *string                                         `json:"suggestion"`
		SuggestionMetadata *FeatureFlagSuggestionMetadata                  `json:"suggestion_metadata,omitempty"`
		UpdatedAt          *time.Time                                      `json:"updated_at,omitempty"`
		UpdatedBy          *uuid.UUID                                      `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.BaseFlagHistoryId == nil {
		return fmt.Errorf("required field base_flag_history_id missing")
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
	if all.CurrentValue == nil {
		return fmt.Errorf("required field current_value missing")
	}
	if all.FeatureFlagId == nil {
		return fmt.Errorf("required field feature_flag_id missing")
	}
	if all.LifecycleStatus == nil {
		return fmt.Errorf("required field lifecycle_status missing")
	}
	if all.Property == nil {
		return fmt.Errorf("required field property missing")
	}
	if all.Suggestion == nil {
		return fmt.Errorf("required field suggestion missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "base_flag_history_id", "comment", "created_at", "created_by", "current_status", "current_value", "deleted_at", "deleted_by", "environment_id", "feature_flag_id", "lifecycle_status", "message", "property", "suggestion", "suggestion_metadata", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Action = *all.Action
	o.BaseFlagHistoryId = *all.BaseFlagHistoryId
	o.Comment = all.Comment
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if !all.CurrentStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.CurrentStatus = *all.CurrentStatus
	}
	o.CurrentValue = *all.CurrentValue
	o.DeletedAt = all.DeletedAt
	o.DeletedBy = all.DeletedBy
	o.EnvironmentId = all.EnvironmentId
	o.FeatureFlagId = *all.FeatureFlagId
	if !all.LifecycleStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.LifecycleStatus = *all.LifecycleStatus
	}
	o.Message = all.Message
	o.Property = *all.Property
	o.Suggestion = *all.Suggestion
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
