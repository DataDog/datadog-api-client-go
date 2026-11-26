// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// BackendDistributionYAxis Test-only generated schema for the dashboard versioning proof of concept.
type BackendDistributionYAxis struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	IncludeZero *bool `json:"include_zero,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Max *datadogV1.BackendDistributionYAxis20270101Max `json:"max,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Min *datadogV1.BackendDistributionYAxis20270101Min `json:"min,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Scale *BackendDistributionYAxisScale `json:"scale,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowUnit *bool `json:"show_unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewBackendDistributionYAxis instantiates a new BackendDistributionYAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackendDistributionYAxis() *BackendDistributionYAxis {
	this := BackendDistributionYAxis{}
	return &this
}

// NewBackendDistributionYAxisWithDefaults instantiates a new BackendDistributionYAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackendDistributionYAxisWithDefaults() *BackendDistributionYAxis {
	this := BackendDistributionYAxis{}
	return &this
}

// GetIncludeZero returns the IncludeZero field value if set, zero value otherwise.
func (o *BackendDistributionYAxis) GetIncludeZero() bool {
	if o == nil || o.IncludeZero == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZero
}

// GetIncludeZeroOk returns a tuple with the IncludeZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendDistributionYAxis) GetIncludeZeroOk() (*bool, bool) {
	if o == nil || o.IncludeZero == nil {
		return nil, false
	}
	return o.IncludeZero, true
}

// HasIncludeZero returns a boolean if a field has been set.
func (o *BackendDistributionYAxis) HasIncludeZero() bool {
	return o != nil && o.IncludeZero != nil
}

// SetIncludeZero gets a reference to the given bool and assigns it to the IncludeZero field.
func (o *BackendDistributionYAxis) SetIncludeZero(v bool) {
	o.IncludeZero = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BackendDistributionYAxis) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendDistributionYAxis) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BackendDistributionYAxis) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BackendDistributionYAxis) SetLabel(v string) {
	o.Label = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *BackendDistributionYAxis) GetMax() datadogV1.BackendDistributionYAxis20270101Max {
	if o == nil || o.Max == nil {
		var ret datadogV1.BackendDistributionYAxis20270101Max
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendDistributionYAxis) GetMaxOk() (*datadogV1.BackendDistributionYAxis20270101Max, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *BackendDistributionYAxis) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given datadogV1.BackendDistributionYAxis20270101Max and assigns it to the Max field.
func (o *BackendDistributionYAxis) SetMax(v datadogV1.BackendDistributionYAxis20270101Max) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *BackendDistributionYAxis) GetMin() datadogV1.BackendDistributionYAxis20270101Min {
	if o == nil || o.Min == nil {
		var ret datadogV1.BackendDistributionYAxis20270101Min
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendDistributionYAxis) GetMinOk() (*datadogV1.BackendDistributionYAxis20270101Min, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *BackendDistributionYAxis) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given datadogV1.BackendDistributionYAxis20270101Min and assigns it to the Min field.
func (o *BackendDistributionYAxis) SetMin(v datadogV1.BackendDistributionYAxis20270101Min) {
	o.Min = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *BackendDistributionYAxis) GetScale() BackendDistributionYAxisScale {
	if o == nil || o.Scale == nil {
		var ret BackendDistributionYAxisScale
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendDistributionYAxis) GetScaleOk() (*BackendDistributionYAxisScale, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *BackendDistributionYAxis) HasScale() bool {
	return o != nil && o.Scale != nil
}

// SetScale gets a reference to the given BackendDistributionYAxisScale and assigns it to the Scale field.
func (o *BackendDistributionYAxis) SetScale(v BackendDistributionYAxisScale) {
	o.Scale = &v
}

// GetShowUnit returns the ShowUnit field value if set, zero value otherwise.
func (o *BackendDistributionYAxis) GetShowUnit() bool {
	if o == nil || o.ShowUnit == nil {
		var ret bool
		return ret
	}
	return *o.ShowUnit
}

// GetShowUnitOk returns a tuple with the ShowUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendDistributionYAxis) GetShowUnitOk() (*bool, bool) {
	if o == nil || o.ShowUnit == nil {
		return nil, false
	}
	return o.ShowUnit, true
}

// HasShowUnit returns a boolean if a field has been set.
func (o *BackendDistributionYAxis) HasShowUnit() bool {
	return o != nil && o.ShowUnit != nil
}

// SetShowUnit gets a reference to the given bool and assigns it to the ShowUnit field.
func (o *BackendDistributionYAxis) SetShowUnit(v bool) {
	o.ShowUnit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackendDistributionYAxis) MarshalJSON() ([]byte, error) {
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
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.ShowUnit != nil {
		toSerialize["show_unit"] = o.ShowUnit
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackendDistributionYAxis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeZero *bool                                          `json:"include_zero,omitempty"`
		Label       *string                                        `json:"label,omitempty"`
		Max         *datadogV1.BackendDistributionYAxis20270101Max `json:"max,omitempty"`
		Min         *datadogV1.BackendDistributionYAxis20270101Min `json:"min,omitempty"`
		Scale       *BackendDistributionYAxisScale                 `json:"scale,omitempty"`
		ShowUnit    *bool                                          `json:"show_unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.IncludeZero = all.IncludeZero
	o.Label = all.Label
	o.Max = all.Max
	o.Min = all.Min
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
