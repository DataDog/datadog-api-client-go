// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FreeLayout Test-only generated schema for the dashboard versioning proof of concept.
type FreeLayout struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Height *float64 `json:"height,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Width *float64 `json:"width,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	X float64 `json:"x"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Y float64 `json:"y"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFreeLayout instantiates a new FreeLayout object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFreeLayout(x float64, y float64) *FreeLayout {
	this := FreeLayout{}
	this.X = x
	this.Y = y
	return &this
}

// NewFreeLayoutWithDefaults instantiates a new FreeLayout object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFreeLayoutWithDefaults() *FreeLayout {
	this := FreeLayout{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *FreeLayout) GetHeight() float64 {
	if o == nil || o.Height == nil {
		var ret float64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeLayout) GetHeightOk() (*float64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *FreeLayout) HasHeight() bool {
	return o != nil && o.Height != nil
}

// SetHeight gets a reference to the given float64 and assigns it to the Height field.
func (o *FreeLayout) SetHeight(v float64) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *FreeLayout) GetWidth() float64 {
	if o == nil || o.Width == nil {
		var ret float64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeLayout) GetWidthOk() (*float64, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *FreeLayout) HasWidth() bool {
	return o != nil && o.Width != nil
}

// SetWidth gets a reference to the given float64 and assigns it to the Width field.
func (o *FreeLayout) SetWidth(v float64) {
	o.Width = &v
}

// GetX returns the X field value.
func (o *FreeLayout) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *FreeLayout) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value.
func (o *FreeLayout) SetX(v float64) {
	o.X = v
}

// GetY returns the Y field value.
func (o *FreeLayout) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *FreeLayout) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value.
func (o *FreeLayout) SetY(v float64) {
	o.Y = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FreeLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FreeLayout) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Height *float64 `json:"height,omitempty"`
		Width  *float64 `json:"width,omitempty"`
		X      *float64 `json:"x"`
		Y      *float64 `json:"y"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.X == nil {
		return fmt.Errorf("required field x missing")
	}
	if all.Y == nil {
		return fmt.Errorf("required field y missing")
	}
	o.Height = all.Height
	o.Width = all.Width
	o.X = *all.X
	o.Y = *all.Y

	return nil
}
