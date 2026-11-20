// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapBackendDefinitionYaxis Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapBackendDefinitionYaxis struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	IncludeZero *bool `json:"include_zero,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Max *string `json:"max,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Min *string `json:"min,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumBuckets *float64 `json:"num_buckets,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Scale *ScaleType `json:"scale,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowUnit *bool `json:"show_unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatmapBackendDefinitionYaxis instantiates a new HeatmapBackendDefinitionYaxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatmapBackendDefinitionYaxis() *HeatmapBackendDefinitionYaxis {
	this := HeatmapBackendDefinitionYaxis{}
	return &this
}

// NewHeatmapBackendDefinitionYaxisWithDefaults instantiates a new HeatmapBackendDefinitionYaxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatmapBackendDefinitionYaxisWithDefaults() *HeatmapBackendDefinitionYaxis {
	this := HeatmapBackendDefinitionYaxis{}
	return &this
}

// GetIncludeZero returns the IncludeZero field value if set, zero value otherwise.
func (o *HeatmapBackendDefinitionYaxis) GetIncludeZero() bool {
	if o == nil || o.IncludeZero == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZero
}

// GetIncludeZeroOk returns a tuple with the IncludeZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendDefinitionYaxis) GetIncludeZeroOk() (*bool, bool) {
	if o == nil || o.IncludeZero == nil {
		return nil, false
	}
	return o.IncludeZero, true
}

// HasIncludeZero returns a boolean if a field has been set.
func (o *HeatmapBackendDefinitionYaxis) HasIncludeZero() bool {
	return o != nil && o.IncludeZero != nil
}

// SetIncludeZero gets a reference to the given bool and assigns it to the IncludeZero field.
func (o *HeatmapBackendDefinitionYaxis) SetIncludeZero(v bool) {
	o.IncludeZero = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *HeatmapBackendDefinitionYaxis) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendDefinitionYaxis) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *HeatmapBackendDefinitionYaxis) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *HeatmapBackendDefinitionYaxis) SetLabel(v string) {
	o.Label = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *HeatmapBackendDefinitionYaxis) GetMax() string {
	if o == nil || o.Max == nil {
		var ret string
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendDefinitionYaxis) GetMaxOk() (*string, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *HeatmapBackendDefinitionYaxis) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given string and assigns it to the Max field.
func (o *HeatmapBackendDefinitionYaxis) SetMax(v string) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *HeatmapBackendDefinitionYaxis) GetMin() string {
	if o == nil || o.Min == nil {
		var ret string
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendDefinitionYaxis) GetMinOk() (*string, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *HeatmapBackendDefinitionYaxis) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given string and assigns it to the Min field.
func (o *HeatmapBackendDefinitionYaxis) SetMin(v string) {
	o.Min = &v
}

// GetNumBuckets returns the NumBuckets field value if set, zero value otherwise.
func (o *HeatmapBackendDefinitionYaxis) GetNumBuckets() float64 {
	if o == nil || o.NumBuckets == nil {
		var ret float64
		return ret
	}
	return *o.NumBuckets
}

// GetNumBucketsOk returns a tuple with the NumBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendDefinitionYaxis) GetNumBucketsOk() (*float64, bool) {
	if o == nil || o.NumBuckets == nil {
		return nil, false
	}
	return o.NumBuckets, true
}

// HasNumBuckets returns a boolean if a field has been set.
func (o *HeatmapBackendDefinitionYaxis) HasNumBuckets() bool {
	return o != nil && o.NumBuckets != nil
}

// SetNumBuckets gets a reference to the given float64 and assigns it to the NumBuckets field.
func (o *HeatmapBackendDefinitionYaxis) SetNumBuckets(v float64) {
	o.NumBuckets = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *HeatmapBackendDefinitionYaxis) GetScale() ScaleType {
	if o == nil || o.Scale == nil {
		var ret ScaleType
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendDefinitionYaxis) GetScaleOk() (*ScaleType, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *HeatmapBackendDefinitionYaxis) HasScale() bool {
	return o != nil && o.Scale != nil
}

// SetScale gets a reference to the given ScaleType and assigns it to the Scale field.
func (o *HeatmapBackendDefinitionYaxis) SetScale(v ScaleType) {
	o.Scale = &v
}

// GetShowUnit returns the ShowUnit field value if set, zero value otherwise.
func (o *HeatmapBackendDefinitionYaxis) GetShowUnit() bool {
	if o == nil || o.ShowUnit == nil {
		var ret bool
		return ret
	}
	return *o.ShowUnit
}

// GetShowUnitOk returns a tuple with the ShowUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapBackendDefinitionYaxis) GetShowUnitOk() (*bool, bool) {
	if o == nil || o.ShowUnit == nil {
		return nil, false
	}
	return o.ShowUnit, true
}

// HasShowUnit returns a boolean if a field has been set.
func (o *HeatmapBackendDefinitionYaxis) HasShowUnit() bool {
	return o != nil && o.ShowUnit != nil
}

// SetShowUnit gets a reference to the given bool and assigns it to the ShowUnit field.
func (o *HeatmapBackendDefinitionYaxis) SetShowUnit(v bool) {
	o.ShowUnit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatmapBackendDefinitionYaxis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeZero != nil {
		toSerialize["include_zero"] = o.IncludeZero
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.NumBuckets != nil {
		toSerialize["num_buckets"] = o.NumBuckets
	}
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.ShowUnit != nil {
		toSerialize["show_unit"] = o.ShowUnit
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatmapBackendDefinitionYaxis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeZero *bool      `json:"include_zero,omitempty"`
		Label       *string    `json:"label,omitempty"`
		Max         *string    `json:"max,omitempty"`
		Min         *string    `json:"min,omitempty"`
		NumBuckets  *float64   `json:"num_buckets,omitempty"`
		Scale       *ScaleType `json:"scale,omitempty"`
		ShowUnit    *bool      `json:"show_unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.IncludeZero = all.IncludeZero
	o.Label = all.Label
	o.Max = all.Max
	o.Min = all.Min
	o.NumBuckets = all.NumBuckets
	if all.Scale != nil && !all.Scale.IsValid() {
		hasInvalidField = true
	} else {
		o.Scale = all.Scale
	}
	o.ShowUnit = all.ShowUnit

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
