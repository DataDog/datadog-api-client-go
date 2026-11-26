// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MSLBackendWidgetLayout Test-only generated schema for the dashboard versioning proof of concept.
type MSLBackendWidgetLayout struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Height float64 `json:"height"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsColumnBreak *bool `json:"is_column_break,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Width float64 `json:"width"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	X float64 `json:"x"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Y float64 `json:"y"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMSLBackendWidgetLayout instantiates a new MSLBackendWidgetLayout object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMSLBackendWidgetLayout(height float64, width float64, x float64, y float64) *MSLBackendWidgetLayout {
	this := MSLBackendWidgetLayout{}
	this.Height = height
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewMSLBackendWidgetLayoutWithDefaults instantiates a new MSLBackendWidgetLayout object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMSLBackendWidgetLayoutWithDefaults() *MSLBackendWidgetLayout {
	this := MSLBackendWidgetLayout{}
	return &this
}

// GetHeight returns the Height field value.
func (o *MSLBackendWidgetLayout) GetHeight() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *MSLBackendWidgetLayout) GetHeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value.
func (o *MSLBackendWidgetLayout) SetHeight(v float64) {
	o.Height = v
}

// GetIsColumnBreak returns the IsColumnBreak field value if set, zero value otherwise.
func (o *MSLBackendWidgetLayout) GetIsColumnBreak() bool {
	if o == nil || o.IsColumnBreak == nil {
		var ret bool
		return ret
	}
	return *o.IsColumnBreak
}

// GetIsColumnBreakOk returns a tuple with the IsColumnBreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MSLBackendWidgetLayout) GetIsColumnBreakOk() (*bool, bool) {
	if o == nil || o.IsColumnBreak == nil {
		return nil, false
	}
	return o.IsColumnBreak, true
}

// HasIsColumnBreak returns a boolean if a field has been set.
func (o *MSLBackendWidgetLayout) HasIsColumnBreak() bool {
	return o != nil && o.IsColumnBreak != nil
}

// SetIsColumnBreak gets a reference to the given bool and assigns it to the IsColumnBreak field.
func (o *MSLBackendWidgetLayout) SetIsColumnBreak(v bool) {
	o.IsColumnBreak = &v
}

// GetWidth returns the Width field value.
func (o *MSLBackendWidgetLayout) GetWidth() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *MSLBackendWidgetLayout) GetWidthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value.
func (o *MSLBackendWidgetLayout) SetWidth(v float64) {
	o.Width = v
}

// GetX returns the X field value.
func (o *MSLBackendWidgetLayout) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *MSLBackendWidgetLayout) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value.
func (o *MSLBackendWidgetLayout) SetX(v float64) {
	o.X = v
}

// GetY returns the Y field value.
func (o *MSLBackendWidgetLayout) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *MSLBackendWidgetLayout) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value.
func (o *MSLBackendWidgetLayout) SetY(v float64) {
	o.Y = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MSLBackendWidgetLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["height"] = o.Height
	if o.IsColumnBreak != nil {
		toSerialize["is_column_break"] = o.IsColumnBreak
	}
	toSerialize["width"] = o.Width
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MSLBackendWidgetLayout) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Height        *float64 `json:"height"`
		IsColumnBreak *bool    `json:"is_column_break,omitempty"`
		Width         *float64 `json:"width"`
		X             *float64 `json:"x"`
		Y             *float64 `json:"y"`
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
	o.IsColumnBreak = all.IsColumnBreak
	o.Width = *all.Width
	o.X = *all.X
	o.Y = *all.Y

	return nil
}
