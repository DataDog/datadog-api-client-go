// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyListQuery Query definition for a journey list request.
type ProductAnalyticsJourneyListQuery struct {
	// Computed columns to add to each row.
	ComputedColumns []ProductAnalyticsJourneyComputedColumn `json:"computed_columns,omitempty"`
	// Whether to return the entities that converted at the target step, or those that dropped off.
	ConversionType *ProductAnalyticsJourneyConversionType `json:"conversion_type,omitempty"`
	// Attribute columns to return for each row, in addition to the identity join key and `timestamp`.
	EntityColumns []string `json:"entity_columns,omitempty"`
	// Additional search query applied to the returned rows.
	EntityFilters *string `json:"entity_filters,omitempty"`
	// Segments the results by the values of one or more facets.
	GroupBy []ProductAnalyticsGraphQueryGroupBy `json:"group_by,omitempty"`
	// Maximum number of rows to return. Omit it to let the service choose.
	Limit *int64 `json:"limit,omitempty"`
	// Defines the steps of the journey and the filters applied to it.
	Search ProductAnalyticsJourneySearch `json:"search"`
	// Sort configuration for the returned rows. The sort is applied only when `facet`
	// is one of the returned columns; otherwise it is ignored.
	Sort *ProductAnalyticsJourneyListSort `json:"sort,omitempty"`
	// A reference to a step, or a range of steps, in the journey.
	// Use a `node` target to name a single step, or a `path` target to name the range
	// between two steps.
	Target *ProductAnalyticsJourneyTarget `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyListQuery instantiates a new ProductAnalyticsJourneyListQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyListQuery(search ProductAnalyticsJourneySearch) *ProductAnalyticsJourneyListQuery {
	this := ProductAnalyticsJourneyListQuery{}
	this.Search = search
	return &this
}

// NewProductAnalyticsJourneyListQueryWithDefaults instantiates a new ProductAnalyticsJourneyListQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyListQueryWithDefaults() *ProductAnalyticsJourneyListQuery {
	this := ProductAnalyticsJourneyListQuery{}
	return &this
}

// GetComputedColumns returns the ComputedColumns field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetComputedColumns() []ProductAnalyticsJourneyComputedColumn {
	if o == nil || o.ComputedColumns == nil {
		var ret []ProductAnalyticsJourneyComputedColumn
		return ret
	}
	return o.ComputedColumns
}

// GetComputedColumnsOk returns a tuple with the ComputedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetComputedColumnsOk() (*[]ProductAnalyticsJourneyComputedColumn, bool) {
	if o == nil || o.ComputedColumns == nil {
		return nil, false
	}
	return &o.ComputedColumns, true
}

// HasComputedColumns returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasComputedColumns() bool {
	return o != nil && o.ComputedColumns != nil
}

// SetComputedColumns gets a reference to the given []ProductAnalyticsJourneyComputedColumn and assigns it to the ComputedColumns field.
func (o *ProductAnalyticsJourneyListQuery) SetComputedColumns(v []ProductAnalyticsJourneyComputedColumn) {
	o.ComputedColumns = v
}

// GetConversionType returns the ConversionType field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetConversionType() ProductAnalyticsJourneyConversionType {
	if o == nil || o.ConversionType == nil {
		var ret ProductAnalyticsJourneyConversionType
		return ret
	}
	return *o.ConversionType
}

// GetConversionTypeOk returns a tuple with the ConversionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetConversionTypeOk() (*ProductAnalyticsJourneyConversionType, bool) {
	if o == nil || o.ConversionType == nil {
		return nil, false
	}
	return o.ConversionType, true
}

// HasConversionType returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasConversionType() bool {
	return o != nil && o.ConversionType != nil
}

// SetConversionType gets a reference to the given ProductAnalyticsJourneyConversionType and assigns it to the ConversionType field.
func (o *ProductAnalyticsJourneyListQuery) SetConversionType(v ProductAnalyticsJourneyConversionType) {
	o.ConversionType = &v
}

// GetEntityColumns returns the EntityColumns field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetEntityColumns() []string {
	if o == nil || o.EntityColumns == nil {
		var ret []string
		return ret
	}
	return o.EntityColumns
}

// GetEntityColumnsOk returns a tuple with the EntityColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetEntityColumnsOk() (*[]string, bool) {
	if o == nil || o.EntityColumns == nil {
		return nil, false
	}
	return &o.EntityColumns, true
}

// HasEntityColumns returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasEntityColumns() bool {
	return o != nil && o.EntityColumns != nil
}

// SetEntityColumns gets a reference to the given []string and assigns it to the EntityColumns field.
func (o *ProductAnalyticsJourneyListQuery) SetEntityColumns(v []string) {
	o.EntityColumns = v
}

// GetEntityFilters returns the EntityFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetEntityFilters() string {
	if o == nil || o.EntityFilters == nil {
		var ret string
		return ret
	}
	return *o.EntityFilters
}

// GetEntityFiltersOk returns a tuple with the EntityFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetEntityFiltersOk() (*string, bool) {
	if o == nil || o.EntityFilters == nil {
		return nil, false
	}
	return o.EntityFilters, true
}

// HasEntityFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasEntityFilters() bool {
	return o != nil && o.EntityFilters != nil
}

// SetEntityFilters gets a reference to the given string and assigns it to the EntityFilters field.
func (o *ProductAnalyticsJourneyListQuery) SetEntityFilters(v string) {
	o.EntityFilters = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetGroupBy() []ProductAnalyticsGraphQueryGroupBy {
	if o == nil || o.GroupBy == nil {
		var ret []ProductAnalyticsGraphQueryGroupBy
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetGroupByOk() (*[]ProductAnalyticsGraphQueryGroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ProductAnalyticsGraphQueryGroupBy and assigns it to the GroupBy field.
func (o *ProductAnalyticsJourneyListQuery) SetGroupBy(v []ProductAnalyticsGraphQueryGroupBy) {
	o.GroupBy = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProductAnalyticsJourneyListQuery) SetLimit(v int64) {
	o.Limit = &v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsJourneyListQuery) GetSearch() ProductAnalyticsJourneySearch {
	if o == nil {
		var ret ProductAnalyticsJourneySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetSearchOk() (*ProductAnalyticsJourneySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsJourneyListQuery) SetSearch(v ProductAnalyticsJourneySearch) {
	o.Search = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetSort() ProductAnalyticsJourneyListSort {
	if o == nil || o.Sort == nil {
		var ret ProductAnalyticsJourneyListSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetSortOk() (*ProductAnalyticsJourneyListSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ProductAnalyticsJourneyListSort and assigns it to the Sort field.
func (o *ProductAnalyticsJourneyListQuery) SetSort(v ProductAnalyticsJourneyListSort) {
	o.Sort = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListQuery) GetTarget() ProductAnalyticsJourneyTarget {
	if o == nil || o.Target == nil {
		var ret ProductAnalyticsJourneyTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListQuery) GetTargetOk() (*ProductAnalyticsJourneyTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListQuery) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given ProductAnalyticsJourneyTarget and assigns it to the Target field.
func (o *ProductAnalyticsJourneyListQuery) SetTarget(v ProductAnalyticsJourneyTarget) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyListQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComputedColumns != nil {
		toSerialize["computed_columns"] = o.ComputedColumns
	}
	if o.ConversionType != nil {
		toSerialize["conversion_type"] = o.ConversionType
	}
	if o.EntityColumns != nil {
		toSerialize["entity_columns"] = o.EntityColumns
	}
	if o.EntityFilters != nil {
		toSerialize["entity_filters"] = o.EntityFilters
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	toSerialize["search"] = o.Search
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyListQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComputedColumns []ProductAnalyticsJourneyComputedColumn `json:"computed_columns,omitempty"`
		ConversionType  *ProductAnalyticsJourneyConversionType  `json:"conversion_type,omitempty"`
		EntityColumns   []string                                `json:"entity_columns,omitempty"`
		EntityFilters   *string                                 `json:"entity_filters,omitempty"`
		GroupBy         []ProductAnalyticsGraphQueryGroupBy     `json:"group_by,omitempty"`
		Limit           *int64                                  `json:"limit,omitempty"`
		Search          *ProductAnalyticsJourneySearch          `json:"search"`
		Sort            *ProductAnalyticsJourneyListSort        `json:"sort,omitempty"`
		Target          *ProductAnalyticsJourneyTarget          `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"computed_columns", "conversion_type", "entity_columns", "entity_filters", "group_by", "limit", "search", "sort", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComputedColumns = all.ComputedColumns
	if all.ConversionType != nil && !all.ConversionType.IsValid() {
		hasInvalidField = true
	} else {
		o.ConversionType = all.ConversionType
	}
	o.EntityColumns = all.EntityColumns
	o.EntityFilters = all.EntityFilters
	o.GroupBy = all.GroupBy
	o.Limit = all.Limit
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
