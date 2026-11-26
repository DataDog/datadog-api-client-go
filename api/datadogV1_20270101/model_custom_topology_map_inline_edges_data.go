// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineEdgesData Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineEdgesData struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []CustomTopologyMapInlineEdgeColumn `json:"columns"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapInlineEdgesData instantiates a new CustomTopologyMapInlineEdgesData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapInlineEdgesData(columns []CustomTopologyMapInlineEdgeColumn) *CustomTopologyMapInlineEdgesData {
	this := CustomTopologyMapInlineEdgesData{}
	this.Columns = columns
	return &this
}

// NewCustomTopologyMapInlineEdgesDataWithDefaults instantiates a new CustomTopologyMapInlineEdgesData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapInlineEdgesDataWithDefaults() *CustomTopologyMapInlineEdgesData {
	this := CustomTopologyMapInlineEdgesData{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CustomTopologyMapInlineEdgesData) GetColumns() []CustomTopologyMapInlineEdgeColumn {
	if o == nil {
		var ret []CustomTopologyMapInlineEdgeColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineEdgesData) GetColumnsOk() (*[]CustomTopologyMapInlineEdgeColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CustomTopologyMapInlineEdgesData) SetColumns(v []CustomTopologyMapInlineEdgeColumn) {
	o.Columns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapInlineEdgesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomTopologyMapInlineEdgesData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns *[]CustomTopologyMapInlineEdgeColumn `json:"columns"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	o.Columns = *all.Columns

	return nil
}
