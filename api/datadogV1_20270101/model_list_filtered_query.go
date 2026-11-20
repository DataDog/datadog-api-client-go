// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListFilteredQuery Test-only generated schema for the dashboard versioning proof of concept.
type ListFilteredQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery CompositeEventsPlatformQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Context *EventsContext `json:"context,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ListFilteredQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	FilterQuery CompositeSourceSubQuery `json:"filter_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filters []ListFilteredQueryFiltersItems `json:"filters"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewListFilteredQuery instantiates a new ListFilteredQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListFilteredQuery(baseQuery CompositeEventsPlatformQuery, dataSource ListFilteredQueryDataSource, filterQuery CompositeSourceSubQuery, filters []ListFilteredQueryFiltersItems, name string) *ListFilteredQuery {
	this := ListFilteredQuery{}
	this.BaseQuery = baseQuery
	this.DataSource = dataSource
	this.FilterQuery = filterQuery
	this.Filters = filters
	this.Name = name
	return &this
}

// NewListFilteredQueryWithDefaults instantiates a new ListFilteredQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListFilteredQueryWithDefaults() *ListFilteredQuery {
	this := ListFilteredQuery{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *ListFilteredQuery) GetBaseQuery() CompositeEventsPlatformQuery {
	if o == nil {
		var ret CompositeEventsPlatformQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *ListFilteredQuery) GetBaseQueryOk() (*CompositeEventsPlatformQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *ListFilteredQuery) SetBaseQuery(v CompositeEventsPlatformQuery) {
	o.BaseQuery = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListFilteredQuery) GetContext() EventsContext {
	if o == nil || o.Context == nil {
		var ret EventsContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFilteredQuery) GetContextOk() (*EventsContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListFilteredQuery) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given EventsContext and assigns it to the Context field.
func (o *ListFilteredQuery) SetContext(v EventsContext) {
	o.Context = &v
}

// GetDataSource returns the DataSource field value.
func (o *ListFilteredQuery) GetDataSource() ListFilteredQueryDataSource {
	if o == nil {
		var ret ListFilteredQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ListFilteredQuery) GetDataSourceOk() (*ListFilteredQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ListFilteredQuery) SetDataSource(v ListFilteredQueryDataSource) {
	o.DataSource = v
}

// GetFilterQuery returns the FilterQuery field value.
func (o *ListFilteredQuery) GetFilterQuery() CompositeSourceSubQuery {
	if o == nil {
		var ret CompositeSourceSubQuery
		return ret
	}
	return o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value
// and a boolean to check if the value has been set.
func (o *ListFilteredQuery) GetFilterQueryOk() (*CompositeSourceSubQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterQuery, true
}

// SetFilterQuery sets field value.
func (o *ListFilteredQuery) SetFilterQuery(v CompositeSourceSubQuery) {
	o.FilterQuery = v
}

// GetFilters returns the Filters field value.
func (o *ListFilteredQuery) GetFilters() []ListFilteredQueryFiltersItems {
	if o == nil {
		var ret []ListFilteredQueryFiltersItems
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ListFilteredQuery) GetFiltersOk() (*[]ListFilteredQueryFiltersItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *ListFilteredQuery) SetFilters(v []ListFilteredQueryFiltersItems) {
	o.Filters = v
}

// GetName returns the Name field value.
func (o *ListFilteredQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListFilteredQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ListFilteredQuery) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListFilteredQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	toSerialize["data_source"] = o.DataSource
	toSerialize["filter_query"] = o.FilterQuery
	toSerialize["filters"] = o.Filters
	toSerialize["name"] = o.Name
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListFilteredQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery   *CompositeEventsPlatformQuery    `json:"base_query"`
		Context     *EventsContext                   `json:"context,omitempty"`
		DataSource  *ListFilteredQueryDataSource     `json:"data_source"`
		FilterQuery *CompositeSourceSubQuery         `json:"filter_query"`
		Filters     *[]ListFilteredQueryFiltersItems `json:"filters"`
		Name        *string                          `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
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
	if all.Context != nil && all.Context.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Context = all.Context
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.FilterQuery = *all.FilterQuery
	o.Filters = *all.Filters
	o.Name = *all.Name

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
