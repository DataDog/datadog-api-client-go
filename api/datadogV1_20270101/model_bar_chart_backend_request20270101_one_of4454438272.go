// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BarChartBackendRequest20270101OneOf4454438272 Test-only generated schema for the dashboard versioning proof of concept.
type BarChartBackendRequest20270101OneOf4454438272 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443927104 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewBarChartBackendRequest20270101OneOf4454438272 instantiates a new BarChartBackendRequest20270101OneOf4454438272 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBarChartBackendRequest20270101OneOf4454438272(queries []FormulaAndFunctionQueryDefinition, responseFormat BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443927104) *BarChartBackendRequest20270101OneOf4454438272 {
	this := BarChartBackendRequest20270101OneOf4454438272{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewBarChartBackendRequest20270101OneOf4454438272WithDefaults instantiates a new BarChartBackendRequest20270101OneOf4454438272 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBarChartBackendRequest20270101OneOf4454438272WithDefaults() *BarChartBackendRequest20270101OneOf4454438272 {
	this := BarChartBackendRequest20270101OneOf4454438272{}
	return &this
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetConditionalFormats() []ThresholdFormattingRule {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetConditionalFormatsOk() (*[]ThresholdFormattingRule, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule and assigns it to the ConditionalFormats field.
func (o *BarChartBackendRequest20270101OneOf4454438272) SetConditionalFormats(v []ThresholdFormattingRule) {
	o.ConditionalFormats = v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *BarChartBackendRequest20270101OneOf4454438272) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *BarChartBackendRequest20270101OneOf4454438272) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetResponseFormat() BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443927104 {
	if o == nil {
		var ret BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443927104
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetResponseFormatOk() (*BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443927104, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *BarChartBackendRequest20270101OneOf4454438272) SetResponseFormat(v BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443927104) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetSort() ScalarFormulaSort {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) GetSortOk() (*ScalarFormulaSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *BarChartBackendRequest20270101OneOf4454438272) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort and assigns it to the Sort field.
func (o *BarChartBackendRequest20270101OneOf4454438272) SetSort(v ScalarFormulaSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BarChartBackendRequest20270101OneOf4454438272) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BarChartBackendRequest20270101OneOf4454438272) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConditionalFormats []ThresholdFormattingRule                                     `json:"conditional_formats,omitempty"`
		Formulas           []WidgetFormula                                               `json:"formulas,omitempty"`
		Queries            *[]FormulaAndFunctionQueryDefinition                          `json:"queries"`
		ResponseFormat     *BarChartBackendRequestAnyOfAnyOfItemResponseFormat4443927104 `json:"response_format"`
		Sort               *ScalarFormulaSort                                            `json:"sort,omitempty"`
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
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
