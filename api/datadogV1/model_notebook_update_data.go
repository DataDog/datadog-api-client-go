// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookUpdateData The data for a notebook update request.
type NotebookUpdateData struct {
	// The data attributes of a notebook.
	Attributes NotebookUpdateDataAttributes `json:"attributes"`
	// Type of the Notebook resource.
	Type NotebookResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewNotebookUpdateData instantiates a new NotebookUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookUpdateData(attributes NotebookUpdateDataAttributes, typeVar NotebookResourceType) *NotebookUpdateData {
	this := NotebookUpdateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewNotebookUpdateDataWithDefaults instantiates a new NotebookUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookUpdateDataWithDefaults() *NotebookUpdateData {
	this := NotebookUpdateData{}
	var typeVar NotebookResourceType = NOTEBOOKRESOURCETYPE_NOTEBOOKS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *NotebookUpdateData) GetAttributes() NotebookUpdateDataAttributes {
	if o == nil {
		var ret NotebookUpdateDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *NotebookUpdateData) GetAttributesOk() (*NotebookUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *NotebookUpdateData) SetAttributes(v NotebookUpdateDataAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *NotebookUpdateData) GetType() NotebookResourceType {
	if o == nil {
		var ret NotebookResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebookUpdateData) GetTypeOk() (*NotebookResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotebookUpdateData) SetType(v NotebookResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *NotebookUpdateDataAttributes `json:"attributes"`
		Type       *NotebookResourceType         `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
