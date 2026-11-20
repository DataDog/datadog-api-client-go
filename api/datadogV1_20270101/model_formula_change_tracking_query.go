// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaChangeTrackingQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaChangeTrackingQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaChangeTrackingQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []FormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search FormulaChangeTrackingQuerySearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaChangeTrackingQuery instantiates a new FormulaChangeTrackingQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaChangeTrackingQuery(compute FormulaAndFunctionEventQueryDefinitionCompute, dataSource FormulaChangeTrackingQueryDataSource, name string, search FormulaChangeTrackingQuerySearch) *FormulaChangeTrackingQuery {
	this := FormulaChangeTrackingQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Search = search
	return &this
}

// NewFormulaChangeTrackingQueryWithDefaults instantiates a new FormulaChangeTrackingQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaChangeTrackingQueryWithDefaults() *FormulaChangeTrackingQuery {
	this := FormulaChangeTrackingQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaChangeTrackingQuery) GetCompute() FormulaAndFunctionEventQueryDefinitionCompute {
	if o == nil {
		var ret FormulaAndFunctionEventQueryDefinitionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaChangeTrackingQuery) GetComputeOk() (*FormulaAndFunctionEventQueryDefinitionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaChangeTrackingQuery) SetCompute(v FormulaAndFunctionEventQueryDefinitionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaChangeTrackingQuery) GetDataSource() FormulaChangeTrackingQueryDataSource {
	if o == nil {
		var ret FormulaChangeTrackingQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaChangeTrackingQuery) GetDataSourceOk() (*FormulaChangeTrackingQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaChangeTrackingQuery) SetDataSource(v FormulaChangeTrackingQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaChangeTrackingQuery) GetGroupBy() []FormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []FormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaChangeTrackingQuery) GetGroupByOk() (*[]FormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaChangeTrackingQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []FormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *FormulaChangeTrackingQuery) SetGroupBy(v []FormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaChangeTrackingQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaChangeTrackingQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaChangeTrackingQuery) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value.
func (o *FormulaChangeTrackingQuery) GetSearch() FormulaChangeTrackingQuerySearch {
	if o == nil {
		var ret FormulaChangeTrackingQuerySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *FormulaChangeTrackingQuery) GetSearchOk() (*FormulaChangeTrackingQuerySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *FormulaChangeTrackingQuery) SetSearch(v FormulaChangeTrackingQuerySearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaChangeTrackingQuery) MarshalJSON() ([]byte, error) {
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
func (o *FormulaChangeTrackingQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute    *FormulaAndFunctionEventQueryDefinitionCompute `json:"compute"`
		DataSource *FormulaChangeTrackingQueryDataSource          `json:"data_source"`
		GroupBy    []FormulaAndFunctionEventQueryGroupBy          `json:"group_by,omitempty"`
		Name       *string                                        `json:"name"`
		Search     *FormulaChangeTrackingQuerySearch              `json:"search"`
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
