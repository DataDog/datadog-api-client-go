// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapView20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type GeomapView20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Focus string `json:"focus"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewGeomapView20270101OneOf instantiates a new GeomapView20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapView20270101OneOf(focus string) *GeomapView20270101OneOf {
	this := GeomapView20270101OneOf{}
	this.Focus = focus
	return &this
}

// NewGeomapView20270101OneOfWithDefaults instantiates a new GeomapView20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapView20270101OneOfWithDefaults() *GeomapView20270101OneOf {
	this := GeomapView20270101OneOf{}
	return &this
}

// GetFocus returns the Focus field value.
func (o *GeomapView20270101OneOf) GetFocus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Focus
}

// GetFocusOk returns a tuple with the Focus field value
// and a boolean to check if the value has been set.
func (o *GeomapView20270101OneOf) GetFocusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Focus, true
}

// SetFocus sets field value.
func (o *GeomapView20270101OneOf) SetFocus(v string) {
	o.Focus = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapView20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["focus"] = o.Focus
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapView20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Focus *string `json:"focus"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Focus == nil {
		return fmt.Errorf("required field focus missing")
	}
	o.Focus = *all.Focus

	return nil
}
