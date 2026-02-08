// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookSearchHighlight Highlighted fields from the notebook search.
type NotebookSearchHighlight struct {
	// Highlighted cell text matches.
	CellsText []string `json:"cells.text,omitempty"`
	// Highlighted cell title matches.
	CellsTitle []string `json:"cells.title,omitempty"`
	// Highlighted notebook name matches.
	Name []string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookSearchHighlight instantiates a new NotebookSearchHighlight object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookSearchHighlight() *NotebookSearchHighlight {
	this := NotebookSearchHighlight{}
	return &this
}

// NewNotebookSearchHighlightWithDefaults instantiates a new NotebookSearchHighlight object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookSearchHighlightWithDefaults() *NotebookSearchHighlight {
	this := NotebookSearchHighlight{}
	return &this
}

// GetCellsText returns the CellsText field value if set, zero value otherwise.
func (o *NotebookSearchHighlight) GetCellsText() []string {
	if o == nil || o.CellsText == nil {
		var ret []string
		return ret
	}
	return o.CellsText
}

// GetCellsTextOk returns a tuple with the CellsText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchHighlight) GetCellsTextOk() (*[]string, bool) {
	if o == nil || o.CellsText == nil {
		return nil, false
	}
	return &o.CellsText, true
}

// HasCellsText returns a boolean if a field has been set.
func (o *NotebookSearchHighlight) HasCellsText() bool {
	return o != nil && o.CellsText != nil
}

// SetCellsText gets a reference to the given []string and assigns it to the CellsText field.
func (o *NotebookSearchHighlight) SetCellsText(v []string) {
	o.CellsText = v
}

// GetCellsTitle returns the CellsTitle field value if set, zero value otherwise.
func (o *NotebookSearchHighlight) GetCellsTitle() []string {
	if o == nil || o.CellsTitle == nil {
		var ret []string
		return ret
	}
	return o.CellsTitle
}

// GetCellsTitleOk returns a tuple with the CellsTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchHighlight) GetCellsTitleOk() (*[]string, bool) {
	if o == nil || o.CellsTitle == nil {
		return nil, false
	}
	return &o.CellsTitle, true
}

// HasCellsTitle returns a boolean if a field has been set.
func (o *NotebookSearchHighlight) HasCellsTitle() bool {
	return o != nil && o.CellsTitle != nil
}

// SetCellsTitle gets a reference to the given []string and assigns it to the CellsTitle field.
func (o *NotebookSearchHighlight) SetCellsTitle(v []string) {
	o.CellsTitle = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotebookSearchHighlight) GetName() []string {
	if o == nil || o.Name == nil {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchHighlight) GetNameOk() (*[]string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotebookSearchHighlight) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *NotebookSearchHighlight) SetName(v []string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookSearchHighlight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CellsText != nil {
		toSerialize["cells.text"] = o.CellsText
	}
	if o.CellsTitle != nil {
		toSerialize["cells.title"] = o.CellsTitle
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookSearchHighlight) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CellsText  []string `json:"cells.text,omitempty"`
		CellsTitle []string `json:"cells.title,omitempty"`
		Name       []string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cells.text", "cells.title", "name"})
	} else {
		return err
	}
	o.CellsText = all.CellsText
	o.CellsTitle = all.CellsTitle
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
