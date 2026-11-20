// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTime20270101OneOf4464105952 Test-only generated schema for the dashboard versioning proof of concept.
type WidgetTime20270101OneOf4464105952 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetTime20270101OneOf4464105952 instantiates a new WidgetTime20270101OneOf4464105952 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetTime20270101OneOf4464105952() *WidgetTime20270101OneOf4464105952 {
	this := WidgetTime20270101OneOf4464105952{}
	return &this
}

// NewWidgetTime20270101OneOf4464105952WithDefaults instantiates a new WidgetTime20270101OneOf4464105952 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetTime20270101OneOf4464105952WithDefaults() *WidgetTime20270101OneOf4464105952 {
	this := WidgetTime20270101OneOf4464105952{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetTime20270101OneOf4464105952) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464105952) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetTime20270101OneOf4464105952) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetTime20270101OneOf4464105952) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetTime20270101OneOf4464105952) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetTime20270101OneOf4464105952) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.HideIncompleteCostData = all.HideIncompleteCostData

	return nil
}
