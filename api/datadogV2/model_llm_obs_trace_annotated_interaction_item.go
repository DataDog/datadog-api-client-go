// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsTraceAnnotatedInteractionItem A trace, experiment trace, or session interaction with its associated annotations.
type LLMObsTraceAnnotatedInteractionItem struct {
	// List of annotations for this interaction.
	Annotations []LLMObsAnnotationItem `json:"annotations"`
	// Upstream entity identifier supplied by the caller.
	ContentId string `json:"content_id"`
	// Unique identifier of the interaction.
	Id string `json:"id"`
	// Type of an upstream-entity interaction.
	Type LLMObsTraceInteractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsTraceAnnotatedInteractionItem instantiates a new LLMObsTraceAnnotatedInteractionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsTraceAnnotatedInteractionItem(annotations []LLMObsAnnotationItem, contentId string, id string, typeVar LLMObsTraceInteractionType) *LLMObsTraceAnnotatedInteractionItem {
	this := LLMObsTraceAnnotatedInteractionItem{}
	this.Annotations = annotations
	this.ContentId = contentId
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewLLMObsTraceAnnotatedInteractionItemWithDefaults instantiates a new LLMObsTraceAnnotatedInteractionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsTraceAnnotatedInteractionItemWithDefaults() *LLMObsTraceAnnotatedInteractionItem {
	this := LLMObsTraceAnnotatedInteractionItem{}
	return &this
}

// GetAnnotations returns the Annotations field value.
func (o *LLMObsTraceAnnotatedInteractionItem) GetAnnotations() []LLMObsAnnotationItem {
	if o == nil {
		var ret []LLMObsAnnotationItem
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *LLMObsTraceAnnotatedInteractionItem) GetAnnotationsOk() (*[]LLMObsAnnotationItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *LLMObsTraceAnnotatedInteractionItem) SetAnnotations(v []LLMObsAnnotationItem) {
	o.Annotations = v
}

// GetContentId returns the ContentId field value.
func (o *LLMObsTraceAnnotatedInteractionItem) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *LLMObsTraceAnnotatedInteractionItem) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value.
func (o *LLMObsTraceAnnotatedInteractionItem) SetContentId(v string) {
	o.ContentId = v
}

// GetId returns the Id field value.
func (o *LLMObsTraceAnnotatedInteractionItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsTraceAnnotatedInteractionItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsTraceAnnotatedInteractionItem) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *LLMObsTraceAnnotatedInteractionItem) GetType() LLMObsTraceInteractionType {
	if o == nil {
		var ret LLMObsTraceInteractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsTraceAnnotatedInteractionItem) GetTypeOk() (*LLMObsTraceInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsTraceAnnotatedInteractionItem) SetType(v LLMObsTraceInteractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsTraceAnnotatedInteractionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotations"] = o.Annotations
	toSerialize["content_id"] = o.ContentId
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsTraceAnnotatedInteractionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations *[]LLMObsAnnotationItem     `json:"annotations"`
		ContentId   *string                     `json:"content_id"`
		Id          *string                     `json:"id"`
		Type        *LLMObsTraceInteractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	if all.ContentId == nil {
		return fmt.Errorf("required field content_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotations", "content_id", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Annotations = *all.Annotations
	o.ContentId = *all.ContentId
	o.Id = *all.Id
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
