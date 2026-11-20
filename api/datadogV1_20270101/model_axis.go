// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Axis Test-only generated schema for the dashboard versioning proof of concept.
type Axis struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	IncludeZero *bool `json:"include_zero,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Max *AxisBound `json:"max,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Min *AxisBound `json:"min,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Scale *ScaleType `json:"scale,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Units *bool `json:"units,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAxis instantiates a new Axis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAxis() *Axis {
	this := Axis{}
	return &this
}

// NewAxisWithDefaults instantiates a new Axis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAxisWithDefaults() *Axis {
	this := Axis{}
	return &this
}

// GetIncludeZero returns the IncludeZero field value if set, zero value otherwise.
func (o *Axis) GetIncludeZero() bool {
	if o == nil || o.IncludeZero == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZero
}

// GetIncludeZeroOk returns a tuple with the IncludeZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetIncludeZeroOk() (*bool, bool) {
	if o == nil || o.IncludeZero == nil {
		return nil, false
	}
	return o.IncludeZero, true
}

// HasIncludeZero returns a boolean if a field has been set.
func (o *Axis) HasIncludeZero() bool {
	return o != nil && o.IncludeZero != nil
}

// SetIncludeZero gets a reference to the given bool and assigns it to the IncludeZero field.
func (o *Axis) SetIncludeZero(v bool) {
	o.IncludeZero = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Axis) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Axis) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Axis) SetLabel(v string) {
	o.Label = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Axis) GetMax() AxisBound {
	if o == nil || o.Max == nil {
		var ret AxisBound
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetMaxOk() (*AxisBound, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Axis) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given AxisBound and assigns it to the Max field.
func (o *Axis) SetMax(v AxisBound) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *Axis) GetMin() AxisBound {
	if o == nil || o.Min == nil {
		var ret AxisBound
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetMinOk() (*AxisBound, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *Axis) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given AxisBound and assigns it to the Min field.
func (o *Axis) SetMin(v AxisBound) {
	o.Min = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *Axis) GetScale() ScaleType {
	if o == nil || o.Scale == nil {
		var ret ScaleType
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetScaleOk() (*ScaleType, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *Axis) HasScale() bool {
	return o != nil && o.Scale != nil
}

// SetScale gets a reference to the given ScaleType and assigns it to the Scale field.
func (o *Axis) SetScale(v ScaleType) {
	o.Scale = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *Axis) GetUnits() bool {
	if o == nil || o.Units == nil {
		var ret bool
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Axis) GetUnitsOk() (*bool, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *Axis) HasUnits() bool {
	return o != nil && o.Units != nil
}

// SetUnits gets a reference to the given bool and assigns it to the Units field.
func (o *Axis) SetUnits(v bool) {
	o.Units = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Axis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeZero != nil {
		toSerialize["include_zero"] = o.IncludeZero
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
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Axis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeZero *bool      `json:"include_zero,omitempty"`
		Label       *string    `json:"label,omitempty"`
		Max         *AxisBound `json:"max,omitempty"`
		Min         *AxisBound `json:"min,omitempty"`
		Scale       *ScaleType `json:"scale,omitempty"`
		Units       *bool      `json:"units,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.IncludeZero = all.IncludeZero
	o.Label = all.Label
	o.Max = all.Max
	o.Min = all.Min
	if all.Scale != nil && !all.Scale.IsValid() {
		hasInvalidField = true
	} else {
		o.Scale = all.Scale
	}
	o.Units = all.Units

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
