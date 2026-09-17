// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionReturnCriteria Defines the event that counts as a return, and the window in which it must occur.
type ProductAnalyticsRetentionReturnCriteria struct {
	// A query definition discriminated by the `data_source` field.
	// Use `product_analytics` for standard event queries, or
	// `product_analytics_occurrence` for occurrence-filtered queries.
	BaseQuery ProductAnalyticsBaseQuery `json:"base_query"`
	// A retention interval, either aligned to calendar boundaries or of a fixed length.
	// Cohort criteria use calendar intervals; return criteria use fixed intervals.
	TimeInterval *ProductAnalyticsRetentionTimeInterval `json:"time_interval,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionReturnCriteria instantiates a new ProductAnalyticsRetentionReturnCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionReturnCriteria(baseQuery ProductAnalyticsBaseQuery) *ProductAnalyticsRetentionReturnCriteria {
	this := ProductAnalyticsRetentionReturnCriteria{}
	this.BaseQuery = baseQuery
	return &this
}

// NewProductAnalyticsRetentionReturnCriteriaWithDefaults instantiates a new ProductAnalyticsRetentionReturnCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionReturnCriteriaWithDefaults() *ProductAnalyticsRetentionReturnCriteria {
	this := ProductAnalyticsRetentionReturnCriteria{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *ProductAnalyticsRetentionReturnCriteria) GetBaseQuery() ProductAnalyticsBaseQuery {
	if o == nil {
		var ret ProductAnalyticsBaseQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionReturnCriteria) GetBaseQueryOk() (*ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *ProductAnalyticsRetentionReturnCriteria) SetBaseQuery(v ProductAnalyticsBaseQuery) {
	o.BaseQuery = v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionReturnCriteria) GetTimeInterval() ProductAnalyticsRetentionTimeInterval {
	if o == nil || o.TimeInterval == nil {
		var ret ProductAnalyticsRetentionTimeInterval
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionReturnCriteria) GetTimeIntervalOk() (*ProductAnalyticsRetentionTimeInterval, bool) {
	if o == nil || o.TimeInterval == nil {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionReturnCriteria) HasTimeInterval() bool {
	return o != nil && o.TimeInterval != nil
}

// SetTimeInterval gets a reference to the given ProductAnalyticsRetentionTimeInterval and assigns it to the TimeInterval field.
func (o *ProductAnalyticsRetentionReturnCriteria) SetTimeInterval(v ProductAnalyticsRetentionTimeInterval) {
	o.TimeInterval = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionReturnCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	if o.TimeInterval != nil {
		toSerialize["time_interval"] = o.TimeInterval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionReturnCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *ProductAnalyticsBaseQuery             `json:"base_query"`
		TimeInterval *ProductAnalyticsRetentionTimeInterval `json:"time_interval,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"base_query", "time_interval"})
	} else {
		return err
	}
	o.BaseQuery = *all.BaseQuery
	o.TimeInterval = all.TimeInterval

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
