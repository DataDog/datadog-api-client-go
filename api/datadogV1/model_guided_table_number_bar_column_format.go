// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableNumberBarColumnFormat Formats a guided table column as a number or a bar.
type GuidedTableNumberBarColumnFormat struct {
	// Conditional formatting rules for a guided table column. Either an array of threshold-based rules or a single range-based rule.
	ConditionalFormats *GuidedTableConditionalFormats `json:"conditional_formats,omitempty"`
	//
	Mode GuidedTableNumberBarColumnFormatMode `json:"mode"`
	// Number display format for a guided table column value.
	NumberFormat *GuidedTableNumberFormat `json:"number_format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableNumberBarColumnFormat instantiates a new GuidedTableNumberBarColumnFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableNumberBarColumnFormat(mode GuidedTableNumberBarColumnFormatMode) *GuidedTableNumberBarColumnFormat {
	this := GuidedTableNumberBarColumnFormat{}
	this.Mode = mode
	return &this
}

// NewGuidedTableNumberBarColumnFormatWithDefaults instantiates a new GuidedTableNumberBarColumnFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableNumberBarColumnFormatWithDefaults() *GuidedTableNumberBarColumnFormat {
	this := GuidedTableNumberBarColumnFormat{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GuidedTableNumberBarColumnFormat) GetConditionalFormats() GuidedTableConditionalFormats {
	if o == nil || o.ConditionalFormats == nil {
		var ret GuidedTableConditionalFormats
		return ret
	}
	return *o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableNumberBarColumnFormat) GetConditionalFormatsOk() (*GuidedTableConditionalFormats, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GuidedTableNumberBarColumnFormat) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given GuidedTableConditionalFormats and assigns it to the ConditionalFormats field.
func (o *GuidedTableNumberBarColumnFormat) SetConditionalFormats(v GuidedTableConditionalFormats) {
	o.ConditionalFormats = &v
}

// GetMode returns the Mode field value.
func (o *GuidedTableNumberBarColumnFormat) GetMode() GuidedTableNumberBarColumnFormatMode {
	if o == nil {
		var ret GuidedTableNumberBarColumnFormatMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *GuidedTableNumberBarColumnFormat) GetModeOk() (*GuidedTableNumberBarColumnFormatMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *GuidedTableNumberBarColumnFormat) SetMode(v GuidedTableNumberBarColumnFormatMode) {
	o.Mode = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *GuidedTableNumberBarColumnFormat) GetNumberFormat() GuidedTableNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret GuidedTableNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableNumberBarColumnFormat) GetNumberFormatOk() (*GuidedTableNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *GuidedTableNumberBarColumnFormat) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given GuidedTableNumberFormat and assigns it to the NumberFormat field.
func (o *GuidedTableNumberBarColumnFormat) SetNumberFormat(v GuidedTableNumberFormat) {
	o.NumberFormat = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableNumberBarColumnFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["mode"] = o.Mode
	if o.NumberFormat != nil {
		toSerialize["number_format"] = o.NumberFormat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableNumberBarColumnFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats *GuidedTableConditionalFormats        `json:"conditional_formats,omitempty"`
		Mode               *GuidedTableNumberBarColumnFormatMode `json:"mode"`
		NumberFormat       *GuidedTableNumberFormat              `json:"number_format,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditional_formats", "mode", "number_format"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConditionalFormats = all.ConditionalFormats
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	if all.NumberFormat != nil && all.NumberFormat.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumberFormat = all.NumberFormat

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
