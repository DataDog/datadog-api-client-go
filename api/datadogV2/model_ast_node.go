// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AstNode A node in the abstract syntax tree of the parsed source code.
type AstNode struct {
	// The tree-sitter node type of this AST node.
	AstType string `json:"ast_type"`
	// The child nodes of this AST node, or null for a leaf node.
	Children datadog.NullableList[AstNode] `json:"children"`
	// A position in source code, identified by line and column numbers.
	End AnalysisPosition `json:"end"`
	// The name of the field this node occupies within its parent node, when the parent addresses it by name.
	FieldName *string `json:"field_name,omitempty"`
	// A position in source code, identified by line and column numbers.
	Start AnalysisPosition `json:"start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAstNode instantiates a new AstNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAstNode(astType string, children datadog.NullableList[AstNode], end AnalysisPosition, start AnalysisPosition) *AstNode {
	this := AstNode{}
	this.AstType = astType
	this.Children = children
	this.End = end
	this.Start = start
	return &this
}

// NewAstNodeWithDefaults instantiates a new AstNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAstNodeWithDefaults() *AstNode {
	this := AstNode{}
	return &this
}

// GetAstType returns the AstType field value.
func (o *AstNode) GetAstType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AstType
}

// GetAstTypeOk returns a tuple with the AstType field value
// and a boolean to check if the value has been set.
func (o *AstNode) GetAstTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AstType, true
}

// SetAstType sets field value.
func (o *AstNode) SetAstType(v string) {
	o.AstType = v
}

// GetChildren returns the Children field value.
// If the value is explicit nil, the zero value for []AstNode will be returned.
func (o *AstNode) GetChildren() []AstNode {
	if o == nil {
		var ret []AstNode
		return ret
	}
	return *o.Children.Get()
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AstNode) GetChildrenOk() (*[]AstNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children.Get(), o.Children.IsSet()
}

// SetChildren sets field value.
func (o *AstNode) SetChildren(v []AstNode) {
	o.Children.Set(&v)
}

// GetEnd returns the End field value.
func (o *AstNode) GetEnd() AnalysisPosition {
	if o == nil {
		var ret AnalysisPosition
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *AstNode) GetEndOk() (*AnalysisPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *AstNode) SetEnd(v AnalysisPosition) {
	o.End = v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *AstNode) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AstNode) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *AstNode) HasFieldName() bool {
	return o != nil && o.FieldName != nil
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *AstNode) SetFieldName(v string) {
	o.FieldName = &v
}

// GetStart returns the Start field value.
func (o *AstNode) GetStart() AnalysisPosition {
	if o == nil {
		var ret AnalysisPosition
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *AstNode) GetStartOk() (*AnalysisPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *AstNode) SetStart(v AnalysisPosition) {
	o.Start = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AstNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ast_type"] = o.AstType
	toSerialize["children"] = o.Children.Get()
	toSerialize["end"] = o.End
	if o.FieldName != nil {
		toSerialize["field_name"] = o.FieldName
	}
	toSerialize["start"] = o.Start

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AstNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AstType   *string                       `json:"ast_type"`
		Children  datadog.NullableList[AstNode] `json:"children"`
		End       *AnalysisPosition             `json:"end"`
		FieldName *string                       `json:"field_name,omitempty"`
		Start     *AnalysisPosition             `json:"start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AstType == nil {
		return fmt.Errorf("required field ast_type missing")
	}
	if !all.Children.IsSet() {
		return fmt.Errorf("required field children missing")
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ast_type", "children", "end", "field_name", "start"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AstType = *all.AstType
	o.Children = all.Children
	if all.End.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.End = *all.End
	o.FieldName = all.FieldName
	if all.Start.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Start = *all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
