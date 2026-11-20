// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapProjectionDimensionsItems Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapProjectionDimensionsItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimension CustomTopologyMapNodeDimension `json:"dimension"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapProjectionDimensionsItems instantiates a new CustomTopologyMapProjectionDimensionsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapProjectionDimensionsItems(column string, dimension CustomTopologyMapNodeDimension) *CustomTopologyMapProjectionDimensionsItems {
	this := CustomTopologyMapProjectionDimensionsItems{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewCustomTopologyMapProjectionDimensionsItemsWithDefaults instantiates a new CustomTopologyMapProjectionDimensionsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapProjectionDimensionsItemsWithDefaults() *CustomTopologyMapProjectionDimensionsItems {
	this := CustomTopologyMapProjectionDimensionsItems{}
	return &this
}

// GetColumn returns the Column field value.
func (o *CustomTopologyMapProjectionDimensionsItems) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapProjectionDimensionsItems) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *CustomTopologyMapProjectionDimensionsItems) SetColumn(v string) {
	o.Column = v
}

// GetDimension returns the Dimension field value.
func (o *CustomTopologyMapProjectionDimensionsItems) GetDimension() CustomTopologyMapNodeDimension {
	if o == nil {
		var ret CustomTopologyMapNodeDimension
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapProjectionDimensionsItems) GetDimensionOk() (*CustomTopologyMapNodeDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *CustomTopologyMapProjectionDimensionsItems) SetDimension(v CustomTopologyMapNodeDimension) {
	o.Dimension = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapProjectionDimensionsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["dimension"] = o.Dimension
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomTopologyMapProjectionDimensionsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column    *string                         `json:"column"`
		Dimension *CustomTopologyMapNodeDimension `json:"dimension"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Dimension == nil {
		return fmt.Errorf("required field dimension missing")
	}
	o.Column = *all.Column
	o.Dimension = *all.Dimension

	return nil
}
