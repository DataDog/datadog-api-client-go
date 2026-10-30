// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TopologyMapInlineNodeValueColumn Column of numeric node values. Nodes display non-null values in order of column `id`.
type TopologyMapInlineNodeValueColumn struct {
	// Position of the value in the node, as `values[N]`. The smallest N appears first.
	Id string `json:"id"`
	// Column of numbers.
	Kind TopologyMapInlineColumnKindScalar `json:"kind"`
	// Display options for a value column.
	Meta *TopologyMapInlineNodeValueColumnMeta `json:"meta,omitempty"`
	// Value for each node. Null values are not displayed.
	Values []*float64 `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTopologyMapInlineNodeValueColumn instantiates a new TopologyMapInlineNodeValueColumn object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTopologyMapInlineNodeValueColumn(id string, kind TopologyMapInlineColumnKindScalar, values []*float64) *TopologyMapInlineNodeValueColumn {
	this := TopologyMapInlineNodeValueColumn{}
	this.Id = id
	this.Kind = kind
	this.Values = values
	return &this
}

// NewTopologyMapInlineNodeValueColumnWithDefaults instantiates a new TopologyMapInlineNodeValueColumn object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTopologyMapInlineNodeValueColumnWithDefaults() *TopologyMapInlineNodeValueColumn {
	this := TopologyMapInlineNodeValueColumn{}
	var kind TopologyMapInlineColumnKindScalar = TOPOLOGYMAPINLINECOLUMNKINDSCALAR_SCALAR
	this.Kind = kind
	return &this
}

// GetId returns the Id field value.
func (o *TopologyMapInlineNodeValueColumn) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineNodeValueColumn) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TopologyMapInlineNodeValueColumn) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value.
func (o *TopologyMapInlineNodeValueColumn) GetKind() TopologyMapInlineColumnKindScalar {
	if o == nil {
		var ret TopologyMapInlineColumnKindScalar
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineNodeValueColumn) GetKindOk() (*TopologyMapInlineColumnKindScalar, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *TopologyMapInlineNodeValueColumn) SetKind(v TopologyMapInlineColumnKindScalar) {
	o.Kind = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TopologyMapInlineNodeValueColumn) GetMeta() TopologyMapInlineNodeValueColumnMeta {
	if o == nil || o.Meta == nil {
		var ret TopologyMapInlineNodeValueColumnMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineNodeValueColumn) GetMetaOk() (*TopologyMapInlineNodeValueColumnMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TopologyMapInlineNodeValueColumn) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given TopologyMapInlineNodeValueColumnMeta and assigns it to the Meta field.
func (o *TopologyMapInlineNodeValueColumn) SetMeta(v TopologyMapInlineNodeValueColumnMeta) {
	o.Meta = &v
}

// GetValues returns the Values field value.
func (o *TopologyMapInlineNodeValueColumn) GetValues() []*float64 {
	if o == nil {
		var ret []*float64
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TopologyMapInlineNodeValueColumn) GetValuesOk() (*[]*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *TopologyMapInlineNodeValueColumn) SetValues(v []*float64) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TopologyMapInlineNodeValueColumn) MarshalJSON() ([]byte, error) {
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
func (o *TopologyMapInlineNodeValueColumn) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string                               `json:"id"`
		Kind   *TopologyMapInlineColumnKindScalar    `json:"kind"`
		Meta   *TopologyMapInlineNodeValueColumnMeta `json:"meta,omitempty"`
		Values *[]*float64                           `json:"values"`
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
