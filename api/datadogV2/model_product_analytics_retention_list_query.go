// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionListQuery Query definition for a retention list request.
type ProductAnalyticsRetentionListQuery struct {
	// The attribute columns to include in each returned row.
	Columns []ProductAnalyticsRetentionListColumn `json:"columns,omitempty"`
	// Narrows a retention query to a single cell, at the intersection of one cohort and one return period.
	ComputationScope ProductAnalyticsRetentionCellScope `json:"computation_scope"`
	// Maximum number of rows to return. Use `0` for no limit.
	Limit *int64 `json:"limit,omitempty"`
	// Defines the cohort and return criteria that make up a retention query.
	Search ProductAnalyticsRetentionSearch `json:"search"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionListQuery instantiates a new ProductAnalyticsRetentionListQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionListQuery(computationScope ProductAnalyticsRetentionCellScope, search ProductAnalyticsRetentionSearch) *ProductAnalyticsRetentionListQuery {
	this := ProductAnalyticsRetentionListQuery{}
	this.ComputationScope = computationScope
	this.Search = search
	return &this
}

// NewProductAnalyticsRetentionListQueryWithDefaults instantiates a new ProductAnalyticsRetentionListQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionListQueryWithDefaults() *ProductAnalyticsRetentionListQuery {
	this := ProductAnalyticsRetentionListQuery{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionListQuery) GetColumns() []ProductAnalyticsRetentionListColumn {
	if o == nil || o.Columns == nil {
		var ret []ProductAnalyticsRetentionListColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionListQuery) GetColumnsOk() (*[]ProductAnalyticsRetentionListColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionListQuery) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []ProductAnalyticsRetentionListColumn and assigns it to the Columns field.
func (o *ProductAnalyticsRetentionListQuery) SetColumns(v []ProductAnalyticsRetentionListColumn) {
	o.Columns = v
}

// GetComputationScope returns the ComputationScope field value.
func (o *ProductAnalyticsRetentionListQuery) GetComputationScope() ProductAnalyticsRetentionCellScope {
	if o == nil {
		var ret ProductAnalyticsRetentionCellScope
		return ret
	}
	return o.ComputationScope
}

// GetComputationScopeOk returns a tuple with the ComputationScope field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionListQuery) GetComputationScopeOk() (*ProductAnalyticsRetentionCellScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComputationScope, true
}

// SetComputationScope sets field value.
func (o *ProductAnalyticsRetentionListQuery) SetComputationScope(v ProductAnalyticsRetentionCellScope) {
	o.ComputationScope = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionListQuery) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionListQuery) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionListQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProductAnalyticsRetentionListQuery) SetLimit(v int64) {
	o.Limit = &v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsRetentionListQuery) GetSearch() ProductAnalyticsRetentionSearch {
	if o == nil {
		var ret ProductAnalyticsRetentionSearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionListQuery) GetSearchOk() (*ProductAnalyticsRetentionSearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsRetentionListQuery) SetSearch(v ProductAnalyticsRetentionSearch) {
	o.Search = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionListQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["computation_scope"] = o.ComputationScope
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["search"] = o.Search

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionListQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns          []ProductAnalyticsRetentionListColumn `json:"columns,omitempty"`
		ComputationScope *ProductAnalyticsRetentionCellScope   `json:"computation_scope"`
		Limit            *int64                                `json:"limit,omitempty"`
		Search           *ProductAnalyticsRetentionSearch      `json:"search"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ComputationScope == nil {
		return fmt.Errorf("required field computation_scope missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "computation_scope", "limit", "search"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = all.Columns
	if all.ComputationScope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ComputationScope = *all.ComputationScope
	o.Limit = all.Limit
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
