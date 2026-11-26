// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserDefinedLink Test-only generated schema for the dashboard versioning proof of concept.
type UserDefinedLink struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label string `json:"label"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Link string `json:"link"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewUserDefinedLink instantiates a new UserDefinedLink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserDefinedLink(label string, link string) *UserDefinedLink {
	this := UserDefinedLink{}
	this.Label = label
	this.Link = link
	return &this
}

// NewUserDefinedLinkWithDefaults instantiates a new UserDefinedLink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserDefinedLinkWithDefaults() *UserDefinedLink {
	this := UserDefinedLink{}
	return &this
}

// GetLabel returns the Label field value.
func (o *UserDefinedLink) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *UserDefinedLink) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *UserDefinedLink) SetLabel(v string) {
	o.Label = v
}

// GetLink returns the Link field value.
func (o *UserDefinedLink) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *UserDefinedLink) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value.
func (o *UserDefinedLink) SetLink(v string) {
	o.Link = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserDefinedLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["label"] = o.Label
	toSerialize["link"] = o.Link
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserDefinedLink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label *string `json:"label"`
		Link  *string `json:"link"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Link == nil {
		return fmt.Errorf("required field link missing")
	}
	o.Label = *all.Label
	o.Link = *all.Link

	return nil
}
