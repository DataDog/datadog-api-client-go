// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsUpdatePromptVersionDataAttributes Attributes for updating an LLM Observability prompt version. At least one of `description`, `labels`, or `env_ids` must be provided; all three attributes are optional individually.
type LLMObsUpdatePromptVersionDataAttributes struct {
	// Optional new description for this version.
	Description *string `json:"description,omitempty"`
	// Optional feature-flag environment UUIDs the service attempts to enable and configure to use this version as their default.
	EnvIds []string `json:"env_ids,omitempty"`
	// Optional new labels for this version. Do not use this attribute for new integrations.
	// Deprecated
	Labels []LLMObsPromptVersionLabel `json:"labels,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLLMObsUpdatePromptVersionDataAttributes instantiates a new LLMObsUpdatePromptVersionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsUpdatePromptVersionDataAttributes() *LLMObsUpdatePromptVersionDataAttributes {
	this := LLMObsUpdatePromptVersionDataAttributes{}
	return &this
}

// NewLLMObsUpdatePromptVersionDataAttributesWithDefaults instantiates a new LLMObsUpdatePromptVersionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsUpdatePromptVersionDataAttributesWithDefaults() *LLMObsUpdatePromptVersionDataAttributes {
	this := LLMObsUpdatePromptVersionDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsUpdatePromptVersionDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsUpdatePromptVersionDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsUpdatePromptVersionDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsUpdatePromptVersionDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnvIds returns the EnvIds field value if set, zero value otherwise.
func (o *LLMObsUpdatePromptVersionDataAttributes) GetEnvIds() []string {
	if o == nil || o.EnvIds == nil {
		var ret []string
		return ret
	}
	return o.EnvIds
}

// GetEnvIdsOk returns a tuple with the EnvIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsUpdatePromptVersionDataAttributes) GetEnvIdsOk() (*[]string, bool) {
	if o == nil || o.EnvIds == nil {
		return nil, false
	}
	return &o.EnvIds, true
}

// HasEnvIds returns a boolean if a field has been set.
func (o *LLMObsUpdatePromptVersionDataAttributes) HasEnvIds() bool {
	return o != nil && o.EnvIds != nil
}

// SetEnvIds gets a reference to the given []string and assigns it to the EnvIds field.
func (o *LLMObsUpdatePromptVersionDataAttributes) SetEnvIds(v []string) {
	o.EnvIds = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
// Deprecated
func (o *LLMObsUpdatePromptVersionDataAttributes) GetLabels() []LLMObsPromptVersionLabel {
	if o == nil || o.Labels == nil {
		var ret []LLMObsPromptVersionLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LLMObsUpdatePromptVersionDataAttributes) GetLabelsOk() (*[]LLMObsPromptVersionLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LLMObsUpdatePromptVersionDataAttributes) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []LLMObsPromptVersionLabel and assigns it to the Labels field.
// Deprecated
func (o *LLMObsUpdatePromptVersionDataAttributes) SetLabels(v []LLMObsPromptVersionLabel) {
	o.Labels = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsUpdatePromptVersionDataAttributes) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsUpdatePromptVersionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                    `json:"description,omitempty"`
		EnvIds      []string                   `json:"env_ids,omitempty"`
		Labels      []LLMObsPromptVersionLabel `json:"labels,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Description = all.Description
	o.EnvIds = all.EnvIds
	o.Labels = all.Labels

	return nil
}
