// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookSearchResultData A notebook search result.
type NotebookSearchResultData struct {
	// Highlighted fields from the notebook search.
	Highlight *NotebookSearchHighlight `json:"highlight,omitempty"`
	// Notebook identifier.
	Id string `json:"id"`
	// Notebook search result attributes.
	NotebookData NotebookSearchAttributes `json:"notebook_data"`
	// Notebook resource type.
	Type MetricNotebookType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookSearchResultData instantiates a new NotebookSearchResultData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookSearchResultData(id string, notebookData NotebookSearchAttributes, typeVar MetricNotebookType) *NotebookSearchResultData {
	this := NotebookSearchResultData{}
	this.Id = id
	this.NotebookData = notebookData
	this.Type = typeVar
	return &this
}

// NewNotebookSearchResultDataWithDefaults instantiates a new NotebookSearchResultData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookSearchResultDataWithDefaults() *NotebookSearchResultData {
	this := NotebookSearchResultData{}
	return &this
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *NotebookSearchResultData) GetHighlight() NotebookSearchHighlight {
	if o == nil || o.Highlight == nil {
		var ret NotebookSearchHighlight
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchResultData) GetHighlightOk() (*NotebookSearchHighlight, bool) {
	if o == nil || o.Highlight == nil {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *NotebookSearchResultData) HasHighlight() bool {
	return o != nil && o.Highlight != nil
}

// SetHighlight gets a reference to the given NotebookSearchHighlight and assigns it to the Highlight field.
func (o *NotebookSearchResultData) SetHighlight(v NotebookSearchHighlight) {
	o.Highlight = &v
}

// GetId returns the Id field value.
func (o *NotebookSearchResultData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchResultData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *NotebookSearchResultData) SetId(v string) {
	o.Id = v
}

// GetNotebookData returns the NotebookData field value.
func (o *NotebookSearchResultData) GetNotebookData() NotebookSearchAttributes {
	if o == nil {
		var ret NotebookSearchAttributes
		return ret
	}
	return o.NotebookData
}

// GetNotebookDataOk returns a tuple with the NotebookData field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchResultData) GetNotebookDataOk() (*NotebookSearchAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotebookData, true
}

// SetNotebookData sets field value.
func (o *NotebookSearchResultData) SetNotebookData(v NotebookSearchAttributes) {
	o.NotebookData = v
}

// GetType returns the Type field value.
func (o *NotebookSearchResultData) GetType() MetricNotebookType {
	if o == nil {
		var ret MetricNotebookType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotebookSearchResultData) GetTypeOk() (*MetricNotebookType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *NotebookSearchResultData) SetType(v MetricNotebookType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookSearchResultData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Highlight != nil {
		toSerialize["highlight"] = o.Highlight
	}
	toSerialize["id"] = o.Id
	toSerialize["notebook_data"] = o.NotebookData
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookSearchResultData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Highlight    *NotebookSearchHighlight  `json:"highlight,omitempty"`
		Id           *string                   `json:"id"`
		NotebookData *NotebookSearchAttributes `json:"notebook_data"`
		Type         *MetricNotebookType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.NotebookData == nil {
		return fmt.Errorf("required field notebook_data missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"highlight", "id", "notebook_data", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Highlight != nil && all.Highlight.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Highlight = all.Highlight
	o.Id = *all.Id
	if all.NotebookData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NotebookData = *all.NotebookData
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
