// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaAndFunctionProcessQueryDefinition Test-only generated schema for the dashboard versioning proof of concept.
type FormulaAndFunctionProcessQueryDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaAndFunctionProcessQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsNormalizedCpu *bool `json:"is_normalized_cpu,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metric string `json:"metric"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *QuerySortOrder `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TagFilters []string `json:"tag_filters,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TextFilter *string `json:"text_filter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaAndFunctionProcessQueryDefinition instantiates a new FormulaAndFunctionProcessQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaAndFunctionProcessQueryDefinition(dataSource FormulaAndFunctionProcessQueryDataSource, metric string, name string) *FormulaAndFunctionProcessQueryDefinition {
	this := FormulaAndFunctionProcessQueryDefinition{}
	this.DataSource = dataSource
	this.Metric = metric
	this.Name = name
	return &this
}

// NewFormulaAndFunctionProcessQueryDefinitionWithDefaults instantiates a new FormulaAndFunctionProcessQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaAndFunctionProcessQueryDefinitionWithDefaults() *FormulaAndFunctionProcessQueryDefinition {
	this := FormulaAndFunctionProcessQueryDefinition{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *FormulaAndFunctionProcessQueryDefinition) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *FormulaAndFunctionProcessQueryDefinition) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaAndFunctionProcessQueryDefinition) GetDataSource() FormulaAndFunctionProcessQueryDataSource {
	if o == nil {
		var ret FormulaAndFunctionProcessQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetDataSourceOk() (*FormulaAndFunctionProcessQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaAndFunctionProcessQueryDefinition) SetDataSource(v FormulaAndFunctionProcessQueryDataSource) {
	o.DataSource = v
}

// GetIsNormalizedCpu returns the IsNormalizedCpu field value if set, zero value otherwise.
func (o *FormulaAndFunctionProcessQueryDefinition) GetIsNormalizedCpu() bool {
	if o == nil || o.IsNormalizedCpu == nil {
		var ret bool
		return ret
	}
	return *o.IsNormalizedCpu
}

// GetIsNormalizedCpuOk returns a tuple with the IsNormalizedCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetIsNormalizedCpuOk() (*bool, bool) {
	if o == nil || o.IsNormalizedCpu == nil {
		return nil, false
	}
	return o.IsNormalizedCpu, true
}

// HasIsNormalizedCpu returns a boolean if a field has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) HasIsNormalizedCpu() bool {
	return o != nil && o.IsNormalizedCpu != nil
}

// SetIsNormalizedCpu gets a reference to the given bool and assigns it to the IsNormalizedCpu field.
func (o *FormulaAndFunctionProcessQueryDefinition) SetIsNormalizedCpu(v bool) {
	o.IsNormalizedCpu = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FormulaAndFunctionProcessQueryDefinition) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *FormulaAndFunctionProcessQueryDefinition) SetLimit(v float64) {
	o.Limit = &v
}

// GetMetric returns the Metric field value.
func (o *FormulaAndFunctionProcessQueryDefinition) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *FormulaAndFunctionProcessQueryDefinition) SetMetric(v string) {
	o.Metric = v
}

// GetName returns the Name field value.
func (o *FormulaAndFunctionProcessQueryDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaAndFunctionProcessQueryDefinition) SetName(v string) {
	o.Name = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *FormulaAndFunctionProcessQueryDefinition) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *FormulaAndFunctionProcessQueryDefinition) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *FormulaAndFunctionProcessQueryDefinition) GetSort() QuerySortOrder {
	if o == nil || o.Sort == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetSortOk() (*QuerySortOrder, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given QuerySortOrder and assigns it to the Sort field.
func (o *FormulaAndFunctionProcessQueryDefinition) SetSort(v QuerySortOrder) {
	o.Sort = &v
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *FormulaAndFunctionProcessQueryDefinition) GetTagFilters() []string {
	if o == nil || o.TagFilters == nil {
		var ret []string
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetTagFiltersOk() (*[]string, bool) {
	if o == nil || o.TagFilters == nil {
		return nil, false
	}
	return &o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) HasTagFilters() bool {
	return o != nil && o.TagFilters != nil
}

// SetTagFilters gets a reference to the given []string and assigns it to the TagFilters field.
func (o *FormulaAndFunctionProcessQueryDefinition) SetTagFilters(v []string) {
	o.TagFilters = v
}

// GetTextFilter returns the TextFilter field value if set, zero value otherwise.
func (o *FormulaAndFunctionProcessQueryDefinition) GetTextFilter() string {
	if o == nil || o.TextFilter == nil {
		var ret string
		return ret
	}
	return *o.TextFilter
}

// GetTextFilterOk returns a tuple with the TextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) GetTextFilterOk() (*string, bool) {
	if o == nil || o.TextFilter == nil {
		return nil, false
	}
	return o.TextFilter, true
}

// HasTextFilter returns a boolean if a field has been set.
func (o *FormulaAndFunctionProcessQueryDefinition) HasTextFilter() bool {
	return o != nil && o.TextFilter != nil
}

// SetTextFilter gets a reference to the given string and assigns it to the TextFilter field.
func (o *FormulaAndFunctionProcessQueryDefinition) SetTextFilter(v string) {
	o.TextFilter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaAndFunctionProcessQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	toSerialize["data_source"] = o.DataSource
	if o.IsNormalizedCpu != nil {
		toSerialize["is_normalized_cpu"] = o.IsNormalizedCpu
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["metric"] = o.Metric
	toSerialize["name"] = o.Name
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.TagFilters != nil {
		toSerialize["tag_filters"] = o.TagFilters
	}
	if o.TextFilter != nil {
		toSerialize["text_filter"] = o.TextFilter
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaAndFunctionProcessQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Aggregator      *ScalarAggregator                         `json:"aggregator,omitempty"`
		DataSource      *FormulaAndFunctionProcessQueryDataSource `json:"data_source"`
		IsNormalizedCpu *bool                                     `json:"is_normalized_cpu,omitempty"`
		Limit           *float64                                  `json:"limit,omitempty"`
		Metric          *string                                   `json:"metric"`
		Name            *string                                   `json:"name"`
		QueryFilter     *string                                   `json:"query_filter,omitempty"`
		Sort            *QuerySortOrder                           `json:"sort,omitempty"`
		TagFilters      []string                                  `json:"tag_filters,omitempty"`
		TextFilter      *string                                   `json:"text_filter,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.IsNormalizedCpu = all.IsNormalizedCpu
	o.Limit = all.Limit
	o.Metric = *all.Metric
	o.Name = *all.Name
	o.QueryFilter = all.QueryFilter
	if all.Sort != nil && !all.Sort.IsValid() {
		hasInvalidField = true
	} else {
		o.Sort = all.Sort
	}
	o.TagFilters = all.TagFilters
	o.TextFilter = all.TextFilter

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
