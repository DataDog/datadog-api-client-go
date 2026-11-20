// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisJoinCondition Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisJoinCondition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	SourceColumn string `json:"source_column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TargetColumn string `json:"target_column"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisJoinCondition instantiates a new StructuredAnalysisJoinCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisJoinCondition(sourceColumn string, targetColumn string) *StructuredAnalysisJoinCondition {
	this := StructuredAnalysisJoinCondition{}
	this.SourceColumn = sourceColumn
	this.TargetColumn = targetColumn
	return &this
}

// NewStructuredAnalysisJoinConditionWithDefaults instantiates a new StructuredAnalysisJoinCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisJoinConditionWithDefaults() *StructuredAnalysisJoinCondition {
	this := StructuredAnalysisJoinCondition{}
	return &this
}

// GetSourceColumn returns the SourceColumn field value.
func (o *StructuredAnalysisJoinCondition) GetSourceColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceColumn
}

// GetSourceColumnOk returns a tuple with the SourceColumn field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisJoinCondition) GetSourceColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceColumn, true
}

// SetSourceColumn sets field value.
func (o *StructuredAnalysisJoinCondition) SetSourceColumn(v string) {
	o.SourceColumn = v
}

// GetTargetColumn returns the TargetColumn field value.
func (o *StructuredAnalysisJoinCondition) GetTargetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetColumn
}

// GetTargetColumnOk returns a tuple with the TargetColumn field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisJoinCondition) GetTargetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetColumn, true
}

// SetTargetColumn sets field value.
func (o *StructuredAnalysisJoinCondition) SetTargetColumn(v string) {
	o.TargetColumn = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisJoinCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["source_column"] = o.SourceColumn
	toSerialize["target_column"] = o.TargetColumn
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisJoinCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SourceColumn *string `json:"source_column"`
		TargetColumn *string `json:"target_column"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SourceColumn == nil {
		return fmt.Errorf("required field source_column missing")
	}
	if all.TargetColumn == nil {
		return fmt.Errorf("required field target_column missing")
	}
	o.SourceColumn = *all.SourceColumn
	o.TargetColumn = *all.TargetColumn

	return nil
}
