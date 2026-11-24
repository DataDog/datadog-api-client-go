// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMExclusionFilterUpdateAttributes The attributes of an exclusion filter that can be updated.
// For the built-in Error Tracking exclusion filter, only `enabled` can be set;
// `name`, `event_type`, and `query` must be omitted.
type RUMExclusionFilterUpdateAttributes struct {
	// Whether the exclusion filter is active.
	Enabled *bool `json:"enabled,omitempty"`
	// The type of RUM events to filter on.
	EventType *RUMExclusionFilterEventType `json:"event_type,omitempty"`
	// The name of the exclusion filter.
	Name *string `json:"name,omitempty"`
	// Additional query used to further restrict which RUM events are excluded.
	// Combined with `event_type` when both are provided.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMExclusionFilterUpdateAttributes instantiates a new RUMExclusionFilterUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMExclusionFilterUpdateAttributes() *RUMExclusionFilterUpdateAttributes {
	this := RUMExclusionFilterUpdateAttributes{}
	return &this
}

// NewRUMExclusionFilterUpdateAttributesWithDefaults instantiates a new RUMExclusionFilterUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMExclusionFilterUpdateAttributesWithDefaults() *RUMExclusionFilterUpdateAttributes {
	this := RUMExclusionFilterUpdateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RUMExclusionFilterUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RUMExclusionFilterUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RUMExclusionFilterUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RUMExclusionFilterUpdateAttributes) GetEventType() RUMExclusionFilterEventType {
	if o == nil || o.EventType == nil {
		var ret RUMExclusionFilterEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterUpdateAttributes) GetEventTypeOk() (*RUMExclusionFilterEventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RUMExclusionFilterUpdateAttributes) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given RUMExclusionFilterEventType and assigns it to the EventType field.
func (o *RUMExclusionFilterUpdateAttributes) SetEventType(v RUMExclusionFilterEventType) {
	o.EventType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RUMExclusionFilterUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RUMExclusionFilterUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RUMExclusionFilterUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RUMExclusionFilterUpdateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterUpdateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RUMExclusionFilterUpdateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RUMExclusionFilterUpdateAttributes) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMExclusionFilterUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EventType != nil {
		toSerialize["event_type"] = o.EventType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMExclusionFilterUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                        `json:"enabled,omitempty"`
		EventType *RUMExclusionFilterEventType `json:"event_type,omitempty"`
		Name      *string                      `json:"name,omitempty"`
		Query     *string                      `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "event_type", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.EventType != nil && !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = all.EventType
	}
	o.Name = all.Name
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
