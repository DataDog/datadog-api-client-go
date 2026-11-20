// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisJoin Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisJoin struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	JoinConditions []StructuredAnalysisJoinCondition `json:"join_conditions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	JoinType StructuredAnalysisJoinJoinType `json:"join_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TargetColumns []string `json:"target_columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TargetDataset string `json:"target_dataset"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisJoinType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisJoin instantiates a new StructuredAnalysisJoin object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisJoin(joinConditions []StructuredAnalysisJoinCondition, joinType StructuredAnalysisJoinJoinType, targetColumns []string, targetDataset string, typeVar StructuredAnalysisJoinType) *StructuredAnalysisJoin {
	this := StructuredAnalysisJoin{}
	this.JoinConditions = joinConditions
	this.JoinType = joinType
	this.TargetColumns = targetColumns
	this.TargetDataset = targetDataset
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisJoinWithDefaults instantiates a new StructuredAnalysisJoin object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisJoinWithDefaults() *StructuredAnalysisJoin {
	this := StructuredAnalysisJoin{}
	return &this
}

// GetJoinConditions returns the JoinConditions field value.
func (o *StructuredAnalysisJoin) GetJoinConditions() []StructuredAnalysisJoinCondition {
	if o == nil {
		var ret []StructuredAnalysisJoinCondition
		return ret
	}
	return o.JoinConditions
}

// GetJoinConditionsOk returns a tuple with the JoinConditions field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisJoin) GetJoinConditionsOk() (*[]StructuredAnalysisJoinCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinConditions, true
}

// SetJoinConditions sets field value.
func (o *StructuredAnalysisJoin) SetJoinConditions(v []StructuredAnalysisJoinCondition) {
	o.JoinConditions = v
}

// GetJoinType returns the JoinType field value.
func (o *StructuredAnalysisJoin) GetJoinType() StructuredAnalysisJoinJoinType {
	if o == nil {
		var ret StructuredAnalysisJoinJoinType
		return ret
	}
	return o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisJoin) GetJoinTypeOk() (*StructuredAnalysisJoinJoinType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinType, true
}

// SetJoinType sets field value.
func (o *StructuredAnalysisJoin) SetJoinType(v StructuredAnalysisJoinJoinType) {
	o.JoinType = v
}

// GetTargetColumns returns the TargetColumns field value.
func (o *StructuredAnalysisJoin) GetTargetColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TargetColumns
}

// GetTargetColumnsOk returns a tuple with the TargetColumns field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisJoin) GetTargetColumnsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetColumns, true
}

// SetTargetColumns sets field value.
func (o *StructuredAnalysisJoin) SetTargetColumns(v []string) {
	o.TargetColumns = v
}

// GetTargetDataset returns the TargetDataset field value.
func (o *StructuredAnalysisJoin) GetTargetDataset() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetDataset
}

// GetTargetDatasetOk returns a tuple with the TargetDataset field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisJoin) GetTargetDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetDataset, true
}

// SetTargetDataset sets field value.
func (o *StructuredAnalysisJoin) SetTargetDataset(v string) {
	o.TargetDataset = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisJoin) GetType() StructuredAnalysisJoinType {
	if o == nil {
		var ret StructuredAnalysisJoinType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisJoin) GetTypeOk() (*StructuredAnalysisJoinType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisJoin) SetType(v StructuredAnalysisJoinType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisJoin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["join_conditions"] = o.JoinConditions
	toSerialize["join_type"] = o.JoinType
	toSerialize["target_columns"] = o.TargetColumns
	toSerialize["target_dataset"] = o.TargetDataset
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisJoin) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JoinConditions *[]StructuredAnalysisJoinCondition `json:"join_conditions"`
		JoinType       *StructuredAnalysisJoinJoinType    `json:"join_type"`
		TargetColumns  *[]string                          `json:"target_columns"`
		TargetDataset  *string                            `json:"target_dataset"`
		Type           *StructuredAnalysisJoinType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JoinConditions == nil {
		return fmt.Errorf("required field join_conditions missing")
	}
	if all.JoinType == nil {
		return fmt.Errorf("required field join_type missing")
	}
	if all.TargetColumns == nil {
		return fmt.Errorf("required field target_columns missing")
	}
	if all.TargetDataset == nil {
		return fmt.Errorf("required field target_dataset missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.JoinConditions = *all.JoinConditions
	if !all.JoinType.IsValid() {
		hasInvalidField = true
	} else {
		o.JoinType = *all.JoinType
	}
	o.TargetColumns = *all.TargetColumns
	o.TargetDataset = *all.TargetDataset
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
