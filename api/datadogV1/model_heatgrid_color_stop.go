// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridColorStop A position and color in a continuous gradient.
type HeatgridColorStop struct {
	// A color string, or two color strings for the light and dark themes, in that order.
	Color HeatgridColor `json:"color"`
	// Position in the gradient, from 0 to 100.
	Position int64 `json:"position"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridColorStop instantiates a new HeatgridColorStop object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridColorStop(color HeatgridColor, position int64) *HeatgridColorStop {
	this := HeatgridColorStop{}
	this.Color = color
	this.Position = position
	return &this
}

// NewHeatgridColorStopWithDefaults instantiates a new HeatgridColorStop object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridColorStopWithDefaults() *HeatgridColorStop {
	this := HeatgridColorStop{}
	return &this
}

// GetColor returns the Color field value.
func (o *HeatgridColorStop) GetColor() HeatgridColor {
	if o == nil {
		var ret HeatgridColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *HeatgridColorStop) GetColorOk() (*HeatgridColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value.
func (o *HeatgridColorStop) SetColor(v HeatgridColor) {
	o.Color = v
}

// GetPosition returns the Position field value.
func (o *HeatgridColorStop) GetPosition() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *HeatgridColorStop) GetPositionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value.
func (o *HeatgridColorStop) SetPosition(v int64) {
	o.Position = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridColorStop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["color"] = o.Color
	toSerialize["position"] = o.Position
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridColorStop) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color    *HeatgridColor `json:"color"`
		Position *int64         `json:"position"`
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
