// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventEmailAddressCreateAttributes Attributes for creating an event email address.
type EventEmailAddressCreateAttributes struct {
	// The alert type of events generated from the email address.
	AlertType *EventEmailAddressAlertType `json:"alert_type,omitempty"`
	// A description of the event email address.
	Description *string `json:"description,omitempty"`
	// The format of events ingested through the email address.
	Format EventEmailAddressFormat `json:"format"`
	// A list of handles to notify when an email is received.
	NotifyHandles []string `json:"notify_handles,omitempty"`
	// A list of tags to apply to events generated from the email address.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventEmailAddressCreateAttributes instantiates a new EventEmailAddressCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventEmailAddressCreateAttributes(format EventEmailAddressFormat) *EventEmailAddressCreateAttributes {
	this := EventEmailAddressCreateAttributes{}
	this.Format = format
	return &this
}

// NewEventEmailAddressCreateAttributesWithDefaults instantiates a new EventEmailAddressCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventEmailAddressCreateAttributesWithDefaults() *EventEmailAddressCreateAttributes {
	this := EventEmailAddressCreateAttributes{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *EventEmailAddressCreateAttributes) GetAlertType() EventEmailAddressAlertType {
	if o == nil || o.AlertType == nil {
		var ret EventEmailAddressAlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressCreateAttributes) GetAlertTypeOk() (*EventEmailAddressAlertType, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *EventEmailAddressCreateAttributes) HasAlertType() bool {
	return o != nil && o.AlertType != nil
}

// SetAlertType gets a reference to the given EventEmailAddressAlertType and assigns it to the AlertType field.
func (o *EventEmailAddressCreateAttributes) SetAlertType(v EventEmailAddressAlertType) {
	o.AlertType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventEmailAddressCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventEmailAddressCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventEmailAddressCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFormat returns the Format field value.
func (o *EventEmailAddressCreateAttributes) GetFormat() EventEmailAddressFormat {
	if o == nil {
		var ret EventEmailAddressFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *EventEmailAddressCreateAttributes) GetFormatOk() (*EventEmailAddressFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *EventEmailAddressCreateAttributes) SetFormat(v EventEmailAddressFormat) {
	o.Format = v
}

// GetNotifyHandles returns the NotifyHandles field value if set, zero value otherwise.
func (o *EventEmailAddressCreateAttributes) GetNotifyHandles() []string {
	if o == nil || o.NotifyHandles == nil {
		var ret []string
		return ret
	}
	return o.NotifyHandles
}

// GetNotifyHandlesOk returns a tuple with the NotifyHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressCreateAttributes) GetNotifyHandlesOk() (*[]string, bool) {
	if o == nil || o.NotifyHandles == nil {
		return nil, false
	}
	return &o.NotifyHandles, true
}

// HasNotifyHandles returns a boolean if a field has been set.
func (o *EventEmailAddressCreateAttributes) HasNotifyHandles() bool {
	return o != nil && o.NotifyHandles != nil
}

// SetNotifyHandles gets a reference to the given []string and assigns it to the NotifyHandles field.
func (o *EventEmailAddressCreateAttributes) SetNotifyHandles(v []string) {
	o.NotifyHandles = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventEmailAddressCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventEmailAddressCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventEmailAddressCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EventEmailAddressCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventEmailAddressCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AlertType != nil {
		toSerialize["alert_type"] = o.AlertType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["format"] = o.Format
	if o.NotifyHandles != nil {
		toSerialize["notify_handles"] = o.NotifyHandles
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
func (o *EventEmailAddressCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertType     *EventEmailAddressAlertType `json:"alert_type,omitempty"`
		Description   *string                     `json:"description,omitempty"`
		Format        *EventEmailAddressFormat    `json:"format"`
		NotifyHandles []string                    `json:"notify_handles,omitempty"`
		Tags          []string                    `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_type", "description", "format", "notify_handles", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AlertType != nil && !all.AlertType.IsValid() {
		hasInvalidField = true
	} else {
		o.AlertType = all.AlertType
	}
	o.Description = all.Description
	if !all.Format.IsValid() {
		hasInvalidField = true
	} else {
		o.Format = *all.Format
	}
	o.NotifyHandles = all.NotifyHandles
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
