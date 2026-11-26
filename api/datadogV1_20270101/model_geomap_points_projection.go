// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapPointsProjection Test-only generated schema for the dashboard versioning proof of concept.
type GeomapPointsProjection struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimensions []GeomapPointsProjectionDimensionsItems `json:"dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type GeomapPointsProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapPointsProjection instantiates a new GeomapPointsProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapPointsProjection(dimensions []GeomapPointsProjectionDimensionsItems, typeVar GeomapPointsProjectionType) *GeomapPointsProjection {
	this := GeomapPointsProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewGeomapPointsProjectionWithDefaults instantiates a new GeomapPointsProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapPointsProjectionWithDefaults() *GeomapPointsProjection {
	this := GeomapPointsProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *GeomapPointsProjection) GetDimensions() []GeomapPointsProjectionDimensionsItems {
	if o == nil {
		var ret []GeomapPointsProjectionDimensionsItems
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *GeomapPointsProjection) GetDimensionsOk() (*[]GeomapPointsProjectionDimensionsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *GeomapPointsProjection) SetDimensions(v []GeomapPointsProjectionDimensionsItems) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *GeomapPointsProjection) GetType() GeomapPointsProjectionType {
	if o == nil {
		var ret GeomapPointsProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeomapPointsProjection) GetTypeOk() (*GeomapPointsProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GeomapPointsProjection) SetType(v GeomapPointsProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapPointsProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapPointsProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]GeomapPointsProjectionDimensionsItems `json:"dimensions"`
		Type       *GeomapPointsProjectionType              `json:"type"`
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
