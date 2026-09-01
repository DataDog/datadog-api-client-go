// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsFrontendAnnotatedInteractionItem A frontend interaction with its associated annotations.
type LLMObsFrontendAnnotatedInteractionItem struct {
	// List of annotations for this interaction.
	Annotations []LLMObsAnnotationItemResponse `json:"annotations"`
	// Whether the current caller can annotate this interaction.
	CanAnnotate bool `json:"can_annotate"`
	// Server-generated deterministic identifier derived from the content.
	ContentId string `json:"content_id"`
	// Web content that makes up a `frontend` interaction.
	Frontend LLMObsFrontendContent `json:"frontend"`
	// Unique identifier of the interaction.
	Id string `json:"id"`
	// Type discriminator for a `frontend` interaction.
	Type LLMObsFrontendInteractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsFrontendAnnotatedInteractionItem instantiates a new LLMObsFrontendAnnotatedInteractionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsFrontendAnnotatedInteractionItem(annotations []LLMObsAnnotationItemResponse, canAnnotate bool, contentId string, frontend LLMObsFrontendContent, id string, typeVar LLMObsFrontendInteractionType) *LLMObsFrontendAnnotatedInteractionItem {
	this := LLMObsFrontendAnnotatedInteractionItem{}
	this.Annotations = annotations
	this.CanAnnotate = canAnnotate
	this.ContentId = contentId
	this.Frontend = frontend
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewLLMObsFrontendAnnotatedInteractionItemWithDefaults instantiates a new LLMObsFrontendAnnotatedInteractionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsFrontendAnnotatedInteractionItemWithDefaults() *LLMObsFrontendAnnotatedInteractionItem {
	this := LLMObsFrontendAnnotatedInteractionItem{}
	return &this
}

// GetAnnotations returns the Annotations field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetAnnotations() []LLMObsAnnotationItemResponse {
	if o == nil {
		var ret []LLMObsAnnotationItemResponse
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetAnnotationsOk() (*[]LLMObsAnnotationItemResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) SetAnnotations(v []LLMObsAnnotationItemResponse) {
	o.Annotations = v
}

// GetCanAnnotate returns the CanAnnotate field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetCanAnnotate() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CanAnnotate
}

// GetCanAnnotateOk returns a tuple with the CanAnnotate field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetCanAnnotateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanAnnotate, true
}

// SetCanAnnotate sets field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) SetCanAnnotate(v bool) {
	o.CanAnnotate = v
}

// GetContentId returns the ContentId field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) SetContentId(v string) {
	o.ContentId = v
}

// GetFrontend returns the Frontend field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetFrontend() LLMObsFrontendContent {
	if o == nil {
		var ret LLMObsFrontendContent
		return ret
	}
	return o.Frontend
}

// GetFrontendOk returns a tuple with the Frontend field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetFrontendOk() (*LLMObsFrontendContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frontend, true
}

// SetFrontend sets field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) SetFrontend(v LLMObsFrontendContent) {
	o.Frontend = v
}

// GetId returns the Id field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetType() LLMObsFrontendInteractionType {
	if o == nil {
		var ret LLMObsFrontendInteractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendAnnotatedInteractionItem) GetTypeOk() (*LLMObsFrontendInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsFrontendAnnotatedInteractionItem) SetType(v LLMObsFrontendInteractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsFrontendAnnotatedInteractionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotations"] = o.Annotations
	toSerialize["can_annotate"] = o.CanAnnotate
	toSerialize["content_id"] = o.ContentId
	toSerialize["frontend"] = o.Frontend
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsFrontendAnnotatedInteractionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations *[]LLMObsAnnotationItemResponse `json:"annotations"`
		CanAnnotate *bool                           `json:"can_annotate"`
		ContentId   *string                         `json:"content_id"`
		Frontend    *LLMObsFrontendContent          `json:"frontend"`
		Id          *string                         `json:"id"`
		Type        *LLMObsFrontendInteractionType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	if all.CanAnnotate == nil {
		return fmt.Errorf("required field can_annotate missing")
	}
	if all.ContentId == nil {
		return fmt.Errorf("required field content_id missing")
	}
	if all.Frontend == nil {
		return fmt.Errorf("required field frontend missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotations", "can_annotate", "content_id", "frontend", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Annotations = *all.Annotations
	o.CanAnnotate = *all.CanAnnotate
	o.ContentId = *all.ContentId
	if all.Frontend.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Frontend = *all.Frontend
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
