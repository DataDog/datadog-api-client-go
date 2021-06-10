/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// EventCreateRequest Object representing an event.
type EventCreateRequest struct {
	// An arbitrary string to use for aggregation. Limited to 100 characters. If you specify a key, all events using that key are grouped together in the Event Stream.
	AggregationKey *string         `json:"aggregation_key,omitempty"`
	AlertType      *EventAlertType `json:"alert_type,omitempty"`
	// POSIX timestamp of the event. Must be sent as an integer (i.e. no quotes). Limited to events no older than 7 days.
	DateHappened *int64 `json:"date_happened,omitempty"`
	// A device name.
	DeviceName *string `json:"device_name,omitempty"`
	// Host name to associate with the event. Any tags associated with the host are also applied to this event.
	Host *string `json:"host,omitempty"`
	// Integer ID of the event.
	Id *int64 `json:"id,omitempty"`
	// Payload of the event.
	Payload  *string        `json:"payload,omitempty"`
	Priority *EventPriority `json:"priority,omitempty"`
	// ID of the parent event. Must be sent as an integer (i.e. no quotes).
	RelatedEventId *int64 `json:"related_event_id,omitempty"`
	// The type of event being posted. Option examples include nagios, hudson, jenkins, my_apps, chef, puppet, git, bitbucket, etc. A complete list of source attribute values [available here](https://docs.datadoghq.com/integrations/faq/list-of-api-source-attribute-value).
	SourceTypeName *string `json:"source_type_name,omitempty"`
	// A list of tags to apply to the event.
	Tags *[]string `json:"tags,omitempty"`
	// The body of the event. Limited to 4000 characters. The text supports markdown. To use markdown in the event text, start the text block with `%%% \\n` and end the text block with `\\n %%%`. Use `msg_text` with the Datadog Ruby library.
	Text string `json:"text"`
	// The event title. Limited to 100 characters. Use `msg_title` with the Datadog Ruby library.
	Title string `json:"title"`
	// URL of the event.
	Url *string `json:"url,omitempty"`
}

// NewEventCreateRequest instantiates a new EventCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCreateRequest(text string, title string) *EventCreateRequest {
	this := EventCreateRequest{}
	this.Text = text
	this.Title = title
	return &this
}

// NewEventCreateRequestWithDefaults instantiates a new EventCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCreateRequestWithDefaults() *EventCreateRequest {
	this := EventCreateRequest{}
	return &this
}

// GetAggregationKey returns the AggregationKey field value if set, zero value otherwise.
func (o *EventCreateRequest) GetAggregationKey() string {
	if o == nil || o.AggregationKey == nil {
		var ret string
		return ret
	}
	return *o.AggregationKey
}

// GetAggregationKeyOk returns a tuple with the AggregationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetAggregationKeyOk() (*string, bool) {
	if o == nil || o.AggregationKey == nil {
		return nil, false
	}
	return o.AggregationKey, true
}

// HasAggregationKey returns a boolean if a field has been set.
func (o *EventCreateRequest) HasAggregationKey() bool {
	if o != nil && o.AggregationKey != nil {
		return true
	}

	return false
}

// SetAggregationKey gets a reference to the given string and assigns it to the AggregationKey field.
func (o *EventCreateRequest) SetAggregationKey(v string) {
	o.AggregationKey = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *EventCreateRequest) GetAlertType() EventAlertType {
	if o == nil || o.AlertType == nil {
		var ret EventAlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetAlertTypeOk() (*EventAlertType, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *EventCreateRequest) HasAlertType() bool {
	if o != nil && o.AlertType != nil {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given EventAlertType and assigns it to the AlertType field.
func (o *EventCreateRequest) SetAlertType(v EventAlertType) {
	o.AlertType = &v
}

// GetDateHappened returns the DateHappened field value if set, zero value otherwise.
func (o *EventCreateRequest) GetDateHappened() int64 {
	if o == nil || o.DateHappened == nil {
		var ret int64
		return ret
	}
	return *o.DateHappened
}

// GetDateHappenedOk returns a tuple with the DateHappened field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetDateHappenedOk() (*int64, bool) {
	if o == nil || o.DateHappened == nil {
		return nil, false
	}
	return o.DateHappened, true
}

// HasDateHappened returns a boolean if a field has been set.
func (o *EventCreateRequest) HasDateHappened() bool {
	if o != nil && o.DateHappened != nil {
		return true
	}

	return false
}

// SetDateHappened gets a reference to the given int64 and assigns it to the DateHappened field.
func (o *EventCreateRequest) SetDateHappened(v int64) {
	o.DateHappened = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *EventCreateRequest) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *EventCreateRequest) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *EventCreateRequest) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EventCreateRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EventCreateRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EventCreateRequest) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventCreateRequest) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventCreateRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *EventCreateRequest) SetId(v int64) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *EventCreateRequest) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *EventCreateRequest) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *EventCreateRequest) SetPayload(v string) {
	o.Payload = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *EventCreateRequest) GetPriority() EventPriority {
	if o == nil || o.Priority == nil {
		var ret EventPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetPriorityOk() (*EventPriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *EventCreateRequest) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given EventPriority and assigns it to the Priority field.
func (o *EventCreateRequest) SetPriority(v EventPriority) {
	o.Priority = &v
}

// GetRelatedEventId returns the RelatedEventId field value if set, zero value otherwise.
func (o *EventCreateRequest) GetRelatedEventId() int64 {
	if o == nil || o.RelatedEventId == nil {
		var ret int64
		return ret
	}
	return *o.RelatedEventId
}

// GetRelatedEventIdOk returns a tuple with the RelatedEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetRelatedEventIdOk() (*int64, bool) {
	if o == nil || o.RelatedEventId == nil {
		return nil, false
	}
	return o.RelatedEventId, true
}

// HasRelatedEventId returns a boolean if a field has been set.
func (o *EventCreateRequest) HasRelatedEventId() bool {
	if o != nil && o.RelatedEventId != nil {
		return true
	}

	return false
}

// SetRelatedEventId gets a reference to the given int64 and assigns it to the RelatedEventId field.
func (o *EventCreateRequest) SetRelatedEventId(v int64) {
	o.RelatedEventId = &v
}

// GetSourceTypeName returns the SourceTypeName field value if set, zero value otherwise.
func (o *EventCreateRequest) GetSourceTypeName() string {
	if o == nil || o.SourceTypeName == nil {
		var ret string
		return ret
	}
	return *o.SourceTypeName
}

// GetSourceTypeNameOk returns a tuple with the SourceTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetSourceTypeNameOk() (*string, bool) {
	if o == nil || o.SourceTypeName == nil {
		return nil, false
	}
	return o.SourceTypeName, true
}

