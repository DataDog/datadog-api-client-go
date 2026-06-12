// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AIGuardContentPart A single part of a multipart message content.
type AIGuardContentPart struct {
	// An image URL reference for multimodal content.
	ImageUrl *AIGuardImageURL `json:"image_url,omitempty"`
	// The text content of this part, required when type is text.
	Text *string `json:"text,omitempty"`
	// The type of content part, either text or image_url.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIGuardContentPart instantiates a new AIGuardContentPart object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIGuardContentPart(typeVar string) *AIGuardContentPart {
	this := AIGuardContentPart{}
	this.Type = typeVar
	return &this
}

// NewAIGuardContentPartWithDefaults instantiates a new AIGuardContentPart object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIGuardContentPartWithDefaults() *AIGuardContentPart {
	this := AIGuardContentPart{}
	return &this
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *AIGuardContentPart) GetImageUrl() AIGuardImageURL {
	if o == nil || o.ImageUrl == nil {
		var ret AIGuardImageURL
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardContentPart) GetImageUrlOk() (*AIGuardImageURL, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *AIGuardContentPart) HasImageUrl() bool {
	return o != nil && o.ImageUrl != nil
}

// SetImageUrl gets a reference to the given AIGuardImageURL and assigns it to the ImageUrl field.
func (o *AIGuardContentPart) SetImageUrl(v AIGuardImageURL) {
	o.ImageUrl = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AIGuardContentPart) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIGuardContentPart) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AIGuardContentPart) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AIGuardContentPart) SetText(v string) {
	o.Text = &v
}

// GetType returns the Type field value.
func (o *AIGuardContentPart) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AIGuardContentPart) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AIGuardContentPart) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIGuardContentPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIGuardContentPart) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ImageUrl *AIGuardImageURL `json:"image_url,omitempty"`
		Text     *string          `json:"text,omitempty"`
		Type     *string          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"image_url", "text", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ImageUrl != nil && all.ImageUrl.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ImageUrl = all.ImageUrl
	o.Text = all.Text
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
