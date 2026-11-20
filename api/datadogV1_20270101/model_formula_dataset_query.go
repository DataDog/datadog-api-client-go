// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaDatasetQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaDatasetQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []DatasetCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaDatasetQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetId string `json:"dataset_id"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DatasetProvider PublishedDatasetProvider `json:"dataset_provider"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []DatasetGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort []DatasetSort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeColumn *string `json:"time_column,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaDatasetQuery instantiates a new FormulaDatasetQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaDatasetQuery(compute []DatasetCompute, dataSource FormulaDatasetQueryDataSource, datasetId string, datasetProvider PublishedDatasetProvider, name string) *FormulaDatasetQuery {
	this := FormulaDatasetQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.DatasetId = datasetId
	this.DatasetProvider = datasetProvider
	this.Name = name
	return &this
}

// NewFormulaDatasetQueryWithDefaults instantiates a new FormulaDatasetQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaDatasetQueryWithDefaults() *FormulaDatasetQuery {
	this := FormulaDatasetQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaDatasetQuery) GetCompute() []DatasetCompute {
	if o == nil {
		var ret []DatasetCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetComputeOk() (*[]DatasetCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaDatasetQuery) SetCompute(v []DatasetCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaDatasetQuery) GetDataSource() FormulaDatasetQueryDataSource {
	if o == nil {
		var ret FormulaDatasetQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetDataSourceOk() (*FormulaDatasetQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaDatasetQuery) SetDataSource(v FormulaDatasetQueryDataSource) {
	o.DataSource = v
}

// GetDatasetId returns the DatasetId field value.
func (o *FormulaDatasetQuery) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *FormulaDatasetQuery) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetDatasetProvider returns the DatasetProvider field value.
func (o *FormulaDatasetQuery) GetDatasetProvider() PublishedDatasetProvider {
	if o == nil {
		var ret PublishedDatasetProvider
		return ret
	}
	return o.DatasetProvider
}

// GetDatasetProviderOk returns a tuple with the DatasetProvider field value
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetDatasetProviderOk() (*PublishedDatasetProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetProvider, true
}

// SetDatasetProvider sets field value.
func (o *FormulaDatasetQuery) SetDatasetProvider(v PublishedDatasetProvider) {
	o.DatasetProvider = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaDatasetQuery) GetGroupBy() []DatasetGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []DatasetGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetGroupByOk() (*[]DatasetGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaDatasetQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []DatasetGroupBy and assigns it to the GroupBy field.
func (o *FormulaDatasetQuery) SetGroupBy(v []DatasetGroupBy) {
	o.GroupBy = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FormulaDatasetQuery) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FormulaDatasetQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *FormulaDatasetQuery) SetLimit(v float64) {
	o.Limit = &v
}

// GetName returns the Name field value.
func (o *FormulaDatasetQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaDatasetQuery) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *FormulaDatasetQuery) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *FormulaDatasetQuery) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *FormulaDatasetQuery) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *FormulaDatasetQuery) GetSort() []DatasetSort {
	if o == nil || o.Sort == nil {
		var ret []DatasetSort
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetSortOk() (*[]DatasetSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return &o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *FormulaDatasetQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given []DatasetSort and assigns it to the Sort field.
func (o *FormulaDatasetQuery) SetSort(v []DatasetSort) {
	o.Sort = v
}

// GetTimeColumn returns the TimeColumn field value if set, zero value otherwise.
func (o *FormulaDatasetQuery) GetTimeColumn() string {
	if o == nil || o.TimeColumn == nil {
		var ret string
		return ret
	}
	return *o.TimeColumn
}

// GetTimeColumnOk returns a tuple with the TimeColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaDatasetQuery) GetTimeColumnOk() (*string, bool) {
	if o == nil || o.TimeColumn == nil {
		return nil, false
	}
	return o.TimeColumn, true
}

// HasTimeColumn returns a boolean if a field has been set.
func (o *FormulaDatasetQuery) HasTimeColumn() bool {
	return o != nil && o.TimeColumn != nil
}

// SetTimeColumn gets a reference to the given string and assigns it to the TimeColumn field.
func (o *FormulaDatasetQuery) SetTimeColumn(v string) {
	o.TimeColumn = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaDatasetQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["dataset_provider"] = o.DatasetProvider
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["name"] = o.Name
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.TimeColumn != nil {
		toSerialize["time_column"] = o.TimeColumn
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaDatasetQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute         *[]DatasetCompute              `json:"compute"`
		DataSource      *FormulaDatasetQueryDataSource `json:"data_source"`
		DatasetId       *string                        `json:"dataset_id"`
		DatasetProvider *PublishedDatasetProvider      `json:"dataset_provider"`
		GroupBy         []DatasetGroupBy               `json:"group_by,omitempty"`
		Limit           *float64                       `json:"limit,omitempty"`
		Name            *string                        `json:"name"`
		QueryFilter     *string                        `json:"query_filter,omitempty"`
		Sort            []DatasetSort                  `json:"sort,omitempty"`
		TimeColumn      *string                        `json:"time_column,omitempty"`
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
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if all.DatasetProvider == nil {
		return fmt.Errorf("required field dataset_provider missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.DatasetId = *all.DatasetId
	if !all.DatasetProvider.IsValid() {
		hasInvalidField = true
	} else {
		o.DatasetProvider = *all.DatasetProvider
	}
	o.GroupBy = all.GroupBy
	o.Limit = all.Limit
	o.Name = *all.Name
	o.QueryFilter = all.QueryFilter
	o.Sort = all.Sort
	o.TimeColumn = all.TimeColumn

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
