// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressResponseAttributes Attributes of an event email address resource.
type EventEmailAddressResponseAttributes struct {
	// The alert type of events generated from the email address.
	AlertType *EventEmailAddressAlertType `json:"alert_type,omitempty"`
	// The timestamp of when the event email address was created.
	CreatedAt time.Time `json:"created_at"`
	// A description of the event email address.
	Description datadog.NullableString `json:"description,omitempty"`
	// The generated email address for ingesting events.
	Email string `json:"email"`
	// The format of events ingested through the email address.
	Format EventEmailAddressFormat `json:"format"`
	// The timestamp of when the event email address was last used.
	LastUsedAt datadog.NullableTime `json:"last_used_at,omitempty"`
	// A list of handles to notify when an email is received.
	NotifyHandles []string `json:"notify_handles,omitempty"`
	// The timestamp of when the event email address was revoked.
	RevokedAt datadog.NullableTime `json:"revoked_at,omitempty"`
	// A list of tags to apply to events generated from the email address.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressResponseAttributes instantiates a new EventEmailAddressResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressResponseAttributes(createdAt time.Time, email string, format EventEmailAddressFormat) *EventEmailAddressResponseAttributes {
	this := EventEmailAddressResponseAttributes{}
	this.CreatedAt = createdAt
	this.Email = email
	this.Format = format
	return &this
}

// NewEventEmailAddressResponseAttributesWithDefaults instantiates a new EventEmailAddressResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressResponseAttributesWithDefaults() *EventEmailAddressResponseAttributes {
	this := EventEmailAddressResponseAttributes{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *EventEmailAddressResponseAttributes) GetAlertType() EventEmailAddressAlertType {
	if o == nil || o.AlertType == nil {
		var ret EventEmailAddressAlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseAttributes) GetAlertTypeOk() (*EventEmailAddressAlertType, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *EventEmailAddressResponseAttributes) HasAlertType() bool {
	return o != nil && o.AlertType != nil
}

// SetAlertType gets a reference to the given EventEmailAddressAlertType and assigns it to the AlertType field.
func (o *EventEmailAddressResponseAttributes) SetAlertType(v EventEmailAddressAlertType) {
	o.AlertType = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *EventEmailAddressResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *EventEmailAddressResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventEmailAddressResponseAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventEmailAddressResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EventEmailAddressResponseAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *EventEmailAddressResponseAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *EventEmailAddressResponseAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *EventEmailAddressResponseAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetEmail returns the Email field value.
func (o *EventEmailAddressResponseAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *EventEmailAddressResponseAttributes) SetEmail(v string) {
	o.Email = v
}

// GetFormat returns the Format field value.
func (o *EventEmailAddressResponseAttributes) GetFormat() EventEmailAddressFormat {
	if o == nil {
		var ret EventEmailAddressFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseAttributes) GetFormatOk() (*EventEmailAddressFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *EventEmailAddressResponseAttributes) SetFormat(v EventEmailAddressFormat) {
	o.Format = v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventEmailAddressResponseAttributes) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventEmailAddressResponseAttributes) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *EventEmailAddressResponseAttributes) HasLastUsedAt() bool {
	return o != nil && o.LastUsedAt.IsSet()
}

// SetLastUsedAt gets a reference to the given datadog.NullableTime and assigns it to the LastUsedAt field.
func (o *EventEmailAddressResponseAttributes) SetLastUsedAt(v time.Time) {
	o.LastUsedAt.Set(&v)
}

// SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil.
func (o *EventEmailAddressResponseAttributes) SetLastUsedAtNil() {
	o.LastUsedAt.Set(nil)
}

// UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil.
func (o *EventEmailAddressResponseAttributes) UnsetLastUsedAt() {
	o.LastUsedAt.Unset()
}

// GetNotifyHandles returns the NotifyHandles field value if set, zero value otherwise.
func (o *EventEmailAddressResponseAttributes) GetNotifyHandles() []string {
	if o == nil || o.NotifyHandles == nil {
		var ret []string
		return ret
	}
	return o.NotifyHandles
}

// GetNotifyHandlesOk returns a tuple with the NotifyHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseAttributes) GetNotifyHandlesOk() (*[]string, bool) {
	if o == nil || o.NotifyHandles == nil {
		return nil, false
	}
	return &o.NotifyHandles, true
}

// HasNotifyHandles returns a boolean if a field has been set.
func (o *EventEmailAddressResponseAttributes) HasNotifyHandles() bool {
	return o != nil && o.NotifyHandles != nil
}

// SetNotifyHandles gets a reference to the given []string and assigns it to the NotifyHandles field.
func (o *EventEmailAddressResponseAttributes) SetNotifyHandles(v []string) {
	o.NotifyHandles = v
}

// GetRevokedAt returns the RevokedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventEmailAddressResponseAttributes) GetRevokedAt() time.Time {
	if o == nil || o.RevokedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RevokedAt.Get()
}

// GetRevokedAtOk returns a tuple with the RevokedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventEmailAddressResponseAttributes) GetRevokedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevokedAt.Get(), o.RevokedAt.IsSet()
}

// HasRevokedAt returns a boolean if a field has been set.
func (o *EventEmailAddressResponseAttributes) HasRevokedAt() bool {
	return o != nil && o.RevokedAt.IsSet()
}

// SetRevokedAt gets a reference to the given datadog.NullableTime and assigns it to the RevokedAt field.
func (o *EventEmailAddressResponseAttributes) SetRevokedAt(v time.Time) {
	o.RevokedAt.Set(&v)
}

// SetRevokedAtNil sets the value for RevokedAt to be an explicit nil.
func (o *EventEmailAddressResponseAttributes) SetRevokedAtNil() {
	o.RevokedAt.Set(nil)
}

// UnsetRevokedAt ensures that no value is present for RevokedAt, not even an explicit nil.
func (o *EventEmailAddressResponseAttributes) UnsetRevokedAt() {
	o.RevokedAt.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventEmailAddressResponseAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventEmailAddressResponseAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EventEmailAddressResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AlertType != nil {
		toSerialize["alert_type"] = o.AlertType
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["email"] = o.Email
	toSerialize["format"] = o.Format
	if o.LastUsedAt.IsSet() {
		toSerialize["last_used_at"] = o.LastUsedAt.Get()
	}
	if o.NotifyHandles != nil {
		toSerialize["notify_handles"] = o.NotifyHandles
	}
	if o.RevokedAt.IsSet() {
		toSerialize["revoked_at"] = o.RevokedAt.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventEmailAddressResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertType     *EventEmailAddressAlertType `json:"alert_type,omitempty"`
		CreatedAt     *time.Time                  `json:"created_at"`
		Description   datadog.NullableString      `json:"description,omitempty"`
		Email         *string                     `json:"email"`
		Format        *EventEmailAddressFormat    `json:"format"`
		LastUsedAt    datadog.NullableTime        `json:"last_used_at,omitempty"`
		NotifyHandles []string                    `json:"notify_handles,omitempty"`
		RevokedAt     datadog.NullableTime        `json:"revoked_at,omitempty"`
		Tags          []string                    `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_type", "created_at", "description", "email", "format", "last_used_at", "notify_handles", "revoked_at", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AlertType != nil && !all.AlertType.IsValid() {
		hasInvalidField = true
	} else {
		o.AlertType = all.AlertType
	}
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.Email = *all.Email
	if !all.Format.IsValid() {
		hasInvalidField = true
	} else {
		o.Format = *all.Format
	}
	o.LastUsedAt = all.LastUsedAt
	o.NotifyHandles = all.NotifyHandles
	o.RevokedAt = all.RevokedAt
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
