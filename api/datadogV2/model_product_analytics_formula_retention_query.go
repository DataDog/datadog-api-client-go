// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFormulaRetentionQuery Query definition for a retention scalar or retention timeseries request.
type ProductAnalyticsFormulaRetentionQuery struct {
	// Restricts a retention query to part of the grid, so that results can be examined in detail.
	// Omit it to compute the whole grid.
	ComputationScope *ProductAnalyticsRetentionScope `json:"computation_scope,omitempty"`
	// The metric and aggregation applied to a retention query.
	Compute ProductAnalyticsRetentionCompute `json:"compute"`
	// Splits the results by the values of one or more facets.
	GroupBy []ProductAnalyticsRetentionGroupBy `json:"group_by,omitempty"`
	// Defines the cohort and return criteria that make up a retention query.
	Search ProductAnalyticsRetentionSearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFormulaRetentionQuery instantiates a new ProductAnalyticsFormulaRetentionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFormulaRetentionQuery(compute ProductAnalyticsRetentionCompute, search ProductAnalyticsRetentionSearch) *ProductAnalyticsFormulaRetentionQuery {
	this := ProductAnalyticsFormulaRetentionQuery{}
	this.Compute = compute
	this.Search = search
	return &this
}

// NewProductAnalyticsFormulaRetentionQueryWithDefaults instantiates a new ProductAnalyticsFormulaRetentionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFormulaRetentionQueryWithDefaults() *ProductAnalyticsFormulaRetentionQuery {
	this := ProductAnalyticsFormulaRetentionQuery{}
	return &this
}

// GetComputationScope returns the ComputationScope field value if set, zero value otherwise.
func (o *ProductAnalyticsFormulaRetentionQuery) GetComputationScope() ProductAnalyticsRetentionScope {
	if o == nil || o.ComputationScope == nil {
		var ret ProductAnalyticsRetentionScope
		return ret
	}
	return *o.ComputationScope
}

// GetComputationScopeOk returns a tuple with the ComputationScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaRetentionQuery) GetComputationScopeOk() (*ProductAnalyticsRetentionScope, bool) {
	if o == nil || o.ComputationScope == nil {
		return nil, false
	}
	return o.ComputationScope, true
}

// HasComputationScope returns a boolean if a field has been set.
func (o *ProductAnalyticsFormulaRetentionQuery) HasComputationScope() bool {
	return o != nil && o.ComputationScope != nil
}

// SetComputationScope gets a reference to the given ProductAnalyticsRetentionScope and assigns it to the ComputationScope field.
func (o *ProductAnalyticsFormulaRetentionQuery) SetComputationScope(v ProductAnalyticsRetentionScope) {
	o.ComputationScope = &v
}

// GetCompute returns the Compute field value.
func (o *ProductAnalyticsFormulaRetentionQuery) GetCompute() ProductAnalyticsRetentionCompute {
	if o == nil {
		var ret ProductAnalyticsRetentionCompute
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaRetentionQuery) GetComputeOk() (*ProductAnalyticsRetentionCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value.
func (o *ProductAnalyticsFormulaRetentionQuery) SetCompute(v ProductAnalyticsRetentionCompute) {
	o.Compute = v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ProductAnalyticsFormulaRetentionQuery) GetGroupBy() []ProductAnalyticsRetentionGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsRetentionGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaRetentionQuery) GetGroupByOk() (*[]ProductAnalyticsRetentionGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ProductAnalyticsFormulaRetentionQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsRetentionGroupBy and assigns it to the GroupBy field.
func (o *ProductAnalyticsFormulaRetentionQuery) SetGroupBy(v []ProductAnalyticsRetentionGroupBy) {
	o.GroupBy = v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsFormulaRetentionQuery) GetSearch() ProductAnalyticsRetentionSearch {
	if o == nil {
		var ret ProductAnalyticsRetentionSearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFormulaRetentionQuery) GetSearchOk() (*ProductAnalyticsRetentionSearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsFormulaRetentionQuery) SetSearch(v ProductAnalyticsRetentionSearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFormulaRetentionQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComputationScope != nil {
		toSerialize["computation_scope"] = o.ComputationScope
	}
	toSerialize["compute"] = o.Compute
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["search"] = o.Search

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFormulaRetentionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComputationScope *ProductAnalyticsRetentionScope    `json:"computation_scope,omitempty"`
		Compute          *ProductAnalyticsRetentionCompute  `json:"compute"`
		GroupBy          []ProductAnalyticsRetentionGroupBy `json:"group_by,omitempty"`
		Search           *ProductAnalyticsRetentionSearch   `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Compute == nil {
		return fmt.Errorf("required field compute missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"computation_scope", "compute", "group_by", "search"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComputationScope = all.ComputationScope
	if all.Compute.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compute = *all.Compute
	o.GroupBy = all.GroupBy
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
