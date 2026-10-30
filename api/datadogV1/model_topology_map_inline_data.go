// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineData Static nodes and edges, stored in columns. Within a table, every column must have the same number of values and a unique `id`.
type TopologyMapInlineData struct {
	// Table of edges, one row per edge.
	Edges *TopologyMapInlineEdgesData `json:"edges,omitempty"`
	// Table of nodes, one row per node.
	Nodes *TopologyMapInlineNodesData `json:"nodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyMapInlineData instantiates a new TopologyMapInlineData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyMapInlineData() *TopologyMapInlineData {
	this := TopologyMapInlineData{}
	return &this
}

// NewTopologyMapInlineDataWithDefaults instantiates a new TopologyMapInlineData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyMapInlineDataWithDefaults() *TopologyMapInlineData {
	this := TopologyMapInlineData{}
	return &this
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *TopologyMapInlineData) GetEdges() TopologyMapInlineEdgesData {
	if o == nil || o.Edges == nil {
		var ret TopologyMapInlineEdgesData
		return ret
	}
	return *o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineData) GetEdgesOk() (*TopologyMapInlineEdgesData, bool) {
	if o == nil || o.Edges == nil {
		return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *TopologyMapInlineData) HasEdges() bool {
	return o != nil && o.Edges != nil
}

// SetEdges gets a reference to the given TopologyMapInlineEdgesData and assigns it to the Edges field.
func (o *TopologyMapInlineData) SetEdges(v TopologyMapInlineEdgesData) {
	o.Edges = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *TopologyMapInlineData) GetNodes() TopologyMapInlineNodesData {
	if o == nil || o.Nodes == nil {
		var ret TopologyMapInlineNodesData
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineData) GetNodesOk() (*TopologyMapInlineNodesData, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *TopologyMapInlineData) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given TopologyMapInlineNodesData and assigns it to the Nodes field.
func (o *TopologyMapInlineData) SetNodes(v TopologyMapInlineNodesData) {
	o.Nodes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyMapInlineData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Edges != nil {
		toSerialize["edges"] = o.Edges
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyMapInlineData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Edges *TopologyMapInlineEdgesData `json:"edges,omitempty"`
		Nodes *TopologyMapInlineNodesData `json:"nodes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Edges != nil && all.Edges.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Edges = all.Edges
	if all.Nodes != nil && all.Nodes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Nodes = all.Nodes

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
