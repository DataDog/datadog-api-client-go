// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressUpdateAttributes Attributes for updating an event email address.
type EventEmailAddressUpdateAttributes struct {
	// The alert type of events generated from the email address.
	AlertType EventEmailAddressAlertType `json:"alert_type"`
	// A description of the event email address.
	Description datadog.NullableString `json:"description"`
	// A list of handles to notify when an email is received.
	NotifyHandles []string `json:"notify_handles"`
	// A list of tags to apply to events generated from the email address.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressUpdateAttributes instantiates a new EventEmailAddressUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressUpdateAttributes(alertType EventEmailAddressAlertType, description datadog.NullableString, notifyHandles []string, tags []string) *EventEmailAddressUpdateAttributes {
	this := EventEmailAddressUpdateAttributes{}
	this.AlertType = alertType
	this.Description = description
	this.NotifyHandles = notifyHandles
	this.Tags = tags
	return &this
}

// NewEventEmailAddressUpdateAttributesWithDefaults instantiates a new EventEmailAddressUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressUpdateAttributesWithDefaults() *EventEmailAddressUpdateAttributes {
	this := EventEmailAddressUpdateAttributes{}
	return &this
}

// GetAlertType returns the AlertType field value.
func (o *EventEmailAddressUpdateAttributes) GetAlertType() EventEmailAddressAlertType {
	if o == nil {
		var ret EventEmailAddressAlertType
		return ret
	}
	return o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressUpdateAttributes) GetAlertTypeOk() (*EventEmailAddressAlertType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertType, true
}

// SetAlertType sets field value.
func (o *EventEmailAddressUpdateAttributes) SetAlertType(v EventEmailAddressAlertType) {
	o.AlertType = v
}

// GetDescription returns the Description field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *EventEmailAddressUpdateAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventEmailAddressUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value.
func (o *EventEmailAddressUpdateAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetNotifyHandles returns the NotifyHandles field value.
func (o *EventEmailAddressUpdateAttributes) GetNotifyHandles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NotifyHandles
}

// GetNotifyHandlesOk returns a tuple with the NotifyHandles field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressUpdateAttributes) GetNotifyHandlesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyHandles, true
}

// SetNotifyHandles sets field value.
func (o *EventEmailAddressUpdateAttributes) SetNotifyHandles(v []string) {
	o.NotifyHandles = v
}

// GetTags returns the Tags field value.
func (o *EventEmailAddressUpdateAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *EventEmailAddressUpdateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["alert_type"] = o.AlertType
	toSerialize["description"] = o.Description.Get()
	toSerialize["notify_handles"] = o.NotifyHandles
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventEmailAddressUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertType     *EventEmailAddressAlertType `json:"alert_type"`
		Description   datadog.NullableString      `json:"description"`
		NotifyHandles *[]string                   `json:"notify_handles"`
		Tags          *[]string                   `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AlertType == nil {
		return fmt.Errorf("required field alert_type missing")
	}
	if !all.Description.IsSet() {
		return fmt.Errorf("required field description missing")
	}
	if all.NotifyHandles == nil {
		return fmt.Errorf("required field notify_handles missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_type", "description", "notify_handles", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AlertType.IsValid() {
		hasInvalidField = true
	} else {
		o.AlertType = *all.AlertType
	}
	o.Description = all.Description
	o.NotifyHandles = *all.NotifyHandles
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
