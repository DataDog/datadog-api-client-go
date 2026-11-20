// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionXAxis Test-only generated schema for the dashboard versioning proof of concept.
type DistributionXAxis struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	IncludeZero *bool `json:"include_zero,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Max *AxisBound `json:"max,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Min *AxisBound `json:"min,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumBuckets *float64 `json:"num_buckets,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Scale *DistributionScaleType `json:"scale,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionXAxis instantiates a new DistributionXAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionXAxis() *DistributionXAxis {
	this := DistributionXAxis{}
	return &this
}

// NewDistributionXAxisWithDefaults instantiates a new DistributionXAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionXAxisWithDefaults() *DistributionXAxis {
	this := DistributionXAxis{}
	return &this
}

// GetIncludeZero returns the IncludeZero field value if set, zero value otherwise.
func (o *DistributionXAxis) GetIncludeZero() bool {
	if o == nil || o.IncludeZero == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZero
}

// GetIncludeZeroOk returns a tuple with the IncludeZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionXAxis) GetIncludeZeroOk() (*bool, bool) {
	if o == nil || o.IncludeZero == nil {
		return nil, false
	}
	return o.IncludeZero, true
}

// HasIncludeZero returns a boolean if a field has been set.
func (o *DistributionXAxis) HasIncludeZero() bool {
	return o != nil && o.IncludeZero != nil
}

// SetIncludeZero gets a reference to the given bool and assigns it to the IncludeZero field.
func (o *DistributionXAxis) SetIncludeZero(v bool) {
	o.IncludeZero = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *DistributionXAxis) GetMax() AxisBound {
	if o == nil || o.Max == nil {
		var ret AxisBound
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionXAxis) GetMaxOk() (*AxisBound, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *DistributionXAxis) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given AxisBound and assigns it to the Max field.
func (o *DistributionXAxis) SetMax(v AxisBound) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *DistributionXAxis) GetMin() AxisBound {
	if o == nil || o.Min == nil {
		var ret AxisBound
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionXAxis) GetMinOk() (*AxisBound, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *DistributionXAxis) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given AxisBound and assigns it to the Min field.
func (o *DistributionXAxis) SetMin(v AxisBound) {
	o.Min = &v
}

// GetNumBuckets returns the NumBuckets field value if set, zero value otherwise.
func (o *DistributionXAxis) GetNumBuckets() float64 {
	if o == nil || o.NumBuckets == nil {
		var ret float64
		return ret
	}
	return *o.NumBuckets
}

// GetNumBucketsOk returns a tuple with the NumBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionXAxis) GetNumBucketsOk() (*float64, bool) {
	if o == nil || o.NumBuckets == nil {
		return nil, false
	}
	return o.NumBuckets, true
}

// HasNumBuckets returns a boolean if a field has been set.
func (o *DistributionXAxis) HasNumBuckets() bool {
	return o != nil && o.NumBuckets != nil
}

// SetNumBuckets gets a reference to the given float64 and assigns it to the NumBuckets field.
func (o *DistributionXAxis) SetNumBuckets(v float64) {
	o.NumBuckets = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *DistributionXAxis) GetScale() DistributionScaleType {
	if o == nil || o.Scale == nil {
		var ret DistributionScaleType
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionXAxis) GetScaleOk() (*DistributionScaleType, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *DistributionXAxis) HasScale() bool {
	return o != nil && o.Scale != nil
}

// SetScale gets a reference to the given DistributionScaleType and assigns it to the Scale field.
func (o *DistributionXAxis) SetScale(v DistributionScaleType) {
	o.Scale = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionXAxis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeZero != nil {
		toSerialize["include_zero"] = o.IncludeZero
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionXAxis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeZero *bool                  `json:"include_zero,omitempty"`
		Max         *AxisBound             `json:"max,omitempty"`
		Min         *AxisBound             `json:"min,omitempty"`
		NumBuckets  *float64               `json:"num_buckets,omitempty"`
		Scale       *DistributionScaleType `json:"scale,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.IncludeZero = all.IncludeZero
	o.Max = all.Max
	o.Min = all.Min
	o.NumBuckets = all.NumBuckets
	if all.Scale != nil && !all.Scale.IsValid() {
		hasInvalidField = true
	} else {
		o.Scale = all.Scale
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
