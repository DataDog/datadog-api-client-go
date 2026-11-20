// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NumberReferenceTableColumnReference Test-only generated schema for the dashboard versioning proof of concept.
type NumberReferenceTableColumnReference struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filter *string `json:"filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNumberReferenceTableColumnReference instantiates a new NumberReferenceTableColumnReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNumberReferenceTableColumnReference(column string, query string) *NumberReferenceTableColumnReference {
	this := NumberReferenceTableColumnReference{}
	this.Column = column
	this.Query = query
	return &this
}

// NewNumberReferenceTableColumnReferenceWithDefaults instantiates a new NumberReferenceTableColumnReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNumberReferenceTableColumnReferenceWithDefaults() *NumberReferenceTableColumnReference {
	this := NumberReferenceTableColumnReference{}
	return &this
}

// GetColumn returns the Column field value.
func (o *NumberReferenceTableColumnReference) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *NumberReferenceTableColumnReference) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *NumberReferenceTableColumnReference) SetColumn(v string) {
	o.Column = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *NumberReferenceTableColumnReference) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberReferenceTableColumnReference) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *NumberReferenceTableColumnReference) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *NumberReferenceTableColumnReference) SetFilter(v string) {
	o.Filter = &v
}

// GetQuery returns the Query field value.
func (o *NumberReferenceTableColumnReference) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *NumberReferenceTableColumnReference) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *NumberReferenceTableColumnReference) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NumberReferenceTableColumnReference) MarshalJSON() ([]byte, error) {
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
func (o *NumberReferenceTableColumnReference) UnmarshalJSON(bytes []byte) (err error) {
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
