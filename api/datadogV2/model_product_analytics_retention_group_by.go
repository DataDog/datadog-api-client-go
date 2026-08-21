// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGroupBy Splits retention results by the values of a facet.
type ProductAnalyticsRetentionGroupBy struct {
	// The attribute path to group by.
	Facet string `json:"facet"`
	// Maximum number of groups to return. Omit it to let the service choose.
	Limit *int64 `json:"limit,omitempty"`
	// Whether to drop entities that have no value for the facet.
	ShouldExcludeMissing *bool `json:"should_exclude_missing,omitempty"`
	// Sort configuration for group-by results.
	Sort *ProductAnalyticsGroupBySort `json:"sort,omitempty"`
	// Audience source backing the group-by, when grouping by an audience rather than a facet.
	Source *string `json:"source,omitempty"`
	// Which axis of the retention grid a group-by applies to.
	Target ProductAnalyticsRetentionGroupByTarget `json:"target"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionGroupBy instantiates a new ProductAnalyticsRetentionGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionGroupBy(facet string, target ProductAnalyticsRetentionGroupByTarget) *ProductAnalyticsRetentionGroupBy {
	this := ProductAnalyticsRetentionGroupBy{}
	this.Facet = facet
	var shouldExcludeMissing bool = false
	this.ShouldExcludeMissing = &shouldExcludeMissing
	this.Target = target
	return &this
}

// NewProductAnalyticsRetentionGroupByWithDefaults instantiates a new ProductAnalyticsRetentionGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionGroupByWithDefaults() *ProductAnalyticsRetentionGroupBy {
	this := ProductAnalyticsRetentionGroupBy{}
	var shouldExcludeMissing bool = false
	this.ShouldExcludeMissing = &shouldExcludeMissing
	return &this
}

// GetFacet returns the Facet field value.
func (o *ProductAnalyticsRetentionGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *ProductAnalyticsRetentionGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGroupBy) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGroupBy) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGroupBy) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProductAnalyticsRetentionGroupBy) SetLimit(v int64) {
	o.Limit = &v
}

// GetShouldExcludeMissing returns the ShouldExcludeMissing field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGroupBy) GetShouldExcludeMissing() bool {
	if o == nil || o.ShouldExcludeMissing == nil {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeMissing
}

// GetShouldExcludeMissingOk returns a tuple with the ShouldExcludeMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGroupBy) GetShouldExcludeMissingOk() (*bool, bool) {
	if o == nil || o.ShouldExcludeMissing == nil {
		return nil, false
	}
	return o.ShouldExcludeMissing, true
}

// HasShouldExcludeMissing returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGroupBy) HasShouldExcludeMissing() bool {
	return o != nil && o.ShouldExcludeMissing != nil
}

// SetShouldExcludeMissing gets a reference to the given bool and assigns it to the ShouldExcludeMissing field.
func (o *ProductAnalyticsRetentionGroupBy) SetShouldExcludeMissing(v bool) {
	o.ShouldExcludeMissing = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGroupBy) GetSort() ProductAnalyticsGroupBySort {
	if o == nil || o.Sort == nil {
		var ret ProductAnalyticsGroupBySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGroupBy) GetSortOk() (*ProductAnalyticsGroupBySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGroupBy) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ProductAnalyticsGroupBySort and assigns it to the Sort field.
func (o *ProductAnalyticsRetentionGroupBy) SetSort(v ProductAnalyticsGroupBySort) {
	o.Sort = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGroupBy) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGroupBy) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGroupBy) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ProductAnalyticsRetentionGroupBy) SetSource(v string) {
	o.Source = &v
}

// GetTarget returns the Target field value.
func (o *ProductAnalyticsRetentionGroupBy) GetTarget() ProductAnalyticsRetentionGroupByTarget {
	if o == nil {
		var ret ProductAnalyticsRetentionGroupByTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGroupBy) GetTargetOk() (*ProductAnalyticsRetentionGroupByTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ProductAnalyticsRetentionGroupBy) SetTarget(v ProductAnalyticsRetentionGroupByTarget) {
	o.Target = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionGroupBy) MarshalJSON() ([]byte, error) {
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
	toSerialize["target"] = o.Target

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet                *string                                 `json:"facet"`
		Limit                *int64                                  `json:"limit,omitempty"`
		ShouldExcludeMissing *bool                                   `json:"should_exclude_missing,omitempty"`
		Sort                 *ProductAnalyticsGroupBySort            `json:"sort,omitempty"`
		Source               *string                                 `json:"source,omitempty"`
		Target               *ProductAnalyticsRetentionGroupByTarget `json:"target"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "limit", "should_exclude_missing", "sort", "source", "target"})
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
	o.Source = all.Source
	if !all.Target.IsValid() {
		hasInvalidField = true
	} else {
		o.Target = *all.Target
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
