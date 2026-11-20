// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendPointsRequest20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendPointsRequest20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []GeomapBackendPointsRequestAnyOfAnyOfItemColumnsItems_20270101 `json:"columns"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Query LogsStreamQuery_20270101 `json:"query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat_20270101 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *GeomapRequestStyle_20270101 `json:"style,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats []TextFormattingRule_20270101 `json:"text_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendPointsRequest20270101OneOf instantiates a new GeomapBackendPointsRequest20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendPointsRequest20270101OneOf(columns []GeomapBackendPointsRequestAnyOfAnyOfItemColumnsItems_20270101, query LogsStreamQuery_20270101, responseFormat GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat_20270101) *GeomapBackendPointsRequest20270101OneOf {
	this := GeomapBackendPointsRequest20270101OneOf{}
	this.Columns = columns
	this.Query = query
	this.ResponseFormat = responseFormat
	return &this
}

// NewGeomapBackendPointsRequest20270101OneOfWithDefaults instantiates a new GeomapBackendPointsRequest20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendPointsRequest20270101OneOfWithDefaults() *GeomapBackendPointsRequest20270101OneOf {
	this := GeomapBackendPointsRequest20270101OneOf{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *GeomapBackendPointsRequest20270101OneOf) GetColumns() []GeomapBackendPointsRequestAnyOfAnyOfItemColumnsItems_20270101 {
	if o == nil {
		var ret []GeomapBackendPointsRequestAnyOfAnyOfItemColumnsItems_20270101
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) GetColumnsOk() (*[]GeomapBackendPointsRequestAnyOfAnyOfItemColumnsItems_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *GeomapBackendPointsRequest20270101OneOf) SetColumns(v []GeomapBackendPointsRequestAnyOfAnyOfItemColumnsItems_20270101) {
	o.Columns = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GeomapBackendPointsRequest20270101OneOf) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *GeomapBackendPointsRequest20270101OneOf) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// GetQuery returns the Query field value.
func (o *GeomapBackendPointsRequest20270101OneOf) GetQuery() LogsStreamQuery_20270101 {
	if o == nil {
		var ret LogsStreamQuery_20270101
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) GetQueryOk() (*LogsStreamQuery_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GeomapBackendPointsRequest20270101OneOf) SetQuery(v LogsStreamQuery_20270101) {
	o.Query = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *GeomapBackendPointsRequest20270101OneOf) GetResponseFormat() GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat_20270101 {
	if o == nil {
		var ret GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) GetResponseFormatOk() (*GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *GeomapBackendPointsRequest20270101OneOf) SetResponseFormat(v GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat_20270101) {
	o.ResponseFormat = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *GeomapBackendPointsRequest20270101OneOf) GetStyle() GeomapRequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret GeomapRequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) GetStyleOk() (*GeomapRequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given GeomapRequestStyle_20270101 and assigns it to the Style field.
func (o *GeomapBackendPointsRequest20270101OneOf) SetStyle(v GeomapRequestStyle_20270101) {
	o.Style = &v
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *GeomapBackendPointsRequest20270101OneOf) GetTextFormats() []TextFormattingRule_20270101 {
	if o == nil || o.TextFormats == nil {
		var ret []TextFormattingRule_20270101
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) GetTextFormatsOk() (*[]TextFormattingRule_20270101, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *GeomapBackendPointsRequest20270101OneOf) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given []TextFormattingRule_20270101 and assigns it to the TextFormats field.
func (o *GeomapBackendPointsRequest20270101OneOf) SetTextFormats(v []TextFormattingRule_20270101) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendPointsRequest20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	toSerialize["query"] = o.Query
	toSerialize["response_format"] = o.ResponseFormat
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendPointsRequest20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns            *[]GeomapBackendPointsRequestAnyOfAnyOfItemColumnsItems_20270101 `json:"columns"`
		ConditionalFormats []ThresholdFormattingRule_20270101                               `json:"conditional_formats,omitempty"`
		Query              *LogsStreamQuery_20270101                                        `json:"query"`
		ResponseFormat     *GeomapBackendPointsRequestAnyOfAnyOfItemResponseFormat_20270101 `json:"response_format"`
		Style              *GeomapRequestStyle_20270101                                     `json:"style,omitempty"`
		TextFormats        []TextFormattingRule_20270101                                    `json:"text_formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	o.ConditionalFormats = all.ConditionalFormats
	o.Query = *all.Query
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}
	o.Style = all.Style
	o.TextFormats = all.TextFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
