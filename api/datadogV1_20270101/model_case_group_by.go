// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseGroupBy Test-only generated schema for the dashboard versioning proof of concept.
type CaseGroupBy struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Field string `json:"field"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCaseGroupBy instantiates a new CaseGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseGroupBy(field string) *CaseGroupBy {
	this := CaseGroupBy{}
	this.Field = field
	return &this
}

// NewCaseGroupByWithDefaults instantiates a new CaseGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseGroupByWithDefaults() *CaseGroupBy {
	this := CaseGroupBy{}
	return &this
}

// GetField returns the Field field value.
func (o *CaseGroupBy) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *CaseGroupBy) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *CaseGroupBy) SetField(v string) {
	o.Field = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field *string `json:"field"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	o.Field = *all.Field

	return nil
}
