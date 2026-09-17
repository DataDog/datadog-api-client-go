// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAstResponseDataAttributes The attributes of the get-AST response, containing the parsed abstract syntax tree.
type GetAstResponseDataAttributes struct {
	// A node in the abstract syntax tree of the parsed source code.
	Result AstNode `json:"result"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetAstResponseDataAttributes instantiates a new GetAstResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetAstResponseDataAttributes(result AstNode) *GetAstResponseDataAttributes {
	this := GetAstResponseDataAttributes{}
	this.Result = result
	return &this
}

// NewGetAstResponseDataAttributesWithDefaults instantiates a new GetAstResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetAstResponseDataAttributesWithDefaults() *GetAstResponseDataAttributes {
	this := GetAstResponseDataAttributes{}
	return &this
}

// GetResult returns the Result field value.
func (o *GetAstResponseDataAttributes) GetResult() AstNode {
	if o == nil {
		var ret AstNode
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetAstResponseDataAttributes) GetResultOk() (*AstNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value.
func (o *GetAstResponseDataAttributes) SetResult(v AstNode) {
	o.Result = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetAstResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["result"] = o.Result

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetAstResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Result *AstNode `json:"result"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Result == nil {
		return fmt.Errorf("required field result missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"result"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = *all.Result

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
