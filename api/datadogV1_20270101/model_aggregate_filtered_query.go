// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateFilteredQuery Test-only generated schema for the dashboard versioning proof of concept.
type AggregateFilteredQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery CompositeEventsPlatformQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute []CompositeCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource AggregateFilteredQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	FilterQuery CompositeSourceSubQuery `json:"filter_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filters []AggregateFilteredQueryFiltersItems `json:"filters"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []CompositeGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAggregateFilteredQuery instantiates a new AggregateFilteredQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregateFilteredQuery(baseQuery CompositeEventsPlatformQuery, compute []CompositeCompute, dataSource AggregateFilteredQueryDataSource, filterQuery CompositeSourceSubQuery, filters []AggregateFilteredQueryFiltersItems, name string) *AggregateFilteredQuery {
	this := AggregateFilteredQuery{}
	this.BaseQuery = baseQuery
	this.Compute = compute
	this.DataSource = dataSource
	this.FilterQuery = filterQuery
	this.Filters = filters
	this.Name = name
	return &this
}

// NewAggregateFilteredQueryWithDefaults instantiates a new AggregateFilteredQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregateFilteredQueryWithDefaults() *AggregateFilteredQuery {
	this := AggregateFilteredQuery{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *AggregateFilteredQuery) GetBaseQuery() CompositeEventsPlatformQuery {
	if o == nil {
		var ret CompositeEventsPlatformQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *AggregateFilteredQuery) GetBaseQueryOk() (*CompositeEventsPlatformQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *AggregateFilteredQuery) SetBaseQuery(v CompositeEventsPlatformQuery) {
	o.BaseQuery = v
}

// GetCompute returns the Compute field value.
func (o *AggregateFilteredQuery) GetCompute() []CompositeCompute {
	if o == nil {
		var ret []CompositeCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *AggregateFilteredQuery) GetComputeOk() (*[]CompositeCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *AggregateFilteredQuery) SetCompute(v []CompositeCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *AggregateFilteredQuery) GetDataSource() AggregateFilteredQueryDataSource {
	if o == nil {
		var ret AggregateFilteredQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *AggregateFilteredQuery) GetDataSourceOk() (*AggregateFilteredQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *AggregateFilteredQuery) SetDataSource(v AggregateFilteredQueryDataSource) {
	o.DataSource = v
}

// GetFilterQuery returns the FilterQuery field value.
func (o *AggregateFilteredQuery) GetFilterQuery() CompositeSourceSubQuery {
	if o == nil {
		var ret CompositeSourceSubQuery
		return ret
	}
	return o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value
// and a boolean to check if the value has been set.
func (o *AggregateFilteredQuery) GetFilterQueryOk() (*CompositeSourceSubQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterQuery, true
}

// SetFilterQuery sets field value.
func (o *AggregateFilteredQuery) SetFilterQuery(v CompositeSourceSubQuery) {
	o.FilterQuery = v
}

// GetFilters returns the Filters field value.
func (o *AggregateFilteredQuery) GetFilters() []AggregateFilteredQueryFiltersItems {
	if o == nil {
		var ret []AggregateFilteredQueryFiltersItems
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *AggregateFilteredQuery) GetFiltersOk() (*[]AggregateFilteredQueryFiltersItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *AggregateFilteredQuery) SetFilters(v []AggregateFilteredQueryFiltersItems) {
	o.Filters = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *AggregateFilteredQuery) GetGroupBy() []CompositeGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []CompositeGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateFilteredQuery) GetGroupByOk() (*[]CompositeGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *AggregateFilteredQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []CompositeGroupBy and assigns it to the GroupBy field.
func (o *AggregateFilteredQuery) SetGroupBy(v []CompositeGroupBy) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *AggregateFilteredQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AggregateFilteredQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AggregateFilteredQuery) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregateFilteredQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	toSerialize["filter_query"] = o.FilterQuery
	toSerialize["filters"] = o.Filters
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregateFilteredQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery   *CompositeEventsPlatformQuery         `json:"base_query"`
		Compute     *[]CompositeCompute                   `json:"compute"`
		DataSource  *AggregateFilteredQueryDataSource     `json:"data_source"`
		FilterQuery *CompositeSourceSubQuery              `json:"filter_query"`
		Filters     *[]AggregateFilteredQueryFiltersItems `json:"filters"`
		GroupBy     []CompositeGroupBy                    `json:"group_by,omitempty"`
		Name        *string                               `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.FilterQuery == nil {
		return fmt.Errorf("required field filter_query missing")
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	if all.BaseQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BaseQuery = *all.BaseQuery
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.FilterQuery = *all.FilterQuery
	o.Filters = *all.Filters
	o.GroupBy = all.GroupBy
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
