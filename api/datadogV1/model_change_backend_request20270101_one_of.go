// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeBackendRequest20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type ChangeBackendRequest20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ChangeType *ChangeTypeOptions_20270101 `json:"change_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Formulas []WidgetFormula_20270101 `json:"formulas,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IncreaseGood *bool `json:"increase_good,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OrderBy *OrderByOptions_20270101 `json:"order_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OrderDir *QuerySortOrder_20270101 `json:"order_dir,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Queries []FormulaAndFunctionQueryDefinition_20270101 `json:"queries"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ResponseFormat ChangeBackendRequestAnyOfAnyOfItemResponseFormat_20270101 `json:"response_format"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowPresent *bool `json:"show_present,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ScalarFormulaSort_20270101 `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChangeBackendRequest20270101OneOf instantiates a new ChangeBackendRequest20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeBackendRequest20270101OneOf(queries []FormulaAndFunctionQueryDefinition_20270101, responseFormat ChangeBackendRequestAnyOfAnyOfItemResponseFormat_20270101) *ChangeBackendRequest20270101OneOf {
	this := ChangeBackendRequest20270101OneOf{}
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewChangeBackendRequest20270101OneOfWithDefaults instantiates a new ChangeBackendRequest20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeBackendRequest20270101OneOfWithDefaults() *ChangeBackendRequest20270101OneOf {
	this := ChangeBackendRequest20270101OneOf{}
	return &this
}

// GetChangeType returns the ChangeType field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetChangeType() ChangeTypeOptions_20270101 {
	if o == nil || o.ChangeType == nil {
		var ret ChangeTypeOptions_20270101
		return ret
	}
	return *o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetChangeTypeOk() (*ChangeTypeOptions_20270101, bool) {
	if o == nil || o.ChangeType == nil {
		return nil, false
	}
	return o.ChangeType, true
}

// HasChangeType returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasChangeType() bool {
	return o != nil && o.ChangeType != nil
}

// SetChangeType gets a reference to the given ChangeTypeOptions_20270101 and assigns it to the ChangeType field.
func (o *ChangeBackendRequest20270101OneOf) SetChangeType(v ChangeTypeOptions_20270101) {
	o.ChangeType = &v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetFormulas() []WidgetFormula_20270101 {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula_20270101
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetFormulasOk() (*[]WidgetFormula_20270101, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula_20270101 and assigns it to the Formulas field.
func (o *ChangeBackendRequest20270101OneOf) SetFormulas(v []WidgetFormula_20270101) {
	o.Formulas = v
}

// GetIncreaseGood returns the IncreaseGood field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetIncreaseGood() bool {
	if o == nil || o.IncreaseGood == nil {
		var ret bool
		return ret
	}
	return *o.IncreaseGood
}

// GetIncreaseGoodOk returns a tuple with the IncreaseGood field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetIncreaseGoodOk() (*bool, bool) {
	if o == nil || o.IncreaseGood == nil {
		return nil, false
	}
	return o.IncreaseGood, true
}

// HasIncreaseGood returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasIncreaseGood() bool {
	return o != nil && o.IncreaseGood != nil
}

