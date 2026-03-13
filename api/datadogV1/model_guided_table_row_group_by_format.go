// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableRowGroupByFormat Text formatting configuration for this group-by column.
type GuidedTableRowGroupByFormat struct {
	//
	TextFormats []GuidedTableTextFormattingRule `json:"text_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableRowGroupByFormat instantiates a new GuidedTableRowGroupByFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableRowGroupByFormat() *GuidedTableRowGroupByFormat {
	this := GuidedTableRowGroupByFormat{}
	return &this
}

// NewGuidedTableRowGroupByFormatWithDefaults instantiates a new GuidedTableRowGroupByFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableRowGroupByFormatWithDefaults() *GuidedTableRowGroupByFormat {
	this := GuidedTableRowGroupByFormat{}
	return &this
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *GuidedTableRowGroupByFormat) GetTextFormats() []GuidedTableTextFormattingRule {
	if o == nil || o.TextFormats == nil {
		var ret []GuidedTableTextFormattingRule
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableRowGroupByFormat) GetTextFormatsOk() (*[]GuidedTableTextFormattingRule, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *GuidedTableRowGroupByFormat) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given []GuidedTableTextFormattingRule and assigns it to the TextFormats field.
func (o *GuidedTableRowGroupByFormat) SetTextFormats(v []GuidedTableTextFormattingRule) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableRowGroupByFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableRowGroupByFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TextFormats []GuidedTableTextFormattingRule `json:"text_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"text_formats"})
	} else {
		return err
	}
	o.TextFormats = all.TextFormats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
