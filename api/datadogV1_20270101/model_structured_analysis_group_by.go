// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisGroupBy Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisGroupBy struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisGroupBy instantiates a new StructuredAnalysisGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisGroupBy(column string) *StructuredAnalysisGroupBy {
	this := StructuredAnalysisGroupBy{}
	this.Column = column
	return &this
}

// NewStructuredAnalysisGroupByWithDefaults instantiates a new StructuredAnalysisGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisGroupByWithDefaults() *StructuredAnalysisGroupBy {
	this := StructuredAnalysisGroupBy{}
	return &this
}

// GetColumn returns the Column field value.
func (o *StructuredAnalysisGroupBy) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisGroupBy) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *StructuredAnalysisGroupBy) SetColumn(v string) {
	o.Column = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StructuredAnalysisGroupBy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisGroupBy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StructuredAnalysisGroupBy) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StructuredAnalysisGroupBy) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string `json:"column"`
		Name   *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	o.Column = *all.Column
	o.Name = all.Name

	return nil
}
