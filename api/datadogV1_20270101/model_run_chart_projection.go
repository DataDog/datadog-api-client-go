// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RunChartProjection Test-only generated schema for the dashboard versioning proof of concept.
type RunChartProjection struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimensions []RunChartProjectionDimension `json:"dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type RunChartProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRunChartProjection instantiates a new RunChartProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRunChartProjection(dimensions []RunChartProjectionDimension, typeVar RunChartProjectionType) *RunChartProjection {
	this := RunChartProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewRunChartProjectionWithDefaults instantiates a new RunChartProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRunChartProjectionWithDefaults() *RunChartProjection {
	this := RunChartProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *RunChartProjection) GetDimensions() []RunChartProjectionDimension {
	if o == nil {
		var ret []RunChartProjectionDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *RunChartProjection) GetDimensionsOk() (*[]RunChartProjectionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *RunChartProjection) SetDimensions(v []RunChartProjectionDimension) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *RunChartProjection) GetType() RunChartProjectionType {
	if o == nil {
		var ret RunChartProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RunChartProjection) GetTypeOk() (*RunChartProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RunChartProjection) SetType(v RunChartProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RunChartProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RunChartProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]RunChartProjectionDimension `json:"dimensions"`
		Type       *RunChartProjectionType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dimensions == nil {
		return fmt.Errorf("required field dimensions missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Dimensions = *all.Dimensions
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
