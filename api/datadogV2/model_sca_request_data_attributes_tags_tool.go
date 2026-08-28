// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesTagsTool Tool metadata included in SCA tags.
type ScaRequestDataAttributesTagsTool struct {
	// Metadata about the tool that generated the SCA tags.
	Generator *ScaRequestDataAttributesTagsToolGenerator `json:"generator,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesTagsTool instantiates a new ScaRequestDataAttributesTagsTool object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesTagsTool() *ScaRequestDataAttributesTagsTool {
	this := ScaRequestDataAttributesTagsTool{}
	return &this
}

// NewScaRequestDataAttributesTagsToolWithDefaults instantiates a new ScaRequestDataAttributesTagsTool object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesTagsToolWithDefaults() *ScaRequestDataAttributesTagsTool {
	this := ScaRequestDataAttributesTagsTool{}
	return &this
}

// GetGenerator returns the Generator field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesTagsTool) GetGenerator() ScaRequestDataAttributesTagsToolGenerator {
	if o == nil || o.Generator == nil {
		var ret ScaRequestDataAttributesTagsToolGenerator
		return ret
	}
	return *o.Generator
}

// GetGeneratorOk returns a tuple with the Generator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesTagsTool) GetGeneratorOk() (*ScaRequestDataAttributesTagsToolGenerator, bool) {
	if o == nil || o.Generator == nil {
		return nil, false
	}
	return o.Generator, true
}

// HasGenerator returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesTagsTool) HasGenerator() bool {
	return o != nil && o.Generator != nil
}

// SetGenerator gets a reference to the given ScaRequestDataAttributesTagsToolGenerator and assigns it to the Generator field.
func (o *ScaRequestDataAttributesTagsTool) SetGenerator(v ScaRequestDataAttributesTagsToolGenerator) {
	o.Generator = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesTagsTool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Generator != nil {
		toSerialize["generator"] = o.Generator
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesTagsTool) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Generator *ScaRequestDataAttributesTagsToolGenerator `json:"generator,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Generator != nil && all.Generator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Generator = all.Generator

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
