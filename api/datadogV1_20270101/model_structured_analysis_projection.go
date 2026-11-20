// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisProjection Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisProjection struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []StructuredAnalysisProjectionColumnsItems `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisProjection instantiates a new StructuredAnalysisProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisProjection(columns []StructuredAnalysisProjectionColumnsItems, typeVar StructuredAnalysisProjectionType) *StructuredAnalysisProjection {
	this := StructuredAnalysisProjection{}
	this.Columns = columns
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisProjectionWithDefaults instantiates a new StructuredAnalysisProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisProjectionWithDefaults() *StructuredAnalysisProjection {
	this := StructuredAnalysisProjection{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *StructuredAnalysisProjection) GetColumns() []StructuredAnalysisProjectionColumnsItems {
	if o == nil {
		var ret []StructuredAnalysisProjectionColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisProjection) GetColumnsOk() (*[]StructuredAnalysisProjectionColumnsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *StructuredAnalysisProjection) SetColumns(v []StructuredAnalysisProjectionColumnsItems) {
	o.Columns = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisProjection) GetType() StructuredAnalysisProjectionType {
	if o == nil {
		var ret StructuredAnalysisProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisProjection) GetTypeOk() (*StructuredAnalysisProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisProjection) SetType(v StructuredAnalysisProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns *[]StructuredAnalysisProjectionColumnsItems `json:"columns"`
		Type    *StructuredAnalysisProjectionType           `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Columns = *all.Columns
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
