// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCreatePromptDataAttributes Attributes for creating an LLM Observability prompt and its first version. `prompt_id` and `template` are required; all other attributes are optional.
type LLMObsCreatePromptDataAttributes struct {
	// Optional description of the prompt.
	Description *string `json:"description,omitempty"`
	// Optional feature-flag environment UUIDs the service attempts to enable and configure to use the first version as their default after creation.
	EnvIds []string `json:"env_ids,omitempty"`
	// Optional labels to attach to the first version. Do not use this attribute for new integrations.
	// Deprecated
	Labels []LLMObsPromptVersionLabel `json:"labels,omitempty"`
	// Customer-provided identifier for the new prompt.
	PromptId string `json:"prompt_id"`
	// A text template or a list of chat messages.
	Template LLMObsPromptTemplate `json:"template"`
	// Optional title of the prompt.
	Title *string `json:"title,omitempty"`
	// Optional user-supplied version identifier for the first version.
	UserVersion *string `json:"user_version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCreatePromptDataAttributes instantiates a new LLMObsCreatePromptDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCreatePromptDataAttributes(promptId string, template LLMObsPromptTemplate) *LLMObsCreatePromptDataAttributes {
	this := LLMObsCreatePromptDataAttributes{}
	this.PromptId = promptId
	this.Template = template
	return &this
}

// NewLLMObsCreatePromptDataAttributesWithDefaults instantiates a new LLMObsCreatePromptDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCreatePromptDataAttributesWithDefaults() *LLMObsCreatePromptDataAttributes {
	this := LLMObsCreatePromptDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsCreatePromptDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsCreatePromptDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsCreatePromptDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnvIds returns the EnvIds field value if set, zero value otherwise.
func (o *LLMObsCreatePromptDataAttributes) GetEnvIds() []string {
	if o == nil || o.EnvIds == nil {
		var ret []string
		return ret
	}
	return o.EnvIds
}

// GetEnvIdsOk returns a tuple with the EnvIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptDataAttributes) GetEnvIdsOk() (*[]string, bool) {
	if o == nil || o.EnvIds == nil {
		return nil, false
	}
	return &o.EnvIds, true
}

// HasEnvIds returns a boolean if a field has been set.
func (o *LLMObsCreatePromptDataAttributes) HasEnvIds() bool {
	return o != nil && o.EnvIds != nil
}

// SetEnvIds gets a reference to the given []string and assigns it to the EnvIds field.
func (o *LLMObsCreatePromptDataAttributes) SetEnvIds(v []string) {
	o.EnvIds = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
// Deprecated
func (o *LLMObsCreatePromptDataAttributes) GetLabels() []LLMObsPromptVersionLabel {
	if o == nil || o.Labels == nil {
		var ret []LLMObsPromptVersionLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LLMObsCreatePromptDataAttributes) GetLabelsOk() (*[]LLMObsPromptVersionLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LLMObsCreatePromptDataAttributes) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []LLMObsPromptVersionLabel and assigns it to the Labels field.
// Deprecated
func (o *LLMObsCreatePromptDataAttributes) SetLabels(v []LLMObsPromptVersionLabel) {
	o.Labels = v
}

// GetPromptId returns the PromptId field value.
func (o *LLMObsCreatePromptDataAttributes) GetPromptId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptDataAttributes) GetPromptIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptId, true
}

// SetPromptId sets field value.
func (o *LLMObsCreatePromptDataAttributes) SetPromptId(v string) {
	o.PromptId = v
}

// GetTemplate returns the Template field value.
func (o *LLMObsCreatePromptDataAttributes) GetTemplate() LLMObsPromptTemplate {
	if o == nil {
		var ret LLMObsPromptTemplate
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptDataAttributes) GetTemplateOk() (*LLMObsPromptTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value.
func (o *LLMObsCreatePromptDataAttributes) SetTemplate(v LLMObsPromptTemplate) {
	o.Template = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LLMObsCreatePromptDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LLMObsCreatePromptDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LLMObsCreatePromptDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetUserVersion returns the UserVersion field value if set, zero value otherwise.
func (o *LLMObsCreatePromptDataAttributes) GetUserVersion() string {
	if o == nil || o.UserVersion == nil {
		var ret string
		return ret
	}
	return *o.UserVersion
}

// GetUserVersionOk returns a tuple with the UserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptDataAttributes) GetUserVersionOk() (*string, bool) {
	if o == nil || o.UserVersion == nil {
		return nil, false
	}
	return o.UserVersion, true
}

// HasUserVersion returns a boolean if a field has been set.
func (o *LLMObsCreatePromptDataAttributes) HasUserVersion() bool {
	return o != nil && o.UserVersion != nil
}

// SetUserVersion gets a reference to the given string and assigns it to the UserVersion field.
func (o *LLMObsCreatePromptDataAttributes) SetUserVersion(v string) {
	o.UserVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCreatePromptDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EnvIds != nil {
		toSerialize["env_ids"] = o.EnvIds
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["prompt_id"] = o.PromptId
	toSerialize["template"] = o.Template
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.UserVersion != nil {
		toSerialize["user_version"] = o.UserVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCreatePromptDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                    `json:"description,omitempty"`
		EnvIds      []string                   `json:"env_ids,omitempty"`
		Labels      []LLMObsPromptVersionLabel `json:"labels,omitempty"`
		PromptId    *string                    `json:"prompt_id"`
		Template    *LLMObsPromptTemplate      `json:"template"`
		Title       *string                    `json:"title,omitempty"`
		UserVersion *string                    `json:"user_version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PromptId == nil {
		return fmt.Errorf("required field prompt_id missing")
	}
	if all.Template == nil {
		return fmt.Errorf("required field template missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "env_ids", "labels", "prompt_id", "template", "title", "user_version"})
	} else {
		return err
	}
	o.Description = all.Description
	o.EnvIds = all.EnvIds
	o.Labels = all.Labels
	o.PromptId = *all.PromptId
	o.Template = *all.Template
	o.Title = all.Title
	o.UserVersion = all.UserVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
