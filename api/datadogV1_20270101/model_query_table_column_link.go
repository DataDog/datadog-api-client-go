// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryTableColumnLink Test-only generated schema for the dashboard versioning proof of concept.
type QueryTableColumnLink struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Label *string `json:"label,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewQueryTableColumnLink instantiates a new QueryTableColumnLink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryTableColumnLink(url string) *QueryTableColumnLink {
	this := QueryTableColumnLink{}
	this.Url = url
	return &this
}

// NewQueryTableColumnLinkWithDefaults instantiates a new QueryTableColumnLink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryTableColumnLinkWithDefaults() *QueryTableColumnLink {
	this := QueryTableColumnLink{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *QueryTableColumnLink) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTableColumnLink) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *QueryTableColumnLink) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *QueryTableColumnLink) SetLabel(v string) {
	o.Label = &v
}

// GetUrl returns the Url field value.
func (o *QueryTableColumnLink) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *QueryTableColumnLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *QueryTableColumnLink) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryTableColumnLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	toSerialize["url"] = o.Url
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryTableColumnLink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label *string `json:"label,omitempty"`
		Url   *string `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	o.Label = all.Label
	o.Url = *all.Url

	return nil
}
