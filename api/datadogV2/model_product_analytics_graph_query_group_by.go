// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsGraphQueryGroupBy Segments journey results by the values of a facet.
type ProductAnalyticsGraphQueryGroupBy struct {
	// Attribute path to group by.
	Facet string `json:"facet"`
	// Maximum number of groups to return. Omit it to let the service choose.
	Limit *int64 `json:"limit,omitempty"`
	// Whether to exclude entities that have no value for this facet.
	ShouldExcludeMissing *bool `json:"should_exclude_missing,omitempty"`
	// Sort configuration for group-by results.
	Sort *ProductAnalyticsGroupBySort `json:"sort,omitempty"`
	// Audience dimension to group by, instead of an event facet.
	Source *ProductAnalyticsGraphQueryGroupBySource `json:"source,omitempty"`
	// A reference to a step, or a range of steps, in the journey.
	// Use a `node` target to name a single step, or a `path` target to name the range
	// between two steps.
	Target *ProductAnalyticsJourneyTarget `json:"target,omitempty"`
	// Restricts the results to these facet values.
	ValueFilters []string `json:"value_filters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsGraphQueryGroupBy instantiates a new ProductAnalyticsGraphQueryGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsGraphQueryGroupBy(facet string) *ProductAnalyticsGraphQueryGroupBy {
	this := ProductAnalyticsGraphQueryGroupBy{}
	this.Facet = facet
	var shouldExcludeMissing bool = false
	this.ShouldExcludeMissing = &shouldExcludeMissing
	return &this
}

// NewProductAnalyticsGraphQueryGroupByWithDefaults instantiates a new ProductAnalyticsGraphQueryGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsGraphQueryGroupByWithDefaults() *ProductAnalyticsGraphQueryGroupBy {
	this := ProductAnalyticsGraphQueryGroupBy{}
	var shouldExcludeMissing bool = false
	this.ShouldExcludeMissing = &shouldExcludeMissing
	return &this
}

// GetFacet returns the Facet field value.
func (o *ProductAnalyticsGraphQueryGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *ProductAnalyticsGraphQueryGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsGraphQueryGroupBy) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProductAnalyticsGraphQueryGroupBy) SetLimit(v int64) {
	o.Limit = &v
}

// GetShouldExcludeMissing returns the ShouldExcludeMissing field value if set, zero value otherwise.
func (o *ProductAnalyticsGraphQueryGroupBy) GetShouldExcludeMissing() bool {
	if o == nil || o.ShouldExcludeMissing == nil {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeMissing
}

// GetShouldExcludeMissingOk returns a tuple with the ShouldExcludeMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) GetShouldExcludeMissingOk() (*bool, bool) {
	if o == nil || o.ShouldExcludeMissing == nil {
		return nil, false
	}
	return o.ShouldExcludeMissing, true
}

// HasShouldExcludeMissing returns a boolean if a field has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) HasShouldExcludeMissing() bool {
	return o != nil && o.ShouldExcludeMissing != nil
}

// SetShouldExcludeMissing gets a reference to the given bool and assigns it to the ShouldExcludeMissing field.
func (o *ProductAnalyticsGraphQueryGroupBy) SetShouldExcludeMissing(v bool) {
	o.ShouldExcludeMissing = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProductAnalyticsGraphQueryGroupBy) GetSort() ProductAnalyticsGroupBySort {
	if o == nil || o.Sort == nil {
		var ret ProductAnalyticsGroupBySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) GetSortOk() (*ProductAnalyticsGroupBySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ProductAnalyticsGroupBySort and assigns it to the Sort field.
func (o *ProductAnalyticsGraphQueryGroupBy) SetSort(v ProductAnalyticsGroupBySort) {
	o.Sort = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProductAnalyticsGraphQueryGroupBy) GetSource() ProductAnalyticsGraphQueryGroupBySource {
	if o == nil || o.Source == nil {
		var ret ProductAnalyticsGraphQueryGroupBySource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) GetSourceOk() (*ProductAnalyticsGraphQueryGroupBySource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given ProductAnalyticsGraphQueryGroupBySource and assigns it to the Source field.
func (o *ProductAnalyticsGraphQueryGroupBy) SetSource(v ProductAnalyticsGraphQueryGroupBySource) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProductAnalyticsGraphQueryGroupBy) GetTarget() ProductAnalyticsJourneyTarget {
	if o == nil || o.Target == nil {
		var ret ProductAnalyticsJourneyTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) GetTargetOk() (*ProductAnalyticsJourneyTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given ProductAnalyticsJourneyTarget and assigns it to the Target field.
func (o *ProductAnalyticsGraphQueryGroupBy) SetTarget(v ProductAnalyticsJourneyTarget) {
	o.Target = &v
}

// GetValueFilters returns the ValueFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsGraphQueryGroupBy) GetValueFilters() []string {
	if o == nil || o.ValueFilters == nil {
		var ret []string
		return ret
	}
	return o.ValueFilters
}

// GetValueFiltersOk returns a tuple with the ValueFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) GetValueFiltersOk() (*[]string, bool) {
	if o == nil || o.ValueFilters == nil {
		return nil, false
	}
	return &o.ValueFilters, true
}

// HasValueFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsGraphQueryGroupBy) HasValueFilters() bool {
	return o != nil && o.ValueFilters != nil
}

// SetValueFilters gets a reference to the given []string and assigns it to the ValueFilters field.
func (o *ProductAnalyticsGraphQueryGroupBy) SetValueFilters(v []string) {
	o.ValueFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsGraphQueryGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.ShouldExcludeMissing != nil {
		toSerialize["should_exclude_missing"] = o.ShouldExcludeMissing
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.ValueFilters != nil {
		toSerialize["value_filters"] = o.ValueFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsGraphQueryGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet                *string                                  `json:"facet"`
		Limit                *int64                                   `json:"limit,omitempty"`
		ShouldExcludeMissing *bool                                    `json:"should_exclude_missing,omitempty"`
		Sort                 *ProductAnalyticsGroupBySort             `json:"sort,omitempty"`
		Source               *ProductAnalyticsGraphQueryGroupBySource `json:"source,omitempty"`
		Target               *ProductAnalyticsJourneyTarget           `json:"target,omitempty"`
		ValueFilters         []string                                 `json:"value_filters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "limit", "should_exclude_missing", "sort", "source", "target", "value_filters"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Facet = *all.Facet
	o.Limit = all.Limit
	o.ShouldExcludeMissing = all.ShouldExcludeMissing
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}
	o.Target = all.Target
	o.ValueFilters = all.ValueFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
