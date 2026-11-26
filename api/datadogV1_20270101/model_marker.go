// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Marker Test-only generated schema for the dashboard versioning proof of concept.
type Marker struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DisplayType *string `json:"display_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	FromTs *string `json:"from_ts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsHovered *bool `json:"is_hovered,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Max *float64 `json:"max,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Min *float64 `json:"min,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OnRightYaxis *bool `json:"on_right_yaxis,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Time *string `json:"time,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ToTs *string `json:"to_ts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Ts *string `json:"ts,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type *string `json:"type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Val *float64 `json:"val,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMarker instantiates a new Marker object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMarker() *Marker {
	this := Marker{}
	return &this
}

// NewMarkerWithDefaults instantiates a new Marker object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMarkerWithDefaults() *Marker {
	this := Marker{}
	return &this
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *Marker) GetDisplayType() string {
	if o == nil || o.DisplayType == nil {
		var ret string
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetDisplayTypeOk() (*string, bool) {
	if o == nil || o.DisplayType == nil {
		return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *Marker) HasDisplayType() bool {
	return o != nil && o.DisplayType != nil
}

// SetDisplayType gets a reference to the given string and assigns it to the DisplayType field.
func (o *Marker) SetDisplayType(v string) {
	o.DisplayType = &v
}

// GetFromTs returns the FromTs field value if set, zero value otherwise.
func (o *Marker) GetFromTs() string {
	if o == nil || o.FromTs == nil {
		var ret string
		return ret
	}
	return *o.FromTs
}

// GetFromTsOk returns a tuple with the FromTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetFromTsOk() (*string, bool) {
	if o == nil || o.FromTs == nil {
		return nil, false
	}
	return o.FromTs, true
}

// HasFromTs returns a boolean if a field has been set.
func (o *Marker) HasFromTs() bool {
	return o != nil && o.FromTs != nil
}

// SetFromTs gets a reference to the given string and assigns it to the FromTs field.
func (o *Marker) SetFromTs(v string) {
	o.FromTs = &v
}

// GetIsHovered returns the IsHovered field value if set, zero value otherwise.
func (o *Marker) GetIsHovered() bool {
	if o == nil || o.IsHovered == nil {
		var ret bool
		return ret
	}
	return *o.IsHovered
}

// GetIsHoveredOk returns a tuple with the IsHovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetIsHoveredOk() (*bool, bool) {
	if o == nil || o.IsHovered == nil {
		return nil, false
	}
	return o.IsHovered, true
}

// HasIsHovered returns a boolean if a field has been set.
func (o *Marker) HasIsHovered() bool {
	return o != nil && o.IsHovered != nil
}

// SetIsHovered gets a reference to the given bool and assigns it to the IsHovered field.
func (o *Marker) SetIsHovered(v bool) {
	o.IsHovered = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Marker) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Marker) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Marker) SetLabel(v string) {
	o.Label = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Marker) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Marker) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *Marker) SetMax(v float64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Marker) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Marker) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *Marker) SetMin(v float64) {
	o.Min = &v
}

// GetOnRightYaxis returns the OnRightYaxis field value if set, zero value otherwise.
func (o *Marker) GetOnRightYaxis() bool {
	if o == nil || o.OnRightYaxis == nil {
		var ret bool
		return ret
	}
	return *o.OnRightYaxis
}

// GetOnRightYaxisOk returns a tuple with the OnRightYaxis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetOnRightYaxisOk() (*bool, bool) {
	if o == nil || o.OnRightYaxis == nil {
		return nil, false
	}
	return o.OnRightYaxis, true
}

// HasOnRightYaxis returns a boolean if a field has been set.
func (o *Marker) HasOnRightYaxis() bool {
	return o != nil && o.OnRightYaxis != nil
}

// SetOnRightYaxis gets a reference to the given bool and assigns it to the OnRightYaxis field.
func (o *Marker) SetOnRightYaxis(v bool) {
	o.OnRightYaxis = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Marker) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Marker) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *Marker) SetTime(v string) {
	o.Time = &v
}

// GetToTs returns the ToTs field value if set, zero value otherwise.
func (o *Marker) GetToTs() string {
	if o == nil || o.ToTs == nil {
		var ret string
		return ret
	}
	return *o.ToTs
}

// GetToTsOk returns a tuple with the ToTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetToTsOk() (*string, bool) {
	if o == nil || o.ToTs == nil {
		return nil, false
	}
	return o.ToTs, true
}

// HasToTs returns a boolean if a field has been set.
func (o *Marker) HasToTs() bool {
	return o != nil && o.ToTs != nil
}

// SetToTs gets a reference to the given string and assigns it to the ToTs field.
func (o *Marker) SetToTs(v string) {
	o.ToTs = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *Marker) GetTs() string {
	if o == nil || o.Ts == nil {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetTsOk() (*string, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *Marker) HasTs() bool {
	return o != nil && o.Ts != nil
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *Marker) SetTs(v string) {
	o.Ts = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Marker) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Marker) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Marker) SetType(v string) {
	o.Type = &v
}

// GetVal returns the Val field value if set, zero value otherwise.
func (o *Marker) GetVal() float64 {
	if o == nil || o.Val == nil {
		var ret float64
		return ret
	}
	return *o.Val
}

// GetValOk returns a tuple with the Val field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetValOk() (*float64, bool) {
	if o == nil || o.Val == nil {
		return nil, false
	}
	return o.Val, true
}

// HasVal returns a boolean if a field has been set.
func (o *Marker) HasVal() bool {
	return o != nil && o.Val != nil
}

// SetVal gets a reference to the given float64 and assigns it to the Val field.
func (o *Marker) SetVal(v float64) {
	o.Val = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Marker) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Marker) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Marker) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Marker) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Marker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayType != nil {
		toSerialize["display_type"] = o.DisplayType
	}
	if o.FromTs != nil {
		toSerialize["from_ts"] = o.FromTs
	}
	if o.IsHovered != nil {
		toSerialize["is_hovered"] = o.IsHovered
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.OnRightYaxis != nil {
		toSerialize["on_right_yaxis"] = o.OnRightYaxis
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.ToTs != nil {
		toSerialize["to_ts"] = o.ToTs
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Val != nil {
		toSerialize["val"] = o.Val
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Marker) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayType  *string  `json:"display_type,omitempty"`
		FromTs       *string  `json:"from_ts,omitempty"`
		IsHovered    *bool    `json:"is_hovered,omitempty"`
		Label        *string  `json:"label,omitempty"`
		Max          *float64 `json:"max,omitempty"`
		Min          *float64 `json:"min,omitempty"`
		OnRightYaxis *bool    `json:"on_right_yaxis,omitempty"`
		Time         *string  `json:"time,omitempty"`
		ToTs         *string  `json:"to_ts,omitempty"`
		Ts           *string  `json:"ts,omitempty"`
		Type         *string  `json:"type,omitempty"`
		Val          *float64 `json:"val,omitempty"`
		Value        *string  `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.DisplayType = all.DisplayType
	o.FromTs = all.FromTs
	o.IsHovered = all.IsHovered
	o.Label = all.Label
	o.Max = all.Max
	o.Min = all.Min
	o.OnRightYaxis = all.OnRightYaxis
	o.Time = all.Time
	o.ToTs = all.ToTs
	o.Ts = all.Ts
	o.Type = all.Type
	o.Val = all.Val
	o.Value = all.Value

	return nil
}
