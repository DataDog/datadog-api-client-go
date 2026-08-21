// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAnalyticsListQuery The analytics list query definition. It selects the events to return with `query`, then
// chooses the columns on each event row, the sort applied to those rows, and a row limit.
// Unlike the scalar and timeseries queries, a list query returns raw event rows rather than
// aggregates, so it takes no compute or group-by rule.
type ProductAnalyticsAnalyticsListQuery struct {
	// Audience filter definitions for targeting specific user segments.
	AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
	// Attribute columns to include in each event row.
	Columns []string `json:"columns,omitempty"`
	// Maximum number of event rows to return.
	Limit *int64 `json:"limit,omitempty"`
	// A query definition discriminated by the `data_source` field.
	// Use `product_analytics` for standard event queries, or
	// `product_analytics_occurrence` for occurrence-filtered queries.
	Query ProductAnalyticsBaseQuery `json:"query"`
	// The sort applied to the returned event rows.
	Sort *ProductAnalyticsAnalyticsListSort `json:"sort,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsAnalyticsListQuery instantiates a new ProductAnalyticsAnalyticsListQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsAnalyticsListQuery(query ProductAnalyticsBaseQuery) *ProductAnalyticsAnalyticsListQuery {
	this := ProductAnalyticsAnalyticsListQuery{}
	this.Query = query
	return &this
}

// NewProductAnalyticsAnalyticsListQueryWithDefaults instantiates a new ProductAnalyticsAnalyticsListQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsAnalyticsListQueryWithDefaults() *ProductAnalyticsAnalyticsListQuery {
	this := ProductAnalyticsAnalyticsListQuery{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsListQuery) GetAudienceFilters() ProductAnalyticsAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret ProductAnalyticsAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListQuery) GetAudienceFiltersOk() (*ProductAnalyticsAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsListQuery) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given ProductAnalyticsAudienceFilters and assigns it to the AudienceFilters field.
func (o *ProductAnalyticsAnalyticsListQuery) SetAudienceFilters(v ProductAnalyticsAudienceFilters) {
	o.AudienceFilters = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsListQuery) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListQuery) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsListQuery) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *ProductAnalyticsAnalyticsListQuery) SetColumns(v []string) {
	o.Columns = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsListQuery) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListQuery) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsListQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProductAnalyticsAnalyticsListQuery) SetLimit(v int64) {
	o.Limit = &v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsAnalyticsListQuery) GetQuery() ProductAnalyticsBaseQuery {
	if o == nil {
		var ret ProductAnalyticsBaseQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListQuery) GetQueryOk() (*ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsAnalyticsListQuery) SetQuery(v ProductAnalyticsBaseQuery) {
	o.Query = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProductAnalyticsAnalyticsListQuery) GetSort() ProductAnalyticsAnalyticsListSort {
	if o == nil || o.Sort == nil {
		var ret ProductAnalyticsAnalyticsListSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListQuery) GetSortOk() (*ProductAnalyticsAnalyticsListSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProductAnalyticsAnalyticsListQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ProductAnalyticsAnalyticsListSort and assigns it to the Sort field.
func (o *ProductAnalyticsAnalyticsListQuery) SetSort(v ProductAnalyticsAnalyticsListSort) {
	o.Sort = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsAnalyticsListQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["query"] = o.Query
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsAnalyticsListQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *ProductAnalyticsAudienceFilters   `json:"audience_filters,omitempty"`
		Columns         []string                           `json:"columns,omitempty"`
		Limit           *int64                             `json:"limit,omitempty"`
		Query           *ProductAnalyticsBaseQuery         `json:"query"`
		Sort            *ProductAnalyticsAnalyticsListSort `json:"sort,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "columns", "limit", "query", "sort"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	o.Columns = all.Columns
	o.Limit = all.Limit
	o.Query = *all.Query
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
