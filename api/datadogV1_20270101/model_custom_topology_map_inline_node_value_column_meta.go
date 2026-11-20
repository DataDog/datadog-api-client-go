// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeValueColumnMeta Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeValueColumnMeta struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Unit *NumberFormatUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapInlineNodeValueColumnMeta instantiates a new CustomTopologyMapInlineNodeValueColumnMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapInlineNodeValueColumnMeta() *CustomTopologyMapInlineNodeValueColumnMeta {
	this := CustomTopologyMapInlineNodeValueColumnMeta{}
	return &this
}

// NewCustomTopologyMapInlineNodeValueColumnMetaWithDefaults instantiates a new CustomTopologyMapInlineNodeValueColumnMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapInlineNodeValueColumnMetaWithDefaults() *CustomTopologyMapInlineNodeValueColumnMeta {
	this := CustomTopologyMapInlineNodeValueColumnMeta{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CustomTopologyMapInlineNodeValueColumnMeta) GetUnit() NumberFormatUnit {
	if o == nil || o.Unit == nil {
		var ret NumberFormatUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeValueColumnMeta) GetUnitOk() (*NumberFormatUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CustomTopologyMapInlineNodeValueColumnMeta) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given NumberFormatUnit and assigns it to the Unit field.
func (o *CustomTopologyMapInlineNodeValueColumnMeta) SetUnit(v NumberFormatUnit) {
	o.Unit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapInlineNodeValueColumnMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomTopologyMapInlineNodeValueColumnMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Unit *NumberFormatUnit `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Unit = all.Unit

	return nil
}
