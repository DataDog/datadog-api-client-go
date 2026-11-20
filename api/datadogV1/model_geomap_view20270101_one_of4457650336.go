// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapView20270101OneOf4457650336 Test-only generated schema for the dashboard versioning proof of concept.
type GeomapView20270101OneOf4457650336 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomExtent []interface{} `json:"custom_extent"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapView20270101OneOf4457650336 instantiates a new GeomapView20270101OneOf4457650336 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapView20270101OneOf4457650336(customExtent []interface{}) *GeomapView20270101OneOf4457650336 {
	this := GeomapView20270101OneOf4457650336{}
	this.CustomExtent = customExtent
	return &this
}

// NewGeomapView20270101OneOf4457650336WithDefaults instantiates a new GeomapView20270101OneOf4457650336 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapView20270101OneOf4457650336WithDefaults() *GeomapView20270101OneOf4457650336 {
	this := GeomapView20270101OneOf4457650336{}
	return &this
}

// GetCustomExtent returns the CustomExtent field value.
func (o *GeomapView20270101OneOf4457650336) GetCustomExtent() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.CustomExtent
}

// GetCustomExtentOk returns a tuple with the CustomExtent field value
// and a boolean to check if the value has been set.
func (o *GeomapView20270101OneOf4457650336) GetCustomExtentOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomExtent, true
}

// SetCustomExtent sets field value.
func (o *GeomapView20270101OneOf4457650336) SetCustomExtent(v []interface{}) {
	o.CustomExtent = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapView20270101OneOf4457650336) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["custom_extent"] = o.CustomExtent
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapView20270101OneOf4457650336) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomExtent *[]interface{} `json:"custom_extent"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CustomExtent == nil {
		return fmt.Errorf("required field custom_extent missing")
	}
	o.CustomExtent = *all.CustomExtent

	return nil
}
