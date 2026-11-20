// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsGroupBy Test-only generated schema for the dashboard versioning proof of concept.
type ProductAnalyticsGroupBy struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Facet string `json:"facet"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShouldExcludeMissing *bool `json:"should_exclude_missing,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *GroupBySort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Source *string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewProductAnalyticsGroupBy instantiates a new ProductAnalyticsGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsGroupBy(facet string) *ProductAnalyticsGroupBy {
	this := ProductAnalyticsGroupBy{}
	this.Facet = facet
	return &this
}

// NewProductAnalyticsGroupByWithDefaults instantiates a new ProductAnalyticsGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsGroupByWithDefaults() *ProductAnalyticsGroupBy {
	this := ProductAnalyticsGroupBy{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *ProductAnalyticsGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *ProductAnalyticsGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProductAnalyticsGroupBy) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBy) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductAnalyticsGroupBy) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *ProductAnalyticsGroupBy) SetLimit(v float64) {
	o.Limit = &v
}

// GetShouldExcludeMissing returns the ShouldExcludeMissing field value if set, zero value otherwise.
func (o *ProductAnalyticsGroupBy) GetShouldExcludeMissing() bool {
	if o == nil || o.ShouldExcludeMissing == nil {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeMissing
}

// GetShouldExcludeMissingOk returns a tuple with the ShouldExcludeMissing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBy) GetShouldExcludeMissingOk() (*bool, bool) {
	if o == nil || o.ShouldExcludeMissing == nil {
		return nil, false
	}
	return o.ShouldExcludeMissing, true
}

// HasShouldExcludeMissing returns a boolean if a field has been set.
func (o *ProductAnalyticsGroupBy) HasShouldExcludeMissing() bool {
	return o != nil && o.ShouldExcludeMissing != nil
}

// SetShouldExcludeMissing gets a reference to the given bool and assigns it to the ShouldExcludeMissing field.
func (o *ProductAnalyticsGroupBy) SetShouldExcludeMissing(v bool) {
	o.ShouldExcludeMissing = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProductAnalyticsGroupBy) GetSort() GroupBySort {
	if o == nil || o.Sort == nil {
		var ret GroupBySort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBy) GetSortOk() (*GroupBySort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProductAnalyticsGroupBy) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given GroupBySort and assigns it to the Sort field.
func (o *ProductAnalyticsGroupBy) SetSort(v GroupBySort) {
	o.Sort = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProductAnalyticsGroupBy) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsGroupBy) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProductAnalyticsGroupBy) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ProductAnalyticsGroupBy) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsGroupBy) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet                *string      `json:"facet"`
		Limit                *float64     `json:"limit,omitempty"`
		ShouldExcludeMissing *bool        `json:"should_exclude_missing,omitempty"`
		Sort                 *GroupBySort `json:"sort,omitempty"`
		Source               *string      `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
