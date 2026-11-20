// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchCohortCriteria Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchCohortCriteria struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery BaseProductAnalyticsQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeInterval RetentionSearchCohortCriteriaTimeInterval `json:"time_interval"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionSearchCohortCriteria instantiates a new RetentionSearchCohortCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionSearchCohortCriteria(baseQuery BaseProductAnalyticsQuery, timeInterval RetentionSearchCohortCriteriaTimeInterval) *RetentionSearchCohortCriteria {
	this := RetentionSearchCohortCriteria{}
	this.BaseQuery = baseQuery
	this.TimeInterval = timeInterval
	return &this
}

// NewRetentionSearchCohortCriteriaWithDefaults instantiates a new RetentionSearchCohortCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionSearchCohortCriteriaWithDefaults() *RetentionSearchCohortCriteria {
	this := RetentionSearchCohortCriteria{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *RetentionSearchCohortCriteria) GetBaseQuery() BaseProductAnalyticsQuery {
	if o == nil {
		var ret BaseProductAnalyticsQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchCohortCriteria) GetBaseQueryOk() (*BaseProductAnalyticsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *RetentionSearchCohortCriteria) SetBaseQuery(v BaseProductAnalyticsQuery) {
	o.BaseQuery = v
}

// GetTimeInterval returns the TimeInterval field value.
func (o *RetentionSearchCohortCriteria) GetTimeInterval() RetentionSearchCohortCriteriaTimeInterval {
	if o == nil {
		var ret RetentionSearchCohortCriteriaTimeInterval
		return ret
	}
	return o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchCohortCriteria) GetTimeIntervalOk() (*RetentionSearchCohortCriteriaTimeInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeInterval, true
}

// SetTimeInterval sets field value.
func (o *RetentionSearchCohortCriteria) SetTimeInterval(v RetentionSearchCohortCriteriaTimeInterval) {
	o.TimeInterval = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionSearchCohortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	toSerialize["time_interval"] = o.TimeInterval
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionSearchCohortCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *BaseProductAnalyticsQuery                 `json:"base_query"`
		TimeInterval *RetentionSearchCohortCriteriaTimeInterval `json:"time_interval"`
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

	hasInvalidField := false
	o.BaseQuery = *all.BaseQuery
	if all.TimeInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeInterval = *all.TimeInterval

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
