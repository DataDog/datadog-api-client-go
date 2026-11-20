// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableBackendRequest20270101OneOf4460519136 Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableBackendRequest20270101OneOf4460519136 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator_20270101 `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Alias *string `json:"alias,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CellDisplayMode []CellDisplayModeOption_20270101 `json:"cell_display_mode,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ConditionalFormats []ThresholdFormattingRule_20270101 `json:"conditional_formats,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula_20270101 `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Order *QuerySortOrder_20270101 `json:"order,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition_20270101 `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256_20270101 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort_20270101 `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFormats [][]TextFormattingRule_20270101 `json:"text_formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryTableBackendRequest20270101OneOf4460519136 instantiates a new QueryTableBackendRequest20270101OneOf4460519136 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryTableBackendRequest20270101OneOf4460519136(queries []FormulaAndFunctionQueryDefinition_20270101, responseFormat QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256_20270101) *QueryTableBackendRequest20270101OneOf4460519136 {
	this := QueryTableBackendRequest20270101OneOf4460519136{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewQueryTableBackendRequest20270101OneOf4460519136WithDefaults instantiates a new QueryTableBackendRequest20270101OneOf4460519136 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryTableBackendRequest20270101OneOf4460519136WithDefaults() *QueryTableBackendRequest20270101OneOf4460519136 {
	this := QueryTableBackendRequest20270101OneOf4460519136{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetAggregator() ScalarAggregator_20270101 {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator_20270101
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetAggregatorOk() (*ScalarAggregator_20270101, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator_20270101 and assigns it to the Aggregator field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetAggregator(v ScalarAggregator_20270101) {
	o.Aggregator = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetAlias(v string) {
	o.Alias = &v
}

// GetCellDisplayMode returns the CellDisplayMode field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetCellDisplayMode() []CellDisplayModeOption_20270101 {
	if o == nil || o.CellDisplayMode == nil {
		var ret []CellDisplayModeOption_20270101
		return ret
	}
	return o.CellDisplayMode
}

// GetCellDisplayModeOk returns a tuple with the CellDisplayMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetCellDisplayModeOk() (*[]CellDisplayModeOption_20270101, bool) {
	if o == nil || o.CellDisplayMode == nil {
		return nil, false
	}
	return &o.CellDisplayMode, true
}

// HasCellDisplayMode returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasCellDisplayMode() bool {
	return o != nil && o.CellDisplayMode != nil
}

// SetCellDisplayMode gets a reference to the given []CellDisplayModeOption_20270101 and assigns it to the CellDisplayMode field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetCellDisplayMode(v []CellDisplayModeOption_20270101) {
	o.CellDisplayMode = v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetConditionalFormats() []ThresholdFormattingRule_20270101 {
	if o == nil || o.ConditionalFormats == nil {
		var ret []ThresholdFormattingRule_20270101
		return ret
	}
	return o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetConditionalFormatsOk() (*[]ThresholdFormattingRule_20270101, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return &o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasConditionalFormats() bool {
	return o != nil && o.ConditionalFormats != nil
}

// SetConditionalFormats gets a reference to the given []ThresholdFormattingRule_20270101 and assigns it to the ConditionalFormats field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetConditionalFormats(v []ThresholdFormattingRule_20270101) {
	o.ConditionalFormats = v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetFormulas() []WidgetFormula_20270101 {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula_20270101
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetFormulasOk() (*[]WidgetFormula_20270101, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula_20270101 and assigns it to the Formulas field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetFormulas(v []WidgetFormula_20270101) {
	o.Formulas = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetLimit(v float64) {
	o.Limit = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetOrder() QuerySortOrder_20270101 {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder_20270101
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetOrderOk() (*QuerySortOrder_20270101, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given QuerySortOrder_20270101 and assigns it to the Order field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetOrder(v QuerySortOrder_20270101) {
	o.Order = &v
}

// GetQueries returns the Queries field value.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetQueries() []FormulaAndFunctionQueryDefinition_20270101 {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition_20270101
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetQueries(v []FormulaAndFunctionQueryDefinition_20270101) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetResponseFormat() QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256_20270101 {
	if o == nil {
		var ret QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetResponseFormatOk() (*QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetResponseFormat(v QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256_20270101) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetSort() ScalarFormulaSort_20270101 {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort_20270101
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetSortOk() (*ScalarFormulaSort_20270101, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort_20270101 and assigns it to the Sort field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetSort(v ScalarFormulaSort_20270101) {
	o.Sort = &v
}

// GetTextFormats returns the TextFormats field value if set, zero value otherwise.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetTextFormats() [][]TextFormattingRule_20270101 {
	if o == nil || o.TextFormats == nil {
		var ret [][]TextFormattingRule_20270101
		return ret
	}
	return o.TextFormats
}

// GetTextFormatsOk returns a tuple with the TextFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) GetTextFormatsOk() (*[][]TextFormattingRule_20270101, bool) {
	if o == nil || o.TextFormats == nil {
		return nil, false
	}
	return &o.TextFormats, true
}

// HasTextFormats returns a boolean if a field has been set.
func (o *QueryTableBackendRequest20270101OneOf4460519136) HasTextFormats() bool {
	return o != nil && o.TextFormats != nil
}

// SetTextFormats gets a reference to the given [][]TextFormattingRule_20270101 and assigns it to the TextFormats field.
func (o *QueryTableBackendRequest20270101OneOf4460519136) SetTextFormats(v [][]TextFormattingRule_20270101) {
	o.TextFormats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryTableBackendRequest20270101OneOf4460519136) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.CellDisplayMode != nil {
		toSerialize["cell_display_mode"] = o.CellDisplayMode
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.TextFormats != nil {
		toSerialize["text_formats"] = o.TextFormats
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryTableBackendRequest20270101OneOf4460519136) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator         *ScalarAggregator_20270101                                               `json:"aggregator,omitempty"`
		Alias              *string                                                                  `json:"alias,omitempty"`
		CellDisplayMode    []CellDisplayModeOption_20270101                                         `json:"cell_display_mode,omitempty"`
		ConditionalFormats []ThresholdFormattingRule_20270101                                       `json:"conditional_formats,omitempty"`
		Formulas           []WidgetFormula_20270101                                                 `json:"formulas,omitempty"`
		Limit              *float64                                                                 `json:"limit,omitempty"`
		Order              *QuerySortOrder_20270101                                                 `json:"order,omitempty"`
		Queries            *[]FormulaAndFunctionQueryDefinition_20270101                            `json:"queries"`
		ResponseFormat     *QueryTableBackendRequestAnyOfAnyOfItemResponseFormat4450272256_20270101 `json:"response_format"`
		Sort               *ScalarFormulaSort_20270101                                              `json:"sort,omitempty"`
		TextFormats        [][]TextFormattingRule_20270101                                          `json:"text_formats,omitempty"`
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
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.Alias = all.Alias
	o.CellDisplayMode = all.CellDisplayMode
	o.ConditionalFormats = all.ConditionalFormats
	o.Formulas = all.Formulas
	o.Limit = all.Limit
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}
	o.Queries = *all.Queries
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}
	o.Sort = all.Sort
	o.TextFormats = all.TextFormats

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
