// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisCast Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisCast struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	SourceColumn string `json:"source_column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TargetColumn string `json:"target_column"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TargetType StructuredAnalysisCastTargetType `json:"target_type"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisCastType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisCast instantiates a new StructuredAnalysisCast object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisCast(sourceColumn string, targetColumn string, targetType StructuredAnalysisCastTargetType, typeVar StructuredAnalysisCastType) *StructuredAnalysisCast {
	this := StructuredAnalysisCast{}
	this.SourceColumn = sourceColumn
	this.TargetColumn = targetColumn
	this.TargetType = targetType
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisCastWithDefaults instantiates a new StructuredAnalysisCast object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisCastWithDefaults() *StructuredAnalysisCast {
	this := StructuredAnalysisCast{}
	return &this
}

// GetSourceColumn returns the SourceColumn field value.
func (o *StructuredAnalysisCast) GetSourceColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceColumn
}

// GetSourceColumnOk returns a tuple with the SourceColumn field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCast) GetSourceColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceColumn, true
}

// SetSourceColumn sets field value.
func (o *StructuredAnalysisCast) SetSourceColumn(v string) {
	o.SourceColumn = v
}

// GetTargetColumn returns the TargetColumn field value.
func (o *StructuredAnalysisCast) GetTargetColumn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetColumn
}

// GetTargetColumnOk returns a tuple with the TargetColumn field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCast) GetTargetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetColumn, true
}

// SetTargetColumn sets field value.
func (o *StructuredAnalysisCast) SetTargetColumn(v string) {
	o.TargetColumn = v
}

// GetTargetType returns the TargetType field value.
func (o *StructuredAnalysisCast) GetTargetType() StructuredAnalysisCastTargetType {
	if o == nil {
		var ret StructuredAnalysisCastTargetType
		return ret
	}
	return o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCast) GetTargetTypeOk() (*StructuredAnalysisCastTargetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetType, true
}

// SetTargetType sets field value.
func (o *StructuredAnalysisCast) SetTargetType(v StructuredAnalysisCastTargetType) {
	o.TargetType = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisCast) GetType() StructuredAnalysisCastType {
	if o == nil {
		var ret StructuredAnalysisCastType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCast) GetTypeOk() (*StructuredAnalysisCastType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisCast) SetType(v StructuredAnalysisCastType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisCast) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["source_column"] = o.SourceColumn
	toSerialize["target_column"] = o.TargetColumn
	toSerialize["target_type"] = o.TargetType
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisCast) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SourceColumn *string                           `json:"source_column"`
		TargetColumn *string                           `json:"target_column"`
		TargetType   *StructuredAnalysisCastTargetType `json:"target_type"`
		Type         *StructuredAnalysisCastType       `json:"type"`
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
	if all.TargetType == nil {
		return fmt.Errorf("required field target_type missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.SourceColumn = *all.SourceColumn
	o.TargetColumn = *all.TargetColumn
	if !all.TargetType.IsValid() {
		hasInvalidField = true
	} else {
		o.TargetType = *all.TargetType
	}
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
