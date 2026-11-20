// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridStop Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridStop struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Color HeatgridColor `json:"color"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Position float64 `json:"position"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridStop instantiates a new HeatgridStop object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridStop(color HeatgridColor, position float64) *HeatgridStop {
	this := HeatgridStop{}
	this.Color = color
	this.Position = position
	return &this
}

// NewHeatgridStopWithDefaults instantiates a new HeatgridStop object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridStopWithDefaults() *HeatgridStop {
	this := HeatgridStop{}
	return &this
}

// GetColor returns the Color field value.
func (o *HeatgridStop) GetColor() HeatgridColor {
	if o == nil {
		var ret HeatgridColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *HeatgridStop) GetColorOk() (*HeatgridColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value.
func (o *HeatgridStop) SetColor(v HeatgridColor) {
	o.Color = v
}

// GetPosition returns the Position field value.
func (o *HeatgridStop) GetPosition() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *HeatgridStop) GetPositionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value.
func (o *HeatgridStop) SetPosition(v float64) {
	o.Position = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridStop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["color"] = o.Color
	toSerialize["position"] = o.Position
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridStop) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color    *HeatgridColor `json:"color"`
		Position *float64       `json:"position"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Color == nil {
		return fmt.Errorf("required field color missing")
	}
	if all.Position == nil {
		return fmt.Errorf("required field position missing")
	}
	o.Color = *all.Color
	o.Position = *all.Position

	return nil
}
