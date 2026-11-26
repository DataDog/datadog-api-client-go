// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeStringColumn Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeStringColumn struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Id CustomTopologyMapInlineNodeStringColumnId `json:"id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Kind CustomTopologyMapInlineNodeStringColumnKind `json:"kind"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Values []*string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapInlineNodeStringColumn instantiates a new CustomTopologyMapInlineNodeStringColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapInlineNodeStringColumn(id CustomTopologyMapInlineNodeStringColumnId, kind CustomTopologyMapInlineNodeStringColumnKind, values []*string) *CustomTopologyMapInlineNodeStringColumn {
	this := CustomTopologyMapInlineNodeStringColumn{}
	this.Id = id
	this.Kind = kind
	this.Values = values
	return &this
}

// NewCustomTopologyMapInlineNodeStringColumnWithDefaults instantiates a new CustomTopologyMapInlineNodeStringColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapInlineNodeStringColumnWithDefaults() *CustomTopologyMapInlineNodeStringColumn {
	this := CustomTopologyMapInlineNodeStringColumn{}
	return &this
}

// GetId returns the Id field value.
func (o *CustomTopologyMapInlineNodeStringColumn) GetId() CustomTopologyMapInlineNodeStringColumnId {
	if o == nil {
		var ret CustomTopologyMapInlineNodeStringColumnId
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeStringColumn) GetIdOk() (*CustomTopologyMapInlineNodeStringColumnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomTopologyMapInlineNodeStringColumn) SetId(v CustomTopologyMapInlineNodeStringColumnId) {
	o.Id = v
}

// GetKind returns the Kind field value.
func (o *CustomTopologyMapInlineNodeStringColumn) GetKind() CustomTopologyMapInlineNodeStringColumnKind {
	if o == nil {
		var ret CustomTopologyMapInlineNodeStringColumnKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeStringColumn) GetKindOk() (*CustomTopologyMapInlineNodeStringColumnKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *CustomTopologyMapInlineNodeStringColumn) SetKind(v CustomTopologyMapInlineNodeStringColumnKind) {
	o.Kind = v
}

// GetValues returns the Values field value.
func (o *CustomTopologyMapInlineNodeStringColumn) GetValues() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeStringColumn) GetValuesOk() (*[]*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *CustomTopologyMapInlineNodeStringColumn) SetValues(v []*string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapInlineNodeStringColumn) MarshalJSON() ([]byte, error) {
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
func (o *CustomTopologyMapInlineNodeStringColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *CustomTopologyMapInlineNodeStringColumnId   `json:"id"`
		Kind   *CustomTopologyMapInlineNodeStringColumnKind `json:"kind"`
		Values *[]*string                                   `json:"values"`
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
