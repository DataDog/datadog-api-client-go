// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitDimension Test-only generated schema for the dashboard versioning proof of concept.
type SplitDimension struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	OneGraphPer string `json:"one_graph_per"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSplitDimension instantiates a new SplitDimension object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSplitDimension(oneGraphPer string) *SplitDimension {
	this := SplitDimension{}
	this.OneGraphPer = oneGraphPer
	return &this
}

// NewSplitDimensionWithDefaults instantiates a new SplitDimension object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSplitDimensionWithDefaults() *SplitDimension {
	this := SplitDimension{}
	return &this
}

// GetOneGraphPer returns the OneGraphPer field value.
func (o *SplitDimension) GetOneGraphPer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OneGraphPer
}

// GetOneGraphPerOk returns a tuple with the OneGraphPer field value
// and a boolean to check if the value has been set.
func (o *SplitDimension) GetOneGraphPerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OneGraphPer, true
}

// SetOneGraphPer sets field value.
func (o *SplitDimension) SetOneGraphPer(v string) {
	o.OneGraphPer = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SplitDimension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["one_graph_per"] = o.OneGraphPer
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SplitDimension) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OneGraphPer *string `json:"one_graph_per"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OneGraphPer == nil {
		return fmt.Errorf("required field one_graph_per missing")
	}
	o.OneGraphPer = *all.OneGraphPer

	return nil
}
