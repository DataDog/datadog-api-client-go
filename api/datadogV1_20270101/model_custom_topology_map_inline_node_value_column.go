// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomTopologyMapInlineNodeValueColumn Test-only generated schema for the dashboard versioning proof of concept.
type CustomTopologyMapInlineNodeValueColumn struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Id string `json:"id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Kind CustomTopologyMapInlineNodeValueColumnKind `json:"kind"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Meta *CustomTopologyMapInlineNodeValueColumnMeta `json:"meta,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Values []*float64 `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomTopologyMapInlineNodeValueColumn instantiates a new CustomTopologyMapInlineNodeValueColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomTopologyMapInlineNodeValueColumn(id string, kind CustomTopologyMapInlineNodeValueColumnKind, values []*float64) *CustomTopologyMapInlineNodeValueColumn {
	this := CustomTopologyMapInlineNodeValueColumn{}
	this.Id = id
	this.Kind = kind
	this.Values = values
	return &this
}

// NewCustomTopologyMapInlineNodeValueColumnWithDefaults instantiates a new CustomTopologyMapInlineNodeValueColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomTopologyMapInlineNodeValueColumnWithDefaults() *CustomTopologyMapInlineNodeValueColumn {
	this := CustomTopologyMapInlineNodeValueColumn{}
	return &this
}

// GetId returns the Id field value.
func (o *CustomTopologyMapInlineNodeValueColumn) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeValueColumn) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomTopologyMapInlineNodeValueColumn) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value.
func (o *CustomTopologyMapInlineNodeValueColumn) GetKind() CustomTopologyMapInlineNodeValueColumnKind {
	if o == nil {
		var ret CustomTopologyMapInlineNodeValueColumnKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeValueColumn) GetKindOk() (*CustomTopologyMapInlineNodeValueColumnKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *CustomTopologyMapInlineNodeValueColumn) SetKind(v CustomTopologyMapInlineNodeValueColumnKind) {
	o.Kind = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CustomTopologyMapInlineNodeValueColumn) GetMeta() CustomTopologyMapInlineNodeValueColumnMeta {
	if o == nil || o.Meta == nil {
		var ret CustomTopologyMapInlineNodeValueColumnMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeValueColumn) GetMetaOk() (*CustomTopologyMapInlineNodeValueColumnMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CustomTopologyMapInlineNodeValueColumn) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CustomTopologyMapInlineNodeValueColumnMeta and assigns it to the Meta field.
func (o *CustomTopologyMapInlineNodeValueColumn) SetMeta(v CustomTopologyMapInlineNodeValueColumnMeta) {
	o.Meta = &v
}

// GetValues returns the Values field value.
func (o *CustomTopologyMapInlineNodeValueColumn) GetValues() []*float64 {
	if o == nil {
		var ret []*float64
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *CustomTopologyMapInlineNodeValueColumn) GetValuesOk() (*[]*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *CustomTopologyMapInlineNodeValueColumn) SetValues(v []*float64) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomTopologyMapInlineNodeValueColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["values"] = o.Values
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomTopologyMapInlineNodeValueColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string                                     `json:"id"`
		Kind   *CustomTopologyMapInlineNodeValueColumnKind `json:"kind"`
		Meta   *CustomTopologyMapInlineNodeValueColumnMeta `json:"meta,omitempty"`
		Values *[]*float64                                 `json:"values"`
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
	o.Id = *all.Id
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	o.Values = *all.Values

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
