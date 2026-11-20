// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeIdColumn Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeIdColumn struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Id CustomTopologyMapInlineNodeIdColumnId `json:"id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Kind CustomTopologyMapInlineNodeIdColumnKind `json:"kind"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapInlineNodeIdColumn instantiates a new CustomTopologyMapInlineNodeIdColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapInlineNodeIdColumn(id CustomTopologyMapInlineNodeIdColumnId, kind CustomTopologyMapInlineNodeIdColumnKind, values []string) *CustomTopologyMapInlineNodeIdColumn {
	this := CustomTopologyMapInlineNodeIdColumn{}
	this.Id = id
	this.Kind = kind
	this.Values = values
	return &this
}

// NewCustomTopologyMapInlineNodeIdColumnWithDefaults instantiates a new CustomTopologyMapInlineNodeIdColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapInlineNodeIdColumnWithDefaults() *CustomTopologyMapInlineNodeIdColumn {
	this := CustomTopologyMapInlineNodeIdColumn{}
	return &this
}

// GetId returns the Id field value.
func (o *CustomTopologyMapInlineNodeIdColumn) GetId() CustomTopologyMapInlineNodeIdColumnId {
	if o == nil {
		var ret CustomTopologyMapInlineNodeIdColumnId
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeIdColumn) GetIdOk() (*CustomTopologyMapInlineNodeIdColumnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomTopologyMapInlineNodeIdColumn) SetId(v CustomTopologyMapInlineNodeIdColumnId) {
	o.Id = v
}

// GetKind returns the Kind field value.
func (o *CustomTopologyMapInlineNodeIdColumn) GetKind() CustomTopologyMapInlineNodeIdColumnKind {
	if o == nil {
		var ret CustomTopologyMapInlineNodeIdColumnKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeIdColumn) GetKindOk() (*CustomTopologyMapInlineNodeIdColumnKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *CustomTopologyMapInlineNodeIdColumn) SetKind(v CustomTopologyMapInlineNodeIdColumnKind) {
	o.Kind = v
}

// GetValues returns the Values field value.
func (o *CustomTopologyMapInlineNodeIdColumn) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeIdColumn) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *CustomTopologyMapInlineNodeIdColumn) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapInlineNodeIdColumn) MarshalJSON() ([]byte, error) {
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
func (o *CustomTopologyMapInlineNodeIdColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *CustomTopologyMapInlineNodeIdColumnId   `json:"id"`
		Kind   *CustomTopologyMapInlineNodeIdColumnKind `json:"kind"`
		Values *[]string                                `json:"values"`
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
