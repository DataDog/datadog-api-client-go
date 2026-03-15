// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeySearch Search parameters for a Sankey query.
type ProductAnalyticsSankeySearch struct {
	// Audience filter definitions for targeting specific user segments.
	AudienceFilters *ProductAnalyticsAudienceFilters `json:"audience_filters,omitempty"`
	// Join key configuration for correlating events.
	JoinKeys *ProductAnalyticsJoinKeys `json:"join_keys,omitempty"`
	// Filter for occurrence-based queries.
	Occurrences *ProductAnalyticsOccurrenceFilter `json:"occurrences,omitempty"`
	// The search query. Cannot be empty.
	Query string `json:"query"`
	//
	SubqueryId *string `json:"subquery_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSankeySearch instantiates a new ProductAnalyticsSankeySearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSankeySearch(query string) *ProductAnalyticsSankeySearch {
	this := ProductAnalyticsSankeySearch{}
	this.Query = query
	return &this
}

// NewProductAnalyticsSankeySearchWithDefaults instantiates a new ProductAnalyticsSankeySearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSankeySearchWithDefaults() *ProductAnalyticsSankeySearch {
	this := ProductAnalyticsSankeySearch{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeySearch) GetAudienceFilters() ProductAnalyticsAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret ProductAnalyticsAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeySearch) GetAudienceFiltersOk() (*ProductAnalyticsAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeySearch) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given ProductAnalyticsAudienceFilters and assigns it to the AudienceFilters field.
func (o *ProductAnalyticsSankeySearch) SetAudienceFilters(v ProductAnalyticsAudienceFilters) {
	o.AudienceFilters = &v
}

// GetJoinKeys returns the JoinKeys field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeySearch) GetJoinKeys() ProductAnalyticsJoinKeys {
	if o == nil || o.JoinKeys == nil {
		var ret ProductAnalyticsJoinKeys
		return ret
	}
	return *o.JoinKeys
}

// GetJoinKeysOk returns a tuple with the JoinKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeySearch) GetJoinKeysOk() (*ProductAnalyticsJoinKeys, bool) {
	if o == nil || o.JoinKeys == nil {
		return nil, false
	}
	return o.JoinKeys, true
}

// HasJoinKeys returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeySearch) HasJoinKeys() bool {
	return o != nil && o.JoinKeys != nil
}

// SetJoinKeys gets a reference to the given ProductAnalyticsJoinKeys and assigns it to the JoinKeys field.
func (o *ProductAnalyticsSankeySearch) SetJoinKeys(v ProductAnalyticsJoinKeys) {
	o.JoinKeys = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeySearch) GetOccurrences() ProductAnalyticsOccurrenceFilter {
	if o == nil || o.Occurrences == nil {
		var ret ProductAnalyticsOccurrenceFilter
		return ret
	}
	return *o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeySearch) GetOccurrencesOk() (*ProductAnalyticsOccurrenceFilter, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeySearch) HasOccurrences() bool {
	return o != nil && o.Occurrences != nil
}

// SetOccurrences gets a reference to the given ProductAnalyticsOccurrenceFilter and assigns it to the Occurrences field.
func (o *ProductAnalyticsSankeySearch) SetOccurrences(v ProductAnalyticsOccurrenceFilter) {
	o.Occurrences = &v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsSankeySearch) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeySearch) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsSankeySearch) SetQuery(v string) {
	o.Query = v
}

// GetSubqueryId returns the SubqueryId field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeySearch) GetSubqueryId() string {
	if o == nil || o.SubqueryId == nil {
		var ret string
		return ret
	}
	return *o.SubqueryId
}

// GetSubqueryIdOk returns a tuple with the SubqueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeySearch) GetSubqueryIdOk() (*string, bool) {
	if o == nil || o.SubqueryId == nil {
		return nil, false
	}
	return o.SubqueryId, true
}

// HasSubqueryId returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeySearch) HasSubqueryId() bool {
	return o != nil && o.SubqueryId != nil
}

// SetSubqueryId gets a reference to the given string and assigns it to the SubqueryId field.
func (o *ProductAnalyticsSankeySearch) SetSubqueryId(v string) {
	o.SubqueryId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSankeySearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.JoinKeys != nil {
		toSerialize["join_keys"] = o.JoinKeys
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	toSerialize["query"] = o.Query
	if o.SubqueryId != nil {
		toSerialize["subquery_id"] = o.SubqueryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsSankeySearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *ProductAnalyticsAudienceFilters  `json:"audience_filters,omitempty"`
		JoinKeys        *ProductAnalyticsJoinKeys         `json:"join_keys,omitempty"`
		Occurrences     *ProductAnalyticsOccurrenceFilter `json:"occurrences,omitempty"`
		Query           *string                           `json:"query"`
		SubqueryId      *string                           `json:"subquery_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "join_keys", "occurrences", "query", "subquery_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	if all.JoinKeys != nil && all.JoinKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinKeys = all.JoinKeys
	if all.Occurrences != nil && all.Occurrences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Occurrences = all.Occurrences
	o.Query = *all.Query
	o.SubqueryId = all.SubqueryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
