// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapLayout Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapLayout struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Direction CustomTopologyMapLayoutDirection `json:"direction"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type CustomTopologyMapLayoutType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapLayout instantiates a new CustomTopologyMapLayout object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapLayout(direction CustomTopologyMapLayoutDirection, typeVar CustomTopologyMapLayoutType) *CustomTopologyMapLayout {
	this := CustomTopologyMapLayout{}
	this.Direction = direction
	this.Type = typeVar
	return &this
}

// NewCustomTopologyMapLayoutWithDefaults instantiates a new CustomTopologyMapLayout object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapLayoutWithDefaults() *CustomTopologyMapLayout {
	this := CustomTopologyMapLayout{}
	return &this
}

// GetDirection returns the Direction field value.
func (o *CustomTopologyMapLayout) GetDirection() CustomTopologyMapLayoutDirection {
	if o == nil {
		var ret CustomTopologyMapLayoutDirection
		return ret
	}
	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapLayout) GetDirectionOk() (*CustomTopologyMapLayoutDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value.
func (o *CustomTopologyMapLayout) SetDirection(v CustomTopologyMapLayoutDirection) {
	o.Direction = v
}

// GetType returns the Type field value.
func (o *CustomTopologyMapLayout) GetType() CustomTopologyMapLayoutType {
	if o == nil {
		var ret CustomTopologyMapLayoutType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapLayout) GetTypeOk() (*CustomTopologyMapLayoutType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomTopologyMapLayout) SetType(v CustomTopologyMapLayoutType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["direction"] = o.Direction
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomTopologyMapLayout) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Direction *CustomTopologyMapLayoutDirection `json:"direction"`
		Type      *CustomTopologyMapLayoutType      `json:"type"`
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
