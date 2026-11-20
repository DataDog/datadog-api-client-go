// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetListQuerySort Test-only generated schema for the dashboard versioning proof of concept.
type DatasetListQuerySort struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Fields []DatasetSortField `json:"fields"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatasetListQuerySort instantiates a new DatasetListQuerySort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetListQuerySort(fields []DatasetSortField) *DatasetListQuerySort {
	this := DatasetListQuerySort{}
	this.Fields = fields
	return &this
}

// NewDatasetListQuerySortWithDefaults instantiates a new DatasetListQuerySort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetListQuerySortWithDefaults() *DatasetListQuerySort {
	this := DatasetListQuerySort{}
	return &this
}

// GetFields returns the Fields field value.
func (o *DatasetListQuerySort) GetFields() []DatasetSortField {
	if o == nil {
		var ret []DatasetSortField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *DatasetListQuerySort) GetFieldsOk() (*[]DatasetSortField, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *DatasetListQuerySort) SetFields(v []DatasetSortField) {
	o.Fields = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetListQuerySort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetListQuerySort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields *[]DatasetSortField `json:"fields"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	o.Fields = *all.Fields

	return nil
}
