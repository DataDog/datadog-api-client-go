// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableRowGroupByFormat Test-only generated schema for the dashboard versioning proof of concept.
type TableRowGroupByFormat struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats []TextFormattingRule `json:"text_formats"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTableRowGroupByFormat instantiates a new TableRowGroupByFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableRowGroupByFormat(textFormats []TextFormattingRule) *TableRowGroupByFormat {
	this := TableRowGroupByFormat{}
	this.TextFormats = textFormats
	return &this
}

// NewTableRowGroupByFormatWithDefaults instantiates a new TableRowGroupByFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableRowGroupByFormatWithDefaults() *TableRowGroupByFormat {
	this := TableRowGroupByFormat{}
	return &this
}

// GetTextFormats returns the TextFormats field value.
func (o *TableRowGroupByFormat) GetTextFormats() []TextFormattingRule {
	if o == nil {
		var ret []TextFormattingRule
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value
// and a boolean to check if the value has been set.
func (o *TableRowGroupByFormat) GetTextFormatsOk() (*[]TextFormattingRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// SetTextFormats sets field value.
func (o *TableRowGroupByFormat) SetTextFormats(v []TextFormattingRule) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableRowGroupByFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["text_formats"] = o.TextFormats
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableRowGroupByFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TextFormats *[]TextFormattingRule `json:"text_formats"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TextFormats == nil {
		return fmt.Errorf("required field text_formats missing")
	}
	o.TextFormats = *all.TextFormats

	return nil
}
