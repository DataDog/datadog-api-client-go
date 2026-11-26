// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetGroupBy Test-only generated schema for the dashboard versioning proof of concept.
type DatasetGroupBy struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatasetGroupBy instantiates a new DatasetGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetGroupBy(column string) *DatasetGroupBy {
	this := DatasetGroupBy{}
	this.Column = column
	return &this
}

// NewDatasetGroupByWithDefaults instantiates a new DatasetGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetGroupByWithDefaults() *DatasetGroupBy {
	this := DatasetGroupBy{}
	return &this
}

// GetColumn returns the Column field value.
func (o *DatasetGroupBy) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *DatasetGroupBy) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *DatasetGroupBy) SetColumn(v string) {
	o.Column = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string `json:"column"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	o.Column = *all.Column

	return nil
}
