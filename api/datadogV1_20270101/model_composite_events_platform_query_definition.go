// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CompositeEventsPlatformQueryDefinition Test-only generated schema for the dashboard versioning proof of concept.
type CompositeEventsPlatformQueryDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Columns []ListColumn `json:"columns,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []AggregateCompute `json:"compute,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ComputeCount *bool `json:"compute_count,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaAndFunctionEventsDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []FormulaAndFunctionEventQueryGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search *CompositeEventsPlatformQueryDefinitionSearch `json:"search,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Storage *StorageType `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCompositeEventsPlatformQueryDefinition instantiates a new CompositeEventsPlatformQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCompositeEventsPlatformQueryDefinition(dataSource FormulaAndFunctionEventsDataSource, name string) *CompositeEventsPlatformQueryDefinition {
	this := CompositeEventsPlatformQueryDefinition{}
	this.DataSource = dataSource
	this.Name = name
	return &this
}

// NewCompositeEventsPlatformQueryDefinitionWithDefaults instantiates a new CompositeEventsPlatformQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCompositeEventsPlatformQueryDefinitionWithDefaults() *CompositeEventsPlatformQueryDefinition {
	this := CompositeEventsPlatformQueryDefinition{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *CompositeEventsPlatformQueryDefinition) GetColumns() []ListColumn {
	if o == nil || o.Columns == nil {
		var ret []ListColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetColumnsOk() (*[]ListColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *CompositeEventsPlatformQueryDefinition) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []ListColumn and assigns it to the Columns field.
func (o *CompositeEventsPlatformQueryDefinition) SetColumns(v []ListColumn) {
	o.Columns = v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *CompositeEventsPlatformQueryDefinition) GetCompute() []AggregateCompute {
	if o == nil || o.Compute == nil {
		var ret []AggregateCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetComputeOk() (*[]AggregateCompute, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *CompositeEventsPlatformQueryDefinition) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []AggregateCompute and assigns it to the Compute field.
func (o *CompositeEventsPlatformQueryDefinition) SetCompute(v []AggregateCompute) {
	o.Compute = v
}

// GetComputeCount returns the ComputeCount field value if set, zero value otherwise.
func (o *CompositeEventsPlatformQueryDefinition) GetComputeCount() bool {
	if o == nil || o.ComputeCount == nil {
		var ret bool
		return ret
	}
	return *o.ComputeCount
}

// GetComputeCountOk returns a tuple with the ComputeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetComputeCountOk() (*bool, bool) {
	if o == nil || o.ComputeCount == nil {
		return nil, false
	}
	return o.ComputeCount, true
}

// HasComputeCount returns a boolean if a field has been set.
func (o *CompositeEventsPlatformQueryDefinition) HasComputeCount() bool {
	return o != nil && o.ComputeCount != nil
}

// SetComputeCount gets a reference to the given bool and assigns it to the ComputeCount field.
func (o *CompositeEventsPlatformQueryDefinition) SetComputeCount(v bool) {
	o.ComputeCount = &v
}

// GetDataSource returns the DataSource field value.
func (o *CompositeEventsPlatformQueryDefinition) GetDataSource() FormulaAndFunctionEventsDataSource {
	if o == nil {
		var ret FormulaAndFunctionEventsDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionEventsDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *CompositeEventsPlatformQueryDefinition) SetDataSource(v FormulaAndFunctionEventsDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *CompositeEventsPlatformQueryDefinition) GetGroupBy() []FormulaAndFunctionEventQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []FormulaAndFunctionEventQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetGroupByOk() (*[]FormulaAndFunctionEventQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *CompositeEventsPlatformQueryDefinition) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []FormulaAndFunctionEventQueryGroupBy and assigns it to the GroupBy field.
func (o *CompositeEventsPlatformQueryDefinition) SetGroupBy(v []FormulaAndFunctionEventQueryGroupBy) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *CompositeEventsPlatformQueryDefinition) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *CompositeEventsPlatformQueryDefinition) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *CompositeEventsPlatformQueryDefinition) SetIndexes(v []string) {
	o.Indexes = v
}

// GetName returns the Name field value.
func (o *CompositeEventsPlatformQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CompositeEventsPlatformQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *CompositeEventsPlatformQueryDefinition) GetSearch() CompositeEventsPlatformQueryDefinitionSearch {
	if o == nil || o.Search == nil {
		var ret CompositeEventsPlatformQueryDefinitionSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetSearchOk() (*CompositeEventsPlatformQueryDefinitionSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *CompositeEventsPlatformQueryDefinition) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given CompositeEventsPlatformQueryDefinitionSearch and assigns it to the Search field.
func (o *CompositeEventsPlatformQueryDefinition) SetSearch(v CompositeEventsPlatformQueryDefinitionSearch) {
	o.Search = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *CompositeEventsPlatformQueryDefinition) GetStorage() StorageType {
	if o == nil || o.Storage == nil {
		var ret StorageType
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeEventsPlatformQueryDefinition) GetStorageOk() (*StorageType, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *CompositeEventsPlatformQueryDefinition) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageType and assigns it to the Storage field.
func (o *CompositeEventsPlatformQueryDefinition) SetStorage(v StorageType) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CompositeEventsPlatformQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	if o.ComputeCount != nil {
		toSerialize["compute_count"] = o.ComputeCount
	}
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
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CompositeEventsPlatformQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns      []ListColumn                                  `json:"columns,omitempty"`
		Compute      []AggregateCompute                            `json:"compute,omitempty"`
		ComputeCount *bool                                         `json:"compute_count,omitempty"`
		DataSource   *FormulaAndFunctionEventsDataSource           `json:"data_source"`
		GroupBy      []FormulaAndFunctionEventQueryGroupBy         `json:"group_by,omitempty"`
		Indexes      []string                                      `json:"indexes,omitempty"`
		Name         *string                                       `json:"name"`
		Search       *CompositeEventsPlatformQueryDefinitionSearch `json:"search,omitempty"`
		Storage      *StorageType                                  `json:"storage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Columns = all.Columns
	o.Compute = all.Compute
	o.ComputeCount = all.ComputeCount
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
	if all.Storage != nil && !all.Storage.IsValid() {
		hasInvalidField = true
	} else {
		o.Storage = all.Storage
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
