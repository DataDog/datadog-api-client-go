// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpenAICodexIntegrationAccountSettingsResponse Settings configured on the OpenAI Codex integration account.
type OpenAICodexIntegrationAccountSettingsResponse struct {
	// ID of the Codex workspace to collect data from.
	CodexWorkspaceId string `json:"codex_workspace_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpenAICodexIntegrationAccountSettingsResponse instantiates a new OpenAICodexIntegrationAccountSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpenAICodexIntegrationAccountSettingsResponse(codexWorkspaceId string) *OpenAICodexIntegrationAccountSettingsResponse {
	this := OpenAICodexIntegrationAccountSettingsResponse{}
	this.CodexWorkspaceId = codexWorkspaceId
	return &this
}

// NewOpenAICodexIntegrationAccountSettingsResponseWithDefaults instantiates a new OpenAICodexIntegrationAccountSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpenAICodexIntegrationAccountSettingsResponseWithDefaults() *OpenAICodexIntegrationAccountSettingsResponse {
	this := OpenAICodexIntegrationAccountSettingsResponse{}
	return &this
}

// GetCodexWorkspaceId returns the CodexWorkspaceId field value.
func (o *OpenAICodexIntegrationAccountSettingsResponse) GetCodexWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CodexWorkspaceId
}

// GetCodexWorkspaceIdOk returns a tuple with the CodexWorkspaceId field value
// and a boolean to check if the value has been set.
func (o *OpenAICodexIntegrationAccountSettingsResponse) GetCodexWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodexWorkspaceId, true
}

// SetCodexWorkspaceId sets field value.
func (o *OpenAICodexIntegrationAccountSettingsResponse) SetCodexWorkspaceId(v string) {
	o.CodexWorkspaceId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpenAICodexIntegrationAccountSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["codex_workspace_id"] = o.CodexWorkspaceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpenAICodexIntegrationAccountSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodexWorkspaceId *string `json:"codex_workspace_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CodexWorkspaceId == nil {
		return fmt.Errorf("required field codex_workspace_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"codex_workspace_id"})
	} else {
		return err
	}
	o.CodexWorkspaceId = *all.CodexWorkspaceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
