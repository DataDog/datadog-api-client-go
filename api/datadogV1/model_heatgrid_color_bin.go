// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridColorBin A color and optional lower threshold for a discrete bin.
type HeatgridColorBin struct {
	// A color string, or two color strings for the light and dark themes, in that order.
	Color HeatgridColor `json:"color"`
	// Inclusive lower bound. Omit for the first bin.
	LowerBound *float64 `json:"lower_bound,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridColorBin instantiates a new HeatgridColorBin object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridColorBin(color HeatgridColor) *HeatgridColorBin {
	this := HeatgridColorBin{}
	this.Color = color
	return &this
}

// NewHeatgridColorBinWithDefaults instantiates a new HeatgridColorBin object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridColorBinWithDefaults() *HeatgridColorBin {
	this := HeatgridColorBin{}
	return &this
}

// GetColor returns the Color field value.
func (o *HeatgridColorBin) GetColor() HeatgridColor {
	if o == nil {
		var ret HeatgridColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *HeatgridColorBin) GetColorOk() (*HeatgridColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value.
func (o *HeatgridColorBin) SetColor(v HeatgridColor) {
	o.Color = v
}

// GetLowerBound returns the LowerBound field value if set, zero value otherwise.
func (o *HeatgridColorBin) GetLowerBound() float64 {
	if o == nil || o.LowerBound == nil {
		var ret float64
		return ret
	}
	return *o.LowerBound
}

// GetLowerBoundOk returns a tuple with the LowerBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatgridColorBin) GetLowerBoundOk() (*float64, bool) {
	if o == nil || o.LowerBound == nil {
		return nil, false
	}
	return o.LowerBound, true
}

// HasLowerBound returns a boolean if a field has been set.
func (o *HeatgridColorBin) HasLowerBound() bool {
	return o != nil && o.LowerBound != nil
}

// SetLowerBound gets a reference to the given float64 and assigns it to the LowerBound field.
func (o *HeatgridColorBin) SetLowerBound(v float64) {
	o.LowerBound = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridColorBin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["color"] = o.Color
	if o.LowerBound != nil {
		toSerialize["lower_bound"] = o.LowerBound
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridColorBin) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color      *HeatgridColor `json:"color"`
		LowerBound *float64       `json:"lower_bound,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Color == nil {
		return fmt.Errorf("required field color missing")
	}
	o.Color = *all.Color
	o.LowerBound = all.LowerBound

	return nil
}
