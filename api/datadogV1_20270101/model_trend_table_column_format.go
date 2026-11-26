// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TrendTableColumnFormat Test-only generated schema for the dashboard versioning proof of concept.
type TrendTableColumnFormat struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats *ConditionalFormats `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Mode TrendTableColumnFormatMode `json:"mode"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	NumberFormat *WidgetNumberFormat `json:"number_format,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TrendOptions TableTrendCellDisplayModeOptions `json:"trend_options"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTrendTableColumnFormat instantiates a new TrendTableColumnFormat object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTrendTableColumnFormat(mode TrendTableColumnFormatMode, trendOptions TableTrendCellDisplayModeOptions) *TrendTableColumnFormat {
	this := TrendTableColumnFormat{}
	this.Mode = mode
	this.TrendOptions = trendOptions
	return &this
}

// NewTrendTableColumnFormatWithDefaults instantiates a new TrendTableColumnFormat object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTrendTableColumnFormatWithDefaults() *TrendTableColumnFormat {
	this := TrendTableColumnFormat{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *TrendTableColumnFormat) GetConditionalFormats() ConditionalFormats {
	if o == nil || o.ConditionalFormats == nil {
		var ret ConditionalFormats
		return ret
	}
	return *o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendTableColumnFormat) GetConditionalFormatsOk() (*ConditionalFormats, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *TrendTableColumnFormat) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given ConditionalFormats and assigns it to the ConditionalFormats field.
func (o *TrendTableColumnFormat) SetConditionalFormats(v ConditionalFormats) {
	o.ConditionalFormats = &v
}

// GetMode returns the Mode field value.
func (o *TrendTableColumnFormat) GetMode() TrendTableColumnFormatMode {
	if o == nil {
		var ret TrendTableColumnFormatMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TrendTableColumnFormat) GetModeOk() (*TrendTableColumnFormatMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *TrendTableColumnFormat) SetMode(v TrendTableColumnFormatMode) {
	o.Mode = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *TrendTableColumnFormat) GetNumberFormat() WidgetNumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret WidgetNumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrendTableColumnFormat) GetNumberFormatOk() (*WidgetNumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *TrendTableColumnFormat) HasNumberFormat() bool {
	return o != nil && o.NumberFormat != nil
}

// SetNumberFormat gets a reference to the given WidgetNumberFormat and assigns it to the NumberFormat field.
func (o *TrendTableColumnFormat) SetNumberFormat(v WidgetNumberFormat) {
	o.NumberFormat = &v
}

// GetTrendOptions returns the TrendOptions field value.
func (o *TrendTableColumnFormat) GetTrendOptions() TableTrendCellDisplayModeOptions {
	if o == nil {
		var ret TableTrendCellDisplayModeOptions
		return ret
	}
	return o.TrendOptions
}

// GetTrendOptionsOk returns a tuple with the TrendOptions field value
// and a boolean to check if the value has been set.
func (o *TrendTableColumnFormat) GetTrendOptionsOk() (*TableTrendCellDisplayModeOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrendOptions, true
}

// SetTrendOptions sets field value.
func (o *TrendTableColumnFormat) SetTrendOptions(v TableTrendCellDisplayModeOptions) {
	o.TrendOptions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TrendTableColumnFormat) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TrendTableColumnFormat) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats *ConditionalFormats               `json:"conditional_formats,omitempty"`
		Mode               *TrendTableColumnFormatMode       `json:"mode"`
		NumberFormat       *WidgetNumberFormat               `json:"number_format,omitempty"`
		TrendOptions       *TableTrendCellDisplayModeOptions `json:"trend_options"`
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
