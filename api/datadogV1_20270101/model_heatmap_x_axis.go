// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HeatmapXAxis Test-only generated schema for the dashboard versioning proof of concept.
type HeatmapXAxis struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumBuckets *float64 `json:"num_buckets,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewHeatmapXAxis instantiates a new HeatmapXAxis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHeatmapXAxis() *HeatmapXAxis {
	this := HeatmapXAxis{}
	return &this
}

// NewHeatmapXAxisWithDefaults instantiates a new HeatmapXAxis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHeatmapXAxisWithDefaults() *HeatmapXAxis {
	this := HeatmapXAxis{}
	return &this
}

// GetNumBuckets returns the NumBuckets field value if set, zero value otherwise.
func (o *HeatmapXAxis) GetNumBuckets() float64 {
	if o == nil || o.NumBuckets == nil {
		var ret float64
		return ret
	}
	return *o.NumBuckets
}

// GetNumBucketsOk returns a tuple with the NumBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeatmapXAxis) GetNumBucketsOk() (*float64, bool) {
	if o == nil || o.NumBuckets == nil {
		return nil, false
	}
	return o.NumBuckets, true
}

// HasNumBuckets returns a boolean if a field has been set.
func (o *HeatmapXAxis) HasNumBuckets() bool {
	return o != nil && o.NumBuckets != nil
}

// SetNumBuckets gets a reference to the given float64 and assigns it to the NumBuckets field.
func (o *HeatmapXAxis) SetNumBuckets(v float64) {
	o.NumBuckets = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HeatmapXAxis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NumBuckets != nil {
		toSerialize["num_buckets"] = o.NumBuckets
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HeatmapXAxis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumBuckets *float64 `json:"num_buckets,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.NumBuckets = all.NumBuckets

	return nil
}
