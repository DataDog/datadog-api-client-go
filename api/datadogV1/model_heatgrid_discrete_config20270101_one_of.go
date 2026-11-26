// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteConfig20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type HeatgridDiscreteConfig20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Bins []HeatgridBin_20270101 `json:"bins"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Mode HeatgridDiscreteConfigAnyOfAnyOfItemMode_20270101 `json:"mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Source HeatgridDiscreteConfigAnyOfAnyOfItemSource_20270101 `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridDiscreteConfig20270101OneOf instantiates a new HeatgridDiscreteConfig20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridDiscreteConfig20270101OneOf(bins []HeatgridBin_20270101, mode HeatgridDiscreteConfigAnyOfAnyOfItemMode_20270101, source HeatgridDiscreteConfigAnyOfAnyOfItemSource_20270101) *HeatgridDiscreteConfig20270101OneOf {
	this := HeatgridDiscreteConfig20270101OneOf{}
	this.Bins = bins
	this.Mode = mode
	this.Source = source
	return &this
}

// NewHeatgridDiscreteConfig20270101OneOfWithDefaults instantiates a new HeatgridDiscreteConfig20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridDiscreteConfig20270101OneOfWithDefaults() *HeatgridDiscreteConfig20270101OneOf {
	this := HeatgridDiscreteConfig20270101OneOf{}
	return &this
}

// GetBins returns the Bins field value.
func (o *HeatgridDiscreteConfig20270101OneOf) GetBins() []HeatgridBin_20270101 {
	if o == nil {
		var ret []HeatgridBin_20270101
		return ret
	}
	return o.Bins
}

// GetBinsOk returns a tuple with the Bins field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteConfig20270101OneOf) GetBinsOk() (*[]HeatgridBin_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bins, true
}

// SetBins sets field value.
func (o *HeatgridDiscreteConfig20270101OneOf) SetBins(v []HeatgridBin_20270101) {
	o.Bins = v
}

// GetMode returns the Mode field value.
func (o *HeatgridDiscreteConfig20270101OneOf) GetMode() HeatgridDiscreteConfigAnyOfAnyOfItemMode_20270101 {
	if o == nil {
		var ret HeatgridDiscreteConfigAnyOfAnyOfItemMode_20270101
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteConfig20270101OneOf) GetModeOk() (*HeatgridDiscreteConfigAnyOfAnyOfItemMode_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *HeatgridDiscreteConfig20270101OneOf) SetMode(v HeatgridDiscreteConfigAnyOfAnyOfItemMode_20270101) {
	o.Mode = v
}

// GetSource returns the Source field value.
func (o *HeatgridDiscreteConfig20270101OneOf) GetSource() HeatgridDiscreteConfigAnyOfAnyOfItemSource_20270101 {
	if o == nil {
		var ret HeatgridDiscreteConfigAnyOfAnyOfItemSource_20270101
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteConfig20270101OneOf) GetSourceOk() (*HeatgridDiscreteConfigAnyOfAnyOfItemSource_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HeatgridDiscreteConfig20270101OneOf) SetSource(v HeatgridDiscreteConfigAnyOfAnyOfItemSource_20270101) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridDiscreteConfig20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bins"] = o.Bins
	toSerialize["mode"] = o.Mode
	toSerialize["source"] = o.Source
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatgridDiscreteConfig20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bins   *[]HeatgridBin_20270101                              `json:"bins"`
		Mode   *HeatgridDiscreteConfigAnyOfAnyOfItemMode_20270101   `json:"mode"`
		Source *HeatgridDiscreteConfigAnyOfAnyOfItemSource_20270101 `json:"source"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Bins == nil {
		return fmt.Errorf("required field bins missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}

	hasInvalidField := false
	o.Bins = *all.Bins
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
