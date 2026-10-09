// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatgridDiscreteCustomColor Discrete thresholds with custom colors.
type HeatgridDiscreteCustomColor struct {
	// Two to six bins. Omit `lower_bound` on the first bin. Subsequent lower bounds must be in
	// ascending order.
	Bins []HeatgridColorBin `json:"bins"`
	// Use discrete color thresholds.
	Mode HeatgridDiscreteMode `json:"mode"`
	// Use custom colors.
	Source HeatgridCustomColorSource `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatgridDiscreteCustomColor instantiates a new HeatgridDiscreteCustomColor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatgridDiscreteCustomColor(bins []HeatgridColorBin, mode HeatgridDiscreteMode, source HeatgridCustomColorSource) *HeatgridDiscreteCustomColor {
	this := HeatgridDiscreteCustomColor{}
	this.Bins = bins
	this.Mode = mode
	this.Source = source
	return &this
}

// NewHeatgridDiscreteCustomColorWithDefaults instantiates a new HeatgridDiscreteCustomColor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatgridDiscreteCustomColorWithDefaults() *HeatgridDiscreteCustomColor {
	this := HeatgridDiscreteCustomColor{}
	return &this
}

// GetBins returns the Bins field value.
func (o *HeatgridDiscreteCustomColor) GetBins() []HeatgridColorBin {
	if o == nil {
		var ret []HeatgridColorBin
		return ret
	}
	return o.Bins
}

// GetBinsOk returns a tuple with the Bins field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteCustomColor) GetBinsOk() (*[]HeatgridColorBin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bins, true
}

// SetBins sets field value.
func (o *HeatgridDiscreteCustomColor) SetBins(v []HeatgridColorBin) {
	o.Bins = v
}

// GetMode returns the Mode field value.
func (o *HeatgridDiscreteCustomColor) GetMode() HeatgridDiscreteMode {
	if o == nil {
		var ret HeatgridDiscreteMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteCustomColor) GetModeOk() (*HeatgridDiscreteMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *HeatgridDiscreteCustomColor) SetMode(v HeatgridDiscreteMode) {
	o.Mode = v
}

// GetSource returns the Source field value.
func (o *HeatgridDiscreteCustomColor) GetSource() HeatgridCustomColorSource {
	if o == nil {
		var ret HeatgridCustomColorSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *HeatgridDiscreteCustomColor) GetSourceOk() (*HeatgridCustomColorSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *HeatgridDiscreteCustomColor) SetSource(v HeatgridCustomColorSource) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatgridDiscreteCustomColor) MarshalJSON() ([]byte, error) {
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
func (o *HeatgridDiscreteCustomColor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bins   *[]HeatgridColorBin        `json:"bins"`
		Mode   *HeatgridDiscreteMode      `json:"mode"`
		Source *HeatgridCustomColorSource `json:"source"`
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
