// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregateSegmentQuerySegmentQuery Test-only generated schema for the dashboard versioning proof of concept.
type AggregateSegmentQuerySegmentQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Id string `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAggregateSegmentQuerySegmentQuery instantiates a new AggregateSegmentQuerySegmentQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregateSegmentQuerySegmentQuery(id string) *AggregateSegmentQuerySegmentQuery {
	this := AggregateSegmentQuerySegmentQuery{}
	this.Id = id
	return &this
}

// NewAggregateSegmentQuerySegmentQueryWithDefaults instantiates a new AggregateSegmentQuerySegmentQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregateSegmentQuerySegmentQueryWithDefaults() *AggregateSegmentQuerySegmentQuery {
	this := AggregateSegmentQuerySegmentQuery{}
	return &this
}

// GetId returns the Id field value.
func (o *AggregateSegmentQuerySegmentQuery) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AggregateSegmentQuerySegmentQuery) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AggregateSegmentQuerySegmentQuery) SetId(v string) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregateSegmentQuerySegmentQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregateSegmentQuerySegmentQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	o.Id = *all.Id

	return nil
}
