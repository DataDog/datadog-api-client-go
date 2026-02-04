// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardSearchAggregationBucketMultiKey Aggregation bucket with multiple key values.
type DashboardSearchAggregationBucketMultiKey struct {
	// Number of results in this bucket.
	Count int64 `json:"count"`
	// Key values for this bucket.
	Value []string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardSearchAggregationBucketMultiKey instantiates a new DashboardSearchAggregationBucketMultiKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardSearchAggregationBucketMultiKey(count int64, value []string) *DashboardSearchAggregationBucketMultiKey {
	this := DashboardSearchAggregationBucketMultiKey{}
	this.Count = count
	this.Value = value
	return &this
}

// NewDashboardSearchAggregationBucketMultiKeyWithDefaults instantiates a new DashboardSearchAggregationBucketMultiKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardSearchAggregationBucketMultiKeyWithDefaults() *DashboardSearchAggregationBucketMultiKey {
	this := DashboardSearchAggregationBucketMultiKey{}
	return &this
}

// GetCount returns the Count field value.
func (o *DashboardSearchAggregationBucketMultiKey) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregationBucketMultiKey) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *DashboardSearchAggregationBucketMultiKey) SetCount(v int64) {
	o.Count = v
}

// GetValue returns the Value field value.
func (o *DashboardSearchAggregationBucketMultiKey) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchAggregationBucketMultiKey) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *DashboardSearchAggregationBucketMultiKey) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardSearchAggregationBucketMultiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["count"] = o.Count
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardSearchAggregationBucketMultiKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int64    `json:"count"`
		Value *[]string `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "value"})
	} else {
		return err
	}
	o.Count = *all.Count
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
