// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapBackendRegionsRequest20270101OneOf4457463328 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapBackendRegionsRequest20270101OneOf4457463328 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula_20270101 `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition_20270101 `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496_20270101 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort_20270101 `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *GeomapRequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapBackendRegionsRequest20270101OneOf4457463328 instantiates a new GeomapBackendRegionsRequest20270101OneOf4457463328 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapBackendRegionsRequest20270101OneOf4457463328(queries []FormulaAndFunctionQueryDefinition_20270101, responseFormat GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496_20270101) *GeomapBackendRegionsRequest20270101OneOf4457463328 {
	this := GeomapBackendRegionsRequest20270101OneOf4457463328{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewGeomapBackendRegionsRequest20270101OneOf4457463328WithDefaults instantiates a new GeomapBackendRegionsRequest20270101OneOf4457463328 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapBackendRegionsRequest20270101OneOf4457463328WithDefaults() *GeomapBackendRegionsRequest20270101OneOf4457463328 {
	this := GeomapBackendRegionsRequest20270101OneOf4457463328{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetFormulas() []WidgetFormula_20270101 {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula_20270101
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetFormulasOk() (*[]WidgetFormula_20270101, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula_20270101 and assigns it to the Formulas field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) SetFormulas(v []WidgetFormula_20270101) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetQueries() []FormulaAndFunctionQueryDefinition_20270101 {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition_20270101
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) SetQueries(v []FormulaAndFunctionQueryDefinition_20270101) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetResponseFormat() GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496_20270101 {
	if o == nil {
		var ret GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetResponseFormatOk() (*GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) SetResponseFormat(v GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496_20270101) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetSort() ScalarFormulaSort_20270101 {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort_20270101
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetSortOk() (*ScalarFormulaSort_20270101, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort_20270101 and assigns it to the Sort field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) SetSort(v ScalarFormulaSort_20270101) {
	o.Sort = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetStyle() GeomapRequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret GeomapRequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) GetStyleOk() (*GeomapRequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given GeomapRequestStyle_20270101 and assigns it to the Style field.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) SetStyle(v GeomapRequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapBackendRegionsRequest20270101OneOf4457463328) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapBackendRegionsRequest20270101OneOf4457463328) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule_20270101                                          `json:"conditional_formats,omitempty"`
		Formulas           []WidgetFormula_20270101                                                    `json:"formulas,omitempty"`
		Queries            *[]FormulaAndFunctionQueryDefinition_20270101                               `json:"queries"`
		ResponseFormat     *GeomapBackendRegionsRequestAnyOfAnyOfItemResponseFormat4447322496_20270101 `json:"response_format"`
		Sort               *ScalarFormulaSort_20270101                                                 `json:"sort,omitempty"`
		Style              *GeomapRequestStyle_20270101                                                `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}

	hasInvalidField := false
	o.ConditionalFormats = all.ConditionalFormats
	o.Formulas = all.Formulas
	o.Queries = *all.Queries
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}
	o.Sort = all.Sort
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
