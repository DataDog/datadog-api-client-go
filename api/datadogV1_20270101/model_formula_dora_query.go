// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDoraQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDoraQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaDoraQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []FormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []interface{} `json:"indexes"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search *FormulaDoraQuerySearch `json:"search,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaDoraQuery instantiates a new FormulaDoraQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaDoraQuery(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaDoraQueryDataSource, indexes []interface{}, name string) *FormulaDoraQuery {
	this := FormulaDoraQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Indexes = indexes
	this.Name = name
	return &this
}

// NewFormulaDoraQueryWithDefaults instantiates a new FormulaDoraQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaDoraQueryWithDefaults() *FormulaDoraQuery {
	this := FormulaDoraQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaDoraQuery) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret FormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaDoraQuery) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaDoraQuery) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaDoraQuery) GetDataSource() FormulaDoraQueryDataSource {
	if o == nil {
		var ret FormulaDoraQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaDoraQuery) GetDataSourceOk() (*FormulaDoraQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaDoraQuery) SetDataSource(v FormulaDoraQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaDoraQuery) GetGroupBy() []FormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []FormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDoraQuery) GetGroupByOk() (*[]FormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaDoraQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []FormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *FormulaDoraQuery) SetGroupBy(v []FormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value.
func (o *FormulaDoraQuery) GetIndexes() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value
// and a boolean to check if the value has been set.
func (o *FormulaDoraQuery) GetIndexesOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// SetIndexes sets field value.
func (o *FormulaDoraQuery) SetIndexes(v []interface{}) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *FormulaDoraQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaDoraQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaDoraQuery) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FormulaDoraQuery) GetSearch() FormulaDoraQuerySearch {
	if o == nil || o.Search == nil {
		var ret FormulaDoraQuerySearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDoraQuery) GetSearchOk() (*FormulaDoraQuerySearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FormulaDoraQuery) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given FormulaDoraQuerySearch and assigns it to the Search field.
func (o *FormulaDoraQuery) SetSearch(v FormulaDoraQuerySearch) {
	o.Search = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaDoraQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["indexes"] = o.Indexes
	toSerialize["name"] = o.Name
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaDoraQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		DataSource *FormulaDoraQueryDataSource                    `json:"data_source"`
		GroupBy    []FormulaAndFunctionEventQueryGroupBy          `json:"group_by,omitempty"`
		Indexes    *[]interface{}                                 `json:"indexes"`
		Name       *string                                        `json:"name"`
		Search     *FormulaDoraQuerySearch                        `json:"search,omitempty"`
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
	if all.Indexes == nil {
		return fmt.Errorf("required field indexes missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
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
	o.Indexes = *all.Indexes
	o.Name = *all.Name
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
