// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatMapWidgetYAxis Y Axis controls for the heat map widget.
type HeatMapWidgetYAxis struct {
	// Set to `true` to include zero.
	IncludeZero *bool `json:"include_zero,omitempty"`
	// The label of the axis to display on the graph. Only usable on Scatterplot Widgets.
	Label *string `json:"label,omitempty"`
	// Specifies maximum numeric value to show on the axis. Defaults to `auto`.
	Max *string `json:"max,omitempty"`
	// Specifies minimum numeric value to show on the axis. Defaults to `auto`.
	Min *string `json:"min,omitempty"`
	// Number of value buckets to target, aka the resolution of the value bins.
	NumBuckets *int64 `json:"num_buckets,omitempty"`
	// Specifies the scale type. Possible values are `linear`, `log`, `sqrt`, and `pow##` (for example `pow2` or `pow0.5`).
	Scale *string `json:"scale,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHeatMapWidgetYAxis instantiates a new HeatMapWidgetYAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatMapWidgetYAxis() *HeatMapWidgetYAxis {
	this := HeatMapWidgetYAxis{}
	var max string = "auto"
	this.Max = &max
	var min string = "auto"
	this.Min = &min
	var scale string = "linear"
	this.Scale = &scale
	return &this
}

// NewHeatMapWidgetYAxisWithDefaults instantiates a new HeatMapWidgetYAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatMapWidgetYAxisWithDefaults() *HeatMapWidgetYAxis {
	this := HeatMapWidgetYAxis{}
	var max string = "auto"
	this.Max = &max
	var min string = "auto"
	this.Min = &min
	var scale string = "linear"
	this.Scale = &scale
	return &this
}

// GetIncludeZero returns the IncludeZero field value if set, zero value otherwise.
func (o *HeatMapWidgetYAxis) GetIncludeZero() bool {
	if o == nil || o.IncludeZero == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZero
}

// GetIncludeZeroOk returns a tuple with the IncludeZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatMapWidgetYAxis) GetIncludeZeroOk() (*bool, bool) {
	if o == nil || o.IncludeZero == nil {
		return nil, false
	}
	return o.IncludeZero, true
}

// HasIncludeZero returns a boolean if a field has been set.
func (o *HeatMapWidgetYAxis) HasIncludeZero() bool {
	return o != nil && o.IncludeZero != nil
}

// SetIncludeZero gets a reference to the given bool and assigns it to the IncludeZero field.
func (o *HeatMapWidgetYAxis) SetIncludeZero(v bool) {
	o.IncludeZero = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *HeatMapWidgetYAxis) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatMapWidgetYAxis) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *HeatMapWidgetYAxis) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *HeatMapWidgetYAxis) SetLabel(v string) {
	o.Label = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *HeatMapWidgetYAxis) GetMax() string {
	if o == nil || o.Max == nil {
		var ret string
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatMapWidgetYAxis) GetMaxOk() (*string, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *HeatMapWidgetYAxis) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given string and assigns it to the Max field.
func (o *HeatMapWidgetYAxis) SetMax(v string) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *HeatMapWidgetYAxis) GetMin() string {
	if o == nil || o.Min == nil {
		var ret string
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatMapWidgetYAxis) GetMinOk() (*string, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *HeatMapWidgetYAxis) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given string and assigns it to the Min field.
func (o *HeatMapWidgetYAxis) SetMin(v string) {
	o.Min = &v
}

// GetNumBuckets returns the NumBuckets field value if set, zero value otherwise.
func (o *HeatMapWidgetYAxis) GetNumBuckets() int64 {
	if o == nil || o.NumBuckets == nil {
		var ret int64
		return ret
	}
	return *o.NumBuckets
}

// GetNumBucketsOk returns a tuple with the NumBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatMapWidgetYAxis) GetNumBucketsOk() (*int64, bool) {
	if o == nil || o.NumBuckets == nil {
		return nil, false
	}
	return o.NumBuckets, true
}

// HasNumBuckets returns a boolean if a field has been set.
func (o *HeatMapWidgetYAxis) HasNumBuckets() bool {
	return o != nil && o.NumBuckets != nil
}

// SetNumBuckets gets a reference to the given int64 and assigns it to the NumBuckets field.
func (o *HeatMapWidgetYAxis) SetNumBuckets(v int64) {
	o.NumBuckets = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *HeatMapWidgetYAxis) GetScale() string {
	if o == nil || o.Scale == nil {
		var ret string
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatMapWidgetYAxis) GetScaleOk() (*string, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *HeatMapWidgetYAxis) HasScale() bool {
	return o != nil && o.Scale != nil
}

// SetScale gets a reference to the given string and assigns it to the Scale field.
func (o *HeatMapWidgetYAxis) SetScale(v string) {
	o.Scale = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatMapWidgetYAxis) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatMapWidgetYAxis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeZero *bool   `json:"include_zero,omitempty"`
		Label       *string `json:"label,omitempty"`
		Max         *string `json:"max,omitempty"`
		Min         *string `json:"min,omitempty"`
		NumBuckets  *int64  `json:"num_buckets,omitempty"`
		Scale       *string `json:"scale,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_zero", "label", "max", "min", "num_buckets", "scale"})
	} else {
		return err
	}
	o.IncludeZero = all.IncludeZero
	o.Label = all.Label
	o.Max = all.Max
	o.Min = all.Min
	o.NumBuckets = all.NumBuckets
	o.Scale = all.Scale

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
