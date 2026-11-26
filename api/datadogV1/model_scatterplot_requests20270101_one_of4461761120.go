// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotRequests20270101OneOf4461761120 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotRequests20270101OneOf4461761120 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Table ScatterplotBackendTableRequest_20270101 `json:"table"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotRequests20270101OneOf4461761120 instantiates a new ScatterplotRequests20270101OneOf4461761120 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotRequests20270101OneOf4461761120(table ScatterplotBackendTableRequest_20270101) *ScatterplotRequests20270101OneOf4461761120 {
	this := ScatterplotRequests20270101OneOf4461761120{}
	this.Table = table
	return &this
}

// NewScatterplotRequests20270101OneOf4461761120WithDefaults instantiates a new ScatterplotRequests20270101OneOf4461761120 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotRequests20270101OneOf4461761120WithDefaults() *ScatterplotRequests20270101OneOf4461761120 {
	this := ScatterplotRequests20270101OneOf4461761120{}
	return &this
}

// GetTable returns the Table field value.
func (o *ScatterplotRequests20270101OneOf4461761120) GetTable() ScatterplotBackendTableRequest_20270101 {
	if o == nil {
		var ret ScatterplotBackendTableRequest_20270101
		return ret
	}
	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *ScatterplotRequests20270101OneOf4461761120) GetTableOk() (*ScatterplotBackendTableRequest_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value.
func (o *ScatterplotRequests20270101OneOf4461761120) SetTable(v ScatterplotBackendTableRequest_20270101) {
	o.Table = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotRequests20270101OneOf4461761120) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["table"] = o.Table
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotRequests20270101OneOf4461761120) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Table *ScatterplotBackendTableRequest_20270101 `json:"table"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Table == nil {
		return fmt.Errorf("required field table missing")
	}
	o.Table = *all.Table

	return nil
}
