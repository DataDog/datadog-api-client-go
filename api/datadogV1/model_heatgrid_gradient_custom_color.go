// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientCustomColor A continuous gradient with custom color stops.
type HeatgridGradientCustomColor struct {
	// Use a continuous color gradient.
	Mode HeatgridGradientMode `json:"mode"`
	// Use custom colors.
	Source HeatgridCustomColorSource `json:"source"`
	// Two to six stops with positions in ascending order.
	Stops []HeatgridColorStop `json:"stops"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridGradientCustomColor instantiates a new HeatgridGradientCustomColor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridGradientCustomColor(mode HeatgridGradientMode, source HeatgridCustomColorSource, stops []HeatgridColorStop) *HeatgridGradientCustomColor {
	this := HeatgridGradientCustomColor{}
	this.Mode = mode
	this.Source = source
	this.Stops = stops
	return &this
}

// NewHeatgridGradientCustomColorWithDefaults instantiates a new HeatgridGradientCustomColor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridGradientCustomColorWithDefaults() *HeatgridGradientCustomColor {
	this := HeatgridGradientCustomColor{}
	return &this
}

// GetMode returns the Mode field value.
func (o *HeatgridGradientCustomColor) GetMode() HeatgridGradientMode {
	if o == nil {
		var ret HeatgridGradientMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientCustomColor) GetModeOk() (*HeatgridGradientMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *HeatgridGradientCustomColor) SetMode(v HeatgridGradientMode) {
	o.Mode = v
}

// GetSource returns the Source field value.
func (o *HeatgridGradientCustomColor) GetSource() HeatgridCustomColorSource {
	if o == nil {
		var ret HeatgridCustomColorSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientCustomColor) GetSourceOk() (*HeatgridCustomColorSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HeatgridGradientCustomColor) SetSource(v HeatgridCustomColorSource) {
	o.Source = v
}

// GetStops returns the Stops field value.
func (o *HeatgridGradientCustomColor) GetStops() []HeatgridColorStop {
	if o == nil {
		var ret []HeatgridColorStop
		return ret
	}
	return o.Stops
}

// GetStopsOk returns a tuple with the Stops field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientCustomColor) GetStopsOk() (*[]HeatgridColorStop, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stops, true
}

// SetStops sets field value.
func (o *HeatgridGradientCustomColor) SetStops(v []HeatgridColorStop) {
	o.Stops = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridGradientCustomColor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	toSerialize["source"] = o.Source
	toSerialize["stops"] = o.Stops
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridGradientCustomColor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode   *HeatgridGradientMode      `json:"mode"`
		Source *HeatgridCustomColorSource `json:"source"`
		Stops  *[]HeatgridColorStop       `json:"stops"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Stops == nil {
		return fmt.Errorf("required field stops missing")
	}

	hasInvalidField := false
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.Stops = *all.Stops

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
