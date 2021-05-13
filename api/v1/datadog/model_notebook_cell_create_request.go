/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// NotebookCellCreateRequest The description of a notebook cell create request.
type NotebookCellCreateRequest struct {
	Attributes NotebookCellCreateRequestAttributes `json:"attributes"`
	Type       NotebookCellResourceType            `json:"type"`
}

// NewNotebookCellCreateRequest instantiates a new NotebookCellCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotebookCellCreateRequest(attributes NotebookCellCreateRequestAttributes, type_ NotebookCellResourceType) *NotebookCellCreateRequest {
	this := NotebookCellCreateRequest{}
	this.Attributes = attributes
	this.Type = type_
	return &this
}

// NewNotebookCellCreateRequestWithDefaults instantiates a new NotebookCellCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotebookCellCreateRequestWithDefaults() *NotebookCellCreateRequest {
	this := NotebookCellCreateRequest{}
	var type_ NotebookCellResourceType = NOTEBOOKCELLRESOURCETYPE_NOTEBOOK_CELLS
	this.Type = type_
	return &this
}

// GetAttributes returns the Attributes field value
func (o *NotebookCellCreateRequest) GetAttributes() NotebookCellCreateRequestAttributes {
	if o == nil {
		var ret NotebookCellCreateRequestAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *NotebookCellCreateRequest) GetAttributesOk() (*NotebookCellCreateRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *NotebookCellCreateRequest) SetAttributes(v NotebookCellCreateRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value
func (o *NotebookCellCreateRequest) GetType() NotebookCellResourceType {
	if o == nil {
		var ret NotebookCellResourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebookCellCreateRequest) GetTypeOk() (*NotebookCellResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotebookCellCreateRequest) SetType(v NotebookCellResourceType) {
	o.Type = v
}

func (o NotebookCellCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

func (o *NotebookCellCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Attributes *NotebookCellCreateRequestAttributes `json:"attributes"`
		Type       *NotebookCellResourceType            `json:"type"`
	}{}
	all := struct {
		Attributes NotebookCellCreateRequestAttributes `json:"attributes"}`
		Type       NotebookCellResourceType            `json:"type"}`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Attributes == nil {
		return fmt.Errorf("Required field attributes missing")
	}
	if required.Type == nil {
		return fmt.Errorf("Required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Attributes = all.Attributes
	o.Type = all.Type
	return nil
}

type NullableNotebookCellCreateRequest struct {
	value *NotebookCellCreateRequest
	isSet bool
}

func (v NullableNotebookCellCreateRequest) Get() *NotebookCellCreateRequest {
	return v.value
}

func (v *NullableNotebookCellCreateRequest) Set(val *NotebookCellCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotebookCellCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotebookCellCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotebookCellCreateRequest(val *NotebookCellCreateRequest) *NullableNotebookCellCreateRequest {
	return &NullableNotebookCellCreateRequest{value: val, isSet: true}
}

func (v NullableNotebookCellCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotebookCellCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
