// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StructuredAnalysisCalculatedField Test-only generated schema for the dashboard versioning proof of concept.
type StructuredAnalysisCalculatedField struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Expression string `json:"expression"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type StructuredAnalysisCalculatedFieldType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewStructuredAnalysisCalculatedField instantiates a new StructuredAnalysisCalculatedField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStructuredAnalysisCalculatedField(expression string, name string, typeVar StructuredAnalysisCalculatedFieldType) *StructuredAnalysisCalculatedField {
	this := StructuredAnalysisCalculatedField{}
	this.Expression = expression
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewStructuredAnalysisCalculatedFieldWithDefaults instantiates a new StructuredAnalysisCalculatedField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStructuredAnalysisCalculatedFieldWithDefaults() *StructuredAnalysisCalculatedField {
	this := StructuredAnalysisCalculatedField{}
	return &this
}

// GetExpression returns the Expression field value.
func (o *StructuredAnalysisCalculatedField) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCalculatedField) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *StructuredAnalysisCalculatedField) SetExpression(v string) {
	o.Expression = v
}

// GetName returns the Name field value.
func (o *StructuredAnalysisCalculatedField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCalculatedField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *StructuredAnalysisCalculatedField) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *StructuredAnalysisCalculatedField) GetType() StructuredAnalysisCalculatedFieldType {
	if o == nil {
		var ret StructuredAnalysisCalculatedFieldType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StructuredAnalysisCalculatedField) GetTypeOk() (*StructuredAnalysisCalculatedFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StructuredAnalysisCalculatedField) SetType(v StructuredAnalysisCalculatedFieldType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StructuredAnalysisCalculatedField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["expression"] = o.Expression
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StructuredAnalysisCalculatedField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression *string                                `json:"expression"`
		Name       *string                                `json:"name"`
		Type       *StructuredAnalysisCalculatedFieldType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Expression = *all.Expression
	o.Name = *all.Name
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
