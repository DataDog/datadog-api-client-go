// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapProjection Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapProjection struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimensions []CustomTopologyMapProjectionDimensionsItems `json:"dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CustomTopologyMapProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapProjection instantiates a new CustomTopologyMapProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapProjection(dimensions []CustomTopologyMapProjectionDimensionsItems, typeVar CustomTopologyMapProjectionType) *CustomTopologyMapProjection {
	this := CustomTopologyMapProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewCustomTopologyMapProjectionWithDefaults instantiates a new CustomTopologyMapProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapProjectionWithDefaults() *CustomTopologyMapProjection {
	this := CustomTopologyMapProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *CustomTopologyMapProjection) GetDimensions() []CustomTopologyMapProjectionDimensionsItems {
	if o == nil {
		var ret []CustomTopologyMapProjectionDimensionsItems
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapProjection) GetDimensionsOk() (*[]CustomTopologyMapProjectionDimensionsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *CustomTopologyMapProjection) SetDimensions(v []CustomTopologyMapProjectionDimensionsItems) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *CustomTopologyMapProjection) GetType() CustomTopologyMapProjectionType {
	if o == nil {
		var ret CustomTopologyMapProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapProjection) GetTypeOk() (*CustomTopologyMapProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomTopologyMapProjection) SetType(v CustomTopologyMapProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomTopologyMapProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]CustomTopologyMapProjectionDimensionsItems `json:"dimensions"`
		Type       *CustomTopologyMapProjectionType              `json:"type"`
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
