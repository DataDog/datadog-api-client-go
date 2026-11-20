// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitVectorTagFilter Test-only generated schema for the dashboard versioning proof of concept.
type SplitVectorTagFilter struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	TagKey string `json:"tag_key"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TagValues []string `json:"tag_values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSplitVectorTagFilter instantiates a new SplitVectorTagFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitVectorTagFilter(tagKey string, tagValues []string) *SplitVectorTagFilter {
	this := SplitVectorTagFilter{}
	this.TagKey = tagKey
	this.TagValues = tagValues
	return &this
}

// NewSplitVectorTagFilterWithDefaults instantiates a new SplitVectorTagFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitVectorTagFilterWithDefaults() *SplitVectorTagFilter {
	this := SplitVectorTagFilter{}
	return &this
}

// GetTagKey returns the TagKey field value.
func (o *SplitVectorTagFilter) GetTagKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TagKey
}

// GetTagKeyOk returns a tuple with the TagKey field value
// and a boolean to check if the value has been set.
func (o *SplitVectorTagFilter) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKey, true
}

// SetTagKey sets field value.
func (o *SplitVectorTagFilter) SetTagKey(v string) {
	o.TagKey = v
}

// GetTagValues returns the TagValues field value.
func (o *SplitVectorTagFilter) GetTagValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value
// and a boolean to check if the value has been set.
func (o *SplitVectorTagFilter) GetTagValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagValues, true
}

// SetTagValues sets field value.
func (o *SplitVectorTagFilter) SetTagValues(v []string) {
	o.TagValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitVectorTagFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["tag_key"] = o.TagKey
	toSerialize["tag_values"] = o.TagValues
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitVectorTagFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TagKey    *string   `json:"tag_key"`
		TagValues *[]string `json:"tag_values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TagKey == nil {
		return fmt.Errorf("required field tag_key missing")
	}
	if all.TagValues == nil {
		return fmt.Errorf("required field tag_values missing")
	}
	o.TagKey = *all.TagKey
	o.TagValues = *all.TagValues

	return nil
}
