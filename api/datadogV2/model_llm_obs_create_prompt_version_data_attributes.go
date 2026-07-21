// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsCreatePromptVersionDataAttributes Attributes for creating a new version of an LLM Observability prompt. `template` is required; all other attributes are optional.
type LLMObsCreatePromptVersionDataAttributes struct {
	// Optional description of this version.
	Description *string `json:"description,omitempty"`
	// Optional feature-flag environment UUIDs the service attempts to enable and configure to use this version as their default after creation.
	EnvIds []string `json:"env_ids,omitempty"`
	// Optional labels to attach to this version. Do not use this attribute for new integrations.
	// Deprecated
	Labels []LLMObsPromptVersionLabel `json:"labels,omitempty"`
	// A text template or a list of chat messages.
	Template LLMObsPromptTemplate `json:"template"`
	// Optional user-supplied version identifier for this version.
	UserVersion *string `json:"user_version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsCreatePromptVersionDataAttributes instantiates a new LLMObsCreatePromptVersionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsCreatePromptVersionDataAttributes(template LLMObsPromptTemplate) *LLMObsCreatePromptVersionDataAttributes {
	this := LLMObsCreatePromptVersionDataAttributes{}
	this.Template = template
	return &this
}

// NewLLMObsCreatePromptVersionDataAttributesWithDefaults instantiates a new LLMObsCreatePromptVersionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsCreatePromptVersionDataAttributesWithDefaults() *LLMObsCreatePromptVersionDataAttributes {
	this := LLMObsCreatePromptVersionDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsCreatePromptVersionDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsCreatePromptVersionDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnvIds returns the EnvIds field value if set, zero value otherwise.
func (o *LLMObsCreatePromptVersionDataAttributes) GetEnvIds() []string {
	if o == nil || o.EnvIds == nil {
		var ret []string
		return ret
	}
	return o.EnvIds
}

// GetEnvIdsOk returns a tuple with the EnvIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) GetEnvIdsOk() (*[]string, bool) {
	if o == nil || o.EnvIds == nil {
		return nil, false
	}
	return &o.EnvIds, true
}

// HasEnvIds returns a boolean if a field has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) HasEnvIds() bool {
	return o != nil && o.EnvIds != nil
}

// SetEnvIds gets a reference to the given []string and assigns it to the EnvIds field.
func (o *LLMObsCreatePromptVersionDataAttributes) SetEnvIds(v []string) {
	o.EnvIds = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
// Deprecated
func (o *LLMObsCreatePromptVersionDataAttributes) GetLabels() []LLMObsPromptVersionLabel {
	if o == nil || o.Labels == nil {
		var ret []LLMObsPromptVersionLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LLMObsCreatePromptVersionDataAttributes) GetLabelsOk() (*[]LLMObsPromptVersionLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []LLMObsPromptVersionLabel and assigns it to the Labels field.
// Deprecated
func (o *LLMObsCreatePromptVersionDataAttributes) SetLabels(v []LLMObsPromptVersionLabel) {
	o.Labels = v
}

// GetTemplate returns the Template field value.
func (o *LLMObsCreatePromptVersionDataAttributes) GetTemplate() LLMObsPromptTemplate {
	if o == nil {
		var ret LLMObsPromptTemplate
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) GetTemplateOk() (*LLMObsPromptTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value.
func (o *LLMObsCreatePromptVersionDataAttributes) SetTemplate(v LLMObsPromptTemplate) {
	o.Template = v
}

// GetUserVersion returns the UserVersion field value if set, zero value otherwise.
func (o *LLMObsCreatePromptVersionDataAttributes) GetUserVersion() string {
	if o == nil || o.UserVersion == nil {
		var ret string
		return ret
	}
	return *o.UserVersion
}

// GetUserVersionOk returns a tuple with the UserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) GetUserVersionOk() (*string, bool) {
	if o == nil || o.UserVersion == nil {
		return nil, false
	}
	return o.UserVersion, true
}

// HasUserVersion returns a boolean if a field has been set.
func (o *LLMObsCreatePromptVersionDataAttributes) HasUserVersion() bool {
	return o != nil && o.UserVersion != nil
}

// SetUserVersion gets a reference to the given string and assigns it to the UserVersion field.
func (o *LLMObsCreatePromptVersionDataAttributes) SetUserVersion(v string) {
	o.UserVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsCreatePromptVersionDataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["template"] = o.Template
	if o.UserVersion != nil {
		toSerialize["user_version"] = o.UserVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsCreatePromptVersionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                    `json:"description,omitempty"`
		EnvIds      []string                   `json:"env_ids,omitempty"`
		Labels      []LLMObsPromptVersionLabel `json:"labels,omitempty"`
		Template    *LLMObsPromptTemplate      `json:"template"`
		UserVersion *string                    `json:"user_version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Template == nil {
		return fmt.Errorf("required field template missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "env_ids", "labels", "template", "user_version"})
	} else {
		return err
	}
	o.Description = all.Description
	o.EnvIds = all.EnvIds
	o.Labels = all.Labels
	o.Template = *all.Template
	o.UserVersion = all.UserVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
