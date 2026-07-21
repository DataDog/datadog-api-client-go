// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptSDKDataAttributes Attributes of a flattened prompt version returned for SDK consumption. Exactly one of `template` and `chat_template` is returned.
type LLMObsPromptSDKDataAttributes struct {
	// Chat template for this prompt version, as a list of role and content messages. Omitted for text templates.
	ChatTemplate []LLMObsPromptChatMessage `json:"chat_template,omitempty"`
	// Labels attached to the selected version.
	// Deprecated
	Labels []string `json:"labels,omitempty"`
	// Customer-provided identifier of the prompt.
	PromptId *string `json:"prompt_id,omitempty"`
	// Unique identifier of this prompt version.
	PromptVersionUuid *string `json:"prompt_version_uuid,omitempty"`
	// Text template for this prompt version. Omitted for chat templates.
	Template *string `json:"template,omitempty"`
	// Version identifier for this prompt version. This is the sequential version number unless a user-supplied version identifier was set, in which case that identifier is used instead.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPromptSDKDataAttributes instantiates a new LLMObsPromptSDKDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptSDKDataAttributes() *LLMObsPromptSDKDataAttributes {
	this := LLMObsPromptSDKDataAttributes{}
	return &this
}

// NewLLMObsPromptSDKDataAttributesWithDefaults instantiates a new LLMObsPromptSDKDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptSDKDataAttributesWithDefaults() *LLMObsPromptSDKDataAttributes {
	this := LLMObsPromptSDKDataAttributes{}
	return &this
}

// GetChatTemplate returns the ChatTemplate field value if set, zero value otherwise.
func (o *LLMObsPromptSDKDataAttributes) GetChatTemplate() []LLMObsPromptChatMessage {
	if o == nil || o.ChatTemplate == nil {
		var ret []LLMObsPromptChatMessage
		return ret
	}
	return o.ChatTemplate
}

// GetChatTemplateOk returns a tuple with the ChatTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptSDKDataAttributes) GetChatTemplateOk() (*[]LLMObsPromptChatMessage, bool) {
	if o == nil || o.ChatTemplate == nil {
		return nil, false
	}
	return &o.ChatTemplate, true
}

// HasChatTemplate returns a boolean if a field has been set.
func (o *LLMObsPromptSDKDataAttributes) HasChatTemplate() bool {
	return o != nil && o.ChatTemplate != nil
}

// SetChatTemplate gets a reference to the given []LLMObsPromptChatMessage and assigns it to the ChatTemplate field.
func (o *LLMObsPromptSDKDataAttributes) SetChatTemplate(v []LLMObsPromptChatMessage) {
	o.ChatTemplate = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
// Deprecated
func (o *LLMObsPromptSDKDataAttributes) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LLMObsPromptSDKDataAttributes) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LLMObsPromptSDKDataAttributes) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
// Deprecated
func (o *LLMObsPromptSDKDataAttributes) SetLabels(v []string) {
	o.Labels = v
}

// GetPromptId returns the PromptId field value if set, zero value otherwise.
func (o *LLMObsPromptSDKDataAttributes) GetPromptId() string {
	if o == nil || o.PromptId == nil {
		var ret string
		return ret
	}
	return *o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptSDKDataAttributes) GetPromptIdOk() (*string, bool) {
	if o == nil || o.PromptId == nil {
		return nil, false
	}
	return o.PromptId, true
}

// HasPromptId returns a boolean if a field has been set.
func (o *LLMObsPromptSDKDataAttributes) HasPromptId() bool {
	return o != nil && o.PromptId != nil
}

// SetPromptId gets a reference to the given string and assigns it to the PromptId field.
func (o *LLMObsPromptSDKDataAttributes) SetPromptId(v string) {
	o.PromptId = &v
}

// GetPromptVersionUuid returns the PromptVersionUuid field value if set, zero value otherwise.
func (o *LLMObsPromptSDKDataAttributes) GetPromptVersionUuid() string {
	if o == nil || o.PromptVersionUuid == nil {
		var ret string
		return ret
	}
	return *o.PromptVersionUuid
}

// GetPromptVersionUuidOk returns a tuple with the PromptVersionUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptSDKDataAttributes) GetPromptVersionUuidOk() (*string, bool) {
	if o == nil || o.PromptVersionUuid == nil {
		return nil, false
	}
	return o.PromptVersionUuid, true
}

// HasPromptVersionUuid returns a boolean if a field has been set.
func (o *LLMObsPromptSDKDataAttributes) HasPromptVersionUuid() bool {
	return o != nil && o.PromptVersionUuid != nil
}

// SetPromptVersionUuid gets a reference to the given string and assigns it to the PromptVersionUuid field.
func (o *LLMObsPromptSDKDataAttributes) SetPromptVersionUuid(v string) {
	o.PromptVersionUuid = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *LLMObsPromptSDKDataAttributes) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptSDKDataAttributes) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *LLMObsPromptSDKDataAttributes) HasTemplate() bool {
	return o != nil && o.Template != nil
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *LLMObsPromptSDKDataAttributes) SetTemplate(v string) {
	o.Template = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LLMObsPromptSDKDataAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptSDKDataAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LLMObsPromptSDKDataAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LLMObsPromptSDKDataAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptSDKDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChatTemplate != nil {
		toSerialize["chat_template"] = o.ChatTemplate
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.PromptId != nil {
		toSerialize["prompt_id"] = o.PromptId
	}
	if o.PromptVersionUuid != nil {
		toSerialize["prompt_version_uuid"] = o.PromptVersionUuid
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptSDKDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChatTemplate      []LLMObsPromptChatMessage `json:"chat_template,omitempty"`
		Labels            []string                  `json:"labels,omitempty"`
		PromptId          *string                   `json:"prompt_id,omitempty"`
		PromptVersionUuid *string                   `json:"prompt_version_uuid,omitempty"`
		Template          *string                   `json:"template,omitempty"`
		Version           *string                   `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chat_template", "labels", "prompt_id", "prompt_version_uuid", "template", "version"})
	} else {
		return err
	}
	o.ChatTemplate = all.ChatTemplate
	o.Labels = all.Labels
	o.PromptId = all.PromptId
	o.PromptVersionUuid = all.PromptVersionUuid
	o.Template = all.Template
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
