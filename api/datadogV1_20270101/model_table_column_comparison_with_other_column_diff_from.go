// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableColumnComparisonWithOtherColumnDiffFrom Test-only generated schema for the dashboard versioning proof of concept.
type TableColumnComparisonWithOtherColumnDiffFrom struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type TableColumnComparisonWithOtherColumnDiffFromType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableColumnComparisonWithOtherColumnDiffFrom instantiates a new TableColumnComparisonWithOtherColumnDiffFrom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableColumnComparisonWithOtherColumnDiffFrom(column string, typeVar TableColumnComparisonWithOtherColumnDiffFromType) *TableColumnComparisonWithOtherColumnDiffFrom {
	this := TableColumnComparisonWithOtherColumnDiffFrom{}
	this.Column = column
	this.Type = typeVar
	return &this
}

// NewTableColumnComparisonWithOtherColumnDiffFromWithDefaults instantiates a new TableColumnComparisonWithOtherColumnDiffFrom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableColumnComparisonWithOtherColumnDiffFromWithDefaults() *TableColumnComparisonWithOtherColumnDiffFrom {
	this := TableColumnComparisonWithOtherColumnDiffFrom{}
	return &this
}

// GetColumn returns the Column field value.
func (o *TableColumnComparisonWithOtherColumnDiffFrom) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *TableColumnComparisonWithOtherColumnDiffFrom) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *TableColumnComparisonWithOtherColumnDiffFrom) SetColumn(v string) {
	o.Column = v
}

// GetType returns the Type field value.
func (o *TableColumnComparisonWithOtherColumnDiffFrom) GetType() TableColumnComparisonWithOtherColumnDiffFromType {
	if o == nil {
		var ret TableColumnComparisonWithOtherColumnDiffFromType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TableColumnComparisonWithOtherColumnDiffFrom) GetTypeOk() (*TableColumnComparisonWithOtherColumnDiffFromType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TableColumnComparisonWithOtherColumnDiffFrom) SetType(v TableColumnComparisonWithOtherColumnDiffFromType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableColumnComparisonWithOtherColumnDiffFrom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableColumnComparisonWithOtherColumnDiffFrom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string                                           `json:"column"`
		Type   *TableColumnComparisonWithOtherColumnDiffFromType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Column = *all.Column
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
