// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455770336 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455770336 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AppsecspanQuery LogQueryDefinition_20270101 `json:"appsecspan_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle_20270101 `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455770336 instantiates a new DistributionBackendRequest20270101OneOf4455770336 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455770336(appsecspanQuery LogQueryDefinition_20270101) *DistributionBackendRequest20270101OneOf4455770336 {
	this := DistributionBackendRequest20270101OneOf4455770336{}
	this.AppsecspanQuery = appsecspanQuery
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455770336WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455770336 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455770336WithDefaults() *DistributionBackendRequest20270101OneOf4455770336 {
	this := DistributionBackendRequest20270101OneOf4455770336{}
	return &this
}

// GetAppsecspanQuery returns the AppsecspanQuery field value.
func (o *DistributionBackendRequest20270101OneOf4455770336) GetAppsecspanQuery() LogQueryDefinition_20270101 {
	if o == nil {
		var ret LogQueryDefinition_20270101
		return ret
	}
	return o.AppsecspanQuery
}

// GetAppsecspanQueryOk returns a tuple with the AppsecspanQuery field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455770336) GetAppsecspanQueryOk() (*LogQueryDefinition_20270101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppsecspanQuery, true
}

// SetAppsecspanQuery sets field value.
func (o *DistributionBackendRequest20270101OneOf4455770336) SetAppsecspanQuery(v LogQueryDefinition_20270101) {
	o.AppsecspanQuery = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455770336) GetStyle() RequestStyle_20270101 {
	if o == nil || o.Style == nil {
		var ret RequestStyle_20270101
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455770336) GetStyleOk() (*RequestStyle_20270101, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455770336) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle_20270101 and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455770336) SetStyle(v RequestStyle_20270101) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455770336) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["appsecspan_query"] = o.AppsecspanQuery
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455770336) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppsecspanQuery *LogQueryDefinition_20270101 `json:"appsecspan_query"`
		Style           *RequestStyle_20270101       `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AppsecspanQuery == nil {
		return fmt.Errorf("required field appsecspan_query missing")
	}
	o.AppsecspanQuery = *all.AppsecspanQuery
	o.Style = all.Style

	return nil
}
