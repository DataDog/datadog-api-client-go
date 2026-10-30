// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyRequestCustomProjectionDimension Mapping from one result column to one node field.
type TopologyRequestCustomProjectionDimension struct {
	// Column name from the result. Must be a group-by tag key or a formula.
	Column string `json:"column"`
	// Node field that receives the column's value: `id`, `title`, `subtitle`, or `values[N]` for the value at position N. When `id` is not set, each node's ID is derived from the group-by tags (for example, `env:prod,service:web`). When `title` is not set, the title is the node's ID. When several rows set the same non-`id` field on a node, the last non-null value wins.
	Dimension string `json:"dimension"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyRequestCustomProjectionDimension instantiates a new TopologyRequestCustomProjectionDimension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyRequestCustomProjectionDimension(column string, dimension string) *TopologyRequestCustomProjectionDimension {
	this := TopologyRequestCustomProjectionDimension{}
	this.Column = column
	this.Dimension = dimension
	return &this
}

// NewTopologyRequestCustomProjectionDimensionWithDefaults instantiates a new TopologyRequestCustomProjectionDimension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyRequestCustomProjectionDimensionWithDefaults() *TopologyRequestCustomProjectionDimension {
	this := TopologyRequestCustomProjectionDimension{}
	return &this
}

// GetColumn returns the Column field value.
func (o *TopologyRequestCustomProjectionDimension) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustomProjectionDimension) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *TopologyRequestCustomProjectionDimension) SetColumn(v string) {
	o.Column = v
}

// GetDimension returns the Dimension field value.
func (o *TopologyRequestCustomProjectionDimension) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustomProjectionDimension) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value.
func (o *TopologyRequestCustomProjectionDimension) SetDimension(v string) {
	o.Dimension = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyRequestCustomProjectionDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["dimension"] = o.Dimension
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyRequestCustomProjectionDimension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column    *string `json:"column"`
		Dimension *string `json:"dimension"`
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
