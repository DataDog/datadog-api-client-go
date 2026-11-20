// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GroupStyleRuleMatch Test-only generated schema for the dashboard versioning proof of concept.
type GroupStyleRuleMatch struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Column string `json:"column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGroupStyleRuleMatch instantiates a new GroupStyleRuleMatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGroupStyleRuleMatch(column string, value string) *GroupStyleRuleMatch {
	this := GroupStyleRuleMatch{}
	this.Column = column
	this.Value = value
	return &this
}

// NewGroupStyleRuleMatchWithDefaults instantiates a new GroupStyleRuleMatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGroupStyleRuleMatchWithDefaults() *GroupStyleRuleMatch {
	this := GroupStyleRuleMatch{}
	return &this
}

// GetColumn returns the Column field value.
func (o *GroupStyleRuleMatch) GetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Column
}

// GetColumnOk returns a tuple with the Column field value
// and a boolean to check if the value has been set.
func (o *GroupStyleRuleMatch) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Column, true
}

// SetColumn sets field value.
func (o *GroupStyleRuleMatch) SetColumn(v string) {
	o.Column = v
}

// GetValue returns the Value field value.
func (o *GroupStyleRuleMatch) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GroupStyleRuleMatch) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *GroupStyleRuleMatch) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GroupStyleRuleMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column"] = o.Column
	toSerialize["value"] = o.Value
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GroupStyleRuleMatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Column *string `json:"column"`
		Value  *string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Column == nil {
		return fmt.Errorf("required field column missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	o.Column = *all.Column
	o.Value = *all.Value

	return nil
}
