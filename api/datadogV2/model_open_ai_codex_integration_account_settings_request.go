// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountSettingsRequest Settings for creating the OpenAI Codex integration account.
type OpenAICodexIntegrationAccountSettingsRequest struct {
	// ID of the Codex workspace to collect data from.
	CodexWorkspaceId string `json:"codex_workspace_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationAccountSettingsRequest instantiates a new OpenAICodexIntegrationAccountSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationAccountSettingsRequest(codexWorkspaceId string) *OpenAICodexIntegrationAccountSettingsRequest {
	this := OpenAICodexIntegrationAccountSettingsRequest{}
	this.CodexWorkspaceId = codexWorkspaceId
	return &this
}

// NewOpenAICodexIntegrationAccountSettingsRequestWithDefaults instantiates a new OpenAICodexIntegrationAccountSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationAccountSettingsRequestWithDefaults() *OpenAICodexIntegrationAccountSettingsRequest {
	this := OpenAICodexIntegrationAccountSettingsRequest{}
	return &this
}

// GetCodexWorkspaceId returns the CodexWorkspaceId field value.
func (o *OpenAICodexIntegrationAccountSettingsRequest) GetCodexWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CodexWorkspaceId
}

// GetCodexWorkspaceIdOk returns a tuple with the CodexWorkspaceId field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountSettingsRequest) GetCodexWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodexWorkspaceId, true
}

// SetCodexWorkspaceId sets field value.
func (o *OpenAICodexIntegrationAccountSettingsRequest) SetCodexWorkspaceId(v string) {
	o.CodexWorkspaceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationAccountSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["codex_workspace_id"] = o.CodexWorkspaceId
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationAccountSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodexWorkspaceId *string `json:"codex_workspace_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CodexWorkspaceId == nil {
		return fmt.Errorf("required field codex_workspace_id missing")
	}
	o.CodexWorkspaceId = *all.CodexWorkspaceId

	return nil
}
