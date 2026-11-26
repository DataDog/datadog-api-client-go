// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProcessQueryDefinition Test-only generated schema for the dashboard versioning proof of concept.
type ProcessQueryDefinition struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	FilterBy []string `json:"filter_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IsNormalizedCpu *bool `json:"is_normalized_cpu,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit float64 `json:"limit"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Metric string `json:"metric"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SearchBy *string `json:"search_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProcessQueryDefinition instantiates a new ProcessQueryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProcessQueryDefinition(limit float64, metric string) *ProcessQueryDefinition {
	this := ProcessQueryDefinition{}
	this.Limit = limit
	this.Metric = metric
	return &this
}

// NewProcessQueryDefinitionWithDefaults instantiates a new ProcessQueryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProcessQueryDefinitionWithDefaults() *ProcessQueryDefinition {
	this := ProcessQueryDefinition{}
	return &this
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise.
func (o *ProcessQueryDefinition) GetFilterBy() []string {
	if o == nil || o.FilterBy == nil {
		var ret []string
		return ret
	}
	return o.FilterBy
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetFilterByOk() (*[]string, bool) {
	if o == nil || o.FilterBy == nil {
		return nil, false
	}
	return &o.FilterBy, true
}

// HasFilterBy returns a boolean if a field has been set.
func (o *ProcessQueryDefinition) HasFilterBy() bool {
	return o != nil && o.FilterBy != nil
}

// SetFilterBy gets a reference to the given []string and assigns it to the FilterBy field.
func (o *ProcessQueryDefinition) SetFilterBy(v []string) {
	o.FilterBy = v
}

// GetIsNormalizedCpu returns the IsNormalizedCpu field value if set, zero value otherwise.
func (o *ProcessQueryDefinition) GetIsNormalizedCpu() bool {
	if o == nil || o.IsNormalizedCpu == nil {
		var ret bool
		return ret
	}
	return *o.IsNormalizedCpu
}

// GetIsNormalizedCpuOk returns a tuple with the IsNormalizedCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetIsNormalizedCpuOk() (*bool, bool) {
	if o == nil || o.IsNormalizedCpu == nil {
		return nil, false
	}
	return o.IsNormalizedCpu, true
}

// HasIsNormalizedCpu returns a boolean if a field has been set.
func (o *ProcessQueryDefinition) HasIsNormalizedCpu() bool {
	return o != nil && o.IsNormalizedCpu != nil
}

// SetIsNormalizedCpu gets a reference to the given bool and assigns it to the IsNormalizedCpu field.
func (o *ProcessQueryDefinition) SetIsNormalizedCpu(v bool) {
	o.IsNormalizedCpu = &v
}

// GetLimit returns the Limit field value.
func (o *ProcessQueryDefinition) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *ProcessQueryDefinition) SetLimit(v float64) {
	o.Limit = v
}

// GetMetric returns the Metric field value.
func (o *ProcessQueryDefinition) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *ProcessQueryDefinition) SetMetric(v string) {
	o.Metric = v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *ProcessQueryDefinition) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *ProcessQueryDefinition) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *ProcessQueryDefinition) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetSearchBy returns the SearchBy field value if set, zero value otherwise.
func (o *ProcessQueryDefinition) GetSearchBy() string {
	if o == nil || o.SearchBy == nil {
		var ret string
		return ret
	}
	return *o.SearchBy
}

// GetSearchByOk returns a tuple with the SearchBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetSearchByOk() (*string, bool) {
	if o == nil || o.SearchBy == nil {
		return nil, false
	}
	return o.SearchBy, true
}

// HasSearchBy returns a boolean if a field has been set.
func (o *ProcessQueryDefinition) HasSearchBy() bool {
	return o != nil && o.SearchBy != nil
}

// SetSearchBy gets a reference to the given string and assigns it to the SearchBy field.
func (o *ProcessQueryDefinition) SetSearchBy(v string) {
	o.SearchBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProcessQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FilterBy != nil {
		toSerialize["filter_by"] = o.FilterBy
	}
	if o.IsNormalizedCpu != nil {
		toSerialize["is_normalized_cpu"] = o.IsNormalizedCpu
	}
	toSerialize["limit"] = o.Limit
	toSerialize["metric"] = o.Metric
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.SearchBy != nil {
		toSerialize["search_by"] = o.SearchBy
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProcessQueryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FilterBy        []string `json:"filter_by,omitempty"`
		IsNormalizedCpu *bool    `json:"is_normalized_cpu,omitempty"`
		Limit           *float64 `json:"limit"`
		Metric          *string  `json:"metric"`
		QueryFilter     *string  `json:"query_filter,omitempty"`
		SearchBy        *string  `json:"search_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	o.FilterBy = all.FilterBy
	o.IsNormalizedCpu = all.IsNormalizedCpu
	o.Limit = *all.Limit
	o.Metric = *all.Metric
	o.QueryFilter = all.QueryFilter
	o.SearchBy = all.SearchBy

	return nil
}