// HasSourceTypeName returns a boolean if a field has been set.
func (o *EventCreateRequest) HasSourceTypeName() bool {
	if o != nil && o.SourceTypeName != nil {
		return true
	}

	return false
}

// SetSourceTypeName gets a reference to the given string and assigns it to the SourceTypeName field.
func (o *EventCreateRequest) SetSourceTypeName(v string) {
	o.SourceTypeName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventCreateRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventCreateRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EventCreateRequest) SetTags(v []string) {
	o.Tags = &v
}

// GetText returns the Text field value
func (o *EventCreateRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *EventCreateRequest) SetText(v string) {
	o.Text = v
}

// GetTitle returns the Title field value
func (o *EventCreateRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *EventCreateRequest) SetTitle(v string) {
	o.Title = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EventCreateRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EventCreateRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EventCreateRequest) SetUrl(v string) {
	o.Url = &v
}

func (o EventCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AggregationKey != nil {
		toSerialize["aggregation_key"] = o.AggregationKey
	}
	if o.AlertType != nil {
		toSerialize["alert_type"] = o.AlertType
	}
	if o.DateHappened != nil {
		toSerialize["date_happened"] = o.DateHappened
	}
	if o.DeviceName != nil {
		toSerialize["device_name"] = o.DeviceName
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.RelatedEventId != nil {
		toSerialize["related_event_id"] = o.RelatedEventId
	}
	if o.SourceTypeName != nil {
		toSerialize["source_type_name"] = o.SourceTypeName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["text"] = o.Text
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

func (o *EventCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Text  *string `json:"text"`
		Title *string `json:"title"`
	}{}
	all := struct {
		AggregationKey *string         `json:"aggregation_key,omitempty"`
		AlertType      *EventAlertType `json:"alert_type,omitempty"`
		DateHappened   *int64          `json:"date_happened,omitempty"`
		DeviceName     *string         `json:"device_name,omitempty"`
		Host           *string         `json:"host,omitempty"`
		Id             *int64          `json:"id,omitempty"`
		Payload        *string         `json:"payload,omitempty"`
		Priority       *EventPriority  `json:"priority,omitempty"`
		RelatedEventId *int64          `json:"related_event_id,omitempty"`
		SourceTypeName *string         `json:"source_type_name,omitempty"`
		Tags           *[]string       `json:"tags,omitempty"`
		Text           string          `json:"text"`
		Title          string          `json:"title"`
		Url            *string         `json:"url,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Text == nil {
		return fmt.Errorf("Required field text missing")
	}
	if required.Title == nil {
		return fmt.Errorf("Required field title missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.AggregationKey = all.AggregationKey
	o.AlertType = all.AlertType
	o.DateHappened = all.DateHappened
	o.DeviceName = all.DeviceName
	o.Host = all.Host
	o.Id = all.Id
	o.Payload = all.Payload
	o.Priority = all.Priority
	o.RelatedEventId = all.RelatedEventId
	o.SourceTypeName = all.SourceTypeName
	o.Tags = all.Tags
	o.Text = all.Text
	o.Title = all.Title
	o.Url = all.Url
	return nil
}

type NullableEventCreateRequest struct {
	value *EventCreateRequest
	isSet bool
}

func (v NullableEventCreateRequest) Get() *EventCreateRequest {
	return v.value
}

func (v *NullableEventCreateRequest) Set(val *EventCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCreateRequest(val *EventCreateRequest) *NullableEventCreateRequest {
	return &NullableEventCreateRequest{value: val, isSet: true}
}

func (v NullableEventCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
