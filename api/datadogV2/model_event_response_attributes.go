// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventResponseAttributes The object description of an event response attribute.
type EventResponseAttributes struct {
	// Object description of attributes from your event.
	Attributes *EventAttributes `json:"attributes,omitempty"`
	// The host associated with the event.
	Host *string `json:"host,omitempty"`
	// The message of the event.
	Message *string `json:"message,omitempty"`
	// The service associated with the event.
	Service *string `json:"service,omitempty"`
	// The event status. Legacy events can use `failure`, `error`, `warning`,
	// `info`, `success`, `user_update`, `recommendation`, or `snapshot`.
	// Alert events can use `error`, `warn`, or `ok`.
	Status *EventStatusType `json:"status,omitempty"`
	// An array of tags associated with the event.
	Tags []string `json:"tags,omitempty"`
	// The timestamp of the event.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventResponseAttributes instantiates a new EventResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventResponseAttributes() *EventResponseAttributes {
	this := EventResponseAttributes{}
	return &this
}

// NewEventResponseAttributesWithDefaults instantiates a new EventResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventResponseAttributesWithDefaults() *EventResponseAttributes {
	this := EventResponseAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EventResponseAttributes) GetAttributes() EventAttributes {
	if o == nil || o.Attributes == nil {
		var ret EventAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAttributes) GetAttributesOk() (*EventAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EventResponseAttributes) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given EventAttributes and assigns it to the Attributes field.
func (o *EventResponseAttributes) SetAttributes(v EventAttributes) {
	o.Attributes = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EventResponseAttributes) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAttributes) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EventResponseAttributes) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EventResponseAttributes) SetHost(v string) {
	o.Host = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventResponseAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventResponseAttributes) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventResponseAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *EventResponseAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *EventResponseAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *EventResponseAttributes) SetService(v string) {
	o.Service = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventResponseAttributes) GetStatus() EventStatusType {
	if o == nil || o.Status == nil {
		var ret EventStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAttributes) GetStatusOk() (*EventStatusType, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventResponseAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given EventStatusType and assigns it to the Status field.
func (o *EventResponseAttributes) SetStatus(v EventStatusType) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventResponseAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventResponseAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EventResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *EventResponseAttributes) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventResponseAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *EventResponseAttributes) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *EventResponseAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		if o.Timestamp.Nanosecond() == 0 {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *EventAttributes `json:"attributes,omitempty"`
		Host       *string          `json:"host,omitempty"`
		Message    *string          `json:"message,omitempty"`
		Service    *string          `json:"service,omitempty"`
		Status     *EventStatusType `json:"status,omitempty"`
		Tags       []string         `json:"tags,omitempty"`
		Timestamp  *time.Time       `json:"timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "host", "message", "service", "status", "tags", "timestamp"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Host = all.Host
	o.Message = all.Message
	o.Service = all.Service
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Tags = all.Tags
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
