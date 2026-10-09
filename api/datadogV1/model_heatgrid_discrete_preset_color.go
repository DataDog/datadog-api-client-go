// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscretePresetColor A preset discrete color palette.
type HeatgridDiscretePresetColor struct {
	// Use discrete color thresholds.
	Mode HeatgridDiscreteMode `json:"mode"`
	// Name of the preset color palette.
	PresetName string `json:"preset_name"`
	// Use a preset color palette.
	Source HeatgridPresetColorSource `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridDiscretePresetColor instantiates a new HeatgridDiscretePresetColor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridDiscretePresetColor(mode HeatgridDiscreteMode, presetName string, source HeatgridPresetColorSource) *HeatgridDiscretePresetColor {
	this := HeatgridDiscretePresetColor{}
	this.Mode = mode
	this.PresetName = presetName
	this.Source = source
	return &this
}

// NewHeatgridDiscretePresetColorWithDefaults instantiates a new HeatgridDiscretePresetColor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridDiscretePresetColorWithDefaults() *HeatgridDiscretePresetColor {
	this := HeatgridDiscretePresetColor{}
	return &this
}

// GetMode returns the Mode field value.
func (o *HeatgridDiscretePresetColor) GetMode() HeatgridDiscreteMode {
	if o == nil {
		var ret HeatgridDiscreteMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscretePresetColor) GetModeOk() (*HeatgridDiscreteMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *HeatgridDiscretePresetColor) SetMode(v HeatgridDiscreteMode) {
	o.Mode = v
}

// GetPresetName returns the PresetName field value.
func (o *HeatgridDiscretePresetColor) GetPresetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PresetName
}

// GetPresetNameOk returns a tuple with the PresetName field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscretePresetColor) GetPresetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetName, true
}

// SetPresetName sets field value.
func (o *HeatgridDiscretePresetColor) SetPresetName(v string) {
	o.PresetName = v
}

// GetSource returns the Source field value.
func (o *HeatgridDiscretePresetColor) GetSource() HeatgridPresetColorSource {
	if o == nil {
		var ret HeatgridPresetColorSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscretePresetColor) GetSourceOk() (*HeatgridPresetColorSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HeatgridDiscretePresetColor) SetSource(v HeatgridPresetColorSource) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridDiscretePresetColor) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridDiscretePresetColor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode       *HeatgridDiscreteMode      `json:"mode"`
		PresetName *string                    `json:"preset_name"`
		Source     *HeatgridPresetColorSource `json:"source"`
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
	o.PresetName = *all.PresetName
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
