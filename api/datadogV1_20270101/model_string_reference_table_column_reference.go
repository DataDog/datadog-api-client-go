// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StringReferenceTableColumnReference Test-only generated schema for the dashboard versioning proof of concept.
type StringReferenceTableColumnReference struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStringReferenceTableColumnReference instantiates a new StringReferenceTableColumnReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStringReferenceTableColumnReference(column string, query string) *StringReferenceTableColumnReference {
	this := StringReferenceTableColumnReference{}
	this.Column = column
	this.Query = query
	return &this
}

// NewStringReferenceTableColumnReferenceWithDefaults instantiates a new StringReferenceTableColumnReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStringReferenceTableColumnReferenceWithDefaults() *StringReferenceTableColumnReference {
	this := StringReferenceTableColumnReference{}
	return &this
}

// GetColumn returns the Column field value.
func (o *StringReferenceTableColumnReference) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumnReference) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *StringReferenceTableColumnReference) SetColumn(v string) {
	o.Column = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *StringReferenceTableColumnReference) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumnReference) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *StringReferenceTableColumnReference) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *StringReferenceTableColumnReference) SetFilter(v string) {
	o.Filter = &v
}

// GetQuery returns the Query field value.
func (o *StringReferenceTableColumnReference) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *StringReferenceTableColumnReference) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *StringReferenceTableColumnReference) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StringReferenceTableColumnReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StringReferenceTableColumnReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string `json:"column"`
		Filter *string `json:"filter,omitempty"`
		Query  *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	o.Column = *all.Column
	o.Filter = all.Filter
	o.Query = *all.Query

	return nil
}
