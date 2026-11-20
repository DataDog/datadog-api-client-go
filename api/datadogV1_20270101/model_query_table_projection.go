// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableProjection Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableProjection struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ColumnOrder *QueryTableProjectionColumnOrder `json:"column_order,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimensions []QueryTableProjectionDimension `json:"dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type QueryTableProjectionType `json:"type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	UnlistedColumns *QueryTableProjectionUnlistedColumns `json:"unlisted_columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryTableProjection instantiates a new QueryTableProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryTableProjection(dimensions []QueryTableProjectionDimension, typeVar QueryTableProjectionType) *QueryTableProjection {
	this := QueryTableProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewQueryTableProjectionWithDefaults instantiates a new QueryTableProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryTableProjectionWithDefaults() *QueryTableProjection {
	this := QueryTableProjection{}
	return &this
}

// GetColumnOrder returns the ColumnOrder field value if set, zero value otherwise.
func (o *QueryTableProjection) GetColumnOrder() QueryTableProjectionColumnOrder {
	if o == nil || o.ColumnOrder == nil {
		var ret QueryTableProjectionColumnOrder
		return ret
	}
	return *o.ColumnOrder
}

// GetColumnOrderOk returns a tuple with the ColumnOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjection) GetColumnOrderOk() (*QueryTableProjectionColumnOrder, bool) {
	if o == nil || o.ColumnOrder == nil {
		return nil, false
	}
	return o.ColumnOrder, true
}

// HasColumnOrder returns a boolean if a field has been set.
func (o *QueryTableProjection) HasColumnOrder() bool {
	return o != nil && o.ColumnOrder != nil
}

// SetColumnOrder gets a reference to the given QueryTableProjectionColumnOrder and assigns it to the ColumnOrder field.
func (o *QueryTableProjection) SetColumnOrder(v QueryTableProjectionColumnOrder) {
	o.ColumnOrder = &v
}

// GetDimensions returns the Dimensions field value.
func (o *QueryTableProjection) GetDimensions() []QueryTableProjectionDimension {
	if o == nil {
		var ret []QueryTableProjectionDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *QueryTableProjection) GetDimensionsOk() (*[]QueryTableProjectionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *QueryTableProjection) SetDimensions(v []QueryTableProjectionDimension) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *QueryTableProjection) GetType() QueryTableProjectionType {
	if o == nil {
		var ret QueryTableProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *QueryTableProjection) GetTypeOk() (*QueryTableProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *QueryTableProjection) SetType(v QueryTableProjectionType) {
	o.Type = v
}

// GetUnlistedColumns returns the UnlistedColumns field value if set, zero value otherwise.
func (o *QueryTableProjection) GetUnlistedColumns() QueryTableProjectionUnlistedColumns {
	if o == nil || o.UnlistedColumns == nil {
		var ret QueryTableProjectionUnlistedColumns
		return ret
	}
	return *o.UnlistedColumns
}

// GetUnlistedColumnsOk returns a tuple with the UnlistedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableProjection) GetUnlistedColumnsOk() (*QueryTableProjectionUnlistedColumns, bool) {
	if o == nil || o.UnlistedColumns == nil {
		return nil, false
	}
	return o.UnlistedColumns, true
}

// HasUnlistedColumns returns a boolean if a field has been set.
func (o *QueryTableProjection) HasUnlistedColumns() bool {
	return o != nil && o.UnlistedColumns != nil
}

// SetUnlistedColumns gets a reference to the given QueryTableProjectionUnlistedColumns and assigns it to the UnlistedColumns field.
func (o *QueryTableProjection) SetUnlistedColumns(v QueryTableProjectionUnlistedColumns) {
	o.UnlistedColumns = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryTableProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnOrder != nil {
		toSerialize["column_order"] = o.ColumnOrder
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type
	if o.UnlistedColumns != nil {
		toSerialize["unlisted_columns"] = o.UnlistedColumns
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryTableProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnOrder     *QueryTableProjectionColumnOrder     `json:"column_order,omitempty"`
		Dimensions      *[]QueryTableProjectionDimension     `json:"dimensions"`
		Type            *QueryTableProjectionType            `json:"type"`
		UnlistedColumns *QueryTableProjectionUnlistedColumns `json:"unlisted_columns,omitempty"`
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
	if all.ColumnOrder != nil && !all.ColumnOrder.IsValid() {
		hasInvalidField = true
	} else {
		o.ColumnOrder = all.ColumnOrder
	}
	o.Dimensions = *all.Dimensions
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.UnlistedColumns != nil && !all.UnlistedColumns.IsValid() {
		hasInvalidField = true
	} else {
		o.UnlistedColumns = all.UnlistedColumns
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
