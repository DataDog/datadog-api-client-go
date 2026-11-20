// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMExclusionFilterCreateAttributes The attributes of an exclusion filter to create.
type RUMExclusionFilterCreateAttributes struct {
	// Whether the exclusion filter is active. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// The type of RUM events to filter on.
	EventType *RUMExclusionFilterEventType `json:"event_type,omitempty"`
	// The name of the exclusion filter.
	Name string `json:"name"`
	// Additional query used to further restrict which RUM events are excluded.
	// Combined with `event_type` when both are provided.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMExclusionFilterCreateAttributes instantiates a new RUMExclusionFilterCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMExclusionFilterCreateAttributes(name string) *RUMExclusionFilterCreateAttributes {
	this := RUMExclusionFilterCreateAttributes{}
	this.Name = name
	return &this
}

// NewRUMExclusionFilterCreateAttributesWithDefaults instantiates a new RUMExclusionFilterCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMExclusionFilterCreateAttributesWithDefaults() *RUMExclusionFilterCreateAttributes {
	this := RUMExclusionFilterCreateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RUMExclusionFilterCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RUMExclusionFilterCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RUMExclusionFilterCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *RUMExclusionFilterCreateAttributes) GetEventType() RUMExclusionFilterEventType {
	if o == nil || o.EventType == nil {
		var ret RUMExclusionFilterEventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterCreateAttributes) GetEventTypeOk() (*RUMExclusionFilterEventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *RUMExclusionFilterCreateAttributes) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given RUMExclusionFilterEventType and assigns it to the EventType field.
func (o *RUMExclusionFilterCreateAttributes) SetEventType(v RUMExclusionFilterEventType) {
	o.EventType = &v
}

// GetName returns the Name field value.
func (o *RUMExclusionFilterCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMExclusionFilterCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *RUMExclusionFilterCreateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMExclusionFilterCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *RUMExclusionFilterCreateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *RUMExclusionFilterCreateAttributes) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMExclusionFilterCreateAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMExclusionFilterCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                        `json:"enabled,omitempty"`
		EventType *RUMExclusionFilterEventType `json:"event_type,omitempty"`
		Name      *string                      `json:"name"`
		Query     *string                      `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
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
	o.Name = *all.Name
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
