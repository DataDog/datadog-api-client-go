// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridBin Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridBin struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Color HeatgridColor `json:"color"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	LowerBound *float64 `json:"lower_bound,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridBin instantiates a new HeatgridBin object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridBin(color HeatgridColor) *HeatgridBin {
	this := HeatgridBin{}
	this.Color = color
	return &this
}

// NewHeatgridBinWithDefaults instantiates a new HeatgridBin object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridBinWithDefaults() *HeatgridBin {
	this := HeatgridBin{}
	return &this
}

// GetColor returns the Color field value.
func (o *HeatgridBin) GetColor() HeatgridColor {
	if o == nil {
		var ret HeatgridColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *HeatgridBin) GetColorOk() (*HeatgridColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value.
func (o *HeatgridBin) SetColor(v HeatgridColor) {
	o.Color = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *HeatgridBin) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridBin) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *HeatgridBin) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *HeatgridBin) SetLabel(v string) {
	o.Label = &v
}

// GetLowerBound returns the LowerBound field value if set, zero value otherwise.
func (o *HeatgridBin) GetLowerBound() float64 {
	if o == nil || o.LowerBound == nil {
		var ret float64
		return ret
	}
	return *o.LowerBound
}

// GetLowerBoundOk returns a tuple with the LowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridBin) GetLowerBoundOk() (*float64, bool) {
	if o == nil || o.LowerBound == nil {
		return nil, false
	}
	return o.LowerBound, true
}

// HasLowerBound returns a boolean if a field has been set.
func (o *HeatgridBin) HasLowerBound() bool {
	return o != nil && o.LowerBound != nil
}

// SetLowerBound gets a reference to the given float64 and assigns it to the LowerBound field.
func (o *HeatgridBin) SetLowerBound(v float64) {
	o.LowerBound = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridBin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["color"] = o.Color
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LowerBound != nil {
		toSerialize["lower_bound"] = o.LowerBound
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridBin) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color      *HeatgridColor `json:"color"`
		Label      *string        `json:"label,omitempty"`
		LowerBound *float64       `json:"lower_bound,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Color == nil {
		return fmt.Errorf("required field color missing")
	}
	o.Color = *all.Color
	o.Label = all.Label
	o.LowerBound = all.LowerBound

	return nil
}
