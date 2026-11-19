// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorCreateAttributes Attributes of the Product Analytics monitor to create.
type ProductAnalyticsMonitorCreateAttributes struct {
	// Notification message sent when the monitor alerts. It can include `@` notification handles.
	Message string `json:"message"`
	// Name of the monitor.
	Name string `json:"name"`
	// Options of a Product Analytics monitor.
	Options *ProductAnalyticsMonitorOptions `json:"options,omitempty"`
	// Priority of the monitor, from 1 (highest) to 5 (lowest).
	Priority *int32 `json:"priority,omitempty"`
	// Monitor query, for example `formula("query").rollup("sum").last("1h") < 1`.
	// Define the variables it references in `options.variables`.
	Query string `json:"query"`
	// Monitor tags as key and value pairs. Each pair is stored as a `key:value` tag.
	// `chart_id` and `dashboard_id` link the monitor to the chart it was created from. `dashboard_id` requires
	// `chart_id`.
	Tags map[string]string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorCreateAttributes instantiates a new ProductAnalyticsMonitorCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorCreateAttributes(message string, name string, query string) *ProductAnalyticsMonitorCreateAttributes {
	this := ProductAnalyticsMonitorCreateAttributes{}
	this.Message = message
	this.Name = name
	this.Query = query
	return &this
}

// NewProductAnalyticsMonitorCreateAttributesWithDefaults instantiates a new ProductAnalyticsMonitorCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorCreateAttributesWithDefaults() *ProductAnalyticsMonitorCreateAttributes {
	this := ProductAnalyticsMonitorCreateAttributes{}
	return &this
}

// GetMessage returns the Message field value.
func (o *ProductAnalyticsMonitorCreateAttributes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *ProductAnalyticsMonitorCreateAttributes) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value.
func (o *ProductAnalyticsMonitorCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsMonitorCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorCreateAttributes) GetOptions() ProductAnalyticsMonitorOptions {
	if o == nil || o.Options == nil {
		var ret ProductAnalyticsMonitorOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) GetOptionsOk() (*ProductAnalyticsMonitorOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given ProductAnalyticsMonitorOptions and assigns it to the Options field.
func (o *ProductAnalyticsMonitorCreateAttributes) SetOptions(v ProductAnalyticsMonitorOptions) {
	o.Options = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorCreateAttributes) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ProductAnalyticsMonitorCreateAttributes) SetPriority(v int32) {
	o.Priority = &v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsMonitorCreateAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsMonitorCreateAttributes) SetQuery(v string) {
	o.Query = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorCreateAttributes) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ProductAnalyticsMonitorCreateAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["query"] = o.Query
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsMonitorCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message  *string                         `json:"message"`
		Name     *string                         `json:"name"`
		Options  *ProductAnalyticsMonitorOptions `json:"options,omitempty"`
		Priority *int32                          `json:"priority,omitempty"`
		Query    *string                         `json:"query"`
		Tags     map[string]string               `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "name", "options", "priority", "query", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Message = *all.Message
	o.Name = *all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	o.Priority = all.Priority
	o.Query = *all.Query
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
