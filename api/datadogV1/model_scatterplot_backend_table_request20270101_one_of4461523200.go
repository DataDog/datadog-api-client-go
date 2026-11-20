// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotBackendTableRequest20270101OneOf4461523200 Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotBackendTableRequest20270101OneOf4461523200 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryOptions map[string]DashboardInlineSchema2_20270101 `json:"_query_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotBackendTableRequest20270101OneOf4461523200 instantiates a new ScatterplotBackendTableRequest20270101OneOf4461523200 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotBackendTableRequest20270101OneOf4461523200(q string) *ScatterplotBackendTableRequest20270101OneOf4461523200 {
	this := ScatterplotBackendTableRequest20270101OneOf4461523200{}
	this.Q = q
	return &this
}

// NewScatterplotBackendTableRequest20270101OneOf4461523200WithDefaults instantiates a new ScatterplotBackendTableRequest20270101OneOf4461523200 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotBackendTableRequest20270101OneOf4461523200WithDefaults() *ScatterplotBackendTableRequest20270101OneOf4461523200 {
	this := ScatterplotBackendTableRequest20270101OneOf4461523200{}
	return &this
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) GetQueryOptions() map[string]DashboardInlineSchema2_20270101 {
	if o == nil || o.QueryOptions == nil {
		var ret map[string]DashboardInlineSchema2_20270101
		return ret
	}
	return o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) GetQueryOptionsOk() (*map[string]DashboardInlineSchema2_20270101, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return &o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given map[string]DashboardInlineSchema2_20270101 and assigns it to the QueryOptions field.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) SetQueryOptions(v map[string]DashboardInlineSchema2_20270101) {
	o.QueryOptions = v
}

// GetQ returns the Q field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) SetQ(v string) {
	o.Q = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotBackendTableRequest20270101OneOf4461523200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryOptions != nil {
		toSerialize["_query_options"] = o.QueryOptions
	}
	toSerialize["q"] = o.Q
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotBackendTableRequest20270101OneOf4461523200) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryOptions map[string]DashboardInlineSchema2_20270101 `json:"_query_options,omitempty"`
		Q            *string                                    `json:"q"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}
	o.QueryOptions = all.QueryOptions
	o.Q = *all.Q

	return nil
}
