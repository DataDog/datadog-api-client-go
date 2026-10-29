// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineEdgeColumn Column of edge endpoint node IDs.
type TopologyMapInlineEdgeColumn struct {
	// Edge endpoint the column holds.
	Id TopologyMapInlineEdgeColumnId `json:"id"`
	// Column of strings.
	Kind TopologyMapInlineColumnKindString `json:"kind"`
	// Node IDs.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyMapInlineEdgeColumn instantiates a new TopologyMapInlineEdgeColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyMapInlineEdgeColumn(id TopologyMapInlineEdgeColumnId, kind TopologyMapInlineColumnKindString, values []string) *TopologyMapInlineEdgeColumn {
	this := TopologyMapInlineEdgeColumn{}
	this.Id = id
	this.Kind = kind
	this.Values = values
	return &this
}

// NewTopologyMapInlineEdgeColumnWithDefaults instantiates a new TopologyMapInlineEdgeColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyMapInlineEdgeColumnWithDefaults() *TopologyMapInlineEdgeColumn {
	this := TopologyMapInlineEdgeColumn{}
	var kind TopologyMapInlineColumnKindString = TOPOLOGYMAPINLINECOLUMNKINDSTRING_STRING
	this.Kind = kind
	return &this
}

// GetId returns the Id field value.
func (o *TopologyMapInlineEdgeColumn) GetId() TopologyMapInlineEdgeColumnId {
	if o == nil {
		var ret TopologyMapInlineEdgeColumnId
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineEdgeColumn) GetIdOk() (*TopologyMapInlineEdgeColumnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TopologyMapInlineEdgeColumn) SetId(v TopologyMapInlineEdgeColumnId) {
	o.Id = v
}

// GetKind returns the Kind field value.
func (o *TopologyMapInlineEdgeColumn) GetKind() TopologyMapInlineColumnKindString {
	if o == nil {
		var ret TopologyMapInlineColumnKindString
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineEdgeColumn) GetKindOk() (*TopologyMapInlineColumnKindString, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *TopologyMapInlineEdgeColumn) SetKind(v TopologyMapInlineColumnKindString) {
	o.Kind = v
}

// GetValues returns the Values field value.
func (o *TopologyMapInlineEdgeColumn) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineEdgeColumn) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *TopologyMapInlineEdgeColumn) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyMapInlineEdgeColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	toSerialize["values"] = o.Values
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyMapInlineEdgeColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *TopologyMapInlineEdgeColumnId     `json:"id"`
		Kind   *TopologyMapInlineColumnKindString `json:"kind"`
		Values *[]string                          `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}

	hasInvalidField := false
	if !all.Id.IsValid() {
		hasInvalidField = true
	} else {
		o.Id = *all.Id
	}
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	o.Values = *all.Values

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
