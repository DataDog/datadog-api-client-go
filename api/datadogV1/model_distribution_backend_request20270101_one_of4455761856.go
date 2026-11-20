// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455761856 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455761856 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryOptions map[string]DashboardInlineSchema7_20270101 `json:"_query_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455761856 instantiates a new DistributionBackendRequest20270101OneOf4455761856 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455761856(q string) *DistributionBackendRequest20270101OneOf4455761856 {
	this := DistributionBackendRequest20270101OneOf4455761856{}
	this.Q = q
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455761856WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455761856 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455761856WithDefaults() *DistributionBackendRequest20270101OneOf4455761856 {
	this := DistributionBackendRequest20270101OneOf4455761856{}
	return &this
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455761856) GetQueryOptions() map[string]DashboardInlineSchema7_20270101 {
	if o == nil || o.QueryOptions == nil {
		var ret map[string]DashboardInlineSchema7_20270101
		return ret
	}
	return o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455761856) GetQueryOptionsOk() (*map[string]DashboardInlineSchema7_20270101, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return &o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455761856) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given map[string]DashboardInlineSchema7_20270101 and assigns it to the QueryOptions field.
func (o *DistributionBackendRequest20270101OneOf4455761856) SetQueryOptions(v map[string]DashboardInlineSchema7_20270101) {
	o.QueryOptions = v
}

// GetQ returns the Q field value.
func (o *DistributionBackendRequest20270101OneOf4455761856) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455761856) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *DistributionBackendRequest20270101OneOf4455761856) SetQ(v string) {
	o.Q = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455761856) GetStyle() RequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret RequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455761856) GetStyleOk() (*RequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455761856) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle_20270101 and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455761856) SetStyle(v RequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455761856) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryOptions != nil {
		toSerialize["_query_options"] = o.QueryOptions
	}
	toSerialize["q"] = o.Q
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455761856) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryOptions map[string]DashboardInlineSchema7_20270101 `json:"_query_options,omitempty"`
		Q            *string                                    `json:"q"`
		Style        *RequestStyle_20270101                     `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}
	o.QueryOptions = all.QueryOptions
	o.Q = *all.Q
	o.Style = all.Style

	return nil
}
