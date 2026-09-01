// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsFrontendInteractionResponseItem A frontend interaction result.
type LLMObsFrontendInteractionResponseItem struct {
	// Whether this interaction already existed in the queue.
	AlreadyExisted bool `json:"already_existed"`
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

// NewLLMObsFrontendInteractionResponseItem instantiates a new LLMObsFrontendInteractionResponseItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsFrontendInteractionResponseItem(alreadyExisted bool, contentId string, frontend LLMObsFrontendContent, id string, typeVar LLMObsFrontendInteractionType) *LLMObsFrontendInteractionResponseItem {
	this := LLMObsFrontendInteractionResponseItem{}
	this.AlreadyExisted = alreadyExisted
	this.ContentId = contentId
	this.Frontend = frontend
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewLLMObsFrontendInteractionResponseItemWithDefaults instantiates a new LLMObsFrontendInteractionResponseItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsFrontendInteractionResponseItemWithDefaults() *LLMObsFrontendInteractionResponseItem {
	this := LLMObsFrontendInteractionResponseItem{}
	return &this
}

// GetAlreadyExisted returns the AlreadyExisted field value.
func (o *LLMObsFrontendInteractionResponseItem) GetAlreadyExisted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AlreadyExisted
}

// GetAlreadyExistedOk returns a tuple with the AlreadyExisted field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendInteractionResponseItem) GetAlreadyExistedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlreadyExisted, true
}

// SetAlreadyExisted sets field value.
func (o *LLMObsFrontendInteractionResponseItem) SetAlreadyExisted(v bool) {
	o.AlreadyExisted = v
}

// GetContentId returns the ContentId field value.
func (o *LLMObsFrontendInteractionResponseItem) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendInteractionResponseItem) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value.
func (o *LLMObsFrontendInteractionResponseItem) SetContentId(v string) {
	o.ContentId = v
}

// GetFrontend returns the Frontend field value.
func (o *LLMObsFrontendInteractionResponseItem) GetFrontend() LLMObsFrontendContent {
	if o == nil {
		var ret LLMObsFrontendContent
		return ret
	}
	return o.Frontend
}

// GetFrontendOk returns a tuple with the Frontend field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendInteractionResponseItem) GetFrontendOk() (*LLMObsFrontendContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frontend, true
}

// SetFrontend sets field value.
func (o *LLMObsFrontendInteractionResponseItem) SetFrontend(v LLMObsFrontendContent) {
	o.Frontend = v
}

// GetId returns the Id field value.
func (o *LLMObsFrontendInteractionResponseItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendInteractionResponseItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsFrontendInteractionResponseItem) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *LLMObsFrontendInteractionResponseItem) GetType() LLMObsFrontendInteractionType {
	if o == nil {
		var ret LLMObsFrontendInteractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsFrontendInteractionResponseItem) GetTypeOk() (*LLMObsFrontendInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsFrontendInteractionResponseItem) SetType(v LLMObsFrontendInteractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsFrontendInteractionResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["already_existed"] = o.AlreadyExisted
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
func (o *LLMObsFrontendInteractionResponseItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlreadyExisted *bool                          `json:"already_existed"`
		ContentId      *string                        `json:"content_id"`
		Frontend       *LLMObsFrontendContent         `json:"frontend"`
		Id             *string                        `json:"id"`
		Type           *LLMObsFrontendInteractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AlreadyExisted == nil {
		return fmt.Errorf("required field already_existed missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"already_existed", "content_id", "frontend", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AlreadyExisted = *all.AlreadyExisted
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
