// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionCohortCriteria Defines the event that places an entity into a cohort, and how cohorts are bucketed over time.
type ProductAnalyticsRetentionCohortCriteria struct {
	// A query definition discriminated by the `data_source` field.
	// Use `product_analytics` for standard event queries, or
	// `product_analytics_occurrence` for occurrence-filtered queries.
	BaseQuery ProductAnalyticsBaseQuery `json:"base_query"`
	// A retention interval, either aligned to calendar boundaries or of a fixed length.
	// Cohort criteria use calendar intervals; return criteria use fixed intervals.
	TimeInterval ProductAnalyticsRetentionTimeInterval `json:"time_interval"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionCohortCriteria instantiates a new ProductAnalyticsRetentionCohortCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionCohortCriteria(baseQuery ProductAnalyticsBaseQuery, timeInterval ProductAnalyticsRetentionTimeInterval) *ProductAnalyticsRetentionCohortCriteria {
	this := ProductAnalyticsRetentionCohortCriteria{}
	this.BaseQuery = baseQuery
	this.TimeInterval = timeInterval
	return &this
}

// NewProductAnalyticsRetentionCohortCriteriaWithDefaults instantiates a new ProductAnalyticsRetentionCohortCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionCohortCriteriaWithDefaults() *ProductAnalyticsRetentionCohortCriteria {
	this := ProductAnalyticsRetentionCohortCriteria{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *ProductAnalyticsRetentionCohortCriteria) GetBaseQuery() ProductAnalyticsBaseQuery {
	if o == nil {
		var ret ProductAnalyticsBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCohortCriteria) GetBaseQueryOk() (*ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *ProductAnalyticsRetentionCohortCriteria) SetBaseQuery(v ProductAnalyticsBaseQuery) {
	o.BaseQuery = v
}

// GetTimeInterval returns the TimeInterval field value.
func (o *ProductAnalyticsRetentionCohortCriteria) GetTimeInterval() ProductAnalyticsRetentionTimeInterval {
	if o == nil {
		var ret ProductAnalyticsRetentionTimeInterval
		return ret
	}
	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionCohortCriteria) GetTimeIntervalOk() (*ProductAnalyticsRetentionTimeInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value.
func (o *ProductAnalyticsRetentionCohortCriteria) SetTimeInterval(v ProductAnalyticsRetentionTimeInterval) {
	o.TimeInterval = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionCohortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	toSerialize["time_interval"] = o.TimeInterval

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionCohortCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *ProductAnalyticsBaseQuery             `json:"base_query"`
		TimeInterval *ProductAnalyticsRetentionTimeInterval `json:"time_interval"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}
	if all.TimeInterval == nil {
		return fmt.Errorf("required field time_interval missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"base_query", "time_interval"})
	} else {
		return err
	}
	o.BaseQuery = *all.BaseQuery
	o.TimeInterval = *all.TimeInterval

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
