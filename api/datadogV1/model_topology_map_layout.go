// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapLayout Algorithm that positions nodes and edges. The `hierarchical` layout places nodes in ranks so edges flow in one direction.
type TopologyMapLayout struct {
	// Direction of ranks. For example, `top-to-bottom` places first-rank nodes at the top and last-rank nodes at the bottom.
	Direction TopologyMapLayoutDirection `json:"direction"`
	// Type of layout algorithm.
	Type TopologyMapLayoutType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyMapLayout instantiates a new TopologyMapLayout object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyMapLayout(direction TopologyMapLayoutDirection, typeVar TopologyMapLayoutType) *TopologyMapLayout {
	this := TopologyMapLayout{}
	this.Direction = direction
	this.Type = typeVar
	return &this
}

// NewTopologyMapLayoutWithDefaults instantiates a new TopologyMapLayout object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyMapLayoutWithDefaults() *TopologyMapLayout {
	this := TopologyMapLayout{}
	var typeVar TopologyMapLayoutType = TOPOLOGYMAPLAYOUTTYPE_HIERARCHICAL
	this.Type = typeVar
	return &this
}

// GetDirection returns the Direction field value.
func (o *TopologyMapLayout) GetDirection() TopologyMapLayoutDirection {
	if o == nil {
		var ret TopologyMapLayoutDirection
		return ret
	}
	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *TopologyMapLayout) GetDirectionOk() (*TopologyMapLayoutDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value.
func (o *TopologyMapLayout) SetDirection(v TopologyMapLayoutDirection) {
	o.Direction = v
}

// GetType returns the Type field value.
func (o *TopologyMapLayout) GetType() TopologyMapLayoutType {
	if o == nil {
		var ret TopologyMapLayoutType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TopologyMapLayout) GetTypeOk() (*TopologyMapLayoutType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TopologyMapLayout) SetType(v TopologyMapLayoutType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyMapLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["direction"] = o.Direction
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TopologyMapLayout) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Direction *TopologyMapLayoutDirection `json:"direction"`
		Type      *TopologyMapLayoutType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Direction == nil {
		return fmt.Errorf("required field direction missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = *all.Direction
	}
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