// SetIncreaseGood gets a reference to the given bool and assigns it to the IncreaseGood field.
func (o *ChangeBackendRequest20270101OneOf) SetIncreaseGood(v bool) {
	o.IncreaseGood = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *ChangeBackendRequest20270101OneOf) SetLimit(v float64) {
	o.Limit = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetOrderBy() OrderByOptions_20270101 {
	if o == nil || o.OrderBy == nil {
		var ret OrderByOptions_20270101
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetOrderByOk() (*OrderByOptions_20270101, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasOrderBy() bool {
	return o != nil && o.OrderBy != nil
}

// SetOrderBy gets a reference to the given OrderByOptions_20270101 and assigns it to the OrderBy field.
func (o *ChangeBackendRequest20270101OneOf) SetOrderBy(v OrderByOptions_20270101) {
	o.OrderBy = &v
}

// GetOrderDir returns the OrderDir field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetOrderDir() QuerySortOrder_20270101 {
	if o == nil || o.OrderDir == nil {
		var ret QuerySortOrder_20270101
		return ret
	}
	return *o.OrderDir
}

// GetOrderDirOk returns a tuple with the OrderDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetOrderDirOk() (*QuerySortOrder_20270101, bool) {
	if o == nil || o.OrderDir == nil {
		return nil, false
	}
	return o.OrderDir, true
}

// HasOrderDir returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasOrderDir() bool {
	return o != nil && o.OrderDir != nil
}

// SetOrderDir gets a reference to the given QuerySortOrder_20270101 and assigns it to the OrderDir field.
func (o *ChangeBackendRequest20270101OneOf) SetOrderDir(v QuerySortOrder_20270101) {
	o.OrderDir = &v
}

// GetQueries returns the Queries field value.
func (o *ChangeBackendRequest20270101OneOf) GetQueries() []FormulaAndFunctionQueryDefinition_20270101 {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition_20270101
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *ChangeBackendRequest20270101OneOf) SetQueries(v []FormulaAndFunctionQueryDefinition_20270101) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *ChangeBackendRequest20270101OneOf) GetResponseFormat() ChangeBackendRequestAnyOfAnyOfItemResponseFormat_20270101 {
	if o == nil {
		var ret ChangeBackendRequestAnyOfAnyOfItemResponseFormat_20270101
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetResponseFormatOk() (*ChangeBackendRequestAnyOfAnyOfItemResponseFormat_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *ChangeBackendRequest20270101OneOf) SetResponseFormat(v ChangeBackendRequestAnyOfAnyOfItemResponseFormat_20270101) {
	o.ResponseFormat = v
}

// GetShowPresent returns the ShowPresent field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetShowPresent() bool {
	if o == nil || o.ShowPresent == nil {
		var ret bool
		return ret
	}
	return *o.ShowPresent
}

// GetShowPresentOk returns a tuple with the ShowPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetShowPresentOk() (*bool, bool) {
	if o == nil || o.ShowPresent == nil {
		return nil, false
	}
	return o.ShowPresent, true
}

// HasShowPresent returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasShowPresent() bool {
	return o != nil && o.ShowPresent != nil
}

// SetShowPresent gets a reference to the given bool and assigns it to the ShowPresent field.
func (o *ChangeBackendRequest20270101OneOf) SetShowPresent(v bool) {
	o.ShowPresent = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf) GetSort() ScalarFormulaSort_20270101 {
	if o == nil || o.Sort == nil {
		var ret ScalarFormulaSort_20270101
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf) GetSortOk() (*ScalarFormulaSort_20270101, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ScalarFormulaSort_20270101 and assigns it to the Sort field.
func (o *ChangeBackendRequest20270101OneOf) SetSort(v ScalarFormulaSort_20270101) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeBackendRequest20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeType != nil {
		toSerialize["change_type"] = o.ChangeType
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.IncreaseGood != nil {
		toSerialize["increase_good"] = o.IncreaseGood
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}
	if o.OrderDir != nil {
		toSerialize["order_dir"] = o.OrderDir
	}
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat
	if o.ShowPresent != nil {
		toSerialize["show_present"] = o.ShowPresent
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeBackendRequest20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeType     *ChangeTypeOptions_20270101                                `json:"change_type,omitempty"`
		Formulas       []WidgetFormula_20270101                                   `json:"formulas,omitempty"`
		IncreaseGood   *bool                                                      `json:"increase_good,omitempty"`
		Limit          *float64                                                   `json:"limit,omitempty"`
		OrderBy        *OrderByOptions_20270101                                   `json:"order_by,omitempty"`
		OrderDir       *QuerySortOrder_20270101                                   `json:"order_dir,omitempty"`
		Queries        *[]FormulaAndFunctionQueryDefinition_20270101              `json:"queries"`
		ResponseFormat *ChangeBackendRequestAnyOfAnyOfItemResponseFormat_20270101 `json:"response_format"`
		ShowPresent    *bool                                                      `json:"show_present,omitempty"`
		Sort           *ScalarFormulaSort_20270101                                `json:"sort,omitempty"`
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
	if all.ChangeType != nil && !all.ChangeType.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeType = all.ChangeType
	}
	o.Formulas = all.Formulas
	o.IncreaseGood = all.IncreaseGood
	o.Limit = all.Limit
	if all.OrderBy != nil && !all.OrderBy.IsValid() {
		hasInvalidField = true
	} else {
		o.OrderBy = all.OrderBy
	}
	if all.OrderDir != nil && !all.OrderDir.IsValid() {
		hasInvalidField = true
	} else {
		o.OrderDir = all.OrderDir
	}
	o.Queries = *all.Queries
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}
	o.ShowPresent = all.ShowPresent
	o.Sort = all.Sort

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
