// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDetectionRulesQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDetectionRulesQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaDetectionRulesQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy *FormulaDetectionRulesQuery20270101GroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search FormulaDetectionRulesQuerySearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaDetectionRulesQuery instantiates a new FormulaDetectionRulesQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaDetectionRulesQuery(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaDetectionRulesQueryDataSource, name string, search FormulaDetectionRulesQuerySearch) *FormulaDetectionRulesQuery {
	this := FormulaDetectionRulesQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Search = search
	return &this
}

// NewFormulaDetectionRulesQueryWithDefaults instantiates a new FormulaDetectionRulesQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaDetectionRulesQueryWithDefaults() *FormulaDetectionRulesQuery {
	this := FormulaDetectionRulesQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaDetectionRulesQuery) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret FormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaDetectionRulesQuery) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaDetectionRulesQuery) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaDetectionRulesQuery) GetDataSource() FormulaDetectionRulesQueryDataSource {
	if o == nil {
		var ret FormulaDetectionRulesQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaDetectionRulesQuery) GetDataSourceOk() (*FormulaDetectionRulesQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaDetectionRulesQuery) SetDataSource(v FormulaDetectionRulesQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaDetectionRulesQuery) GetGroupBy() FormulaDetectionRulesQuery20270101GroupBy {
	if o == nil || o.GroupBy == nil {
		var ret FormulaDetectionRulesQuery20270101GroupBy
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDetectionRulesQuery) GetGroupByOk() (*FormulaDetectionRulesQuery20270101GroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaDetectionRulesQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given FormulaDetectionRulesQuery20270101GroupBy and assigns it to the GroupBy field.
func (o *FormulaDetectionRulesQuery) SetGroupBy(v FormulaDetectionRulesQuery20270101GroupBy) {
	o.GroupBy = &v
}

// GetName returns the Name field value.
func (o *FormulaDetectionRulesQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaDetectionRulesQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaDetectionRulesQuery) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value.
func (o *FormulaDetectionRulesQuery) GetSearch() FormulaDetectionRulesQuerySearch {
	if o == nil {
		var ret FormulaDetectionRulesQuerySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *FormulaDetectionRulesQuery) GetSearchOk() (*FormulaDetectionRulesQuerySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *FormulaDetectionRulesQuery) SetSearch(v FormulaDetectionRulesQuerySearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaDetectionRulesQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	toSerialize["search"] = o.Search
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaDetectionRulesQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		DataSource *FormulaDetectionRulesQueryDataSource          `json:"data_source"`
		GroupBy    *FormulaDetectionRulesQuery20270101GroupBy     `json:"group_by,omitempty"`
		Name       *string                                        `json:"name"`
		Search     *FormulaDetectionRulesQuerySearch              `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}

	hasInvalidField := false
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
