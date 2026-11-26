// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteConfig20270101OneOf4457881312 Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDiscreteConfig20270101OneOf4457881312 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Mode HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 `json:"mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	PresetName HeatgridPresetName `json:"preset_name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Source HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridDiscreteConfig20270101OneOf4457881312 instantiates a new HeatgridDiscreteConfig20270101OneOf4457881312 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridDiscreteConfig20270101OneOf4457881312(mode HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848, presetName HeatgridPresetName, source HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168) *HeatgridDiscreteConfig20270101OneOf4457881312 {
	this := HeatgridDiscreteConfig20270101OneOf4457881312{}
	this.Mode = mode
	this.PresetName = presetName
	this.Source = source
	return &this
}

// NewHeatgridDiscreteConfig20270101OneOf4457881312WithDefaults instantiates a new HeatgridDiscreteConfig20270101OneOf4457881312 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridDiscreteConfig20270101OneOf4457881312WithDefaults() *HeatgridDiscreteConfig20270101OneOf4457881312 {
	this := HeatgridDiscreteConfig20270101OneOf4457881312{}
	return &this
}

// GetMode returns the Mode field value.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) GetMode() HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848 {
	if o == nil {
		var ret HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) GetModeOk() (*HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) SetMode(v HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848) {
	o.Mode = v
}

// GetPresetName returns the PresetName field value.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) GetPresetName() HeatgridPresetName {
	if o == nil {
		var ret HeatgridPresetName
		return ret
	}
	return o.PresetName
}

// GetPresetNameOk returns a tuple with the PresetName field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) GetPresetNameOk() (*HeatgridPresetName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresetName, true
}

// SetPresetName sets field value.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) SetPresetName(v HeatgridPresetName) {
	o.PresetName = v
}

// GetSource returns the Source field value.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) GetSource() HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 {
	if o == nil {
		var ret HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) GetSourceOk() (*HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) SetSource(v HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridDiscreteConfig20270101OneOf4457881312) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridDiscreteConfig20270101OneOf4457881312) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode       *HeatgridDiscreteConfigAnyOfAnyOfItemMode4447618848   `json:"mode"`
		PresetName *HeatgridPresetName                                   `json:"preset_name"`
		Source     *HeatgridDiscreteConfigAnyOfAnyOfItemSource4447619168 `json:"source"`
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
