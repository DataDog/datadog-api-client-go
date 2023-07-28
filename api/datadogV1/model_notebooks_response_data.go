// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebooksResponseData The data for a notebook in get all response.
type NotebooksResponseData struct {
	// The attributes of a notebook in get all response.
	Attributes NotebooksResponseDataAttributes `json:"attributes"`
	// Unique notebook ID, assigned when you create the notebook.
	Id int64 `json:"id"`
	// Type of the Notebook resource.
	Type NotebookResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewNotebooksResponseData instantiates a new NotebooksResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebooksResponseData(attributes NotebooksResponseDataAttributes, id int64, typeVar NotebookResourceType) *NotebooksResponseData {
	this := NotebooksResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewNotebooksResponseDataWithDefaults instantiates a new NotebooksResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebooksResponseDataWithDefaults() *NotebooksResponseData {
	this := NotebooksResponseData{}
	var typeVar NotebookResourceType = NOTEBOOKRESOURCETYPE_NOTEBOOKS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *NotebooksResponseData) GetAttributes() NotebooksResponseDataAttributes {
	if o == nil {
		var ret NotebooksResponseDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *NotebooksResponseData) GetAttributesOk() (*NotebooksResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *NotebooksResponseData) SetAttributes(v NotebooksResponseDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *NotebooksResponseData) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotebooksResponseData) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *NotebooksResponseData) SetId(v int64) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *NotebooksResponseData) GetType() NotebookResourceType {
	if o == nil {
		var ret NotebookResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebooksResponseData) GetTypeOk() (*NotebookResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotebooksResponseData) SetType(v NotebookResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebooksResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebooksResponseData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes *NotebooksResponseDataAttributes `json:"attributes"`
		Id         *int64                           `json:"id"`
		Type       *NotebookResourceType            `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	var hasInvalidField bool
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	} else {
		o.Attributes = *all.Attributes
	}
	o.Id = *all.Id
	if v := all.Type; !v.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}

	return nil
}
