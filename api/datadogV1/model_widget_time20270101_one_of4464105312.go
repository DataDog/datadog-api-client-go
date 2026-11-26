// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTime20270101OneOf4464105312 Test-only generated schema for the dashboard versioning proof of concept.
type WidgetTime20270101OneOf4464105312 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LiveSpan LegacyLiveSpan_20270101 `json:"live_span"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetTime20270101OneOf4464105312 instantiates a new WidgetTime20270101OneOf4464105312 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetTime20270101OneOf4464105312(liveSpan LegacyLiveSpan_20270101) *WidgetTime20270101OneOf4464105312 {
	this := WidgetTime20270101OneOf4464105312{}
	this.LiveSpan = liveSpan
	return &this
}

// NewWidgetTime20270101OneOf4464105312WithDefaults instantiates a new WidgetTime20270101OneOf4464105312 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetTime20270101OneOf4464105312WithDefaults() *WidgetTime20270101OneOf4464105312 {
	this := WidgetTime20270101OneOf4464105312{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetTime20270101OneOf4464105312) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464105312) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetTime20270101OneOf4464105312) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetTime20270101OneOf4464105312) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetLiveSpan returns the LiveSpan field value.
func (o *WidgetTime20270101OneOf4464105312) GetLiveSpan() LegacyLiveSpan_20270101 {
	if o == nil {
		var ret LegacyLiveSpan_20270101
		return ret
	}
	return o.LiveSpan
}

// GetLiveSpanOk returns a tuple with the LiveSpan field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464105312) GetLiveSpanOk() (*LegacyLiveSpan_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiveSpan, true
}

// SetLiveSpan sets field value.
func (o *WidgetTime20270101OneOf4464105312) SetLiveSpan(v LegacyLiveSpan_20270101) {
	o.LiveSpan = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetTime20270101OneOf4464105312) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	toSerialize["live_span"] = o.LiveSpan
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetTime20270101OneOf4464105312) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool                    `json:"hide_incomplete_cost_data,omitempty"`
		LiveSpan               *LegacyLiveSpan_20270101 `json:"live_span"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LiveSpan == nil {
		return fmt.Errorf("required field live_span missing")
	}
	o.HideIncompleteCostData = all.HideIncompleteCostData
	o.LiveSpan = *all.LiveSpan

	return nil
}
