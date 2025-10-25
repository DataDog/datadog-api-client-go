// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressResponseDataAttributes
type EventEmailAddressResponseDataAttributes struct {
	//
	AlertType *string `json:"alert_type,omitempty"`
	//
	CreatedAt time.Time `json:"created_at"`
	//
	Description *string `json:"description,omitempty"`
	//
	Email string `json:"email"`
	//
	Format string `json:"format"`
	//
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	//
	NotifyHandles []string `json:"notify_handles"`
	//
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	//
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressResponseDataAttributes instantiates a new EventEmailAddressResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressResponseDataAttributes(createdAt time.Time, email string, format string, notifyHandles []string, tags []string) *EventEmailAddressResponseDataAttributes {
	this := EventEmailAddressResponseDataAttributes{}
	this.CreatedAt = createdAt
	this.Email = email
	this.Format = format
	this.NotifyHandles = notifyHandles
	this.Tags = tags
	return &this
}

// NewEventEmailAddressResponseDataAttributesWithDefaults instantiates a new EventEmailAddressResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressResponseDataAttributesWithDefaults() *EventEmailAddressResponseDataAttributes {
	this := EventEmailAddressResponseDataAttributes{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *EventEmailAddressResponseDataAttributes) GetAlertType() string {
	if o == nil || o.AlertType == nil {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetAlertTypeOk() (*string, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *EventEmailAddressResponseDataAttributes) HasAlertType() bool {
	return o != nil && o.AlertType != nil
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *EventEmailAddressResponseDataAttributes) SetAlertType(v string) {
	o.AlertType = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *EventEmailAddressResponseDataAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *EventEmailAddressResponseDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventEmailAddressResponseDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventEmailAddressResponseDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventEmailAddressResponseDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value.
func (o *EventEmailAddressResponseDataAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *EventEmailAddressResponseDataAttributes) SetEmail(v string) {
	o.Email = v
}

// GetFormat returns the Format field value.
func (o *EventEmailAddressResponseDataAttributes) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *EventEmailAddressResponseDataAttributes) SetFormat(v string) {
	o.Format = v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *EventEmailAddressResponseDataAttributes) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *EventEmailAddressResponseDataAttributes) HasLastUsedAt() bool {
	return o != nil && o.LastUsedAt != nil
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *EventEmailAddressResponseDataAttributes) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetNotifyHandles returns the NotifyHandles field value.
func (o *EventEmailAddressResponseDataAttributes) GetNotifyHandles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NotifyHandles
}

// GetNotifyHandlesOk returns a tuple with the NotifyHandles field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetNotifyHandlesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyHandles, true
}

// SetNotifyHandles sets field value.
func (o *EventEmailAddressResponseDataAttributes) SetNotifyHandles(v []string) {
	o.NotifyHandles = v
}

// GetRevokedAt returns the RevokedAt field value if set, zero value otherwise.
func (o *EventEmailAddressResponseDataAttributes) GetRevokedAt() time.Time {
	if o == nil || o.RevokedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RevokedAt
}

// GetRevokedAtOk returns a tuple with the RevokedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetRevokedAtOk() (*time.Time, bool) {
	if o == nil || o.RevokedAt == nil {
		return nil, false
	}
	return o.RevokedAt, true
}

// HasRevokedAt returns a boolean if a field has been set.
func (o *EventEmailAddressResponseDataAttributes) HasRevokedAt() bool {
	return o != nil && o.RevokedAt != nil
}

// SetRevokedAt gets a reference to the given time.Time and assigns it to the RevokedAt field.
func (o *EventEmailAddressResponseDataAttributes) SetRevokedAt(v time.Time) {
	o.RevokedAt = &v
}

// GetTags returns the Tags field value.
func (o *EventEmailAddressResponseDataAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressResponseDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *EventEmailAddressResponseDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressResponseDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["email"] = o.Email
	toSerialize["format"] = o.Format
	if o.LastUsedAt != nil {
		if o.LastUsedAt.Nanosecond() == 0 {
			toSerialize["last_used_at"] = o.LastUsedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_used_at"] = o.LastUsedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["notify_handles"] = o.NotifyHandles
	if o.RevokedAt != nil {
		if o.RevokedAt.Nanosecond() == 0 {
			toSerialize["revoked_at"] = o.RevokedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["revoked_at"] = o.RevokedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventEmailAddressResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertType     *string    `json:"alert_type,omitempty"`
		CreatedAt     *time.Time `json:"created_at"`
		Description   *string    `json:"description,omitempty"`
		Email         *string    `json:"email"`
		Format        *string    `json:"format"`
		LastUsedAt    *time.Time `json:"last_used_at,omitempty"`
		NotifyHandles *[]string  `json:"notify_handles"`
		RevokedAt     *time.Time `json:"revoked_at,omitempty"`
		Tags          *[]string  `json:"tags"`
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
	if all.NotifyHandles == nil {
		return fmt.Errorf("required field notify_handles missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_type", "created_at", "description", "email", "format", "last_used_at", "notify_handles", "revoked_at", "tags"})
	} else {
		return err
	}
	o.AlertType = all.AlertType
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.Email = *all.Email
	o.Format = *all.Format
	o.LastUsedAt = all.LastUsedAt
	o.NotifyHandles = *all.NotifyHandles
	o.RevokedAt = all.RevokedAt
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
