// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequest20270101OneOf4463750144 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequest20270101OneOf4463750144 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TreemapRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreemapBackendRequest20270101OneOf4463750144 instantiates a new TreemapBackendRequest20270101OneOf4463750144 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreemapBackendRequest20270101OneOf4463750144(queries []FormulaAndFunctionQueryDefinition, responseFormat TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168) *TreemapBackendRequest20270101OneOf4463750144 {
	this := TreemapBackendRequest20270101OneOf4463750144{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewTreemapBackendRequest20270101OneOf4463750144WithDefaults instantiates a new TreemapBackendRequest20270101OneOf4463750144 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreemapBackendRequest20270101OneOf4463750144WithDefaults() *TreemapBackendRequest20270101OneOf4463750144 {
	this := TreemapBackendRequest20270101OneOf4463750144{}
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *TreemapBackendRequest20270101OneOf4463750144) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *TreemapBackendRequest20270101OneOf4463750144) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetResponseFormat() TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 {
	if o == nil {
		var ret TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetResponseFormatOk() (*TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *TreemapBackendRequest20270101OneOf4463750144) SetResponseFormat(v TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168) {
	o.ResponseFormat = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetSort() ScalarFormulaSort {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetSortOk() (*ScalarFormulaSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort and assigns it to the Sort field.
func (o *TreemapBackendRequest20270101OneOf4463750144) SetSort(v ScalarFormulaSort) {
	o.Sort = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetStyle() TreemapRequestStyle {
	if o == nil || o.Style == nil {
		var ret TreemapRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) GetStyleOk() (*TreemapRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TreemapBackendRequest20270101OneOf4463750144) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TreemapRequestStyle and assigns it to the Style field.
func (o *TreemapBackendRequest20270101OneOf4463750144) SetStyle(v TreemapRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreemapBackendRequest20270101OneOf4463750144) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *TreemapBackendRequest20270101OneOf4463750144) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []WidgetFormula                                              `json:"formulas,omitempty"`
		Queries        *[]FormulaAndFunctionQueryDefinition                         `json:"queries"`
		ResponseFormat *TreemapBackendRequestAnyOfAnyOfItemResponseFormat4453527168 `json:"response_format"`
		Sort           *ScalarFormulaSort                                           `json:"sort,omitempty"`
		Style          *TreemapRequestStyle                                         `json:"style,omitempty"`
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
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
