// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaCaseQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaCaseQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute CaseCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaCaseQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []CaseGroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter string `json:"query_filter"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *QuerySortOrder `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaCaseQuery instantiates a new FormulaCaseQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaCaseQuery(compute CaseCompute, dataSource FormulaCaseQueryDataSource, name string, queryFilter string) *FormulaCaseQuery {
	this := FormulaCaseQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.QueryFilter = queryFilter
	return &this
}

// NewFormulaCaseQueryWithDefaults instantiates a new FormulaCaseQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaCaseQueryWithDefaults() *FormulaCaseQuery {
	this := FormulaCaseQuery{}
	return &this
}

// GetCompute returns the Compute field value.
func (o *FormulaCaseQuery) GetCompute() CaseCompute {
	if o == nil {
		var ret CaseCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaCaseQuery) GetComputeOk() (*CaseCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaCaseQuery) SetCompute(v CaseCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaCaseQuery) GetDataSource() FormulaCaseQueryDataSource {
	if o == nil {
		var ret FormulaCaseQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaCaseQuery) GetDataSourceOk() (*FormulaCaseQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaCaseQuery) SetDataSource(v FormulaCaseQueryDataSource) {
	o.DataSource = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaCaseQuery) GetGroupBy() []CaseGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []CaseGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaCaseQuery) GetGroupByOk() (*[]CaseGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaCaseQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []CaseGroupBy and assigns it to the GroupBy field.
func (o *FormulaCaseQuery) SetGroupBy(v []CaseGroupBy) {
	o.GroupBy = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FormulaCaseQuery) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaCaseQuery) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FormulaCaseQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *FormulaCaseQuery) SetLimit(v float64) {
	o.Limit = &v
}

// GetName returns the Name field value.
func (o *FormulaCaseQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaCaseQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaCaseQuery) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value.
func (o *FormulaCaseQuery) GetQueryFilter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value
// and a boolean to check if the value has been set.
func (o *FormulaCaseQuery) GetQueryFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryFilter, true
}

// SetQueryFilter sets field value.
func (o *FormulaCaseQuery) SetQueryFilter(v string) {
	o.QueryFilter = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *FormulaCaseQuery) GetSort() QuerySortOrder {
	if o == nil || o.Sort == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaCaseQuery) GetSortOk() (*QuerySortOrder, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *FormulaCaseQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given QuerySortOrder and assigns it to the Sort field.
func (o *FormulaCaseQuery) SetSort(v QuerySortOrder) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaCaseQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["name"] = o.Name
	toSerialize["query_filter"] = o.QueryFilter
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaCaseQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Compute     *CaseCompute                `json:"compute"`
		DataSource  *FormulaCaseQueryDataSource `json:"data_source"`
		GroupBy     []CaseGroupBy               `json:"group_by,omitempty"`
		Limit       *float64                    `json:"limit,omitempty"`
		Name        *string                     `json:"name"`
		QueryFilter *string                     `json:"query_filter"`
		Sort        *QuerySortOrder             `json:"sort,omitempty"`
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
	if all.QueryFilter == nil {
		return fmt.Errorf("required field query_filter missing")
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
	o.Limit = all.Limit
	o.Name = *all.Name
	o.QueryFilter = *all.QueryFilter
	if all.Sort != nil && !all.Sort.IsValid() {
		hasInvalidField = true
	} else {
		o.Sort = all.Sort
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
