// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaIncidentsQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaIncidentsQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaIncidentsQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []FormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search *FormulaIncidentsQuerySearch `json:"search,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaIncidentsQuery instantiates a new FormulaIncidentsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaIncidentsQuery(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaIncidentsQueryDataSource, name string) *FormulaIncidentsQuery {
	this := FormulaIncidentsQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewFormulaIncidentsQueryWithDefaults instantiates a new FormulaIncidentsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaIncidentsQueryWithDefaults() *FormulaIncidentsQuery {
	this := FormulaIncidentsQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaIncidentsQuery) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret FormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaIncidentsQuery) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaIncidentsQuery) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaIncidentsQuery) GetDataSource() FormulaIncidentsQueryDataSource {
	if o == nil {
		var ret FormulaIncidentsQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaIncidentsQuery) GetDataSourceOk() (*FormulaIncidentsQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaIncidentsQuery) SetDataSource(v FormulaIncidentsQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaIncidentsQuery) GetGroupBy() []FormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []FormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaIncidentsQuery) GetGroupByOk() (*[]FormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaIncidentsQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []FormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *FormulaIncidentsQuery) SetGroupBy(v []FormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *FormulaIncidentsQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaIncidentsQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *FormulaIncidentsQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *FormulaIncidentsQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *FormulaIncidentsQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaIncidentsQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaIncidentsQuery) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FormulaIncidentsQuery) GetSearch() FormulaIncidentsQuerySearch {
	if o == nil || o.Search == nil {
		var ret FormulaIncidentsQuerySearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaIncidentsQuery) GetSearchOk() (*FormulaIncidentsQuerySearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FormulaIncidentsQuery) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given FormulaIncidentsQuerySearch and assigns it to the Search field.
func (o *FormulaIncidentsQuery) SetSearch(v FormulaIncidentsQuerySearch) {
	o.Search = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaIncidentsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["name"] = o.Name
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaIncidentsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		DataSource *FormulaIncidentsQueryDataSource               `json:"data_source"`
		GroupBy    []FormulaAndFunctionEventQueryGroupBy          `json:"group_by,omitempty"`
		Indexes    []string                                       `json:"indexes,omitempty"`
		Name       *string                                        `json:"name"`
		Search     *FormulaIncidentsQuerySearch                   `json:"search,omitempty"`
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
	o.Indexes = all.Indexes
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
