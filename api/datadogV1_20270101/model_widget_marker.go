// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetMarker Test-only generated schema for the dashboard versioning proof of concept.
type WidgetMarker struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayType *string `json:"display_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OnRightYaxis *bool `json:"on_right_yaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *string `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetMarker instantiates a new WidgetMarker object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetMarker(value string) *WidgetMarker {
	this := WidgetMarker{}
	this.Value = value
	return &this
}

// NewWidgetMarkerWithDefaults instantiates a new WidgetMarker object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetMarkerWithDefaults() *WidgetMarker {
	this := WidgetMarker{}
	return &this
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *WidgetMarker) GetDisplayType() string {
	if o == nil || o.DisplayType == nil {
		var ret string
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetMarker) GetDisplayTypeOk() (*string, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *WidgetMarker) HasDisplayType() bool {
	return o != nil && o.DisplayType != nil
}

// SetDisplayType gets a reference to the given string and assigns it to the DisplayType field.
func (o *WidgetMarker) SetDisplayType(v string) {
	o.DisplayType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WidgetMarker) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetMarker) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WidgetMarker) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WidgetMarker) SetLabel(v string) {
	o.Label = &v
}

// GetOnRightYaxis returns the OnRightYaxis field value if set, zero value otherwise.
func (o *WidgetMarker) GetOnRightYaxis() bool {
	if o == nil || o.OnRightYaxis == nil {
		var ret bool
		return ret
	}
	return *o.OnRightYaxis
}

// GetOnRightYaxisOk returns a tuple with the OnRightYaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetMarker) GetOnRightYaxisOk() (*bool, bool) {
	if o == nil || o.OnRightYaxis == nil {
		return nil, false
	}
	return o.OnRightYaxis, true
}

// HasOnRightYaxis returns a boolean if a field has been set.
func (o *WidgetMarker) HasOnRightYaxis() bool {
	return o != nil && o.OnRightYaxis != nil
}

// SetOnRightYaxis gets a reference to the given bool and assigns it to the OnRightYaxis field.
func (o *WidgetMarker) SetOnRightYaxis(v bool) {
	o.OnRightYaxis = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *WidgetMarker) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetMarker) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *WidgetMarker) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *WidgetMarker) SetTime(v string) {
	o.Time = &v
}

// GetValue returns the Value field value.
func (o *WidgetMarker) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WidgetMarker) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *WidgetMarker) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetMarker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayType != nil {
		toSerialize["display_type"] = o.DisplayType
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.OnRightYaxis != nil {
		toSerialize["on_right_yaxis"] = o.OnRightYaxis
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetMarker) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayType  *string `json:"display_type,omitempty"`
		Label        *string `json:"label,omitempty"`
		OnRightYaxis *bool   `json:"on_right_yaxis,omitempty"`
		Time         *string `json:"time,omitempty"`
		Value        *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	o.DisplayType = all.DisplayType
	o.Label = all.Label
	o.OnRightYaxis = all.OnRightYaxis
	o.Time = all.Time
	o.Value = *all.Value

	return nil
}
