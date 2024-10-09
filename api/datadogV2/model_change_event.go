// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEvent Object representing a change event.
type ChangeEvent struct {
	// An arbitrary string to use for aggregation. Limited to 100 characters.
	// If you specify a key, all events using that key are grouped together in the Event Stream.
	AggregationKey *string `json:"aggregation_key,omitempty"`
	// Object representing custom event attributes.
	Attributes ChangeEventCustomAttributes `json:"attributes"`
	// Event category to identify the type of event. Only the value `change` is supported.
	Category ChangeEventCategory `json:"category"`
	// The body of the event. Limited to 4000 characters.
	Message *string `json:"message,omitempty"`
	// A list of tags to apply to the event.
	// Refer to [Tags docs](https://docs.datadoghq.com/getting_started/tagging/).
	Tags []string `json:"tags,omitempty"`
	// Timestamp in which the event occurred. Must follow [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format.
	// For example `"2017-01-15T01:30:15.010000Z"`.
	// Defaults to now. Limited to values no older than 18 hours.
	Timestamp *string `json:"timestamp,omitempty"`
	// The event title. Limited to 500 characters.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeEvent instantiates a new ChangeEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEvent(attributes ChangeEventCustomAttributes, category ChangeEventCategory, title string) *ChangeEvent {
	this := ChangeEvent{}
	this.Attributes = attributes
	this.Category = category
	this.Title = title
	return &this
}

// NewChangeEventWithDefaults instantiates a new ChangeEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventWithDefaults() *ChangeEvent {
	this := ChangeEvent{}
	return &this
}

// GetAggregationKey returns the AggregationKey field value if set, zero value otherwise.
func (o *ChangeEvent) GetAggregationKey() string {
	if o == nil || o.AggregationKey == nil {
		var ret string
		return ret
	}
	return *o.AggregationKey
}

// GetAggregationKeyOk returns a tuple with the AggregationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEvent) GetAggregationKeyOk() (*string, bool) {
	if o == nil || o.AggregationKey == nil {
		return nil, false
	}
	return o.AggregationKey, true
}

// HasAggregationKey returns a boolean if a field has been set.
func (o *ChangeEvent) HasAggregationKey() bool {
	return o != nil && o.AggregationKey != nil
}

// SetAggregationKey gets a reference to the given string and assigns it to the AggregationKey field.
func (o *ChangeEvent) SetAggregationKey(v string) {
	o.AggregationKey = &v
}

// GetAttributes returns the Attributes field value.
func (o *ChangeEvent) GetAttributes() ChangeEventCustomAttributes {
	if o == nil {
		var ret ChangeEventCustomAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ChangeEvent) GetAttributesOk() (*ChangeEventCustomAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ChangeEvent) SetAttributes(v ChangeEventCustomAttributes) {
	o.Attributes = v
}

// GetCategory returns the Category field value.
func (o *ChangeEvent) GetCategory() ChangeEventCategory {
	if o == nil {
		var ret ChangeEventCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ChangeEvent) GetCategoryOk() (*ChangeEventCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *ChangeEvent) SetCategory(v ChangeEventCategory) {
	o.Category = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ChangeEvent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEvent) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ChangeEvent) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ChangeEvent) SetMessage(v string) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ChangeEvent) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEvent) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ChangeEvent) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ChangeEvent) SetTags(v []string) {
	o.Tags = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ChangeEvent) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEvent) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ChangeEvent) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ChangeEvent) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value.
func (o *ChangeEvent) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ChangeEvent) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ChangeEvent) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregationKey != nil {
		toSerialize["aggregation_key"] = o.AggregationKey
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["category"] = o.Category
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationKey *string                      `json:"aggregation_key,omitempty"`
		Attributes     *ChangeEventCustomAttributes `json:"attributes"`
		Category       *ChangeEventCategory         `json:"category"`
		Message        *string                      `json:"message,omitempty"`
		Tags           []string                     `json:"tags,omitempty"`
		Timestamp      *string                      `json:"timestamp,omitempty"`
		Title          *string                      `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation_key", "attributes", "category", "message", "tags", "timestamp", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AggregationKey = all.AggregationKey
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Message = all.Message
	o.Tags = all.Tags
	o.Timestamp = all.Timestamp
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
