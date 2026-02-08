// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookSearchAggregations Aggregations of notebook search results.
type NotebookSearchAggregations struct {
	// Aggregation by author.
	Author []NotebookSearchAggregationBucketMultiKey `json:"author,omitempty"`
	// Aggregation by tags.
	Tags []NotebookSearchAggregationBucketKey `json:"tags,omitempty"`
	// Aggregation by template variable names.
	TemplateVariablesName []NotebookSearchAggregationBucketKey `json:"template_variables.name,omitempty"`
	// Aggregation by notebook type.
	Type []NotebookSearchAggregationBucketKey `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotebookSearchAggregations instantiates a new NotebookSearchAggregations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookSearchAggregations() *NotebookSearchAggregations {
	this := NotebookSearchAggregations{}
	return &this
}

// NewNotebookSearchAggregationsWithDefaults instantiates a new NotebookSearchAggregations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookSearchAggregationsWithDefaults() *NotebookSearchAggregations {
	this := NotebookSearchAggregations{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *NotebookSearchAggregations) GetAuthor() []NotebookSearchAggregationBucketMultiKey {
	if o == nil || o.Author == nil {
		var ret []NotebookSearchAggregationBucketMultiKey
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchAggregations) GetAuthorOk() (*[]NotebookSearchAggregationBucketMultiKey, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return &o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *NotebookSearchAggregations) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given []NotebookSearchAggregationBucketMultiKey and assigns it to the Author field.
func (o *NotebookSearchAggregations) SetAuthor(v []NotebookSearchAggregationBucketMultiKey) {
	o.Author = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NotebookSearchAggregations) GetTags() []NotebookSearchAggregationBucketKey {
	if o == nil || o.Tags == nil {
		var ret []NotebookSearchAggregationBucketKey
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchAggregations) GetTagsOk() (*[]NotebookSearchAggregationBucketKey, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NotebookSearchAggregations) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []NotebookSearchAggregationBucketKey and assigns it to the Tags field.
func (o *NotebookSearchAggregations) SetTags(v []NotebookSearchAggregationBucketKey) {
	o.Tags = v
}

// GetTemplateVariablesName returns the TemplateVariablesName field value if set, zero value otherwise.
func (o *NotebookSearchAggregations) GetTemplateVariablesName() []NotebookSearchAggregationBucketKey {
	if o == nil || o.TemplateVariablesName == nil {
		var ret []NotebookSearchAggregationBucketKey
		return ret
	}
	return o.TemplateVariablesName
}

// GetTemplateVariablesNameOk returns a tuple with the TemplateVariablesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchAggregations) GetTemplateVariablesNameOk() (*[]NotebookSearchAggregationBucketKey, bool) {
	if o == nil || o.TemplateVariablesName == nil {
		return nil, false
	}
	return &o.TemplateVariablesName, true
}

// HasTemplateVariablesName returns a boolean if a field has been set.
func (o *NotebookSearchAggregations) HasTemplateVariablesName() bool {
	return o != nil && o.TemplateVariablesName != nil
}

// SetTemplateVariablesName gets a reference to the given []NotebookSearchAggregationBucketKey and assigns it to the TemplateVariablesName field.
func (o *NotebookSearchAggregations) SetTemplateVariablesName(v []NotebookSearchAggregationBucketKey) {
	o.TemplateVariablesName = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotebookSearchAggregations) GetType() []NotebookSearchAggregationBucketKey {
	if o == nil || o.Type == nil {
		var ret []NotebookSearchAggregationBucketKey
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotebookSearchAggregations) GetTypeOk() (*[]NotebookSearchAggregationBucketKey, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotebookSearchAggregations) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given []NotebookSearchAggregationBucketKey and assigns it to the Type field.
func (o *NotebookSearchAggregations) SetType(v []NotebookSearchAggregationBucketKey) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookSearchAggregations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TemplateVariablesName != nil {
		toSerialize["template_variables.name"] = o.TemplateVariablesName
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookSearchAggregations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author                []NotebookSearchAggregationBucketMultiKey `json:"author,omitempty"`
		Tags                  []NotebookSearchAggregationBucketKey      `json:"tags,omitempty"`
		TemplateVariablesName []NotebookSearchAggregationBucketKey      `json:"template_variables.name,omitempty"`
		Type                  []NotebookSearchAggregationBucketKey      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "tags", "template_variables.name", "type"})
	} else {
		return err
	}
	o.Author = all.Author
	o.Tags = all.Tags
	o.TemplateVariablesName = all.TemplateVariablesName
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
