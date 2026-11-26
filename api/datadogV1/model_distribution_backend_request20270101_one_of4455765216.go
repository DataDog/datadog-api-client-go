// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455765216 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455765216 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	RumQuery LogQueryDefinition_20270101 `json:"rum_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455765216 instantiates a new DistributionBackendRequest20270101OneOf4455765216 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455765216(rumQuery LogQueryDefinition_20270101) *DistributionBackendRequest20270101OneOf4455765216 {
	this := DistributionBackendRequest20270101OneOf4455765216{}
	this.RumQuery = rumQuery
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455765216WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455765216 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455765216WithDefaults() *DistributionBackendRequest20270101OneOf4455765216 {
	this := DistributionBackendRequest20270101OneOf4455765216{}
	return &this
}

// GetRumQuery returns the RumQuery field value.
func (o *DistributionBackendRequest20270101OneOf4455765216) GetRumQuery() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455765216) GetRumQueryOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RumQuery, true
}

// SetRumQuery sets field value.
func (o *DistributionBackendRequest20270101OneOf4455765216) SetRumQuery(v LogQueryDefinition_20270101) {
	o.RumQuery = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455765216) GetStyle() RequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret RequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455765216) GetStyleOk() (*RequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455765216) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle_20270101 and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455765216) SetStyle(v RequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455765216) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["rum_query"] = o.RumQuery
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455765216) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RumQuery *LogQueryDefinition_20270101 `json:"rum_query"`
		Style    *RequestStyle_20270101       `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RumQuery == nil {
		return fmt.Errorf("required field rum_query missing")
	}
	o.RumQuery = *all.RumQuery
	o.Style = all.Style

	return nil
}
