// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendAxisRequest Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendAxisRequest struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryOptions map[string]DashboardInlineSchema45 `json:"_query_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Aggregator *ScalarAggregator `json:"aggregator,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendAxisRequest instantiates a new ScatterplotBackendAxisRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendAxisRequest(q string) *ScatterplotBackendAxisRequest {
	this := ScatterplotBackendAxisRequest{}
	this.Q = q
	return &this
}

// NewScatterplotBackendAxisRequestWithDefaults instantiates a new ScatterplotBackendAxisRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendAxisRequestWithDefaults() *ScatterplotBackendAxisRequest {
	this := ScatterplotBackendAxisRequest{}
	return &this
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *ScatterplotBackendAxisRequest) GetQueryOptions() map[string]DashboardInlineSchema45 {
	if o == nil || o.QueryOptions == nil {
		var ret map[string]DashboardInlineSchema45
		return ret
	}
	return o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendAxisRequest) GetQueryOptionsOk() (*map[string]DashboardInlineSchema45, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return &o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *ScatterplotBackendAxisRequest) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given map[string]DashboardInlineSchema45 and assigns it to the QueryOptions field.
func (o *ScatterplotBackendAxisRequest) SetQueryOptions(v map[string]DashboardInlineSchema45) {
	o.QueryOptions = v
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *ScatterplotBackendAxisRequest) GetAggregator() ScalarAggregator {
	if o == nil || o.Aggregator == nil {
		var ret ScalarAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendAxisRequest) GetAggregatorOk() (*ScalarAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *ScatterplotBackendAxisRequest) HasAggregator() bool {
	return o != nil && o.Aggregator != nil
}

// SetAggregator gets a reference to the given ScalarAggregator and assigns it to the Aggregator field.
func (o *ScatterplotBackendAxisRequest) SetAggregator(v ScalarAggregator) {
	o.Aggregator = &v
}

// GetQ returns the Q field value.
func (o *ScatterplotBackendAxisRequest) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendAxisRequest) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *ScatterplotBackendAxisRequest) SetQ(v string) {
	o.Q = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendAxisRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryOptions != nil {
		toSerialize["_query_options"] = o.QueryOptions
	}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	toSerialize["q"] = o.Q
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotBackendAxisRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryOptions map[string]DashboardInlineSchema45 `json:"_query_options,omitempty"`
		Aggregator   *ScalarAggregator                  `json:"aggregator,omitempty"`
		Q            *string                            `json:"q"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}

	hasInvalidField := false
	o.QueryOptions = all.QueryOptions
	if all.Aggregator != nil && !all.Aggregator.IsValid() {
		hasInvalidField = true
	} else {
		o.Aggregator = all.Aggregator
	}
	o.Q = *all.Q

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
