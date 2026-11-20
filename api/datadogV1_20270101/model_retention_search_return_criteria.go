// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchReturnCriteria Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchReturnCriteria struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	BaseQuery BaseProductAnalyticsQuery `json:"base_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TimeInterval *RetentionSearchReturnCriteriaTimeInterval `json:"time_interval,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionSearchReturnCriteria instantiates a new RetentionSearchReturnCriteria object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionSearchReturnCriteria(baseQuery BaseProductAnalyticsQuery) *RetentionSearchReturnCriteria {
	this := RetentionSearchReturnCriteria{}
	this.BaseQuery = baseQuery
	return &this
}

// NewRetentionSearchReturnCriteriaWithDefaults instantiates a new RetentionSearchReturnCriteria object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionSearchReturnCriteriaWithDefaults() *RetentionSearchReturnCriteria {
	this := RetentionSearchReturnCriteria{}
	return &this
}

// GetBaseQuery returns the BaseQuery field value.
func (o *RetentionSearchReturnCriteria) GetBaseQuery() BaseProductAnalyticsQuery {
	if o == nil {
		var ret BaseProductAnalyticsQuery
		return ret
	}
	return o.BaseQuery
}

// GetBaseQueryOk returns a tuple with the BaseQuery field value
// and a boolean to check if the value has been set.
func (o *RetentionSearchReturnCriteria) GetBaseQueryOk() (*BaseProductAnalyticsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseQuery, true
}

// SetBaseQuery sets field value.
func (o *RetentionSearchReturnCriteria) SetBaseQuery(v BaseProductAnalyticsQuery) {
	o.BaseQuery = v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *RetentionSearchReturnCriteria) GetTimeInterval() RetentionSearchReturnCriteriaTimeInterval {
	if o == nil || o.TimeInterval == nil {
		var ret RetentionSearchReturnCriteriaTimeInterval
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearchReturnCriteria) GetTimeIntervalOk() (*RetentionSearchReturnCriteriaTimeInterval, bool) {
	if o == nil || o.TimeInterval == nil {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *RetentionSearchReturnCriteria) HasTimeInterval() bool {
	return o != nil && o.TimeInterval != nil
}

// SetTimeInterval gets a reference to the given RetentionSearchReturnCriteriaTimeInterval and assigns it to the TimeInterval field.
func (o *RetentionSearchReturnCriteria) SetTimeInterval(v RetentionSearchReturnCriteriaTimeInterval) {
	o.TimeInterval = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionSearchReturnCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_query"] = o.BaseQuery
	if o.TimeInterval != nil {
		toSerialize["time_interval"] = o.TimeInterval
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionSearchReturnCriteria) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseQuery    *BaseProductAnalyticsQuery                 `json:"base_query"`
		TimeInterval *RetentionSearchReturnCriteriaTimeInterval `json:"time_interval,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseQuery == nil {
		return fmt.Errorf("required field base_query missing")
	}

	hasInvalidField := false
	o.BaseQuery = *all.BaseQuery
	if all.TimeInterval != nil && all.TimeInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeInterval = all.TimeInterval

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
