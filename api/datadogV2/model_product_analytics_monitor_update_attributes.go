// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorUpdateAttributes Attributes to update on a Product Analytics monitor. Fields that are not set are left unchanged.
// When set, `options` or `tags` replaces the stored value.
type ProductAnalyticsMonitorUpdateAttributes struct {
	// Notification message sent when the monitor alerts.
	Message *string `json:"message,omitempty"`
	// Name of the monitor.
	Name *string `json:"name,omitempty"`
	// Options of a Product Analytics monitor.
	Options *ProductAnalyticsMonitorOptions `json:"options,omitempty"`
	// Priority of the monitor, from 1 (highest) to 5 (lowest).
	Priority *int32 `json:"priority,omitempty"`
	// Monitor query.
	Query *string `json:"query,omitempty"`
	// Monitor tags as key and value pairs. Each pair is stored as a `key:value` tag.
	// `chart_id` and `dashboard_id` link the monitor to the chart it was created from. `dashboard_id` requires
	// `chart_id`.
	Tags map[string]string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorUpdateAttributes instantiates a new ProductAnalyticsMonitorUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorUpdateAttributes() *ProductAnalyticsMonitorUpdateAttributes {
	this := ProductAnalyticsMonitorUpdateAttributes{}
	return &this
}

// NewProductAnalyticsMonitorUpdateAttributesWithDefaults instantiates a new ProductAnalyticsMonitorUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorUpdateAttributesWithDefaults() *ProductAnalyticsMonitorUpdateAttributes {
	this := ProductAnalyticsMonitorUpdateAttributes{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProductAnalyticsMonitorUpdateAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsMonitorUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetOptions() ProductAnalyticsMonitorOptions {
	if o == nil || o.Options == nil {
		var ret ProductAnalyticsMonitorOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetOptionsOk() (*ProductAnalyticsMonitorOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given ProductAnalyticsMonitorOptions and assigns it to the Options field.
func (o *ProductAnalyticsMonitorUpdateAttributes) SetOptions(v ProductAnalyticsMonitorOptions) {
	o.Options = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ProductAnalyticsMonitorUpdateAttributes) SetPriority(v int32) {
	o.Priority = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ProductAnalyticsMonitorUpdateAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ProductAnalyticsMonitorUpdateAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
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
func (o *ProductAnalyticsMonitorUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message  *string                         `json:"message,omitempty"`
		Name     *string                         `json:"name,omitempty"`
		Options  *ProductAnalyticsMonitorOptions `json:"options,omitempty"`
		Priority *int32                          `json:"priority,omitempty"`
		Query    *string                         `json:"query,omitempty"`
		Tags     map[string]string               `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "name", "options", "priority", "query", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Message = all.Message
	o.Name = all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	o.Priority = all.Priority
	o.Query = all.Query
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
