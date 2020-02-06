/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// LogsIndexesOrder Object containing the ordered list of log index names.
type LogsIndexesOrder struct {
	// Array of strings identifying by their name(s) the index(es) of your organisation. Logs are tested against the query filter of each index one by one, following the order of the array. Logs are eventually stored in the first matching index.
	IndexNames []string `json:"index_names"`
}

// NewLogsIndexesOrder instantiates a new LogsIndexesOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsIndexesOrder(indexNames []string) *LogsIndexesOrder {
	this := LogsIndexesOrder{}
	this.IndexNames = indexNames
	return &this
}

// NewLogsIndexesOrderWithDefaults instantiates a new LogsIndexesOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsIndexesOrderWithDefaults() *LogsIndexesOrder {
	this := LogsIndexesOrder{}
	return &this
}

// GetIndexNames returns the IndexNames field value
func (o *LogsIndexesOrder) GetIndexNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IndexNames
}

// SetIndexNames sets field value
func (o *LogsIndexesOrder) SetIndexNames(v []string) {
	o.IndexNames = v
}

type NullableLogsIndexesOrder struct {
	Value        LogsIndexesOrder
	ExplicitNull bool
}

func (v NullableLogsIndexesOrder) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLogsIndexesOrder) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
