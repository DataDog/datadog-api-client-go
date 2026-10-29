// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyRequestCustomProjection Mapping from result columns to node fields. Each result row feeds one node.
type TopologyRequestCustomProjection struct {
	// List of mappings from a result column to a node field. Each `dimension` must be unique.
	Dimensions []TopologyRequestCustomProjectionDimension `json:"dimensions"`
	// Projection type. Each result row is a node.
	Type TopologyRequestCustomProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyRequestCustomProjection instantiates a new TopologyRequestCustomProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyRequestCustomProjection(dimensions []TopologyRequestCustomProjectionDimension, typeVar TopologyRequestCustomProjectionType) *TopologyRequestCustomProjection {
	this := TopologyRequestCustomProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewTopologyRequestCustomProjectionWithDefaults instantiates a new TopologyRequestCustomProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyRequestCustomProjectionWithDefaults() *TopologyRequestCustomProjection {
	this := TopologyRequestCustomProjection{}
	var typeVar TopologyRequestCustomProjectionType = TOPOLOGYREQUESTCUSTOMPROJECTIONTYPE_CUSTOM_TOPOLOGY_MAP_NODE
	this.Type = typeVar
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *TopologyRequestCustomProjection) GetDimensions() []TopologyRequestCustomProjectionDimension {
	if o == nil {
		var ret []TopologyRequestCustomProjectionDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustomProjection) GetDimensionsOk() (*[]TopologyRequestCustomProjectionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *TopologyRequestCustomProjection) SetDimensions(v []TopologyRequestCustomProjectionDimension) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *TopologyRequestCustomProjection) GetType() TopologyRequestCustomProjectionType {
	if o == nil {
		var ret TopologyRequestCustomProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TopologyRequestCustomProjection) GetTypeOk() (*TopologyRequestCustomProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TopologyRequestCustomProjection) SetType(v TopologyRequestCustomProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyRequestCustomProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyRequestCustomProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]TopologyRequestCustomProjectionDimension `json:"dimensions"`
		Type       *TopologyRequestCustomProjectionType        `json:"type"`
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
