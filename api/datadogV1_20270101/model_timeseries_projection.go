// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesProjection Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesProjection struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimensions []TimeseriesProjectionDimensionsItems `json:"dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type TimeseriesProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTimeseriesProjection instantiates a new TimeseriesProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesProjection(dimensions []TimeseriesProjectionDimensionsItems, typeVar TimeseriesProjectionType) *TimeseriesProjection {
	this := TimeseriesProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewTimeseriesProjectionWithDefaults instantiates a new TimeseriesProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesProjectionWithDefaults() *TimeseriesProjection {
	this := TimeseriesProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *TimeseriesProjection) GetDimensions() []TimeseriesProjectionDimensionsItems {
	if o == nil {
		var ret []TimeseriesProjectionDimensionsItems
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *TimeseriesProjection) GetDimensionsOk() (*[]TimeseriesProjectionDimensionsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *TimeseriesProjection) SetDimensions(v []TimeseriesProjectionDimensionsItems) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *TimeseriesProjection) GetType() TimeseriesProjectionType {
	if o == nil {
		var ret TimeseriesProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimeseriesProjection) GetTypeOk() (*TimeseriesProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TimeseriesProjection) SetType(v TimeseriesProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]TimeseriesProjectionDimensionsItems `json:"dimensions"`
		Type       *TimeseriesProjectionType              `json:"type"`
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
