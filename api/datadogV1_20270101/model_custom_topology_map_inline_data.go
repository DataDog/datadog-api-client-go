// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineData Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineData struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Edges *CustomTopologyMapInlineEdgesData `json:"edges,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Nodes *CustomTopologyMapInlineNodesData `json:"nodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapInlineData instantiates a new CustomTopologyMapInlineData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapInlineData() *CustomTopologyMapInlineData {
	this := CustomTopologyMapInlineData{}
	return &this
}

// NewCustomTopologyMapInlineDataWithDefaults instantiates a new CustomTopologyMapInlineData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapInlineDataWithDefaults() *CustomTopologyMapInlineData {
	this := CustomTopologyMapInlineData{}
	return &this
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *CustomTopologyMapInlineData) GetEdges() CustomTopologyMapInlineEdgesData {
	if o == nil || o.Edges == nil {
		var ret CustomTopologyMapInlineEdgesData
		return ret
	}
	return *o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineData) GetEdgesOk() (*CustomTopologyMapInlineEdgesData, bool) {
	if o == nil || o.Edges == nil {
		return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *CustomTopologyMapInlineData) HasEdges() bool {
	return o != nil && o.Edges != nil
}

// SetEdges gets a reference to the given CustomTopologyMapInlineEdgesData and assigns it to the Edges field.
func (o *CustomTopologyMapInlineData) SetEdges(v CustomTopologyMapInlineEdgesData) {
	o.Edges = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *CustomTopologyMapInlineData) GetNodes() CustomTopologyMapInlineNodesData {
	if o == nil || o.Nodes == nil {
		var ret CustomTopologyMapInlineNodesData
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineData) GetNodesOk() (*CustomTopologyMapInlineNodesData, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *CustomTopologyMapInlineData) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given CustomTopologyMapInlineNodesData and assigns it to the Nodes field.
func (o *CustomTopologyMapInlineData) SetNodes(v CustomTopologyMapInlineNodesData) {
	o.Nodes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapInlineData) MarshalJSON() ([]byte, error) {
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
func (o *CustomTopologyMapInlineData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Edges *CustomTopologyMapInlineEdgesData `json:"edges,omitempty"`
		Nodes *CustomTopologyMapInlineNodesData `json:"nodes,omitempty"`
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
