// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GuidedTableTrendColumnFormat Formats a guided table column as a trend sparkline.
type GuidedTableTrendColumnFormat struct {
	// Conditional formatting rules for a guided table column. Either an array of threshold-based rules or a single range-based rule.
	ConditionalFormats *GuidedTableConditionalFormats `json:"conditional_formats,omitempty"`
	//
	Mode GuidedTableTrendColumnFormatMode `json:"mode"`
	// Number display format for a guided table column value.
	NumberFormat *GuidedTableNumberFormat `json:"number_format,omitempty"`
	// Cell display mode options for the widget formula. (only if `cell_display_mode` is set to `trend`).
	TrendOptions WidgetFormulaCellDisplayModeOptions `json:"trend_options"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGuidedTableTrendColumnFormat instantiates a new GuidedTableTrendColumnFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGuidedTableTrendColumnFormat(mode GuidedTableTrendColumnFormatMode, trendOptions WidgetFormulaCellDisplayModeOptions) *GuidedTableTrendColumnFormat {
	this := GuidedTableTrendColumnFormat{}
	this.Mode = mode
	this.TrendOptions = trendOptions
	return &this
}

// NewGuidedTableTrendColumnFormatWithDefaults instantiates a new GuidedTableTrendColumnFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGuidedTableTrendColumnFormatWithDefaults() *GuidedTableTrendColumnFormat {
	this := GuidedTableTrendColumnFormat{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GuidedTableTrendColumnFormat) GetConditionalFormats() GuidedTableConditionalFormats {
	if o == nil || o.ConditionalFormats == nil {
		var ret GuidedTableConditionalFormats
		return ret
	}
	return *o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableTrendColumnFormat) GetConditionalFormatsOk() (*GuidedTableConditionalFormats, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GuidedTableTrendColumnFormat) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given GuidedTableConditionalFormats and assigns it to the ConditionalFormats field.
func (o *GuidedTableTrendColumnFormat) SetConditionalFormats(v GuidedTableConditionalFormats) {
	o.ConditionalFormats = &v
}

// GetMode returns the Mode field value.
func (o *GuidedTableTrendColumnFormat) GetMode() GuidedTableTrendColumnFormatMode {
	if o == nil {
		var ret GuidedTableTrendColumnFormatMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *GuidedTableTrendColumnFormat) GetModeOk() (*GuidedTableTrendColumnFormatMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *GuidedTableTrendColumnFormat) SetMode(v GuidedTableTrendColumnFormatMode) {
	o.Mode = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *GuidedTableTrendColumnFormat) GetNumberFormat() GuidedTableNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret GuidedTableNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidedTableTrendColumnFormat) GetNumberFormatOk() (*GuidedTableNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *GuidedTableTrendColumnFormat) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given GuidedTableNumberFormat and assigns it to the NumberFormat field.
func (o *GuidedTableTrendColumnFormat) SetNumberFormat(v GuidedTableNumberFormat) {
	o.NumberFormat = &v
}

// GetTrendOptions returns the TrendOptions field value.
func (o *GuidedTableTrendColumnFormat) GetTrendOptions() WidgetFormulaCellDisplayModeOptions {
	if o == nil {
		var ret WidgetFormulaCellDisplayModeOptions
		return ret
	}
	return o.TrendOptions
}

// GetTrendOptionsOk returns a tuple with the TrendOptions field value
// and a boolean to check if the value has been set.
func (o *GuidedTableTrendColumnFormat) GetTrendOptionsOk() (*WidgetFormulaCellDisplayModeOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrendOptions, true
}

// SetTrendOptions sets field value.
func (o *GuidedTableTrendColumnFormat) SetTrendOptions(v WidgetFormulaCellDisplayModeOptions) {
	o.TrendOptions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GuidedTableTrendColumnFormat) MarshalJSON() ([]byte, error) {
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
	toSerialize["trend_options"] = o.TrendOptions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GuidedTableTrendColumnFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats *GuidedTableConditionalFormats       `json:"conditional_formats,omitempty"`
		Mode               *GuidedTableTrendColumnFormatMode    `json:"mode"`
		NumberFormat       *GuidedTableNumberFormat             `json:"number_format,omitempty"`
		TrendOptions       *WidgetFormulaCellDisplayModeOptions `json:"trend_options"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.TrendOptions == nil {
		return fmt.Errorf("required field trend_options missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditional_formats", "mode", "number_format", "trend_options"})
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
	if all.TrendOptions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TrendOptions = *all.TrendOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
