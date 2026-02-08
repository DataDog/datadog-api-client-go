// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookSearchAttributes Notebook search result attributes.
type NotebookSearchAttributes struct {
	// Metadata about the notebook.
	Meta NotebookSearchMetadata `json:"meta"`
	// Name of the notebook.
	Name string `json:"name"`
	// List of tags for the notebook.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookSearchAttributes instantiates a new NotebookSearchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookSearchAttributes(meta NotebookSearchMetadata, name string, tags []string) *NotebookSearchAttributes {
	this := NotebookSearchAttributes{}
	this.Meta = meta
	this.Name = name
	this.Tags = tags
	return &this
}

// NewNotebookSearchAttributesWithDefaults instantiates a new NotebookSearchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookSearchAttributesWithDefaults() *NotebookSearchAttributes {
	this := NotebookSearchAttributes{}
	return &this
}

// GetMeta returns the Meta field value.
func (o *NotebookSearchAttributes) GetMeta() NotebookSearchMetadata {
	if o == nil {
		var ret NotebookSearchMetadata
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchAttributes) GetMetaOk() (*NotebookSearchMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *NotebookSearchAttributes) SetMeta(v NotebookSearchMetadata) {
	o.Meta = v
}

// GetName returns the Name field value.
func (o *NotebookSearchAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *NotebookSearchAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value.
func (o *NotebookSearchAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *NotebookSearchAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookSearchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["meta"] = o.Meta
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookSearchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Meta *NotebookSearchMetadata `json:"meta"`
		Name *string                 `json:"name"`
		Tags *[]string               `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"meta", "name", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta
	o.Name = *all.Name
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
