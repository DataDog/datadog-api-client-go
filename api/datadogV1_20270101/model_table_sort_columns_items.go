// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableSortColumnsItems Test-only generated schema for the dashboard versioning proof of concept.
type TableSortColumnsItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order QuerySortOrder `json:"order"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableSortColumnsItems instantiates a new TableSortColumnsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableSortColumnsItems(column string, order QuerySortOrder) *TableSortColumnsItems {
	this := TableSortColumnsItems{}
	this.Column = column
	this.Order = order
	return &this
}

// NewTableSortColumnsItemsWithDefaults instantiates a new TableSortColumnsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableSortColumnsItemsWithDefaults() *TableSortColumnsItems {
	this := TableSortColumnsItems{}
	return &this
}

// GetColumn returns the Column field value.
func (o *TableSortColumnsItems) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *TableSortColumnsItems) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *TableSortColumnsItems) SetColumn(v string) {
	o.Column = v
}

// GetOrder returns the Order field value.
func (o *TableSortColumnsItems) GetOrder() QuerySortOrder {
	if o == nil {
		var ret QuerySortOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *TableSortColumnsItems) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value.
func (o *TableSortColumnsItems) SetOrder(v QuerySortOrder) {
	o.Order = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableSortColumnsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["order"] = o.Order
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableSortColumnsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string         `json:"column"`
		Order  *QuerySortOrder `json:"order"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Order == nil {
		return fmt.Errorf("required field order missing")
	}

	hasInvalidField := false
	o.Column = *all.Column
	if !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = *all.Order
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
