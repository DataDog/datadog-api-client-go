// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAILLMEvaluationsIntegrationAccountSettingsRequest Settings for creating the OpenAI LLM Evaluations integration account.
type OpenAILLMEvaluationsIntegrationAccountSettingsRequest struct {
	// Comma-separated list of custom tags applied to the data collected for this account, for example `team:ai,env:prod`.
	CustomTags *string `json:"custom_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAILLMEvaluationsIntegrationAccountSettingsRequest instantiates a new OpenAILLMEvaluationsIntegrationAccountSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAILLMEvaluationsIntegrationAccountSettingsRequest() *OpenAILLMEvaluationsIntegrationAccountSettingsRequest {
	this := OpenAILLMEvaluationsIntegrationAccountSettingsRequest{}
	return &this
}

// NewOpenAILLMEvaluationsIntegrationAccountSettingsRequestWithDefaults instantiates a new OpenAILLMEvaluationsIntegrationAccountSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAILLMEvaluationsIntegrationAccountSettingsRequestWithDefaults() *OpenAILLMEvaluationsIntegrationAccountSettingsRequest {
	this := OpenAILLMEvaluationsIntegrationAccountSettingsRequest{}
	return &this
}

// GetCustomTags returns the CustomTags field value if set, zero value otherwise.
func (o *OpenAILLMEvaluationsIntegrationAccountSettingsRequest) GetCustomTags() string {
	if o == nil || o.CustomTags == nil {
		var ret string
		return ret
	}
	return *o.CustomTags
}

// GetCustomTagsOk returns a tuple with the CustomTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountSettingsRequest) GetCustomTagsOk() (*string, bool) {
	if o == nil || o.CustomTags == nil {
		return nil, false
	}
	return o.CustomTags, true
}

// HasCustomTags returns a boolean if a field has been set.
func (o *OpenAILLMEvaluationsIntegrationAccountSettingsRequest) HasCustomTags() bool {
	return o != nil && o.CustomTags != nil
}

// SetCustomTags gets a reference to the given string and assigns it to the CustomTags field.
func (o *OpenAILLMEvaluationsIntegrationAccountSettingsRequest) SetCustomTags(v string) {
	o.CustomTags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAILLMEvaluationsIntegrationAccountSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomTags != nil {
		toSerialize["custom_tags"] = o.CustomTags
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAILLMEvaluationsIntegrationAccountSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTags *string `json:"custom_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.CustomTags = all.CustomTags

	return nil
}
