// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridGradientConfig20270101OneOf4457883232 Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridGradientConfig20270101OneOf4457883232 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Mode HeatgridGradientConfigAnyOfAnyOfItemMode4447621888_20270101 `json:"mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PresetName HeatgridPresetName_20270101 `json:"preset_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Source HeatgridGradientConfigAnyOfAnyOfItemSource4447622208_20270101 `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridGradientConfig20270101OneOf4457883232 instantiates a new HeatgridGradientConfig20270101OneOf4457883232 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridGradientConfig20270101OneOf4457883232(mode HeatgridGradientConfigAnyOfAnyOfItemMode4447621888_20270101, presetName HeatgridPresetName_20270101, source HeatgridGradientConfigAnyOfAnyOfItemSource4447622208_20270101) *HeatgridGradientConfig20270101OneOf4457883232 {
	this := HeatgridGradientConfig20270101OneOf4457883232{}
	this.Mode = mode
	this.PresetName = presetName
	this.Source = source
	return &this
}

// NewHeatgridGradientConfig20270101OneOf4457883232WithDefaults instantiates a new HeatgridGradientConfig20270101OneOf4457883232 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridGradientConfig20270101OneOf4457883232WithDefaults() *HeatgridGradientConfig20270101OneOf4457883232 {
	this := HeatgridGradientConfig20270101OneOf4457883232{}
	return &this
}

// GetMode returns the Mode field value.
func (o *HeatgridGradientConfig20270101OneOf4457883232) GetMode() HeatgridGradientConfigAnyOfAnyOfItemMode4447621888_20270101 {
	if o == nil {
		var ret HeatgridGradientConfigAnyOfAnyOfItemMode4447621888_20270101
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientConfig20270101OneOf4457883232) GetModeOk() (*HeatgridGradientConfigAnyOfAnyOfItemMode4447621888_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *HeatgridGradientConfig20270101OneOf4457883232) SetMode(v HeatgridGradientConfigAnyOfAnyOfItemMode4447621888_20270101) {
	o.Mode = v
}

// GetPresetName returns the PresetName field value.
func (o *HeatgridGradientConfig20270101OneOf4457883232) GetPresetName() HeatgridPresetName_20270101 {
	if o == nil {
		var ret HeatgridPresetName_20270101
		return ret
	}
	return o.PresetName
}

// GetPresetNameOk returns a tuple with the PresetName field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientConfig20270101OneOf4457883232) GetPresetNameOk() (*HeatgridPresetName_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetName, true
}

// SetPresetName sets field value.
func (o *HeatgridGradientConfig20270101OneOf4457883232) SetPresetName(v HeatgridPresetName_20270101) {
	o.PresetName = v
}

// GetSource returns the Source field value.
func (o *HeatgridGradientConfig20270101OneOf4457883232) GetSource() HeatgridGradientConfigAnyOfAnyOfItemSource4447622208_20270101 {
	if o == nil {
		var ret HeatgridGradientConfigAnyOfAnyOfItemSource4447622208_20270101
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HeatgridGradientConfig20270101OneOf4457883232) GetSourceOk() (*HeatgridGradientConfigAnyOfAnyOfItemSource4447622208_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HeatgridGradientConfig20270101OneOf4457883232) SetSource(v HeatgridGradientConfigAnyOfAnyOfItemSource4447622208_20270101) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridGradientConfig20270101OneOf4457883232) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	toSerialize["preset_name"] = o.PresetName
	toSerialize["source"] = o.Source
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridGradientConfig20270101OneOf4457883232) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode       *HeatgridGradientConfigAnyOfAnyOfItemMode4447621888_20270101   `json:"mode"`
		PresetName *HeatgridPresetName_20270101                                   `json:"preset_name"`
		Source     *HeatgridGradientConfigAnyOfAnyOfItemSource4447622208_20270101 `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.PresetName == nil {
		return fmt.Errorf("required field preset_name missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}

	hasInvalidField := false
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	if !all.PresetName.IsValid() {
		hasInvalidField = true
	} else {
		o.PresetName = *all.PresetName
	}
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
