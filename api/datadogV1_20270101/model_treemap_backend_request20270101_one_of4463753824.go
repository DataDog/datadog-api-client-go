// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreemapBackendRequest20270101OneOf4463753824 Test-only generated schema for the dashboard versioning proof of concept.
type TreemapBackendRequest20270101OneOf4463753824 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *TreemapRequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreemapBackendRequest20270101OneOf4463753824 instantiates a new TreemapBackendRequest20270101OneOf4463753824 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreemapBackendRequest20270101OneOf4463753824(q string) *TreemapBackendRequest20270101OneOf4463753824 {
	this := TreemapBackendRequest20270101OneOf4463753824{}
	this.Q = q
	return &this
}

// NewTreemapBackendRequest20270101OneOf4463753824WithDefaults instantiates a new TreemapBackendRequest20270101OneOf4463753824 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreemapBackendRequest20270101OneOf4463753824WithDefaults() *TreemapBackendRequest20270101OneOf4463753824 {
	this := TreemapBackendRequest20270101OneOf4463753824{}
	return &this
}

// GetQ returns the Q field value.
func (o *TreemapBackendRequest20270101OneOf4463753824) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463753824) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *TreemapBackendRequest20270101OneOf4463753824) SetQ(v string) {
	o.Q = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *TreemapBackendRequest20270101OneOf4463753824) GetStyle() TreemapRequestStyle {
	if o == nil || o.Style == nil {
		var ret TreemapRequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreemapBackendRequest20270101OneOf4463753824) GetStyleOk() (*TreemapRequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *TreemapBackendRequest20270101OneOf4463753824) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given TreemapRequestStyle and assigns it to the Style field.
func (o *TreemapBackendRequest20270101OneOf4463753824) SetStyle(v TreemapRequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreemapBackendRequest20270101OneOf4463753824) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["q"] = o.Q
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TreemapBackendRequest20270101OneOf4463753824) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Q     *string              `json:"q"`
		Style *TreemapRequestStyle `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}

	hasInvalidField := false
	o.Q = *all.Q
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
