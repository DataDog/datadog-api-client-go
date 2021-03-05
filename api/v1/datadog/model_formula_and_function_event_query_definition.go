/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// FormulaAndFunctionEventQueryDefinition A formula and functions events query.
type FormulaAndFunctionEventQueryDefinition struct {
	Compute    FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	DataSource FormulaAndFunctionEventsDataSource            `json:"data_source"`
	// Group by options.
	GroupBy *[]FormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// An array of index names to query in the stream. Omit or use `[]` to query all indexes at once.
	Indexes *[]string `json:"indexes,omitempty"`
	// Name of the query for use in formulas.
	Name   *string                                       `json:"name,omitempty"`
	Search *FormulaAndFunctionEventQueryDefinitionSearch `json:"search,omitempty"`
}

// NewFormulaAndFunctionEventQueryDefinition instantiates a new FormulaAndFunctionEventQueryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormulaAndFunctionEventQueryDefinition(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaAndFunctionEventsDataSource) *FormulaAndFunctionEventQueryDefinition {
	this := FormulaAndFunctionEventQueryDefinition{}
	this.Compute = compute
	this.DataSource = dataSource
	return &this
}

// NewFormulaAndFunctionEventQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionEventQueryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormulaAndFunctionEventQueryDefinitionWithDefaults() *FormulaAndFunctionEventQueryDefinition {
	this := FormulaAndFunctionEventQueryDefinition{}
	return &this
}

// GetCompute returns the Compute field value
func (o *FormulaAndFunctionEventQueryDefinition) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret FormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}

	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value
func (o *FormulaAndFunctionEventQueryDefinition) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value
func (o *FormulaAndFunctionEventQueryDefinition) GetDataSource() FormulaAndFunctionEventsDataSource {
	if o == nil {
		var ret FormulaAndFunctionEventsDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionEventsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *FormulaAndFunctionEventQueryDefinition) SetDataSource(v FormulaAndFunctionEventsDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetGroupBy() []FormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []FormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetGroupByOk() (*[]FormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []FormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *FormulaAndFunctionEventQueryDefinition) SetGroupBy(v []FormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return *o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasIndexes() bool {
	if o != nil && o.Indexes != nil {
		return true
	}

	return false
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *FormulaAndFunctionEventQueryDefinition) SetIndexes(v []string) {
	o.Indexes = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormulaAndFunctionEventQueryDefinition) SetName(v string) {
	o.Name = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FormulaAndFunctionEventQueryDefinition) GetSearch() FormulaAndFunctionEventQueryDefinitionSearch {
	if o == nil || o.Search == nil {
		var ret FormulaAndFunctionEventQueryDefinitionSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionEventQueryDefinition) GetSearchOk() (*FormulaAndFunctionEventQueryDefinitionSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FormulaAndFunctionEventQueryDefinition) HasSearch() bool {
	if o != nil && o.Search != nil {
		return true
	}

	return false
}

// SetSearch gets a reference to the given FormulaAndFunctionEventQueryDefinitionSearch and assigns it to the Search field.
func (o *FormulaAndFunctionEventQueryDefinition) SetSearch(v FormulaAndFunctionEventQueryDefinitionSearch) {
	o.Search = &v
}

func (o FormulaAndFunctionEventQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["compute"] = o.Compute
	}
	if true {
		toSerialize["data_source"] = o.DataSource
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	return json.Marshal(toSerialize)
}

type NullableFormulaAndFunctionEventQueryDefinition struct {
	value *FormulaAndFunctionEventQueryDefinition
	isSet bool
}

func (v NullableFormulaAndFunctionEventQueryDefinition) Get() *FormulaAndFunctionEventQueryDefinition {
	return v.value
}

func (v *NullableFormulaAndFunctionEventQueryDefinition) Set(val *FormulaAndFunctionEventQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableFormulaAndFunctionEventQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableFormulaAndFunctionEventQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormulaAndFunctionEventQueryDefinition(val *FormulaAndFunctionEventQueryDefinition) *NullableFormulaAndFunctionEventQueryDefinition {
	return &NullableFormulaAndFunctionEventQueryDefinition{value: val, isSet: true}
}

func (v NullableFormulaAndFunctionEventQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormulaAndFunctionEventQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
