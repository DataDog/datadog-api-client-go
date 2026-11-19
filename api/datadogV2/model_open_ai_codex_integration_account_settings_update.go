// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountSettingsUpdate Settings for updating the OpenAI Codex integration account. Only the fields provided are changed.
type OpenAICodexIntegrationAccountSettingsUpdate struct {
	// ID of the Codex workspace to collect data from.
	CodexWorkspaceId *string `json:"codex_workspace_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationAccountSettingsUpdate instantiates a new OpenAICodexIntegrationAccountSettingsUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationAccountSettingsUpdate() *OpenAICodexIntegrationAccountSettingsUpdate {
	this := OpenAICodexIntegrationAccountSettingsUpdate{}
	return &this
}

// NewOpenAICodexIntegrationAccountSettingsUpdateWithDefaults instantiates a new OpenAICodexIntegrationAccountSettingsUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationAccountSettingsUpdateWithDefaults() *OpenAICodexIntegrationAccountSettingsUpdate {
	this := OpenAICodexIntegrationAccountSettingsUpdate{}
	return &this
}

// GetCodexWorkspaceId returns the CodexWorkspaceId field value if set, zero value otherwise.
func (o *OpenAICodexIntegrationAccountSettingsUpdate) GetCodexWorkspaceId() string {
	if o == nil || o.CodexWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.CodexWorkspaceId
}

// GetCodexWorkspaceIdOk returns a tuple with the CodexWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountSettingsUpdate) GetCodexWorkspaceIdOk() (*string, bool) {
	if o == nil || o.CodexWorkspaceId == nil {
		return nil, false
	}
	return o.CodexWorkspaceId, true
}

// HasCodexWorkspaceId returns a boolean if a field has been set.
func (o *OpenAICodexIntegrationAccountSettingsUpdate) HasCodexWorkspaceId() bool {
	return o != nil && o.CodexWorkspaceId != nil
}

// SetCodexWorkspaceId gets a reference to the given string and assigns it to the CodexWorkspaceId field.
func (o *OpenAICodexIntegrationAccountSettingsUpdate) SetCodexWorkspaceId(v string) {
	o.CodexWorkspaceId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationAccountSettingsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CodexWorkspaceId != nil {
		toSerialize["codex_workspace_id"] = o.CodexWorkspaceId
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationAccountSettingsUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodexWorkspaceId *string `json:"codex_workspace_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.CodexWorkspaceId = all.CodexWorkspaceId

	return nil
}
