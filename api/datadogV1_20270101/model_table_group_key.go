// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableGroupKey Test-only generated schema for the dashboard versioning proof of concept.
type TableGroupKey struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupKey string `json:"group_key"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableGroupKey instantiates a new TableGroupKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableGroupKey(groupKey string, query string) *TableGroupKey {
	this := TableGroupKey{}
	this.GroupKey = groupKey
	this.Query = query
	return &this
}

// NewTableGroupKeyWithDefaults instantiates a new TableGroupKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableGroupKeyWithDefaults() *TableGroupKey {
	this := TableGroupKey{}
	return &this
}

// GetGroupKey returns the GroupKey field value.
func (o *TableGroupKey) GetGroupKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupKey
}

// GetGroupKeyOk returns a tuple with the GroupKey field value
// and a boolean to check if the value has been set.
func (o *TableGroupKey) GetGroupKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupKey, true
}

// SetGroupKey sets field value.
func (o *TableGroupKey) SetGroupKey(v string) {
	o.GroupKey = v
}

// GetQuery returns the Query field value.
func (o *TableGroupKey) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TableGroupKey) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *TableGroupKey) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableGroupKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group_key"] = o.GroupKey
	toSerialize["query"] = o.Query
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableGroupKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupKey *string `json:"group_key"`
		Query    *string `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupKey == nil {
		return fmt.Errorf("required field group_key missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	o.GroupKey = *all.GroupKey
	o.Query = *all.Query

	return nil
}
