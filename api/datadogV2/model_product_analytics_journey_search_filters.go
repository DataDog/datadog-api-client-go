// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneySearchFilters Filters applied on top of the journey step expression.
type ProductAnalyticsJourneySearchFilters struct {
	// Restricts the journey to an audience built from named sub-queries.
	// Sub-query names must be unique across `users`, `segments`, and `accounts`.
	AudienceFilters *ProductAnalyticsJourneyAudienceFilters `json:"audience_filters,omitempty"`
	// Filters on journey-level metrics such as time to convert.
	GraphFilters []ProductAnalyticsJourneySearchGraphFilter `json:"graph_filters,omitempty"`
	// Free-text search query applied to the whole journey.
	StringFilter *string `json:"string_filter,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneySearchFilters instantiates a new ProductAnalyticsJourneySearchFilters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneySearchFilters() *ProductAnalyticsJourneySearchFilters {
	this := ProductAnalyticsJourneySearchFilters{}
	return &this
}

// NewProductAnalyticsJourneySearchFiltersWithDefaults instantiates a new ProductAnalyticsJourneySearchFilters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneySearchFiltersWithDefaults() *ProductAnalyticsJourneySearchFilters {
	this := ProductAnalyticsJourneySearchFilters{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneySearchFilters) GetAudienceFilters() ProductAnalyticsJourneyAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret ProductAnalyticsJourneyAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearchFilters) GetAudienceFiltersOk() (*ProductAnalyticsJourneyAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneySearchFilters) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given ProductAnalyticsJourneyAudienceFilters and assigns it to the AudienceFilters field.
func (o *ProductAnalyticsJourneySearchFilters) SetAudienceFilters(v ProductAnalyticsJourneyAudienceFilters) {
	o.AudienceFilters = &v
}

// GetGraphFilters returns the GraphFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneySearchFilters) GetGraphFilters() []ProductAnalyticsJourneySearchGraphFilter {
	if o == nil || o.GraphFilters == nil {
		var ret []ProductAnalyticsJourneySearchGraphFilter
		return ret
	}
	return o.GraphFilters
}

// GetGraphFiltersOk returns a tuple with the GraphFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearchFilters) GetGraphFiltersOk() (*[]ProductAnalyticsJourneySearchGraphFilter, bool) {
	if o == nil || o.GraphFilters == nil {
		return nil, false
	}
	return &o.GraphFilters, true
}

// HasGraphFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneySearchFilters) HasGraphFilters() bool {
	return o != nil && o.GraphFilters != nil
}

// SetGraphFilters gets a reference to the given []ProductAnalyticsJourneySearchGraphFilter and assigns it to the GraphFilters field.
func (o *ProductAnalyticsJourneySearchFilters) SetGraphFilters(v []ProductAnalyticsJourneySearchGraphFilter) {
	o.GraphFilters = v
}

// GetStringFilter returns the StringFilter field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneySearchFilters) GetStringFilter() string {
	if o == nil || o.StringFilter == nil {
		var ret string
		return ret
	}
	return *o.StringFilter
}

// GetStringFilterOk returns a tuple with the StringFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearchFilters) GetStringFilterOk() (*string, bool) {
	if o == nil || o.StringFilter == nil {
		return nil, false
	}
	return o.StringFilter, true
}

// HasStringFilter returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneySearchFilters) HasStringFilter() bool {
	return o != nil && o.StringFilter != nil
}

// SetStringFilter gets a reference to the given string and assigns it to the StringFilter field.
func (o *ProductAnalyticsJourneySearchFilters) SetStringFilter(v string) {
	o.StringFilter = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneySearchFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.GraphFilters != nil {
		toSerialize["graph_filters"] = o.GraphFilters
	}
	if o.StringFilter != nil {
		toSerialize["string_filter"] = o.StringFilter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneySearchFilters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *ProductAnalyticsJourneyAudienceFilters    `json:"audience_filters,omitempty"`
		GraphFilters    []ProductAnalyticsJourneySearchGraphFilter `json:"graph_filters,omitempty"`
		StringFilter    *string                                    `json:"string_filter,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "graph_filters", "string_filter"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	o.GraphFilters = all.GraphFilters
	o.StringFilter = all.StringFilter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
