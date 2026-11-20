// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetTime20270101OneOf4464103552 Test-only generated schema for the dashboard versioning proof of concept.
type WidgetTime20270101OneOf4464103552 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	HideIncompleteCostData *bool `json:"hide_incomplete_cost_data,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type WidgetNewLiveSpanType_20270101 `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Unit string `json:"unit"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewWidgetTime20270101OneOf4464103552 instantiates a new WidgetTime20270101OneOf4464103552 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetTime20270101OneOf4464103552(typeVar WidgetNewLiveSpanType_20270101, unit string, value float64) *WidgetTime20270101OneOf4464103552 {
	this := WidgetTime20270101OneOf4464103552{}
	this.Type = typeVar
	this.Unit = unit
	this.Value = value
	return &this
}

// NewWidgetTime20270101OneOf4464103552WithDefaults instantiates a new WidgetTime20270101OneOf4464103552 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetTime20270101OneOf4464103552WithDefaults() *WidgetTime20270101OneOf4464103552 {
	this := WidgetTime20270101OneOf4464103552{}
	return &this
}

// GetHideIncompleteCostData returns the HideIncompleteCostData field value if set, zero value otherwise.
func (o *WidgetTime20270101OneOf4464103552) GetHideIncompleteCostData() bool {
	if o == nil || o.HideIncompleteCostData == nil {
		var ret bool
		return ret
	}
	return *o.HideIncompleteCostData
}

// GetHideIncompleteCostDataOk returns a tuple with the HideIncompleteCostData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464103552) GetHideIncompleteCostDataOk() (*bool, bool) {
	if o == nil || o.HideIncompleteCostData == nil {
		return nil, false
	}
	return o.HideIncompleteCostData, true
}

// HasHideIncompleteCostData returns a boolean if a field has been set.
func (o *WidgetTime20270101OneOf4464103552) HasHideIncompleteCostData() bool {
	return o != nil && o.HideIncompleteCostData != nil
}

// SetHideIncompleteCostData gets a reference to the given bool and assigns it to the HideIncompleteCostData field.
func (o *WidgetTime20270101OneOf4464103552) SetHideIncompleteCostData(v bool) {
	o.HideIncompleteCostData = &v
}

// GetType returns the Type field value.
func (o *WidgetTime20270101OneOf4464103552) GetType() WidgetNewLiveSpanType_20270101 {
	if o == nil {
		var ret WidgetNewLiveSpanType_20270101
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464103552) GetTypeOk() (*WidgetNewLiveSpanType_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *WidgetTime20270101OneOf4464103552) SetType(v WidgetNewLiveSpanType_20270101) {
	o.Type = v
}

// GetUnit returns the Unit field value.
func (o *WidgetTime20270101OneOf4464103552) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464103552) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *WidgetTime20270101OneOf4464103552) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value.
func (o *WidgetTime20270101OneOf4464103552) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WidgetTime20270101OneOf4464103552) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *WidgetTime20270101OneOf4464103552) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetTime20270101OneOf4464103552) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HideIncompleteCostData != nil {
		toSerialize["hide_incomplete_cost_data"] = o.HideIncompleteCostData
	}
	toSerialize["type"] = o.Type
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetTime20270101OneOf4464103552) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HideIncompleteCostData *bool                           `json:"hide_incomplete_cost_data,omitempty"`
		Type                   *WidgetNewLiveSpanType_20270101 `json:"type"`
		Unit                   *string                         `json:"unit"`
		Value                  *float64                        `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}

	hasInvalidField := false
	o.HideIncompleteCostData = all.HideIncompleteCostData
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Unit = *all.Unit
	o.Value = *all.Value

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
