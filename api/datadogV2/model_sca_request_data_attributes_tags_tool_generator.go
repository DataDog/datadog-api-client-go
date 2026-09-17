// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesTagsToolGenerator Metadata about the tool that generated the SCA tags.
type ScaRequestDataAttributesTagsToolGenerator struct {
	// The name of the tag generator.
	Name *string `json:"name,omitempty"`
	// The version of the tag generator.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesTagsToolGenerator instantiates a new ScaRequestDataAttributesTagsToolGenerator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesTagsToolGenerator() *ScaRequestDataAttributesTagsToolGenerator {
	this := ScaRequestDataAttributesTagsToolGenerator{}
	return &this
}

// NewScaRequestDataAttributesTagsToolGeneratorWithDefaults instantiates a new ScaRequestDataAttributesTagsToolGenerator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesTagsToolGeneratorWithDefaults() *ScaRequestDataAttributesTagsToolGenerator {
	this := ScaRequestDataAttributesTagsToolGenerator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesTagsToolGenerator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesTagsToolGenerator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesTagsToolGenerator) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScaRequestDataAttributesTagsToolGenerator) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesTagsToolGenerator) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesTagsToolGenerator) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesTagsToolGenerator) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ScaRequestDataAttributesTagsToolGenerator) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesTagsToolGenerator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesTagsToolGenerator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string `json:"name,omitempty"`
		Version *string `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Name = all.Name
	o.Version = all.Version

	return nil
}
