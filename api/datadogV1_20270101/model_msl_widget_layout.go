// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MSLWidgetLayout Test-only generated schema for the dashboard versioning proof of concept.
type MSLWidgetLayout struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Height float64 `json:"height"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Parent datadog.NullableFloat64 `json:"parent,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Width float64 `json:"width"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	X float64 `json:"x"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Y float64 `json:"y"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMSLWidgetLayout instantiates a new MSLWidgetLayout object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMSLWidgetLayout(height float64, width float64, x float64, y float64) *MSLWidgetLayout {
	this := MSLWidgetLayout{}
	this.Height = height
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewMSLWidgetLayoutWithDefaults instantiates a new MSLWidgetLayout object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMSLWidgetLayoutWithDefaults() *MSLWidgetLayout {
	this := MSLWidgetLayout{}
	return &this
}

// GetHeight returns the Height field value.
func (o *MSLWidgetLayout) GetHeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *MSLWidgetLayout) GetHeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value.
func (o *MSLWidgetLayout) SetHeight(v float64) {
	o.Height = v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSLWidgetLayout) GetParent() float64 {
	if o == nil || o.Parent.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MSLWidgetLayout) GetParentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *MSLWidgetLayout) HasParent() bool {
	return o != nil && o.Parent.IsSet()
}

// SetParent gets a reference to the given datadog.NullableFloat64 and assigns it to the Parent field.
func (o *MSLWidgetLayout) SetParent(v float64) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil.
func (o *MSLWidgetLayout) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil.
func (o *MSLWidgetLayout) UnsetParent() {
	o.Parent.Unset()
}

// GetWidth returns the Width field value.
func (o *MSLWidgetLayout) GetWidth() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *MSLWidgetLayout) GetWidthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value.
func (o *MSLWidgetLayout) SetWidth(v float64) {
	o.Width = v
}

// GetX returns the X field value.
func (o *MSLWidgetLayout) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *MSLWidgetLayout) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value.
func (o *MSLWidgetLayout) SetX(v float64) {
	o.X = v
}

// GetY returns the Y field value.
func (o *MSLWidgetLayout) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *MSLWidgetLayout) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value.
func (o *MSLWidgetLayout) SetY(v float64) {
	o.Y = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MSLWidgetLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["height"] = o.Height
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	toSerialize["width"] = o.Width
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MSLWidgetLayout) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Height *float64                `json:"height"`
		Parent datadog.NullableFloat64 `json:"parent,omitempty"`
		Width  *float64                `json:"width"`
		X      *float64                `json:"x"`
		Y      *float64                `json:"y"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Height == nil {
		return fmt.Errorf("required field height missing")
	}
	if all.Width == nil {
		return fmt.Errorf("required field width missing")
	}
	if all.X == nil {
		return fmt.Errorf("required field x missing")
	}
	if all.Y == nil {
		return fmt.Errorf("required field y missing")
	}
	o.Height = *all.Height
	o.Parent = all.Parent
	o.Width = *all.Width
	o.X = *all.X
	o.Y = *all.Y

	return nil
}
