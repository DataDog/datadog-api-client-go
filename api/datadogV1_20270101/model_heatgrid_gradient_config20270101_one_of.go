// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientConfig20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridGradientConfig20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Mode HeatgridGradientConfigAnyOfAnyOfItemMode `json:"mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Source HeatgridGradientConfigAnyOfAnyOfItemSource `json:"source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Stops []HeatgridStop `json:"stops"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridGradientConfig20270101OneOf instantiates a new HeatgridGradientConfig20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridGradientConfig20270101OneOf(mode HeatgridGradientConfigAnyOfAnyOfItemMode, source HeatgridGradientConfigAnyOfAnyOfItemSource, stops []HeatgridStop) *HeatgridGradientConfig20270101OneOf {
	this := HeatgridGradientConfig20270101OneOf{}
	this.Mode = mode
	this.Source = source
	this.Stops = stops
	return &this
}

// NewHeatgridGradientConfig20270101OneOfWithDefaults instantiates a new HeatgridGradientConfig20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridGradientConfig20270101OneOfWithDefaults() *HeatgridGradientConfig20270101OneOf {
	this := HeatgridGradientConfig20270101OneOf{}
	return &this
}

// GetMode returns the Mode field value.
func (o *HeatgridGradientConfig20270101OneOf) GetMode() HeatgridGradientConfigAnyOfAnyOfItemMode {
	if o == nil {
		var ret HeatgridGradientConfigAnyOfAnyOfItemMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientConfig20270101OneOf) GetModeOk() (*HeatgridGradientConfigAnyOfAnyOfItemMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *HeatgridGradientConfig20270101OneOf) SetMode(v HeatgridGradientConfigAnyOfAnyOfItemMode) {
	o.Mode = v
}

// GetSource returns the Source field value.
func (o *HeatgridGradientConfig20270101OneOf) GetSource() HeatgridGradientConfigAnyOfAnyOfItemSource {
	if o == nil {
		var ret HeatgridGradientConfigAnyOfAnyOfItemSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientConfig20270101OneOf) GetSourceOk() (*HeatgridGradientConfigAnyOfAnyOfItemSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HeatgridGradientConfig20270101OneOf) SetSource(v HeatgridGradientConfigAnyOfAnyOfItemSource) {
	o.Source = v
}

// GetStops returns the Stops field value.
func (o *HeatgridGradientConfig20270101OneOf) GetStops() []HeatgridStop {
	if o == nil {
		var ret []HeatgridStop
		return ret
	}
	return o.Stops
}

// GetStopsOk returns a tuple with the Stops field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientConfig20270101OneOf) GetStopsOk() (*[]HeatgridStop, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stops, true
}

// SetStops sets field value.
func (o *HeatgridGradientConfig20270101OneOf) SetStops(v []HeatgridStop) {
	o.Stops = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridGradientConfig20270101OneOf) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridGradientConfig20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode   *HeatgridGradientConfigAnyOfAnyOfItemMode   `json:"mode"`
		Source *HeatgridGradientConfigAnyOfAnyOfItemSource `json:"source"`
		Stops  *[]HeatgridStop                             `json:"stops"`
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
