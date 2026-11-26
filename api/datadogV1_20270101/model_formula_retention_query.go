// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormulaRetentionQuery Test-only generated schema for the dashboard versioning proof of concept.
type FormulaRetentionQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ComputationScope *ComputationScope `json:"computation_scope,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Compute RetentionCompute `json:"compute"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource FormulaRetentionQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ExcludeAnonymousTraffic *bool `json:"exclude_anonymous_traffic,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy []RetentionGroupByItems `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sampling *FormulaRetentionQuerySampling `json:"sampling,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Search RetentionSearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFormulaRetentionQuery instantiates a new FormulaRetentionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormulaRetentionQuery(compute RetentionCompute, dataSource FormulaRetentionQueryDataSource, name string, search RetentionSearch) *FormulaRetentionQuery {
	this := FormulaRetentionQuery{}
	this.Compute = compute
	this.DataSource = dataSource
	this.Name = name
	this.Search = search
	return &this
}

// NewFormulaRetentionQueryWithDefaults instantiates a new FormulaRetentionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormulaRetentionQueryWithDefaults() *FormulaRetentionQuery {
	this := FormulaRetentionQuery{}
	return &this
}

// GetComputationScope returns the ComputationScope field value if set, zero value otherwise.
func (o *FormulaRetentionQuery) GetComputationScope() ComputationScope {
	if o == nil || o.ComputationScope == nil {
		var ret ComputationScope
		return ret
	}
	return *o.ComputationScope
}

// GetComputationScopeOk returns a tuple with the ComputationScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetComputationScopeOk() (*ComputationScope, bool) {
	if o == nil || o.ComputationScope == nil {
		return nil, false
	}
	return o.ComputationScope, true
}

// HasComputationScope returns a boolean if a field has been set.
func (o *FormulaRetentionQuery) HasComputationScope() bool {
	return o != nil && o.ComputationScope != nil
}

// SetComputationScope gets a reference to the given ComputationScope and assigns it to the ComputationScope field.
func (o *FormulaRetentionQuery) SetComputationScope(v ComputationScope) {
	o.ComputationScope = &v
}

// GetCompute returns the Compute field value.
func (o *FormulaRetentionQuery) GetCompute() RetentionCompute {
	if o == nil {
		var ret RetentionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetComputeOk() (*RetentionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *FormulaRetentionQuery) SetCompute(v RetentionCompute) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *FormulaRetentionQuery) GetDataSource() FormulaRetentionQueryDataSource {
	if o == nil {
		var ret FormulaRetentionQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetDataSourceOk() (*FormulaRetentionQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *FormulaRetentionQuery) SetDataSource(v FormulaRetentionQueryDataSource) {
	o.DataSource = v
}

// GetExcludeAnonymousTraffic returns the ExcludeAnonymousTraffic field value if set, zero value otherwise.
func (o *FormulaRetentionQuery) GetExcludeAnonymousTraffic() bool {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeAnonymousTraffic
}

// GetExcludeAnonymousTrafficOk returns a tuple with the ExcludeAnonymousTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetExcludeAnonymousTrafficOk() (*bool, bool) {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		return nil, false
	}
	return o.ExcludeAnonymousTraffic, true
}

// HasExcludeAnonymousTraffic returns a boolean if a field has been set.
func (o *FormulaRetentionQuery) HasExcludeAnonymousTraffic() bool {
	return o != nil && o.ExcludeAnonymousTraffic != nil
}

// SetExcludeAnonymousTraffic gets a reference to the given bool and assigns it to the ExcludeAnonymousTraffic field.
func (o *FormulaRetentionQuery) SetExcludeAnonymousTraffic(v bool) {
	o.ExcludeAnonymousTraffic = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *FormulaRetentionQuery) GetGroupBy() []RetentionGroupByItems {
	if o == nil || o.GroupBy == nil {
		var ret []RetentionGroupByItems
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetGroupByOk() (*[]RetentionGroupByItems, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *FormulaRetentionQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []RetentionGroupByItems and assigns it to the GroupBy field.
func (o *FormulaRetentionQuery) SetGroupBy(v []RetentionGroupByItems) {
	o.GroupBy = v
}

// GetName returns the Name field value.
func (o *FormulaRetentionQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormulaRetentionQuery) SetName(v string) {
	o.Name = v
}

// GetSampling returns the Sampling field value if set, zero value otherwise.
func (o *FormulaRetentionQuery) GetSampling() FormulaRetentionQuerySampling {
	if o == nil || o.Sampling == nil {
		var ret FormulaRetentionQuerySampling
		return ret
	}
	return *o.Sampling
}

// GetSamplingOk returns a tuple with the Sampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetSamplingOk() (*FormulaRetentionQuerySampling, bool) {
	if o == nil || o.Sampling == nil {
		return nil, false
	}
	return o.Sampling, true
}

// HasSampling returns a boolean if a field has been set.
func (o *FormulaRetentionQuery) HasSampling() bool {
	return o != nil && o.Sampling != nil
}

// SetSampling gets a reference to the given FormulaRetentionQuerySampling and assigns it to the Sampling field.
func (o *FormulaRetentionQuery) SetSampling(v FormulaRetentionQuerySampling) {
	o.Sampling = &v
}

// GetSearch returns the Search field value.
func (o *FormulaRetentionQuery) GetSearch() RetentionSearch {
	if o == nil {
		var ret RetentionSearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *FormulaRetentionQuery) GetSearchOk() (*RetentionSearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *FormulaRetentionQuery) SetSearch(v RetentionSearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormulaRetentionQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComputationScope != nil {
		toSerialize["computation_scope"] = o.ComputationScope
	}
	toSerialize["compute"] = o.Compute
	toSerialize["data_source"] = o.DataSource
	if o.ExcludeAnonymousTraffic != nil {
		toSerialize["exclude_anonymous_traffic"] = o.ExcludeAnonymousTraffic
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["name"] = o.Name
	if o.Sampling != nil {
		toSerialize["sampling"] = o.Sampling
	}
	toSerialize["search"] = o.Search
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormulaRetentionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComputationScope        *ComputationScope                `json:"computation_scope,omitempty"`
		Compute                 *RetentionCompute                `json:"compute"`
		DataSource              *FormulaRetentionQueryDataSource `json:"data_source"`
		ExcludeAnonymousTraffic *bool                            `json:"exclude_anonymous_traffic,omitempty"`
		GroupBy                 []RetentionGroupByItems          `json:"group_by,omitempty"`
		Name                    *string                          `json:"name"`
		Sampling                *FormulaRetentionQuerySampling   `json:"sampling,omitempty"`
		Search                  *RetentionSearch                 `json:"search"`
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
	o.ComputationScope = all.ComputationScope
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.ExcludeAnonymousTraffic = all.ExcludeAnonymousTraffic
	o.GroupBy = all.GroupBy
	o.Name = *all.Name
	if all.Sampling != nil && all.Sampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sampling = all.Sampling
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
