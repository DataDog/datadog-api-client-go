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

// SyntheticsDeleteTestsResponse struct for SyntheticsDeleteTestsResponse
type SyntheticsDeleteTestsResponse struct {
	DeletedTests *[]SyntheticsDeleteTestsResponseDeletedTests `json:"deleted_tests,omitempty"`
}

// GetDeletedTests returns the DeletedTests field value if set, zero value otherwise.
func (o *SyntheticsDeleteTestsResponse) GetDeletedTests() []SyntheticsDeleteTestsResponseDeletedTests {
	if o == nil || o.DeletedTests == nil {
		var ret []SyntheticsDeleteTestsResponseDeletedTests
		return ret
	}
	return *o.DeletedTests
}

// GetDeletedTestsOk returns a tuple with the DeletedTests field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsDeleteTestsResponse) GetDeletedTestsOk() ([]SyntheticsDeleteTestsResponseDeletedTests, bool) {
	if o == nil || o.DeletedTests == nil {
		var ret []SyntheticsDeleteTestsResponseDeletedTests
		return ret, false
	}
	return *o.DeletedTests, true
}

// HasDeletedTests returns a boolean if a field has been set.
func (o *SyntheticsDeleteTestsResponse) HasDeletedTests() bool {
	if o != nil && o.DeletedTests != nil {
		return true
	}

	return false
}

// SetDeletedTests gets a reference to the given []SyntheticsDeleteTestsResponseDeletedTests and assigns it to the DeletedTests field.
func (o *SyntheticsDeleteTestsResponse) SetDeletedTests(v []SyntheticsDeleteTestsResponseDeletedTests) {
	o.DeletedTests = &v
}

type NullableSyntheticsDeleteTestsResponse struct {
	Value        SyntheticsDeleteTestsResponse
	ExplicitNull bool
}

func (v NullableSyntheticsDeleteTestsResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsDeleteTestsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
