// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableSort Test-only generated schema for the dashboard versioning proof of concept.
type TableSort struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []TableSortColumnsItems `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit float64 `json:"limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableSort instantiates a new TableSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableSort(columns []TableSortColumnsItems, limit float64) *TableSort {
	this := TableSort{}
	this.Columns = columns
	this.Limit = limit
	return &this
}

// NewTableSortWithDefaults instantiates a new TableSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableSortWithDefaults() *TableSort {
	this := TableSort{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *TableSort) GetColumns() []TableSortColumnsItems {
	if o == nil {
		var ret []TableSortColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *TableSort) GetColumnsOk() (*[]TableSortColumnsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *TableSort) SetColumns(v []TableSortColumnsItems) {
	o.Columns = v
}

// GetLimit returns the Limit field value.
func (o *TableSort) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TableSort) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *TableSort) SetLimit(v float64) {
	o.Limit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["limit"] = o.Limit
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns *[]TableSortColumnsItems `json:"columns"`
		Limit   *float64                 `json:"limit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	o.Columns = *all.Columns
	o.Limit = *all.Limit

	return nil
}
