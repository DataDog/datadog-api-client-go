// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyDropOffQuery Query definition for a journey drop-off analysis request.
type ProductAnalyticsJourneyDropOffQuery struct {
	// Sampling configuration for a query.
	Sampling *ProductAnalyticsSampling `json:"sampling,omitempty"`
	// Defines the steps of the journey and the filters applied to it.
	Search ProductAnalyticsJourneySearch `json:"search"`
	// A reference to a single step of the journey.
	Target *ProductAnalyticsJourneyNodeTarget `json:"target,omitempty"`
	// Maximum number of issues to return.
	TopIssueCount *int64 `json:"top_issue_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyDropOffQuery instantiates a new ProductAnalyticsJourneyDropOffQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyDropOffQuery(search ProductAnalyticsJourneySearch) *ProductAnalyticsJourneyDropOffQuery {
	this := ProductAnalyticsJourneyDropOffQuery{}
	this.Search = search
	var topIssueCount int64 = 5
	this.TopIssueCount = &topIssueCount
	return &this
}

// NewProductAnalyticsJourneyDropOffQueryWithDefaults instantiates a new ProductAnalyticsJourneyDropOffQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyDropOffQueryWithDefaults() *ProductAnalyticsJourneyDropOffQuery {
	this := ProductAnalyticsJourneyDropOffQuery{}
	var topIssueCount int64 = 5
	this.TopIssueCount = &topIssueCount
	return &this
}

// GetSampling returns the Sampling field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyDropOffQuery) GetSampling() ProductAnalyticsSampling {
	if o == nil || o.Sampling == nil {
		var ret ProductAnalyticsSampling
		return ret
	}
	return *o.Sampling
}

// GetSamplingOk returns a tuple with the Sampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffQuery) GetSamplingOk() (*ProductAnalyticsSampling, bool) {
	if o == nil || o.Sampling == nil {
		return nil, false
	}
	return o.Sampling, true
}

// HasSampling returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyDropOffQuery) HasSampling() bool {
	return o != nil && o.Sampling != nil
}

// SetSampling gets a reference to the given ProductAnalyticsSampling and assigns it to the Sampling field.
func (o *ProductAnalyticsJourneyDropOffQuery) SetSampling(v ProductAnalyticsSampling) {
	o.Sampling = &v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsJourneyDropOffQuery) GetSearch() ProductAnalyticsJourneySearch {
	if o == nil {
		var ret ProductAnalyticsJourneySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffQuery) GetSearchOk() (*ProductAnalyticsJourneySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsJourneyDropOffQuery) SetSearch(v ProductAnalyticsJourneySearch) {
	o.Search = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyDropOffQuery) GetTarget() ProductAnalyticsJourneyNodeTarget {
	if o == nil || o.Target == nil {
		var ret ProductAnalyticsJourneyNodeTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffQuery) GetTargetOk() (*ProductAnalyticsJourneyNodeTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyDropOffQuery) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given ProductAnalyticsJourneyNodeTarget and assigns it to the Target field.
func (o *ProductAnalyticsJourneyDropOffQuery) SetTarget(v ProductAnalyticsJourneyNodeTarget) {
	o.Target = &v
}

// GetTopIssueCount returns the TopIssueCount field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyDropOffQuery) GetTopIssueCount() int64 {
	if o == nil || o.TopIssueCount == nil {
		var ret int64
		return ret
	}
	return *o.TopIssueCount
}

// GetTopIssueCountOk returns a tuple with the TopIssueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffQuery) GetTopIssueCountOk() (*int64, bool) {
	if o == nil || o.TopIssueCount == nil {
		return nil, false
	}
	return o.TopIssueCount, true
}

// HasTopIssueCount returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyDropOffQuery) HasTopIssueCount() bool {
	return o != nil && o.TopIssueCount != nil
}

// SetTopIssueCount gets a reference to the given int64 and assigns it to the TopIssueCount field.
func (o *ProductAnalyticsJourneyDropOffQuery) SetTopIssueCount(v int64) {
	o.TopIssueCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyDropOffQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Sampling != nil {
		toSerialize["sampling"] = o.Sampling
	}
	toSerialize["search"] = o.Search
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.TopIssueCount != nil {
		toSerialize["top_issue_count"] = o.TopIssueCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyDropOffQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sampling      *ProductAnalyticsSampling          `json:"sampling,omitempty"`
		Search        *ProductAnalyticsJourneySearch     `json:"search"`
		Target        *ProductAnalyticsJourneyNodeTarget `json:"target,omitempty"`
		TopIssueCount *int64                             `json:"top_issue_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"sampling", "search", "target", "top_issue_count"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Sampling != nil && all.Sampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sampling = all.Sampling
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	o.TopIssueCount = all.TopIssueCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
