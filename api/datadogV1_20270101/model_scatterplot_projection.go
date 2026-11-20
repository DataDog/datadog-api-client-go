// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotProjection Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotProjection struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimensions []ScatterplotProjectionDimensionsItems `json:"dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type ScatterplotProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotProjection instantiates a new ScatterplotProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotProjection(dimensions []ScatterplotProjectionDimensionsItems, typeVar ScatterplotProjectionType) *ScatterplotProjection {
	this := ScatterplotProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewScatterplotProjectionWithDefaults instantiates a new ScatterplotProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotProjectionWithDefaults() *ScatterplotProjection {
	this := ScatterplotProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *ScatterplotProjection) GetDimensions() []ScatterplotProjectionDimensionsItems {
	if o == nil {
		var ret []ScatterplotProjectionDimensionsItems
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *ScatterplotProjection) GetDimensionsOk() (*[]ScatterplotProjectionDimensionsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *ScatterplotProjection) SetDimensions(v []ScatterplotProjectionDimensionsItems) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *ScatterplotProjection) GetType() ScatterplotProjectionType {
	if o == nil {
		var ret ScatterplotProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScatterplotProjection) GetTypeOk() (*ScatterplotProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScatterplotProjection) SetType(v ScatterplotProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]ScatterplotProjectionDimensionsItems `json:"dimensions"`
		Type       *ScatterplotProjectionType              `json:"type"`
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
